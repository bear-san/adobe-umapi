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
	"fmt"
	"github.com/bear-san/adobe-umapi/pkg/user"
	"os"
	"testing"
	"time"
)

func TestExec(t *testing.T) {
	caller := NewCaller(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))

	federatedIdDomain := os.Getenv("FEDERATED_ID_DOMAIN")
	enterpriseIdDomain := os.Getenv("ENTERPRISE_ID_DOMAIN")

	t.Run("Create User", func(t *testing.T) {
		// 作成系のテスト
		userRequests := []user.Request{
			// EnterpriseIDの作成
			{
				User:       fmt.Sprintf("tools-test-enterprise-id@%s", enterpriseIdDomain),
				RequestID:  "Create_Enterprise_ID",
				Domain:     nil,
				UseAdobeID: false,
				Do: []user.Command{
					{
						CreateEnterpriseID: &user.CreateEnterpriseIDRequest{
							Email:     fmt.Sprintf("tools-test-enterprise-id@%s", enterpriseIdDomain),
							Country:   "JP",
							FirstName: "Test",
							LastName:  "EnterpriseID",
							Option:    "updateIfAlreadyExists",
						},
					},
				},
			},
			// FederatedIDの作成
			{
				User:       "tools-test",
				RequestID:  "Create_Federated_ID",
				Domain:     &federatedIdDomain,
				UseAdobeID: false,
				Do: []user.Command{
					{
						CreateFederatedID: &user.CreateFederatedIDRequest{
							Email:     fmt.Sprintf("tools-test-federated-id@%s", federatedIdDomain),
							Country:   "JP",
							FirstName: "Test",
							LastName:  "FederatedID",
							Option:    "updateIfAlreadyExists",
						},
					},
				},
			},
		}

		result, err := caller.Exec(&userRequests, os.Getenv("ORG_ID"))
		if err != nil {
			t.Errorf("failed to create user: %v", err)
		}

		if result.Completed != 2 {
			t.Errorf("failed to create user: %v", result)
		}
	})

	t.Run("Add Group", func(t *testing.T) {
		// グループ追加のテスト
		userRequests := []user.Request{
			{
				User:       fmt.Sprintf("tools-test-enterprise-id@%s", enterpriseIdDomain),
				RequestID:  "Add_Group_Enterprise_ID",
				Domain:     nil,
				UseAdobeID: false,
				Do: []user.Command{
					{
						AddGroup: &user.AddGroupRequest{
							Group: []string{"test-group"},
						},
					},
				},
			},
			{
				User:       fmt.Sprintf("tools-test-federated-id@%s", federatedIdDomain),
				RequestID:  "Add_Group_Federated_ID",
				Domain:     nil,
				UseAdobeID: false,
				Do: []user.Command{
					{
						AddGroup: &user.AddGroupRequest{
							Group: []string{"test-group"},
						},
					},
				},
			},
		}

		result, err := caller.Exec(&userRequests, os.Getenv("ORG_ID"))
		if err != nil {
			t.Errorf("failed to add group: %v", err)
		}

		if result.Completed != 2 {
			t.Errorf("failed to add group: %v", result)
		}
	})

	time.Sleep(5 * time.Second)

	t.Run("Update User", func(t *testing.T) {
		// 更新系のテスト
		userRequests := []user.Request{
			{
				User:       fmt.Sprintf("tools-test-enterprise-id@%s", enterpriseIdDomain),
				RequestID:  "Update_Enterprise_ID",
				Domain:     nil,
				UseAdobeID: false,
				Do: []user.Command{
					{
						Update: &user.UpdateRequest{
							Email:     fmt.Sprintf("tools-test-enterprise-id@%s", enterpriseIdDomain),
							FirstName: "Updated",
							LastName:  "EnterpriseID",
							UserName:  fmt.Sprintf("tools-test-enterprise-id@%s", enterpriseIdDomain),
						},
					},
				},
			},
			{
				User:       fmt.Sprintf("tools-test-federated-id@%s", federatedIdDomain),
				RequestID:  "Update_Federated_ID",
				Domain:     nil,
				UseAdobeID: false,
				Do: []user.Command{
					{
						Update: &user.UpdateRequest{
							Email:     fmt.Sprintf("tools-test-federated-id@%s", federatedIdDomain),
							FirstName: "Updated",
							LastName:  "FederatedID",
							UserName:  fmt.Sprintf("tools-test-federated-id@%s", federatedIdDomain),
						},
					},
				},
			},
		}

		result, err := caller.Exec(&userRequests, os.Getenv("ORG_ID"))
		if err != nil {
			t.Errorf("failed to update user: %v", err)
		}

		if result.Completed != 2 {
			t.Errorf("failed to update user: %v", result)
		}
	})

	time.Sleep(5 * time.Second)

	t.Run("Remove User from Group", func(t *testing.T) {
		// グループからの削除のテスト
		userRequests := []user.Request{
			{
				User:       fmt.Sprintf("tools-test-enterprise-id@%s", enterpriseIdDomain),
				RequestID:  "Remove_Group_Enterprise_ID",
				Domain:     nil,
				UseAdobeID: false,
				Do: []user.Command{
					{
						RemoveFromGroup: &user.RemoveFromGroupRequest{
							Group: []string{"test-group"},
						},
					},
				},
			},
			{
				User:       fmt.Sprintf("tools-test-federated-id@%s", federatedIdDomain),
				RequestID:  "Remove_Group_Federated_ID",
				Domain:     nil,
				UseAdobeID: false,
				Do: []user.Command{
					{
						RemoveFromGroup: &user.RemoveFromGroupRequest{
							Group: []string{"test-group"},
						},
					},
				},
			},
		}

		result, err := caller.Exec(&userRequests, os.Getenv("ORG_ID"))
		if err != nil {
			t.Errorf("failed to remove user from group: %v", err)
		}

		if result.Completed != 2 {
			t.Errorf("failed to remove user from group: %v", result)
		}
	})

	time.Sleep(5 * time.Second)

	t.Run("Remove User from Org", func(t *testing.T) {
		// 組織からの削除のテスト
		userRequests := []user.Request{
			{
				User:       fmt.Sprintf("tools-test-enterprise-id@%s", enterpriseIdDomain),
				RequestID:  "Remove_Org_Enterprise_ID",
				Domain:     nil,
				UseAdobeID: false,
				Do: []user.Command{
					{
						RemoveFromOrg: &user.RemoveFromOrgRequest{
							DeleteAccount: true,
						},
					},
				},
			},
			{
				User:       fmt.Sprintf("tools-test-federated-id@%s", federatedIdDomain),
				RequestID:  "Remove_Org_Federated_ID",
				Domain:     nil,
				UseAdobeID: false,
				Do: []user.Command{
					{
						RemoveFromOrg: &user.RemoveFromOrgRequest{
							DeleteAccount: true,
						},
					},
				},
			},
		}

		result, err := caller.Exec(&userRequests, os.Getenv("ORG_ID"))
		if err != nil {
			t.Errorf("failed to remove user from org: %v", err)
		}

		if result.Completed != 2 {
			t.Errorf("failed to remove user from org: %v", result)
		}
	})
}
