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

package vpcnetworks

import (
	"errors"
	"github.com/go-the-way/cputil"
	q "github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

type (
	ListReq struct {
		Search    string `url:"search,omitempty"`     // 搜索
		Node      string `url:"node,omitempty"`       // 节点ID
		Area      string `url:"area,omitempty"`       // 区域ID
		NodeGroup string `url:"node_group,omitempty"` // 节点分组ID
		User      string `url:"user,omitempty"`       // 用户ID
		ListType  string `url:"list_type,omitempty"`  // 传all获取所有

		Page    string `url:"page,omitempty"`
		PerPage string `url:"per_page,omitempty"`
		Orderby string `url:"orderby,omitempty"` // 排序(id,size,name,type)
		Sort    string `url:"sort,omitempty"`
	}
	ListResp struct {
		Data []struct {
			Id       int    `json:"id"`       // vpc网络ID
			Name     string `json:"name"`     // vpc网络名称
			Ips      string `json:"ips"`      // 网段
			Uid      int    `json:"uid"`      // 用户ID
			Username string `json:"username"` // 用户名
			Host     []struct {
				Id       int    `json:"id"`       // 使用该VPC实例ID
				Hostname string `json:"hostname"` // 使用该VPC实例主机名
			} `json:"host"`
			VpcMac []struct {
				NodeName string `json:"node_name"` //	节点名称(v2.3.9+)
				Mac      string `json:"mac"`       // vpcMAC地址(v2.3.9+)
			} `json:"vpc_mac"`
		} `json:"data"`
		Meta struct {
			Total     int `json:"total"`
			TotalPage int `json:"total_page"`
			Page      int `json:"page"`
			PerPage   int `json:"per_page"`
		} `json:"meta"`
	}
)

func (r *ListReq) Url() string                 { return "/v1/vpc_networks" }
func (r *ListReq) Method() string              { return http.MethodGet }
func (r *ListReq) Header() http.Header         { return nil }
func (r *ListReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *ListReq) Form() (form url.Values)     { return }
func (r *ListReq) Body() any                   { return nil }

// List vpc网络列表
func List(ctx *cputil.Context, req *ListReq) (*ListResp, error) {
	if resp, err := cputil.Execute[*ListReq, *ListResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
