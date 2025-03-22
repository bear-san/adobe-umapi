// Copyright 2024 Kentaro Abe
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

func Setup(clientId string, clientSecret string) (*AccessTokenPayload, error) {
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
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		errPayload, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("failed to get access token: %v", string(errPayload))
	}

	Credential = AccessTokenPayload{}
	err = json.NewDecoder(res.Body).Decode(&Credential)
	if err != nil {
		return nil, err
	}

	return &Credential, nil
}

type AccessTokenPayload struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}
