package models
type TokenResponse struct {
	AcessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn int    `json:"expires_in"`
}