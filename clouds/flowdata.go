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

package clouds

import (
	"errors"
	"fmt"
	"github.com/rwscode/cputil"
	"net/http"
	"net/url"

	q "github.com/google/go-querystring/query"
)

type (
	FlowDataReq struct {
		ID uint // 实例ID

		Type      string `url:"type,omitempty"`       // 1按小时,2按天,3按月,其他默认根据时间范围自动判断
		StartTime string `url:"start_time,omitempty"` // 开始毫秒时间戳
		EndTime   string `url:"end_time,omitempty"`   // 结束毫秒时间戳(默认现在)
		Unit      string `url:"unit,omitempty"`       // 转换到单位(KB,MB,GB,TB)
	}
	FlowDataResp struct {
		Data []struct {
			Time string  `json:"time"` // 时间(根据类型时间返回不一样)
			In   float64 `json:"in"`   // 进流量
			Out  float64 `json:"out"`  // 出流量
		} `json:"data"`
		Unit        string `json:"unit"`         // 当前单位
		Type        int    `json:"type"`         // 当前类型
		TrafficType int    `json:"traffic_type"` // 统计方向(1进2出3合计)
	}
)

func (r *FlowDataReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/flow_data", r.ID) }
func (r *FlowDataReq) Method() string              { return http.MethodGet }
func (r *FlowDataReq) Header() http.Header         { return nil }
func (r *FlowDataReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *FlowDataReq) Form() (form url.Values)     { return }
func (r *FlowDataReq) Body() any                   { return nil }

// FlowData 实例流量数据
func FlowData(ctx *cputil.Context, req *FlowDataReq) (*FlowDataResp, error) {
	if resp, err := cputil.Execute[*FlowDataReq, *FlowDataResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
