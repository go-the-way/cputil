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
	PanelPassReq struct {
		ID        uint   `json:"-"`          // 实例ID
		PanelPass string `json:"panel_pass"` //	面板管理密码
	}
	PanelPassResp struct{}
)

func (r *PanelPassReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/panel_pass", r.ID) }
func (r *PanelPassReq) Method() string              { return http.MethodPut }
func (r *PanelPassReq) Header() http.Header         { return nil }
func (r *PanelPassReq) Values() (values url.Values) { return }
func (r *PanelPassReq) Form() (form url.Values)     { return }
func (r *PanelPassReq) Body() any                   { return r }

// PanelPass 修改面板管理密码
// 修改面板管理密码(v2.2.2+)
func PanelPass(ctx *cputil.Context, req *PanelPassReq) (*PanelPassResp, error) {
	if resp, err := cputil.Execute[*PanelPassReq, *PanelPassResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
