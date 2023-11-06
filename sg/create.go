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
	"github.com/rwscode/cputil"
	"net/http"
	"net/url"
)

type (
	CreateReq struct {
		Name        string `json:"name"`        // 名称
		Description string `json:"description"` // 描述
		Uid         int    `json:"uid"`         // 用户ID
		Hostid      int    `json:"hostid"`      // 实例ID(如果有添加并关联到该实例)
		Rid         int    `json:"rid"`         // 资源包ID
		Type        string `json:"type"`        // 安全组类型(host=专业魔方云,lightHost=轻量魔方云,hyperv=Hyper-V,v2.5.2+)
	}
	CreateResp struct {
		ID string `json:"id"` // 安全组ID
	}
)

func (r *CreateReq) Url() string                 { return "/v1/security_groups" }
func (r *CreateReq) Method() string              { return http.MethodPost }
func (r *CreateReq) Header() http.Header         { return nil }
func (r *CreateReq) Values() (values url.Values) { return }
func (r *CreateReq) Form() (form url.Values)     { return }
func (r *CreateReq) Body() any                   { return r }

// Create 创建安全组
func Create(ctx *cputil.Context, req *CreateReq) (*CreateResp, error) {
	if resp, err := cputil.Execute[*CreateReq, *CreateResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
