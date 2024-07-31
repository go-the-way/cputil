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
	LinksReq struct {
		Cloud []uint `json:"cloud"` // 实例ID
		ID    uint   `json:"id"`    // 安全组ID
		Type  uint   `json:"type"`  // 类型(0其他该安全组实例会被移除,1不影响该安全组其他实例关联)
	}
	LinksResp struct{}
)

func (r *LinksReq) Url() string                 { return fmt.Sprintf("/v1/security_groups/%d/links", r.ID) }
func (r *LinksReq) Method() string              { return http.MethodPost }
func (r *LinksReq) Header() http.Header         { return nil }
func (r *LinksReq) Values() (values url.Values) { return }
func (r *LinksReq) Form() (form url.Values)     { return }
func (r *LinksReq) Body() any                   { return r }

// Links 关联安全组
func Links(ctx *cputil.Context, req *LinksReq) (*LinksResp, error) {
	if resp, err := cputil.Execute[*LinksReq, *LinksResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
