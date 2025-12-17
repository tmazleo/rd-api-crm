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
func GetAccessToken(tokenURL string) (*models.TokenResponse, error) {
	payload := map[string]string{
		"client_id":     os.Getenv("RD_CLIENT_ID"),
		"client_secret": os.Getenv("RD_CLIENT_SECRET"),
		"code":          os.Getenv("RD_AUTH_CODE"),
		"redirect_uri":  os.Getenv("RD_REDIRECT_URI"),
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

	// Leia o body cru (debug)
	rawBody, _ := io.ReadAll(resp.Body)
	fmt.Println("Token raw response:", string(rawBody))

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("token error: %s", rawBody)
	}

	var token models.TokenResponse
	json.Unmarshal(rawBody, &token)

	return &token, nil
}