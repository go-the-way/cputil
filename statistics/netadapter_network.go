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
	"fmt"
	"github.com/go-the-way/cputil"
	"github.com/go-the-way/cputil/clouds"
)

type (
	// NetAdapterNetworkReq 网卡带宽
	NetAdapterNetworkReq struct {
		Id            uint  // cloud id
		NetworkOffset uint8 // 网卡顺序 从0开始
		StartTime     int64 // 开始毫秒时间戳(1699344000000) UTC+8
		EndTime       int64 // 结束毫秒时间戳(1699344060000) UTC+8
	}
	NetAdapterNetworkResp struct {
		List []NetAdapterRespList `json:"list"` // 流量列表
		Stat NetDetailResp        `json:"stat"` // 汇总数据
	}
)

func NetAdapterNetwork(ctx *cputil.Context, req *NetAdapterNetworkReq) (*NetAdapterNetworkResp, error) {
	detail, err := clouds.Detail(ctx, &clouds.DetailReq{ID: req.Id})
	if err != nil {
		return nil, err
	}
	netAdapterResp, err := NetAdapter(ctx, &NetAdapterReq{Id: req.Id, Kvm: detail.Kvmid, NetworkOffset: req.NetworkOffset, StartTime: req.StartTime, EndTime: req.EndTime})
	if err != nil {
		return nil, err
	}
	netDetailResp, err := NetDetail(ctx, &NetDetailReq{NodeId: uint(detail.NodeId), KvmIfName: fmt.Sprintf("%s.%d", detail.Kvmid, req.NetworkOffset), StartTime: req.StartTime, EndTime: req.EndTime})
	if err != nil {
		return nil, err
	}
	return &NetAdapterNetworkResp{netAdapterResp.List, *netDetailResp}, nil
}
