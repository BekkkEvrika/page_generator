package keycloak

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func GetServiceToken(
	baseURL string, // https://admin.obukanal.ru/keycloak
	realm string, // core
	clientID string, // api-gateway
	secret string, // client secret
) (string, int, error) {

	form := url.Values{}
	form.Set("grant_type", "client_credentials")
	form.Set("client_id", clientID)
	form.Set("client_secret", secret)

	tokenURL := fmt.Sprintf(
		"%s/realms/%s/protocol/openid-connect/token",
		baseURL,
		realm,
	)

	req, err := http.NewRequest(
		http.MethodPost,
		tokenURL,
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return "", 0, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", 0, fmt.Errorf(
			"token request failed: %d %s",
			resp.StatusCode,
			string(body),
		)
	}

	var tr TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return "", 0, err
	}

	return tr.AccessToken, tr.ExpiresIn, nil
}
