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
	"github.com/rwscode/cputil"
	"github.com/rwscode/cputil/clouds"
)

type (
	// KvmInfoReq CPU+内存使用量
	KvmInfoReq struct {
		Id        uint  // cloud id
		StartTime int64 // 开始毫秒时间戳(1699344000000) UTC+8
		EndTime   int64 // 结束毫秒时间戳(1699344060000) UTC+8
	}
	KvmInfoRespList struct {
		UtcTime  string  `json:"utc_time"`  // UTC time(2023-11-07T00:48:00Z)
		CPU      float64 `json:"cpu"`       // CPU使用率(n.xy)
		MemTotal int64   `json:"mem_total"` // 总内存(Bytes)
		MemUsed  int64   `json:"mem_used"`  // 已使用内存(Bytes)
	}
	KvmInfoResp struct {
		List []KvmInfoRespList `json:"list"`
	}
)

// info CPU+Mem实时数据
/*
response
---
[
	[
		"2023-11-07T00:48:00Z", // utc time
 		7.790737, // CPU usage (%)
		1900793856, // total mem in bytes
		404992000  // used mem in bytes
	]
]
*/

func KvmInfo(ctx *cputil.Context, req *KvmInfoReq) (*KvmInfoResp, error) {
	kvmId, err := clouds.KvmId(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if kvmId == "" {
		return nil, errors.New("kvm id为空，查询失败")
	}
	arr, err := Statistics(ctx, &Req{Kvm: kvmId, Type: "kvm_info", StartTime: req.StartTime, EndTime: req.EndTime})
	if err != nil {
		return nil, err
	}
	list := make([]KvmInfoRespList, 0)
	if arr != nil && len(arr) > 0 {
		for _, aaa := range arr {
			// "2023-11-07T00:48:00Z", // utc time
			// 7.790737, // CPU usage (%)
			// 1900793856, // total mem in bytes
			// 404992000  // used mem in bytes
			var resp KvmInfoRespList
			if aaa != nil && len(aaa) >= 4 {
				if utcTime := aaa[0]; utcTime != nil {
					resp.UtcTime = utcTime.(string)
				}
				if cpu := aaa[1]; cpu != nil {
					resp.CPU = cpu.(float64)
				}
				if memTotal := aaa[2]; memTotal != nil {
					resp.MemTotal = memTotal.(int64)
				}
				if memUsed := aaa[3]; memUsed != nil {
					resp.MemUsed = memUsed.(int64)
				}
				list = append(list, resp)
			}
		}
	}
	return &KvmInfoResp{list}, nil
}
