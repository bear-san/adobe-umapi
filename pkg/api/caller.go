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
	"github.com/bear-san/adobe-umapi/pkg/auth"
	"github.com/bear-san/adobe-umapi/pkg/user"
	"io"
	"net/http"
	"strings"
)

var BaseUrl = "https://usermanagement.adobe.io/v2/usermanagement"

func Exec(userRequests *[]user.Request, auth auth.AccessTokenPayload, orgId string, apiKey string) (*Result, error) {
	payload, err := json.Marshal(userRequests)
	if err != nil {
		return nil, err
	}

	httpClient := http.DefaultClient
	req, err := http.NewRequest("POST", BaseUrl+"/action/"+orgId, strings.NewReader(string(payload)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+auth.AccessToken)
	req.Header.Add("X-Api-Key", apiKey)

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
