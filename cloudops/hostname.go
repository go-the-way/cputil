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
	HostnameReq struct {
		ID       uint   `json:"-"`        // 实例ID
		Hostname string `json:"hostname"` // 主机名
	}
	HostnameResp struct{}
)

func (r *HostnameReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/hostname", r.ID) }
func (r *HostnameReq) Method() string              { return http.MethodPut }
func (r *HostnameReq) Header() http.Header         { return nil }
func (r *HostnameReq) Values() (values url.Values) { return }
func (r *HostnameReq) Form() (form url.Values)     { return }
func (r *HostnameReq) Body() any                   { return r }

// Hostname 修改主机名
func Hostname(ctx *cputil.Context, req *HostnameReq) (*HostnameResp, error) {
	if resp, err := cputil.Execute[*HostnameReq, *HostnameResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
