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
	"github.com/rwscode/cputil"
	"net/http"
	"net/url"
)

type (
	RealDataReq struct {
		ID []uint `json:"id"`
	}
	RealDataResp struct {
		Id               int     `json:"id"`                 // 实例ID
		Kvm              string  `json:"kvm"`                // 实例KVMID
		CpuUsage         string  `json:"cpu_usage"`          // CPU使用率
		MemoryUsage      string  `json:"memory_usage"`       // 内存使用百分比(v2.1.7改,-1未获取到v2.1.8改)
		MemoryTotal      string  `json:"memory_total"`       // 总内存(v2.1.7+)
		MemoryUsable     string  `json:"memory_usable"`      // 已用内存(v2.1.7+,v2.2.6+改为已用内存)
		CurrentInBw      string  `json:"current_in_bw"`      // 进带宽(v2.1.8+)
		CurrentOutBw     string  `json:"current_out_bw"`     // 出带宽(v2.1.8+)
		CurrentReadByte  int     `json:"current_read_byte"`  // 磁盘读取速度(MB/s,v2.4.7+)
		CurrentWriteByte float64 `json:"current_write_byte"` // 磁盘写入速度(MB/s,v2.4.7+)
	}
)

func (r *RealDataReq) Url() string                 { return "/v1/clouds/real_data" }
func (r *RealDataReq) Method() string              { return http.MethodGet }
func (r *RealDataReq) Header() http.Header         { return nil }
func (r *RealDataReq) Values() (values url.Values) { return }
func (r *RealDataReq) Form() (form url.Values)     { return }
func (r *RealDataReq) Body() any                   { return r }

// RealData 获取实例实时CPU/内存
func RealData(ctx *cloudplatform.Context, req *RealDataReq) (*RealDataResp, error) {
	if resp, err := cloudplatform.Execute[*RealDataReq, *RealDataResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
