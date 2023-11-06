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

package sg

import (
	"errors"
	q "github.com/google/go-querystring/query"
	"github.com/rwscode/cputil"
	"net/http"
	"net/url"
)

type (
	ListReq struct {
		Search   string `url:"search,omitempty"`    // 搜索
		ListType string `url:"list_type,omitempty"` // 获取类型(all,page),all会忽略页数直接返回所有
		Type     string `url:"list_type,omitempty"` // 安全组类型(host=专业魔方云,lightHost=轻量魔方云,hyperv=Hyper-V,v2.5.2+)
		User     string `url:"user,omitempty"`      // 用户ID

		Page    string `url:"page,omitempty"`
		PerPage string `url:"per_page,omitempty"`
		Orderby string `url:"orderby,omitempty"` // 排序(id,name,ext_port,int_port,protocol)
		Sort    string `url:"sort,omitempty"`
	}
	ListResp struct {
		Data []struct {
			Id          int    `json:"id"`          // 安全组ID
			Name        string `json:"name"`        // 名称
			Description string `json:"description"` // 描述
			Uid         int    `json:"uid"`         // 用户ID
			CloudNum    int    `json:"cloud_num"`   // 关联数
			Username    string `json:"username"`    // 用户名
			RuleNum     int    `json:"rule_num"`    // 安全组规则数量
			CreateTime  string `json:"create_time"` // 创建时间
			Type        string `json:"type"`        // 安全组类型(host = 专业魔方云, lightHost = 轻量魔方云, hyperv = Hyper-V, v2.5.2+)
		} `json:"data"`
		Meta struct {
			Total     int `json:"total"`
			TotalPage int `json:"total_page"`
			Page      int `json:"page"`
			PerPage   int `json:"per_page"`
		} `json:"meta"`
	}
)

func (r *ListReq) Url() string                 { return "/v1/security_groups" }
func (r *ListReq) Method() string              { return http.MethodGet }
func (r *ListReq) Header() http.Header         { return nil }
func (r *ListReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *ListReq) Form() (form url.Values)     { return }
func (r *ListReq) Body() any                   { return nil }

// List 安全组列表
func List(ctx *cputil.Context, req *ListReq) (*ListResp, error) {
	if resp, err := cputil.Execute[*ListReq, *ListResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
