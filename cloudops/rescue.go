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
	RescueEnterReq struct {
		ID       uint   `json:"-"`         // 实例ID
		Type     string `json:"type"`      // 指定救援系统类型(0=跟随实例,1=windows,2=linux)
		TempPass string `json:"temp_pass"` // 救援系统临时密码，当前系统密码
	}
	RescueEnterResp struct{}
)

func (r *RescueEnterReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/rescue", r.ID) }
func (r *RescueEnterReq) Method() string              { return http.MethodPost }
func (r *RescueEnterReq) Header() http.Header         { return nil }
func (r *RescueEnterReq) Values() (values url.Values) { return }
func (r *RescueEnterReq) Form() (form url.Values)     { return }
func (r *RescueEnterReq) Body() any                   { return nil }

func RescueEnter(ctx *cloudplatform.Context, req *RescueEnterReq) (*RescueEnterResp, error) {
	if resp, err := cloudplatform.Execute[*RescueEnterReq, *RescueEnterResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}

type (
	RescueExitReq struct {
		ID uint `json:"-"` // 实例ID
	}
	RescueExitResp struct{}
)

func (r *RescueExitReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/rescue", r.ID) }
func (r *RescueExitReq) Method() string              { return http.MethodDelete }
func (r *RescueExitReq) Header() http.Header         { return nil }
func (r *RescueExitReq) Values() (values url.Values) { return }
func (r *RescueExitReq) Form() (form url.Values)     { return }
func (r *RescueExitReq) Body() any                   { return nil }

func RescueExit(ctx *cloudplatform.Context, req *RescueExitReq) (*RescueExitResp, error) {
	if resp, err := cloudplatform.Execute[*RescueExitReq, *RescueExitResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
