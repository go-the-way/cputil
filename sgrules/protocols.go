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

package sgrules

import (
	"errors"
	"github.com/go-the-way/cputil"
	"net/http"
	"net/url"
)

type (
	ProtocolsReq struct{}
	data         struct {
		Name  string `json:"name"`  // 名称
		Value string `json:"value"` // 值
		Port  string `json:"port"`  // 端口范围
	}
	ProtocolsResp struct {
		data
	}
)

func (r *ProtocolsReq) Url() string                 { return "/v1/security_group_rule_protocols" }
func (r *ProtocolsReq) Method() string              { return http.MethodGet }
func (r *ProtocolsReq) Header() http.Header         { return nil }
func (r *ProtocolsReq) Values() (values url.Values) { return }
func (r *ProtocolsReq) Form() (form url.Values)     { return }
func (r *ProtocolsReq) Body() any                   { return nil }

// Protocols 获取规则可用协议
func Protocols(ctx *cputil.Context, req *ProtocolsReq) ([]*ProtocolsResp, error) {
	if resp, err := cputil.Execute[*ProtocolsReq, []*ProtocolsResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
