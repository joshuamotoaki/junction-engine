package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"github.com/tigerappsorg/junction-engine/config"
	"github.com/tigerappsorg/junction-engine/models"
)

type casService struct {
	loginURL  string
	logoutURL string
	appURL    string
	casURL    string
	jwtSecret string
	jwtExpiry time.Duration
}

type CASService interface {
	GetLoginURL() string
	GetLogoutURL() string
	ValidateTicket(ticket string) (*models.User, error)
	GenerateJWT(user *models.User) (string, error)
	ValidateJWT(tokenString string) (*models.JWTClaims, error)
}

func NewCASService(cfg *config.Config) CASService {
	loginUrl := fmt.Sprintf("%s/login?service=%s", cfg.CASServerURL, url.QueryEscape(cfg.AppURL+"/auth/callback"))
	logoutUrl := fmt.Sprintf("%s/logout", cfg.CASServerURL)

	return &casService{
		loginURL:  loginUrl,
		logoutURL: logoutUrl,
		appURL:    cfg.AppURL,
		casURL:    cfg.CASServerURL,
		jwtSecret: cfg.JWTSecret,
		jwtExpiry: cfg.JWTExpiry,
	}
}

func (c *casService) GetLoginURL() string {
	return c.loginURL
}

func (c *casService) GetLogoutURL() string {
	return c.logoutURL
}

// Validate a CAS ticket and returns user information
func (c *casService) ValidateTicket(ticket string) (*models.User, error) {
	serviceURL := c.appURL + "/auth/callback"
	validateURL := fmt.Sprintf("%s/p3/serviceValidate?service=%s&ticket=%s&format=json",
		c.casURL,
		url.QueryEscape(serviceURL),
		url.QueryEscape(ticket))

	log.Debug().
		Str("ticket", ticket).
		Str("service_url", serviceURL).
		Str("validate_url", validateURL).
		Msg("Validating CAS ticket")

	resp, err := http.Get(validateURL)
	if err != nil {
		return nil, fmt.Errorf("failed to validate ticket: %w", err)
	}
	defer resp.Body.Close()

	log.Debug().Msgf("CAS validation response status: %s", resp.Status)

	var casResp models.CASResponse
	if err := json.NewDecoder(resp.Body).Decode(&casResp); err != nil {
		return nil, fmt.Errorf("failed to decode CAS response: %w", err)
	}

	if casResp.ServiceResponse.AuthenticationFailure != nil {
		return nil, fmt.Errorf("CAS authentication failed: %s",
			casResp.ServiceResponse.AuthenticationFailure.Message)
	}

	if casResp.ServiceResponse.AuthenticationSuccess == nil {
		return nil, fmt.Errorf("no authentication success in CAS response")
	}

	auth := casResp.ServiceResponse.AuthenticationSuccess

	// Extract class year 
	year := "Graduate"
	for _, group := range auth.Attributes.GrouperGroups {
		if strings.Contains(group, "PU:basis:classyear:") {
			parts := strings.Split(group, ":")
			if len(parts) > 3 {
				year = parts[3]
				break
			}
		}
	}

	// Get display name and email
	displayName := "Student"
	if len(auth.Attributes.DisplayName) > 0 {
		displayName = auth.Attributes.DisplayName[0]
	}

	email := ""
	if len(auth.Attributes.Mail) > 0 {
		email = auth.Attributes.Mail[0]
	}

	return &models.User{
		NetID:     auth.User,
		Name:      displayName,
		Email:     email,
		ClassYear: year,
		CreatedAt: time.Now(),
	}, nil
}

// Create a JWT token for the user
func (c *casService) GenerateJWT(user *models.User) (string, error) {
	claims := models.JWTClaims{
		NetID: user.NetID,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(c.jwtExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "junction-engine",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(c.jwtSecret))
}

// Validate a JWT token and returns the user claims
func (c *casService) ValidateJWT(tokenString string) (*models.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(c.jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
