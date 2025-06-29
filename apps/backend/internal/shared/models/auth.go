package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CasUser struct {
	NetID     string    `json:"netid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ClassYear string    `json:"class_year"`
	CreatedAt time.Time `json:"created_at"`
}

type JWTClaims struct {
	NetID string `json:"netid"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type CASResponse struct {
	ServiceResponse struct {
		AuthenticationSuccess *struct {
			User       string `json:"user"`
			Attributes struct {
				DisplayName   []string `json:"displayname"`
				Mail          []string `json:"mail"`
				GrouperGroups []string `json:"grouperGroups"`
			} `json:"attributes"`
		} `json:"authenticationSuccess,omitempty"`
		AuthenticationFailure *struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"authenticationFailure,omitempty"`
	} `json:"serviceResponse"`
}
