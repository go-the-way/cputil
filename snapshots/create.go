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
	"github.com/rwscode/cputil"
	"net/http"
	"net/url"
)

type (
	CreateReq struct {
		ID      uint   `json:"-"`        // 实例ID
		Type    string `json:"type"`     // 备份类型(snap快照,backup备份)
		Name    string `json:"name"`     // 备注名
		StoreId int    `json:"store_id"` // 存储ID(创建备份可以指定)
	}
	CreateResp struct {
		TaskID string `json:"taskid"`
	}
)

func (r *CreateReq) Url() string                 { return fmt.Sprintf("/v1/disks/%d/snapshots", r.ID) }
func (r *CreateReq) Method() string              { return http.MethodPost }
func (r *CreateReq) Header() http.Header         { return nil }
func (r *CreateReq) Values() (values url.Values) { return }
func (r *CreateReq) Form() (form url.Values)     { return }
func (r *CreateReq) Body() any                   { return r }

// Create 创建备份/快照
func Create(ctx *cputil.Context, req *CreateReq) (*CreateResp, error) {
	if resp, err := cputil.Execute[*CreateReq, *CreateResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
