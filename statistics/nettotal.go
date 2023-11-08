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
	// NetTotalReq 查询总流量
	// http://~/v1/net_total?node_id=14&kvm=kvm1146&type=net_adapter&kvm_ifname=kvm1146.0&status=1&st=1699432018000&et=1699439971121
	NetTotalReq struct {
		NodeId    uint   `url:"node_id,omitempty"`    // node_id
		KvmIfName string `url:"kvm_ifname,omitempty"` // $kvmid.$NetworkOffset
	}
	NetTotalResp struct {
		UtcTime string `json:"utc_time"` // utc time
		Accept  int64  `json:"accept"`   // 接收流量(Byte)
		Send    int64  `json:"send"`     // 发送流量(Byte)
	}
)

func (r *NetTotalReq) Url() string                 { return "/v1/net_total" }
func (r *NetTotalReq) Method() string              { return http.MethodGet }
func (r *NetTotalReq) Header() http.Header         { return nil }
func (r *NetTotalReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *NetTotalReq) Form() (form url.Values)     { return }
func (r *NetTotalReq) Body() any                   { return nil }

func netTotal(ctx *cputil.Context, req *NetTotalReq) ([]any, error) {
	if resp, err := cputil.Execute[*NetTotalReq, []any](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}

func NetTotal(ctx *cputil.Context, req *NetTotalReq) (*NetTotalResp, error) {
	totals, err := netTotal(ctx, req)
	if err != nil {
		return nil, err
	}
	var (
		utcTime string
		accept  int64
		send    int64
	)
	if totals != nil && len(totals) >= 3 {
		if n := totals[0]; n != nil {
			utcTime = n.(string)
		}
		if n := totals[1]; n != nil {
			accept = int64(n.(float64))
		}
		if n := totals[2]; n != nil {
			send = int64(n.(float64))
		}
	}
	return &NetTotalResp{utcTime, accept, send}, err
}
