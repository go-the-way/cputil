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
	PasswordReq struct {
		ID       uint   `json:"-"`
		Password string `json:"password"`
	}
	PasswordResp struct {
		TaskID string `json:"taskid"` // 任务ID(v2.1.8+)
	}
)

func (r *PasswordReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/password", r.ID) }
func (r *PasswordReq) Method() string              { return http.MethodPut }
func (r *PasswordReq) Header() http.Header         { return nil }
func (r *PasswordReq) Values() (values url.Values) { return }
func (r *PasswordReq) Form() (form url.Values)     { return }
func (r *PasswordReq) Body() any                   { return r }

// Password 重置密码
func Password(ctx *cputil.Context, req *PasswordReq) (*PasswordResp, error) {
	if resp, err := cputil.Execute[*PasswordReq, *PasswordResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
