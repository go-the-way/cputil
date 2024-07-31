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
	"github.com/go-the-way/cputil"
	"net/http"
	"net/url"
)

type (
	UnSuspendReq struct {
		ID uint `json:"-"` // 实例ID
	}
	UnSuspendResp struct{}
)

func (r *UnSuspendReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/unsuspend", r.ID) }
func (r *UnSuspendReq) Method() string              { return http.MethodPost }
func (r *UnSuspendReq) Header() http.Header         { return nil }
func (r *UnSuspendReq) Values() (values url.Values) { return }
func (r *UnSuspendReq) Form() (form url.Values)     { return }
func (r *UnSuspendReq) Body() any                   { return nil }

// UnSuspend 解除暂停
func UnSuspend(ctx *cputil.Context, req *UnSuspendReq) (*UnSuspendResp, error) {
	if resp, err := cputil.Execute[*UnSuspendReq, *UnSuspendResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
