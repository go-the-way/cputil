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
	"github.com/rwscode/cputil"
	"net/http"
	"net/url"
)

type (
	CreateReq struct {
		ID          uint   `json:"-"`           // 安全组ID
		Description string `json:"description"` // 描述
		Direction   string `json:"direction"`   // 规则方向(in进方向,out出方向)
		Protocol    string `json:"protocol"`    // 协议(all,all_tcp,all_udp,tcp,udp,icmp,ssh,telnet,http,https,mssql,oracle,mysql,rdp,postgresql,redis,gre)
		Port        string `json:"port"`        // 端口范围(专业版)
		Ip          string `json:"ip"`          // 授权IP(专业版,Hyper-V)
		Lock        int    `json:"lock"`        // 是否锁定(0不锁定,1锁定,v2.3.0+)
		StartIp     string `json:"start_ip"`    // 起始IP(轻量版)
		EndIp       string `json:"end_ip"`      // 结束IP(轻量版)
		StartPort   int    `json:"start_port"`  // 起始端口(轻量版,Hyper-V)
		EndPort     int    `json:"end_port"`    // 结束端口(轻量版,Hyper-V)
		Priority    int    `json:"priority"`    // 优先级(轻量版)
		Action      string `json:"action"`      // 授权策略(accept=允许,drop=拒绝,轻量版,Hyper-V)
	}
	CreateResp struct{}
)

func (r *CreateReq) Url() string                 { return fmt.Sprintf("/v1/security_groups/%d/rules", r.ID) }
func (r *CreateReq) Method() string              { return http.MethodPost }
func (r *CreateReq) Header() http.Header         { return nil }
func (r *CreateReq) Values() (values url.Values) { return }
func (r *CreateReq) Form() (form url.Values)     { return }
func (r *CreateReq) Body() any                   { return r }

// Create 添加安全组规则
func Create(ctx *cputil.Context, req *CreateReq) (*CreateResp, error) {
	if resp, err := cputil.Execute[*CreateReq, *CreateResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
