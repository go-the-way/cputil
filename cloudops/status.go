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
	StatusReq struct {
		ID uint `json:"-"` // 实例ID
	}
	StatusResp struct {
		Status   string `json:"status"`    // 实例状态(unknown未知,wait_reboot等待重启,on开机,off关机,task任务中,suspend暂停,recycle回收站中,paused挂起,cold_migrate冷迁移,hot_migrate热迁移)
		Task     string `json:"task"`      // 任务类型(当status=task时返回)
		TaskName string `json:"task_name"` // 任务类型名称(当status=task时返回)
	}
)

func (r *StatusReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/status", r.ID) }
func (r *StatusReq) Method() string              { return http.MethodGet }
func (r *StatusReq) Header() http.Header         { return nil }
func (r *StatusReq) Values() (values url.Values) { return }
func (r *StatusReq) Form() (form url.Values)     { return }
func (r *StatusReq) Body() any                   { return nil }

// Status 获取实例状态
func Status(ctx *cputil.Context, req *StatusReq) (*StatusResp, error) {
	if resp, err := cputil.Execute[*StatusReq, *StatusResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
