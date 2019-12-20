// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"time"
)

type AuthResponse struct {
	TokenToStore string `json:"tokenToStore"`
	TTL          int    `json:"ttl"`
	RefreshToken string `json:"refreshToken"`
}

type NewUser struct {
	Email        string `json:"email"`
	Token        string `json:"token"`
	AuthMethodID int    `json:"authMethodId"`
}

type UpdatedProfile struct {
	FirstName   *string    `json:"firstName"`
	LastName    *string    `json:"lastName"`
	DateOfBirth *time.Time `json:"dateOfBirth"`
}

type User struct {
	Email     string     `json:"email"`
	Role      int        `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type UserProfile struct {
	FirstName   *string    `json:"firstName"`
	LastName    *string    `json:"lastName"`
	DateOfBirth *time.Time `json:"dateOfBirth"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

type LoginDetails struct {
	Identification string `json:"identification"`
	Token          string `json:"token"`
	AuthMethodID   int    `json:"authMethodId"`
}
