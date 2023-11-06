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

package image

import (
	"errors"
	q "github.com/google/go-querystring/query"
	"github.com/rwscode/cputil"
	"net/http"
	"net/url"
)

type (
	ListReq struct {
		ImageGroupId string `url:"image_group_id,omitempty"` // 镜像组id
		Rescue       string `url:"rescue,omitempty"`         // 是否救援系统(0不包含救援系统,1救援系统)

		Page    string `url:"page,omitempty"`
		PerPage string `url:"per_page,omitempty"`
		Orderby string `url:"orderby,omitempty"` // 排序(id,name,use,image_group_id,status)
		Sort    string `url:"sort,omitempty"`
	}
	ListResp struct {
		Data []struct {
			Id                int           `json:"id"`
			ImageGroupId      int           `json:"image_group_id"`
			Name              string        `json:"name"`
			Filename          string        `json:"filename"`
			StoreId           []interface{} `json:"store_id"`
			Status            int           `json:"status"`
			Type              int           `json:"type"`
			BootType          string        `json:"boot_type"`
			Url               string        `json:"url"`
			UploadType        int           `json:"upload_type"`
			OnScript          string        `json:"on_script"`
			DiskDriver        string        `json:"disk_driver"`
			CpuModel          int           `json:"cpu_model"`
			Console           int           `json:"console"`
			LastUpdateTime    string        `json:"last_update_time"`
			Size              string        `json:"size"`
			Key               string        `json:"key"`
			ClientHidden      int           `json:"client_hidden"`
			ExternalExpansion int           `json:"external_expansion"`
			Use               int           `json:"use"`
			Info              []struct {
				NodeId     int    `json:"node_id"`
				StoreId    int    `json:"store_id"`
				StoreName  string `json:"store_name"`
				Status     int    `json:"status"`
				HistoryNum int    `json:"history_num"`
				CanImport  int    `json:"can_import"`
			} `json:"info"`
			RealPath []string `json:"real_path"`
			Group    struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
				Svg  int    `json:"svg"`
			} `json:"group"`
		} `json:"data"`
		Meta struct {
			Total     int `json:"total"`
			TotalPage int `json:"total_page"`
			PerPage   int `json:"per_page"`
			Page      int `json:"page"`
		} `json:"meta"`
		Area []struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
			Node []struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"node"`
		} `json:"area"`
	}
)

func (r *ListReq) Url() string                 { return "/v1/image" }
func (r *ListReq) Method() string              { return http.MethodGet }
func (r *ListReq) Header() http.Header         { return nil }
func (r *ListReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *ListReq) Form() (form url.Values)     { return }
func (r *ListReq) Body() any                   { return nil }

// List 实例列表
func List(ctx *cputil.Context, req *ListReq) (*ListResp, error) {
	if resp, err := cputil.Execute[*ListReq, *ListResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
