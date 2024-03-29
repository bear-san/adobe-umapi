package auth

import (
	"encoding/json"
	"net/http"
)

var baseUrl = "https://ims-na1.adobelogin.com/ims/token/v2"
var Credential AccessTokenPayload

func Setup(clientId string, clientSecret string) error {
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	httpClient := http.DefaultClient
	req, err := http.NewRequest("POST", baseUrl, nil)
	if err != nil {
		return err
	}

	req.Form.Add("grant_type", "client_credentials")
	req.Form.Add("client_id", clientId)
	req.Form.Add("client_secret", clientSecret)
	req.Form.Add("scope", "openid,AdobeID,user_management_sdk")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return err
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
