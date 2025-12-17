package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchContacts(accessToken string, page int) (map[string]interface{}, error) {
	url := fmt.Sprintf(
		"https://api.rd.services/platform/contacts?page=%d",
		page,
	)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	fmt.Println("API raw response:", string(raw))

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, raw)
	}

	var result map[string]interface{}
	json.Unmarshal(raw, &result)

	return result, nil
}
