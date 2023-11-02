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
	"github.com/rwscode/cputil"
	"net/http"
	"net/url"
)

type (
	UpdateReq struct {
		ID          uint   `json:"-"`           // 安全组ID
		Name        string `json:"name"`        // 名称
		Description string `json:"description"` // 描述
	}
	UpdateResp struct{}
)

func (r *UpdateReq) Url() string                 { return fmt.Sprintf("/v1/security_groups/%d", r.ID) }
func (r *UpdateReq) Method() string              { return http.MethodPut }
func (r *UpdateReq) Header() http.Header         { return nil }
func (r *UpdateReq) Values() (values url.Values) { return }
func (r *UpdateReq) Form() (form url.Values)     { return }
func (r *UpdateReq) Body() any                   { return r }

// Update 修改安全组
func Update(ctx *cloudplatform.Context, req *UpdateReq) (*UpdateResp, error) {
	if resp, err := cloudplatform.Execute[*UpdateReq, *UpdateResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
