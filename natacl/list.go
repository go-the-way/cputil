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

package natacl

import (
	"errors"
	"fmt"
	"github.com/go-the-way/cputil"
	q "github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

type (
	ListReq struct {
		ID       uint   `url:"-"`                   // 实例ID
		Search   string `url:"search,omitempty"`    // 搜索
		ListType string `url:"list_type,omitempty"` // 获取类型(all,page),all会忽略页数直接返回所有

		Page    string `url:"page,omitempty"`
		PerPage string `url:"per_page,omitempty"`
		Orderby string `url:"orderby,omitempty"` // 排序(id,name,ext_port,int_port,protocol)
		Sort    string `url:"sort,omitempty"`
	}
	ListResp struct {
		NatHostIp string `json:"nat_host_ip"`
		Data      []struct {
			ID       uint   `json:"id"`
			Name     string `json:"name"`
			ExtPort  uint   `json:"ext_port"`
			IntPort  uint   `json:"int_port"`
			Protocol uint   `json:"protocol"`
		} `json:"data"`
		Meta struct {
			Total     int `json:"total"`
			TotalPage int `json:"total_page"`
			Page      int `json:"page"`
			PerPage   int `json:"per_page"`
		} `json:"meta"`
	}
)

func (r *ListReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/nat_acl", r.ID) }
func (r *ListReq) Method() string              { return http.MethodGet }
func (r *ListReq) Header() http.Header         { return nil }
func (r *ListReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *ListReq) Form() (form url.Values)     { return }
func (r *ListReq) Body() any                   { return nil }

// List NAT转发列表
func List(ctx *cputil.Context, req *ListReq) (*ListResp, error) {
	if resp, err := cputil.Execute[*ListReq, *ListResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
