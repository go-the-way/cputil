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
	"github.com/rwscode/cputil"
	"net/http"
	"net/url"
)

type (
	DetailReq  struct{ TaskID uint }
	DetailResp struct {
		Id         int           `json:"id"`
		Data       []interface{} `json:"data"`
		Type       string        `json:"type"`
		Hostid     int           `json:"hostid"`
		Status     int           `json:"status"`
		CreateTime string        `json:"create_time"`
		StartTime  string        `json:"start_time"`
		EndTime    string        `json:"end_time"`
		Progress   int           `json:"progress"`
		Msg        string        `json:"msg"`
		UserType   int           `json:"user_type"`
		Uid        int           `json:"uid"`
		Ip         string        `json:"ip"`
		TypeDesc   string        `json:"type_desc"`
	}
)

func (r *DetailReq) Url() string                 { return fmt.Sprintf("/v1/tasks/%d", r.TaskID) }
func (r *DetailReq) Method() string              { return http.MethodGet }
func (r *DetailReq) Header() http.Header         { return nil }
func (r *DetailReq) Values() (values url.Values) { return }
func (r *DetailReq) Form() (form url.Values)     { return }
func (r *DetailReq) Body() any                   { return nil }

// Detail 任务详情
func Detail(ctx *cputil.Context, req *DetailReq) (*DetailResp, error) {
	if resp, err := cputil.Execute[*DetailReq, *DetailResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
