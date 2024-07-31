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

package snapshots

import (
	"errors"
	"fmt"
	"github.com/go-the-way/cputil"
	"net/http"
	"net/url"
)

type (
	RestoreReq struct {
		ID uint // 备份/快照ID
	}
	RestoreResp struct {
		TaskID string `json:"taskid"`
		Os     struct {
			Id       int    `json:"id"`       // 还原后镜像ID
			Name     string `json:"name"`     // 还原后镜像名称
			User     string `json:"user"`     // 还原后用户名
			Password string `json:"password"` // 还原后密码
			Port     int    `json:"port"`     // 还原后端口
		} `json:"os"`
	}
)

func (r *RestoreReq) Url() string                 { return fmt.Sprintf("/v1/snapshots/%d/restore", r.ID) }
func (r *RestoreReq) Method() string              { return http.MethodPost }
func (r *RestoreReq) Header() http.Header         { return nil }
func (r *RestoreReq) Values() (values url.Values) { return }
func (r *RestoreReq) Form() (form url.Values)     { return }
func (r *RestoreReq) Body() any                   { return nil }

// Restore 备份/快照还原
func Restore(ctx *cputil.Context, req *RestoreReq) (*RestoreResp, error) {
	if resp, err := cputil.Execute[*RestoreReq, *RestoreResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
