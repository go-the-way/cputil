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

package statistics

import (
	"errors"
	q "github.com/google/go-querystring/query"
	"github.com/rwscode/cputil"
	"net/http"
	"net/url"
)

type (
	// Req
	// http://~/v1/statistics?node_id=13&kvm=kvm1145&type=kvm_info&st=1699344000000&et=1699344060000
	Req struct {
		Kvm       string `url:"kvm,omitempty"`  // kvmid
		Type      string `url:"type,omitempty"` // kvm_info
		StartTime int64  `url:"st,omitempty"`   // 开始毫秒时间戳(1699344000000) UTC+8
		EndTime   int64  `url:"et,omitempty"`   // 结束毫秒时间戳(1699344060000) UTC+8

		DevName   string `url:"dev_name,omitempty"`   // for disk io
		KvmIfName string `url:"kvm_ifname,omitempty"` // for net adapter
		Status    string `url:"status,omitempty"`     // for net adapter
	}
)

func (r *Req) Url() string                 { return "/v1/statistics" }
func (r *Req) Method() string              { return http.MethodGet }
func (r *Req) Header() http.Header         { return nil }
func (r *Req) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *Req) Form() (form url.Values)     { return }
func (r *Req) Body() any                   { return nil }

func Statistics(ctx *cputil.Context, req *Req) ([][]any, error) {
	if resp, err := cputil.Execute[*Req, [][]any](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
