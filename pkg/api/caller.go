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

package api

import (
	"encoding/json"
	"fmt"
	"github.com/bear-san/adobe-umapi/pkg/user"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	authBaseUrl = "https://ims-na1.adobelogin.com/ims/token/v2"
	baseUrl     = "https://usermanagement.adobe.io/v2/usermanagement"
)

type Caller struct {
	clientID     string
	clientSecret string
}

func (caller Caller) Exec(userRequests *[]user.Request, orgId string) (*Result, error) {
	credential, err := caller.authSetup(caller.clientID, caller.clientSecret)
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(userRequests)
	if err != nil {
		return nil, err
	}

	httpClient := http.DefaultClient
	req, err := http.NewRequest("POST", baseUrl+"/action/"+orgId, strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+credential.AccessToken)
	req.Header.Add("X-Api-Key", caller.clientID)

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		errorPayload, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("failed to execute request: %v", string(errorPayload))
	}

	resultText, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	result := Result{}
	err = json.Unmarshal(resultText, &result)

	return &result, err
}

func (caller Caller) authSetup(clientId string, clientSecret string) (*AccessTokenPayload, error) {
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	httpClient := http.DefaultClient
	values := url.Values{}

	values.Set("grant_type", "client_credentials")
	values.Set("client_id", clientId)
	values.Set("client_secret", clientSecret)
	values.Set("scope", "openid,AdobeID,user_management_sdk")

	req, err := http.NewRequest("POST", authBaseUrl, strings.NewReader(values.Encode()))
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

	credential := AccessTokenPayload{}
	err = json.NewDecoder(res.Body).Decode(&credential)
	if err != nil {
		return nil, err
	}

	return &credential, nil
}

func NewCaller(clientID string, clientSecret string) *Caller {
	return &Caller{
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

type AccessTokenPayload struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type Result struct {
	Completed           int        `json:"completed"`
	NotCompleted        int        `json:"notCompleted"`
	CompletedInTestMode int        `json:"completedInTestMode"`
	Errors              *[]Error   `json:"errors,omitempty"`
	Warnings            *[]Warning `json:"warnings,omitempty"`
	Result              string     `json:"result"`
}

type Error struct {
	Index     int    `json:"index"`
	Step      int    `json:"step"`
	RequestID string `json:"requestID"`
	Message   string `json:"message"`
	User      string `json:"user"`
	ErrorCode string `json:"errorCode"`
}

type Warning struct {
	Index       int    `json:"index"`
	Step        int    `json:"step"`
	RequestID   string `json:"requestID"`
	Message     string `json:"message"`
	User        string `json:"user"`
	WarningCode string `json:"warningCode"`
}
