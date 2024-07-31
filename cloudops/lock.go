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
	LockReq struct {
		ID uint `json:"-"` // 实例ID
	}
	LockResp struct{}
)

func (r *LockReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/lock", r.ID) }
func (r *LockReq) Method() string              { return http.MethodPost }
func (r *LockReq) Header() http.Header         { return nil }
func (r *LockReq) Values() (values url.Values) { return }
func (r *LockReq) Form() (form url.Values)     { return }
func (r *LockReq) Body() any                   { return nil }

// Lock 实例锁定
func Lock(ctx *cputil.Context, req *LockReq) (*LockResp, error) {
	if resp, err := cputil.Execute[*LockReq, *LockResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
