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

package tasks

import (
	"errors"
	"fmt"
	"github.com/go-the-way/cputil"
	"net/http"
	"net/url"
)

type (
	CancelReq  struct{ TaskID uint }
	CancelResp struct{}
)

func (r *CancelReq) Url() string                 { return fmt.Sprintf("/v1/tasks/%d/cancel", r.TaskID) }
func (r *CancelReq) Method() string              { return http.MethodPost }
func (r *CancelReq) Header() http.Header         { return nil }
func (r *CancelReq) Values() (values url.Values) { return }
func (r *CancelReq) Form() (form url.Values)     { return }
func (r *CancelReq) Body() any                   { return nil }

// Cancel 取消任务
func Cancel(ctx *cputil.Context, req *CancelReq) (*CancelResp, error) {
	if resp, err := cputil.Execute[*CancelReq, *CancelResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
