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
	SuspendReq struct {
		ID   uint   `json:"-"`    // 实例ID
		Type string `json:"type"` // 暂停类型(traffic=流量超额,due=到期,v2.3.9+)
	}
	SuspendResp struct{}
)

func (r *SuspendReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/suspend", r.ID) }
func (r *SuspendReq) Method() string              { return http.MethodPost }
func (r *SuspendReq) Header() http.Header         { return nil }
func (r *SuspendReq) Values() (values url.Values) { return }
func (r *SuspendReq) Form() (form url.Values)     { return }
func (r *SuspendReq) Body() any                   { return r }

// Suspend 暂停
func Suspend(ctx *cputil.Context, req *SuspendReq) (*SuspendResp, error) {
	if resp, err := cputil.Execute[*SuspendReq, *SuspendResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
