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
	OnReq struct {
		ID uint `json:"-"` // 实例ID
	}
	OnResp struct {
		TaskID string `json:"taskid"`
	}
)

func (r *OnReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/on", r.ID) }
func (r *OnReq) Method() string              { return http.MethodPost }
func (r *OnReq) Header() http.Header         { return nil }
func (r *OnReq) Values() (values url.Values) { return }
func (r *OnReq) Form() (form url.Values)     { return }
func (r *OnReq) Body() any                   { return nil }

// On 开机
func On(ctx *cputil.Context, req *OnReq) (*OnResp, error) {
	if resp, err := cputil.Execute[*OnReq, *OnResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
