package keycloak

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

type KeycloakClient struct {
	BaseURL    string // https://admin.obukanal.ru/keycloak
	Realm      string // core
	ClientUUID string // UUID клиента api-gateway
	Token      string // ervice access tokens
	HTTP       *http.Client
}

func NewKeycloakClient(baseURL string, realm string, clientId string, clientUUID string, secret string) *KeycloakClient {
	token, ttl, err := GetServiceToken(
		baseURL,
		realm,
		clientId,
		secret,
	)
	if err != nil {
		log.Fatal("failed to get keycloak token:", err)
	}
	kcClient := KeycloakClient{
		BaseURL:    baseURL,
		Realm:      realm,
		ClientUUID: clientUUID,
		Token:      token,
		HTTP:       &http.Client{Timeout: 10 * time.Second},
	}

	log.Printf("keycloak token received, ttl=%ds", ttl)
	return &kcClient
}

func (c *KeycloakClient) auth(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")
}

func (c *KeycloakClient) url(path string) string {
	return fmt.Sprintf("%s/admin/realms/%s/clients/%s/authz/resource-server%s",
		c.BaseURL, c.Realm, c.ClientUUID, path,
	)
}

func (c *KeycloakClient) EnsureScope(name string) error {
	// 1. check
	req, _ := http.NewRequest(
		"GET",
		c.url("/scope?name="+url.QueryEscape(name)),
		nil,
	)
	c.auth(req)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var scopes []map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&scopes)

	if len(scopes) > 0 {
		return nil // ✅ already exists
	}
	// 2. create
	payload := map[string]string{
		"name": name,
	}
	body, _ := json.Marshal(payload)
	createReq, _ := http.NewRequest(
		"POST",
		c.url("/scope"),
		bytes.NewReader(body),
	)
	c.auth(createReq)
	createResp, err := c.HTTP.Do(createReq)
	if err != nil {
		return err
	}
	defer createResp.Body.Close()
	if createResp.StatusCode == http.StatusConflict {
		return nil // race-condition safe
	}
	if createResp.StatusCode >= 300 {
		return fmt.Errorf("ensure scope failed (%s): %d", name, createResp.StatusCode)
	}
	return nil
}

func (c *KeycloakClient) EnsureResource(name string, uris []string) error {
	req, _ := http.NewRequest(
		"GET",
		c.url("/resource?name="+url.QueryEscape(name)),
		nil,
	)
	c.auth(req)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var resources []map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&resources)

	if len(resources) > 0 {
		return nil // ✅ exists
	}

	// 2. create
	payload := map[string]interface{}{
		"name": name,
		"type": "model",
		"uris": uris,
	}

	body, _ := json.Marshal(payload)

	createReq, _ := http.NewRequest(
		"POST",
		c.url("/resource"),
		bytes.NewReader(body),
	)
	c.auth(createReq)

	createResp, err := c.HTTP.Do(createReq)
	if err != nil {
		return err
	}
	defer createResp.Body.Close()

	if createResp.StatusCode == http.StatusConflict {
		return nil
	}

	if createResp.StatusCode >= 300 {
		return fmt.Errorf("ensure resource failed (%s): %d", name, createResp.StatusCode)
	}

	return nil
}

func (c *KeycloakClient) EnsureScopeLinked(resourceName, scopeName string) error {
	// 1. get resource
	req, _ := http.NewRequest(
		"GET",
		c.url("/resource?name="+url.QueryEscape(resourceName)),
		nil,
	)
	c.auth(req)

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var resources []struct {
		ID     string `json:"_id"`
		Name   string `json:"name"`
		Scopes []struct {
			Name string `json:"name"`
		} `json:"scopes"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&resources); err != nil {
		return err
	}

	if len(resources) == 0 {
		return fmt.Errorf("resource not found: %s", resourceName)
	}

	res := resources[0]

	// 2. check scope linked
	for _, s := range res.Scopes {
		if s.Name == scopeName {
			return nil // ✅ already linked
		}
	}

	// 3. append scope
	var scopes []map[string]string
	for _, s := range res.Scopes {
		scopes = append(scopes, map[string]string{"name": s.Name})
	}
	scopes = append(scopes, map[string]string{"name": scopeName})

	payload := map[string]interface{}{
		"name":   res.Name,
		"scopes": scopes,
	}

	body, _ := json.Marshal(payload)

	updateReq, _ := http.NewRequest(
		"PUT",
		c.url("/resource/"+res.ID),
		bytes.NewReader(body),
	)
	c.auth(updateReq)

	updateResp, err := c.HTTP.Do(updateReq)
	if err != nil {
		return err
	}
	defer updateResp.Body.Close()

	if updateResp.StatusCode >= 300 {
		return fmt.Errorf("link scope failed (%s -> %s): %d",
			resourceName, scopeName, updateResp.StatusCode)
	}

	return nil
}
