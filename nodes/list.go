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

package nodes

import (
	"errors"
	"github.com/rwscode/cputil"
	"net/http"
	"net/url"

	q "github.com/google/go-querystring/query"
)

type (
	ListReq struct {
		Search    string   `url:"search,omitempty"`     // 搜索
		Area      string   `url:"area,omitempty"`       // 搜索区域id
		NodeGroup string   `url:"node_group,omitempty"` // 搜索节点分组ID
		Enable    string   `url:"enable,omitempty"`     // 搜索是否启用
		ListType  string   `url:"list_type,omitempty"`  // 获取类型(all,page),all会忽略页数直接返回所有
		Rid       string   `url:"rid,omitempty"`        // 资源包ID(前台用户)
		Type      []string `url:"type,omitempty"`       // 节点类型(v2.5.2+)

		Page    string `url:"page,omitempty"`
		PerPage string `url:"per_page,omitempty"`
		Orderby string `url:"orderby,omitempty"` // 排序(id,name,ip,area_id,group_id,cloud_num,group_name)
		Sort    string `url:"sort,omitempty"`
	}
	ListResp struct {
		Data []struct {
			Id              int         `json:"id"`
			Name            string      `json:"name"`
			Ip              string      `json:"ip"`
			AreaId          int         `json:"area_id"`
			Status          int         `json:"status"`
			Port            int         `json:"port"`
			SshPass         string      `json:"ssh_pass"`
			Username        string      `json:"username"`
			Password        string      `json:"password"`
			Dns1            string      `json:"dns1"`
			Dns2            string      `json:"dns2"`
			NatStatus       int         `json:"nat_status"`
			UpdateFailedMsg string      `json:"update_failed_msg"`
			UpgradeStatus   int         `json:"upgrade_status"`
			UpgradeTime     string      `json:"upgrade_time"`
			Enable          int         `json:"enable"`
			AutoBootHost    int         `json:"auto_boot_host"`
			LastLiveTime    string      `json:"last_live_time"`
			GroupId         int         `json:"group_id"`
			SingleIpNat     int         `json:"single_ip_nat"`
			ChooseStoreRule int         `json:"choose_store_rule"`
			Evacuate        int         `json:"evacuate"`
			NatPortRange    string      `json:"nat_port_range"`
			MaxMemory       int         `json:"max_memory"`
			Type            string      `json:"type"`
			Trunk           int         `json:"trunk"`
			Remark          string      `json:"remark"`
			AreaName        string      `json:"area_name"`
			CloudNum        int         `json:"cloud_num"`
			GroupName       interface{} `json:"group_name"`
			Removable       bool        `json:"removable"`
			Alive           bool        `json:"alive"`
		} `json:"data"`
		Meta struct {
			Total     int `json:"total"`
			TotalPage int `json:"total_page"`
			Page      int `json:"page"`
			PerPage   int `json:"per_page"`
		} `json:"meta"`
	}
)

func (r *ListReq) Url() string                 { return "/v1/nodes" }
func (r *ListReq) Method() string              { return http.MethodGet }
func (r *ListReq) Header() http.Header         { return nil }
func (r *ListReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *ListReq) Form() (form url.Values)     { return }
func (r *ListReq) Body() any                   { return nil }

// List 实例列表
func List(ctx *cloudplatform.Context, req *ListReq) (*ListResp, error) {
	if resp, err := cloudplatform.Execute[*ListReq, *ListResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
