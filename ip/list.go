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

package ip

import (
	"errors"
	"github.com/go-the-way/cputil"
	q "github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

type (
	ListReq struct {
		IpSegmentId string `url:"ip_segment_id,omitempty"` // 所属ip段id
		Ip          string `url:"ip,omitempty"`            // 模糊搜索ip/备注
		Type        string `url:"type,omitempty"`          // 类型搜索(cloud=实例,load_balance=负载均衡)
		HostId      string `url:"hostid,omitempty"`        // 关联id搜索
		Lock        string `url:"lock,omitempty"`          // 锁定状态(0=未锁定,1=锁定)

		Page    string `url:"page,omitempty"`
		PerPage string `url:"per_page,omitempty"`
		Orderby string `url:"orderby,omitempty"` // 排序(id,ip,ip_int,hostid,vlanid,uid,hostname,node_name,area_name,username,mac)
		Sort    string `url:"sort,omitempty"`
	}
	ListResp struct {
		Data []struct {
			Id              int         `json:"id"`
			Ip              string      `json:"ip"`
			IpSegmentId     int         `json:"ip_segment_id"`
			Hostid          int         `json:"hostid"`
			ServerMainip    string      `json:"server_mainip"`
			Vlanid          int         `json:"vlanid"`
			Remark          string      `json:"remark"`
			Lock            int         `json:"lock"`
			IpMac           string      `json:"ip_mac"`
			Type            string      `json:"type"`
			Hostname        interface{} `json:"hostname"`
			Username        interface{} `json:"username"`
			Parent          interface{} `json:"parent"`
			ParentName      string      `json:"parent_name"`
			Mac             interface{} `json:"mac"`
			ResourcePackage interface{} `json:"resource_package"`
			NodeName        string      `json:"node_name"`
			AreaName        string      `json:"area_name"`
			RCount          interface{} `json:"r_count"`
		} `json:"data"`
		Meta struct {
			Total     int `json:"total"`
			TotalPage int `json:"total_page"`
			PerPage   int `json:"per_page"`
			Page      int `json:"page"`
		} `json:"meta"`
	}
)

func (r *ListReq) Url() string                 { return "/v1/ip" }
func (r *ListReq) Method() string              { return http.MethodGet }
func (r *ListReq) Header() http.Header         { return nil }
func (r *ListReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *ListReq) Form() (form url.Values)     { return }
func (r *ListReq) Body() any                   { return nil }

// List 实例列表
func List(ctx *cputil.Context, req *ListReq) (*ListResp, error) {
	if resp, err := cputil.Execute[*ListReq, *ListResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
