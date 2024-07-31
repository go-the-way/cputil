// Copyright 2023 cputil Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloudops

import (
	"errors"
	"fmt"
	"github.com/go-the-way/cputil"
	"net/http"
	"net/url"
)

type (
	SwitchVpcReq struct {
		ID     int    `json:"-"`
		Vpc    uint   `json:"vpc"`     // VPCID
		VpcIps string `json:"vpc_ips"` // 要新建的VPC IP段(没有vpc时生效)
	}
	SwitchVpcResp struct {
		TaskID string `json:"taskid"`
	}
)

func (r *SwitchVpcReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/network_type", r.ID) }
func (r *SwitchVpcReq) Method() string              { return http.MethodPut }
func (r *SwitchVpcReq) Header() http.Header         { return nil }
func (r *SwitchVpcReq) Values() (values url.Values) { return }
func (r *SwitchVpcReq) Form() (form url.Values)     { return }
func (r *SwitchVpcReq) Body() any                   { return r }

func SwitchVpc(ctx *cputil.Context, req *SwitchVpcReq) (*SwitchVpcResp, error) {
	if resp, err := cputil.Execute[*SwitchVpcReq, *SwitchVpcResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
