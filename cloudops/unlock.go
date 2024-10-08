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
	UnlockReq struct {
		ID uint `json:"-"` // 实例ID
	}
	UnlockResp struct{}
)

func (r *UnlockReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/unlock", r.ID) }
func (r *UnlockReq) Method() string              { return http.MethodPost }
func (r *UnlockReq) Header() http.Header         { return nil }
func (r *UnlockReq) Values() (values url.Values) { return }
func (r *UnlockReq) Form() (form url.Values)     { return }
func (r *UnlockReq) Body() any                   { return nil }

// Unlock 实例解除锁定
func Unlock(ctx *cputil.Context, req *UnlockReq) (*UnlockResp, error) {
	if resp, err := cputil.Execute[*UnlockReq, *UnlockResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
