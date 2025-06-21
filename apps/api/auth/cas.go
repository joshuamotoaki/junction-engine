package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tigerappsorg/junction-engine/config"
	"github.com/tigerappsorg/junction-engine/models"
)

type CASService struct {
	config *config.Config
}

func NewCASService(cfg *config.Config) *CASService {
	return &CASService{
		config: cfg,
	}
}

// GetLoginURL returns the CAS login URL
func (c *CASService) GetLoginURL() string {
	serviceURL := c.config.AppURL + "/auth/callback"
	return fmt.Sprintf("%slogin?service=%s",
		c.config.CASServerURL,
		url.QueryEscape(serviceURL))
}

// GetLogoutURL returns the CAS logout URL
func (c *CASService) GetLogoutURL() string {
	return fmt.Sprintf("%slogout", c.config.CASServerURL)
}

// ValidateTicket validates a CAS ticket and returns user information
func (c *CASService) ValidateTicket(ticket string) (*models.User, error) {
	serviceURL := c.config.AppURL + "/auth/callback"
	validateURL := fmt.Sprintf("%sp3/serviceValidate?service=%s&ticket=%s&format=json",
		c.config.CASServerURL,
		url.QueryEscape(serviceURL),
		url.QueryEscape(ticket))

	resp, err := http.Get(validateURL)
	if err != nil {
		return nil, fmt.Errorf("failed to validate ticket: %w", err)
	}
	defer resp.Body.Close()

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

	// Extract class year from grouper groups
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

// GenerateJWT creates a JWT token for the user
func (c *CASService) GenerateJWT(user *models.User) (string, error) {
	claims := models.JWTClaims{
		NetID: user.NetID,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(c.config.JWTExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "junction-engine",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(c.config.JWTSecret))
}

// ValidateJWT validates a JWT token and returns the user claims
func (c *CASService) ValidateJWT(tokenString string) (*models.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(c.config.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
