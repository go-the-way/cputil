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

package cloudops

import (
	"errors"
	"fmt"
	"github.com/rwscode/cputil"
	"net/http"
	"net/url"
)

type (
	VncReq struct {
		ID uint `json:"-"` // 实例ID
	}
	VncResp struct {
		Path        int    `json:"path"`          // 路径
		Token       string `json:"token"`         // vnc凭证
		VncPass     string `json:"vnc_pass"`      // vnc密码
		Password    string `json:"password"`      // 实例密码
		VncUrl      string `json:"vnc_url"`       // vnc wss地址
		VncUrlHttp  string `json:"vnc_url_http"`  // 外部vnc http地址,未启用时没有
		VncUrlHttps string `json:"vnc_url_https"` // 外部vnc https页面地址,未启用时没有
	}
)

func (r *VncReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/vnc", r.ID) }
func (r *VncReq) Method() string              { return http.MethodPost }
func (r *VncReq) Header() http.Header         { return nil }
func (r *VncReq) Values() (values url.Values) { return }
func (r *VncReq) Form() (form url.Values)     { return }
func (r *VncReq) Body() any                   { return nil }

// Vnc VNC
func Vnc(ctx *cputil.Context, req *VncReq) (*VncResp, error) {
	if resp, err := cputil.Execute[*VncReq, *VncResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
