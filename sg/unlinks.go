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
	"net/http"
	"net/url"
)

type (
	UnlinksReq struct {
		SgID uint `json:"-"`  // 安全组ID
		ID   uint `json:"id"` // 实例ID
	}
	UnlinksResp struct{}
)

func (r *UnlinksReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/security_groups", r.SgID) }
func (r *UnlinksReq) Method() string              { return http.MethodDelete }
func (r *UnlinksReq) Header() http.Header         { return nil }
func (r *UnlinksReq) Values() (values url.Values) { return }
func (r *UnlinksReq) Form() (form url.Values)     { return }
func (r *UnlinksReq) Body() any                   { return r }

// Unlinks 解除关联安全组
func Unlinks(ctx *cputil.Context, req *UnlinksReq) (*UnlinksResp, error) {
	if resp, err := cputil.Execute[*UnlinksReq, *UnlinksResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
