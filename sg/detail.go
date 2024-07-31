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
	"fmt"
	"github.com/go-the-way/cputil"
	q "github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

type (
	DetailReq struct {
		ID         uint   `url:"-"`
		GetAllRule string `url:"get_all_rule,omitempty"` // 是否获取所有规则(0不,1获取)
	}
	DetailResp struct {
		Id          int    `json:"id"`          // 规则ID
		Name        string `json:"name"`        // 名称
		Description string `json:"description"` // 描述
		Uid         int    `json:"uid"`         // 用户ID
		CreateTime  string `json:"create_time"` // 创建时间
		UpdateTime  string `json:"update_time"` // 修改时间
		CloudNum    int    `json:"cloud_num"`   // 关联实例数量
		Username    string `json:"username"`    // 用户名
		Type        string `json:"type"`        // 安全组类型(host=专业魔方云,lightHost=轻量魔方云,hyperv=Hyper-V,v2.5.2+)
		Rule        []struct {
			Id          int    `json:"id"`          // 安全组规则ID
			Description string `json:"description"` // 描述
			Direction   string `json:"direction"`   // 规则方向
			Protocol    string `json:"protocol"`    // 协议
			Port        string `json:"port"`        // 端口范围
			Ip          string `json:"ip"`          // 授权IP
			CreateTime  string `json:"create_time"` // 创建时间
		} `json:"rule"`
	}
)

func (r *DetailReq) Url() string                 { return fmt.Sprintf("/v1/security_groups/%d", r.ID) }
func (r *DetailReq) Method() string              { return http.MethodGet }
func (r *DetailReq) Header() http.Header         { return nil }
func (r *DetailReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *DetailReq) Form() (form url.Values)     { return }
func (r *DetailReq) Body() any                   { return nil }

// Detail 安全组详情
func Detail(ctx *cputil.Context, req *DetailReq) (*DetailResp, error) {
	if resp, err := cputil.Execute[*DetailReq, *DetailResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
