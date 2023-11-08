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
	"github.com/rwscode/cputil"
	"net/http"
	"net/url"

	q "github.com/google/go-querystring/query"
)

type (
	ListReq struct {
		Status  string `url:"status,omitempty"`   // 状态(0未开始,1正在执行,2执行成功,3执行失败,4强制结束,5已取消)
		Cloud   string `url:"cloud,omitempty"`    // 实例ID(传入ID获取某个实例的任务列表)
		RelType string `url:"rel_type,omitempty"` // 任务关联类型(cloud=实例,load_balance=负载均衡)

		Page    string `url:"page,omitempty"`
		PerPage string `url:"per_page,omitempty"`
		Orderby string `url:"orderby,omitempty"` // 排序(id,status,start_time,end_time,progress)
		Sort    string `url:"sort,omitempty"`
	}
	ListResp struct {
		Data []struct {
			Id         int    `json:"id"`
			Type       string `json:"type"`
			Status     int    `json:"status"`
			CreateTime string `json:"create_time"`
			StartTime  string `json:"start_time"`
			EndTime    string `json:"end_time"`
			Progress   int    `json:"progress"`
			Msg        string `json:"msg"`
			RelType    string `json:"rel_type"`
			Hostid     int    `json:"hostid"`
			Hostname   string `json:"hostname"`
		} `json:"data"`
		Meta struct {
			Total     int `json:"total"`
			TotalPage int `json:"total_page"`
			Page      int `json:"page"`
			PerPage   int `json:"per_page"`
		} `json:"meta"`
	}
)

func (r *ListReq) Url() string                 { return "/v1/tasks" }
func (r *ListReq) Method() string              { return http.MethodGet }
func (r *ListReq) Header() http.Header         { return nil }
func (r *ListReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *ListReq) Form() (form url.Values)     { return }
func (r *ListReq) Body() any                   { return nil }

// List 任务列表
func List(ctx *cputil.Context, req *ListReq) (*ListResp, error) {
	if resp, err := cputil.Execute[*ListReq, *ListResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
