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

package snapshots

import (
	"errors"
	"fmt"
	q "github.com/google/go-querystring/query"
	"github.com/rwscode/cputil"
	"net/http"
	"net/url"
)

type (
	ListReq struct {
		ID     uint   `url:"-"`                // 实例ID
		Search string `url:"search,omitempty"` // 搜索
		Type   string `url:"type,omitempty"`   // 按备份类型搜索(snap快照,backup备份)

		Page    string `url:"page,omitempty"`
		PerPage string `url:"per_page,omitempty"`
		Orderby string `url:"orderby,omitempty"` // 排序(id,size,name,type)
		Sort    string `url:"sort,omitempty"`
	}
	ListResp struct {
		Data []struct {
			Id          int    `json:"id"`           // 备份ID
			Size        int    `json:"size"`         // 大小(G)
			Name        string `json:"name"`         // 文件名
			Type        string `json:"type"`         // 备份类型(snap快照,backup备份)
			CreateTime  string `json:"create_time"`  // 创建时间
			Remarks     string `json:"remarks"`      // 备注名
			DiskName    string `json:"disk_name"`    // 关联磁盘文件
			DiskRemarks string `json:"disk_remarks"` // 关联磁盘备注
			RealFile    string `json:"real_file"`    // 当前备份完整路径(v2.4.3+)
			Status      int    `json:"status"`       // 状态(0=创建中,1=创建完成,v2.4.7+)
			Next        []struct {
				Id   int    `json:"id"`   // 下级快照ID
				Name string `json:"name"` // 下级快照文件名
			}
		} `json:"data"`
		Meta struct {
			Total     int `json:"total"`
			TotalPage int `json:"total_page"`
			Page      int `json:"page"`
			PerPage   int `json:"per_page"`
		} `json:"meta"`
	}
)

func (r *ListReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/snapshots", r.ID) }
func (r *ListReq) Method() string              { return http.MethodGet }
func (r *ListReq) Header() http.Header         { return nil }
func (r *ListReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *ListReq) Form() (form url.Values)     { return }
func (r *ListReq) Body() any                   { return nil }

// List 实例备份/快照列表
func List(ctx *cloudplatform.Context, req *ListReq) (*ListResp, error) {
	if resp, err := cloudplatform.Execute[*ListReq, *ListResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
