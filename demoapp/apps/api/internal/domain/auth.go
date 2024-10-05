package domain

import "time"

type UserInfo struct {
	Sub           string    `json:"sub"`
	GivenName     string    `json:"given_name"`
	FamilyName    string    `json:"family_name"`
	Nickname      string    `json:"nickname"`
	Name          string    `json:"name"`
	Picture       string    `json:"picture"`
	UpdatedAt     time.Time `json:"updated_at"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"email_verified"`
}
