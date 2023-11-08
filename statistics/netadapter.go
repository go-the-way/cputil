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
	"fmt"
	"github.com/rwscode/cputil"
	"github.com/rwscode/cputil/clouds"
)

type (
	// NetAdapterReq 网卡
	NetAdapterReq struct {
		Id            uint   // cloud id
		Kvm           string // kvm id
		NetworkOffset uint8  // 网卡顺序 从0开始
		Usage         bool   // 是否查询用量（流量），默认 查询网卡/带宽
		StartTime     int64  // 开始毫秒时间戳(1699344000000) UTC+8
		EndTime       int64  // 结束毫秒时间戳(1699344060000) UTC+8
	}
	NetAdapterRespList struct {
		UtcTime string  `json:"utc_time"` // UTC time(2023-11-07T00:48:00Z)
		In      float64 `json:"in"`       // 进带宽(bps)
		Out     float64 `json:"out"`      // 出带宽(bps)
	}
	NetAdapterResp struct {
		List []NetAdapterRespList `json:"list"`
	}
)

/*
response
---
[
	[
		"2023-11-08T09:39:00Z", // utc time
		7, // 进带宽(bps)
		0 // 出带宽(bps)
	]
]
*/

func NetAdapter(ctx *cputil.Context, req *NetAdapterReq) (*NetAdapterResp, error) {
	kvmId := req.Kvm
	if kvmId == "" {
		kid, err := clouds.KvmId(ctx, req.Id)
		if err != nil {
			return nil, err
		}
		kvmId = kid
	}
	if kvmId == "" {
		return nil, errors.New("kvm id为空，查询失败")
	}
	r := &Req{Kvm: kvmId, Type: "net_adapter", KvmIfName: fmt.Sprintf("%s.%d", kvmId, req.NetworkOffset), StartTime: req.StartTime, EndTime: req.EndTime}
	if req.Usage {
		r.Status = "1"
	}
	arr, err := Statistics(ctx, r)
	if err != nil {
		return nil, err
	}
	list := make([]NetAdapterRespList, 0)
	if arr != nil && len(arr) > 0 {
		for _, aaa := range arr {
			// "2023-11-08T09:39:00Z", // utc time
			// 	7, // 进带宽(bps)
			// 	0 // 出带宽(bps)
			if aaa != nil && len(aaa) >= 3 {
				var resp NetAdapterRespList
				if utcTime := aaa[0]; utcTime != nil {
					resp.UtcTime = utcTime.(string)
				}
				if in := aaa[1]; in != nil {
					resp.In = in.(float64)
				}
				if out := aaa[2]; out != nil {
					resp.Out = out.(float64)
				}
				list = append(list, resp)
			}
		}
	}
	return &NetAdapterResp{list}, nil
}
