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

package user

type Request struct {
	User       string    `json:"user"`
	RequestID  string    `json:"requestID"`
	Domain     *string   `json:"domain,omitempty"`
	UseAdobeID bool      `json:"useAdobeID"`
	Do         []Command `json:"do"`
}

type Command struct {
	AddAdobeID         *AddAdobeIDRequest         `json:"addAdobeID,omitempty"`
	CreateEnterpriseID *CreateEnterpriseIDRequest `json:"createEnterpriseID,omitempty"`
	CreateFederatedID  *CreateFederatedIDRequest  `json:"createFederatedID,omitempty"`
	AddGroup           *AddGroupRequest           `json:"add,omitempty"`
	Update             *UpdateRequest             `json:"update,omitempty"`
	RemoveFromGroup    *RemoveFromGroupRequest    `json:"remove,omitempty"`
	RemoveFromOrg      *RemoveFromOrgRequest      `json:"removeFromOrg,omitempty"`
}
