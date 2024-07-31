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
	MainipReq struct {
		ID uint `json:"id"` // 实例ID
		Ip uint `json:"ip"` // IP ID
	}
	MainipResp struct{}
)

func (r *MainipReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/mainip", r.ID) }
func (r *MainipReq) Method() string              { return http.MethodPost }
func (r *MainipReq) Header() http.Header         { return nil }
func (r *MainipReq) Values() (values url.Values) { return }
func (r *MainipReq) Form() (form url.Values)     { return }
func (r *MainipReq) Body() any                   { return r }

// Mainip 设置主IP
func Mainip(ctx *cputil.Context, req *MainipReq) (*MainipResp, error) {
	if resp, err := cputil.Execute[*MainipReq, *MainipResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
