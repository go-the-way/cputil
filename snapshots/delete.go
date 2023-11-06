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
	DeleteReq struct {
		ID uint // 备份/快照ID
	}
	DeleteResp struct {
		TaskID string `json:"taskid"`
	}
)

func (r *DeleteReq) Url() string                 { return fmt.Sprintf("/v1/snapshots/%d", r.ID) }
func (r *DeleteReq) Method() string              { return http.MethodDelete }
func (r *DeleteReq) Header() http.Header         { return nil }
func (r *DeleteReq) Values() (values url.Values) { return }
func (r *DeleteReq) Form() (form url.Values)     { return }
func (r *DeleteReq) Body() any                   { return nil }

// Delete 删除备份/快照
func Delete(ctx *cputil.Context, req *DeleteReq) (*DeleteResp, error) {
	if resp, err := cputil.Execute[*DeleteReq, *DeleteResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
