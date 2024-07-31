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

package ipsegment

import (
	"errors"
	"github.com/go-the-way/cputil"
	q "github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

type (
	ListReq struct {
		Status  string `url:"status,omitempty"`   // 状态搜索
		Node    string `url:"node,omitempty"`     // 节点id(v2.2.2+)
		IpGroup string `url:"ip_group,omitempty"` // IP分组id(v2.2.2+)
		IpName  string `url:"ip_name,omitempty"`  // 搜索名称
		Area    string `url:"area,omitempty"`     // 区域id(v2.4.3+)
		Rid     string `url:"rid,omitempty"`      // 资源包ID(前台用户)

		Page    string `url:"page,omitempty"`     // 分页数,指定当前第几页
		PerPage string `url:"per_page,omitempty"` // 指定当前页面显示条数
		Orderby string `url:"orderby,omitempty"`  // 排序
		Sort    string `url:"sort,omitempty"`     // 指定当前页面显示条数
	}
	ListResp struct {
		Data []struct {
			Id         int    `json:"id"`
			IpName     string `json:"ip_name"`
			Gateway    string `json:"gateway"`
			SubnetMask string `json:"subnet_mask"`
			Dns1       string `json:"dns1"`
			Dns2       string `json:"dns2"`
			CreateTime string `json:"create_time"`
			Status     int    `json:"status"`
			NodeIds    string `json:"node_ids"`
			GroupId    int    `json:"group_id"`
			GroupName  string `json:"group_name"`
			Count      struct {
				Free  int `json:"free"`
				Used  int `json:"used"`
				Total int `json:"total"`
				Lock  int `json:"lock"`
			} `json:"count"`
			Node []struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
			} `json:"node"`
		} `json:"data"`
		Meta struct {
			Total     int `json:"total"`
			TotalPage int `json:"total_page"`
			PerPage   int `json:"per_page"`
			Page      int `json:"page"`
		} `json:"meta"`
	}
)

func (r *ListReq) Url() string                 { return "/v1/ipSegment" }
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
