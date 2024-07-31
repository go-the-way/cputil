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
	"fmt"
	"github.com/go-the-way/cputil"
	q "github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

type (
	ListReq struct {
		ID        uint   `url:"-"`                   // 安全组ID
		Search    string `url:"search,omitempty"`    // 搜索
		Direction string `url:"direction,omitempty"` // 规则方向筛选(in=进方向,out=出方向,v2.4.1+)
		PerPage   string `url:"per_page,omitempty"`
	}
	ListResp struct {
		Data []struct {
			Id          int    `json:"id"`          // 安全组规则ID
			Description string `json:"description"` // 描述
			Direction   string `json:"direction"`   // 规则方向
			Protocol    string `json:"protocol"`    // 协议
			Port        string `json:"port"`        // 端口范围(专业版)
			Ip          string `json:"ip"`          // 授权IP(专业版,Hyper-V)
			CreateTime  string `json:"create_time"` // 创建时间
			Lock        int    `json:"lock"`        // 是否锁定(0不锁定,1锁定,v2.3.0+)
			StartIp     string `json:"start_ip"`    // 起始IP(轻量版)
			EndIp       string `json:"end_ip"`      // 结束IP(轻量版)
			StartPort   int    `json:"start_port"`  // 起始端口(轻量版,Hyper-V)
			EndPort     int    `json:"end_port"`    // 结束端口(轻量版,Hyper-V)
			Priority    int    `json:"priority"`    // 优先级(轻量版)
			Action      string `json:"action"`      // 授权策略(accept=允许,drop=拒绝,轻量版,Hyper-V)
		} `json:"data"`
		Meta struct {
			Total           int `json:"total"`
			TotalPage       int `json:"total_page"`
			Page            int `json:"page"`
			PerPage         int `json:"per_page"`
			DirectionInNum  int `json:"direction_in_num"`  // 进方向规则总数(v2.4.1+)
			DirectionOutNum int `json:"direction_out_num"` // 出方向规则总数(v2.4.1+)
		} `json:"meta"`
	}
)

func (r *ListReq) Url() string                 { return fmt.Sprintf("/v1/security_groups/%d/rules", r.ID) }
func (r *ListReq) Method() string              { return http.MethodGet }
func (r *ListReq) Header() http.Header         { return nil }
func (r *ListReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *ListReq) Form() (form url.Values)     { return }
func (r *ListReq) Body() any                   { return nil }

// List 安全组规则列表
func List(ctx *cputil.Context, req *ListReq) (*ListResp, error) {
	if resp, err := cputil.Execute[*ListReq, *ListResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
