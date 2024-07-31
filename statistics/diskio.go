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
	"github.com/go-the-way/cputil"
	"github.com/go-the-way/cputil/clouds"
	"github.com/go-the-way/cputil/pkg/timefmt"
)

type (
	// DiskIoReq 硬盘IO
	DiskIoReq struct {
		Id        uint   // cloud id
		Dev       string // vda vdd
		StartTime int64  // 开始毫秒时间戳(1699344000000) UTC+8
		EndTime   int64  // 结束毫秒时间戳(1699344060000) UTC+8
	}
	DiskIoRespList struct {
		Time      string  `json:"time"`       // 时间
		Read      float64 `json:"read"`       // 读取速度(Byte)
		Write     float64 `json:"write"`      // 写入速度(Byte)
		ReadIops  float64 `json:"read_iops"`  // 读取IOPS
		WriteIops float64 `json:"write_iops"` // 写入IOPS
	}
	DiskIoResp struct {
		List []DiskIoRespList `json:"list"`
	}
)

/*
response
---
[
	[
		"2023-11-07T08:00:00Z", // utc time
		0, // 读取速度(KB/s)
		978.2346368715064, // 写入速度(KB/s)
		0, // 读取 IOPS
		0.14525139664804 //读取 IOPS
	]
]
*/

func DiskIo(ctx *cputil.Context, req *DiskIoReq) (*DiskIoResp, error) {
	kvmId, err := clouds.KvmId(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if kvmId == "" {
		return nil, errors.New("kvm id为空，查询失败")
	}
	arr, err := Statistics(ctx, &Req{Kvm: kvmId, Type: "disk_io", Dev: req.Dev, StartTime: req.StartTime, EndTime: req.EndTime})
	if err != nil {
		return nil, err
	}
	list := make([]DiskIoRespList, 0)
	if arr != nil && len(arr) > 0 {
		for _, aaa := range arr {
			// "2023-11-07T08:00:00Z", // utc time
			// 	0, // 读取速度(KB/s)
			// 	978.2346368715064, // 写入速度(KB/s)
			// 	0, // 读取 IOPS
			// 	0.14525139664804 //读取 IOPS
			if aaa != nil && len(aaa) >= 5 {
				var resp DiskIoRespList
				if utcTime := aaa[0]; utcTime != nil {
					resp.Time = timefmt.Utc2Gmt8(utcTime.(string))
				}
				if read := aaa[1]; read != nil {
					resp.Read = read.(float64)
				}
				if write := aaa[2]; write != nil {
					resp.Write = write.(float64)
				}
				if readIops := aaa[3]; readIops != nil {
					resp.ReadIops = readIops.(float64)
				}
				if writeIops := aaa[4]; writeIops != nil {
					resp.WriteIops = writeIops.(float64)
				}
				list = append(list, resp)
			}
		}
	}
	return &DiskIoResp{list}, nil
}
