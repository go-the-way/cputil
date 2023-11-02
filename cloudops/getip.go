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
	"github.com/rwscode/cputil"
	"net/http"
	"net/url"
)

type (
	GetIpReq struct {
		ID uint `json:"-"` // 实例ID
	}
	GetIpResp struct {
		NetworkType string `json:"network_type"` // 网络类型(normal=经典网络,vpc=VPC网络)
		Interface   []struct {
			Name     string `json:"name"` // 网卡名称
			PublicIp []struct {
				Id int    `json:"id"` // 网卡公网IP ID
				Ip string `json:"ip"` // 网卡公网IP地址
			} `json:"public_ip"`
			PrivateIp string `json:"private_ip"` // 网卡内网IP地址
		} `json:"interface"`
		PublicIp []struct {
			Id int    `json:"id"` // 网卡公网IP ID
			Ip string `json:"ip"` // 网卡公网IP地址
		} `json:"public_ip"`
		PrivateIp []string `json:"private_ip"` // 网卡内网IP地址
	}
)

func (r *GetIpReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/ip", r.ID) }
func (r *GetIpReq) Method() string              { return http.MethodGet }
func (r *GetIpReq) Header() http.Header         { return nil }
func (r *GetIpReq) Values() (values url.Values) { return }
func (r *GetIpReq) Form() (form url.Values)     { return }
func (r *GetIpReq) Body() any                   { return nil }

// GetIp 获取实例IP
func GetIp(ctx *cloudplatform.Context, req *GetIpReq) (*GetIpResp, error) {
	if resp, err := cloudplatform.Execute[*GetIpReq, *GetIpResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
