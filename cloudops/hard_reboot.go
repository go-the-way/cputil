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
	HardRebootReq struct {
		ID uint `json:"-"` // 实例ID
	}
	HardRebootResp struct {
		TaskID string `json:"taskid"`
	}
)

func (r *HardRebootReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/hard_reboot", r.ID) }
func (r *HardRebootReq) Method() string              { return http.MethodPost }
func (r *HardRebootReq) Header() http.Header         { return nil }
func (r *HardRebootReq) Values() (values url.Values) { return }
func (r *HardRebootReq) Form() (form url.Values)     { return }
func (r *HardRebootReq) Body() any                   { return nil }

// HardReboot 硬重启
func HardReboot(ctx *cputil.Context, req *HardRebootReq) (*HardRebootResp, error) {
	if resp, err := cputil.Execute[*HardRebootReq, *HardRebootResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
