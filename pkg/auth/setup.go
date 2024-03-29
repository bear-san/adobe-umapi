package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var baseUrl = "https://ims-na1.adobelogin.com/ims/token/v2"
var Credential AccessTokenPayload

func Setup(clientId string, clientSecret string) error {
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	httpClient := http.DefaultClient
	values := url.Values{}

	values.Set("grant_type", "client_credentials")
	values.Set("client_id", clientId)
	values.Set("client_secret", clientSecret)
	values.Set("scope", "openid,AdobeID,user_management_sdk")

	req, err := http.NewRequest("POST", baseUrl, strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		errPayload, _ := io.ReadAll(res.Body)
		return fmt.Errorf("failed to get access token: %v", string(errPayload))
	}

	Credential = AccessTokenPayload{}
	err = json.NewDecoder(res.Body).Decode(&Credential)
	if err != nil {
		return err
	}

	return nil
}

type AccessTokenPayload struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}
