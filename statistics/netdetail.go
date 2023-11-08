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
	// NetDetailReq
	// http://~/v1/net_detail?node_id=14&kvm_ifname=kvm1146.0&st=1699432018000&et=1699438740820
	NetDetailReq struct {
		NodeId    uint   `url:"node_id,omitempty"`    // node_id
		KvmIfName string `url:"kvm_ifname,omitempty"` // $kvmid.$NetworkOffset
		StartTime int64  `url:"st,omitempty"`         // 开始毫秒时间戳(1699344000000) UTC+8
		EndTime   int64  `url:"et,omitempty"`         // 结束毫秒时间戳(1699344060000) UTC+8
	}
	NetDetailResp struct {
		Accept struct {
			Total   float64 `json:"total"`   // 总流量(Byte)
			Max     float64 `json:"max"`     // 最大流量(Byte)
			New     float64 `json:"new"`     // 最新流量(Byte)
			Average float64 `json:"average"` // 平均流量(Byte)
		} `json:"accept"` // 流入流量(Byte)
		Send struct {
			Total   float64 `json:"total"`   // 总流量(Byte)
			Max     float64 `json:"max"`     // 最大流量(Byte)
			New     float64 `json:"new"`     // 最新流量(Byte)
			Average float64 `json:"average"` // 平均流量(Byte)
		} `json:"send"` // 流出流量(Byte)
		Max struct {
			Total  float64 `json:"total"`
			Send   float64 `json:"send"`
			Accept float64 `json:"accept"`
		} `json:"95"` // ignored
	}
)

func (r *NetDetailReq) Url() string                 { return "/v1/net_detail" }
func (r *NetDetailReq) Method() string              { return http.MethodGet }
func (r *NetDetailReq) Header() http.Header         { return nil }
func (r *NetDetailReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *NetDetailReq) Form() (form url.Values)     { return }
func (r *NetDetailReq) Body() any                   { return nil }

func NetDetail(ctx *cputil.Context, req *NetDetailReq) (*NetDetailResp, error) {
	if resp, err := cputil.Execute[*NetDetailReq, *NetDetailResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
