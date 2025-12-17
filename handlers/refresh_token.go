package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"rdstation-crm/models"
)

func RefreshAccessToken(refreshToken string) (*models.TokenResponse, error) {
	tokenURL := "https://api.rd.services/auth/token"

	payload := map[string]string{
		"client_id":     os.Getenv("RD_CLIENT_ID"),
		"client_secret": os.Getenv("RD_CLIENT_SECRET"),
		"refresh_token": refreshToken,
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", tokenURL, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	fmt.Println("Refresh raw response:", string(raw))

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("refresh token error (%d): %s", resp.StatusCode, raw)
	}

	var token models.TokenResponse
	json.Unmarshal(raw, &token)

	return &token, nil
}
