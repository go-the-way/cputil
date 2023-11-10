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
		Search      string `url:"search,omitempty"`       // 搜索
		ListType    string `url:"list_type,omitempty"`    // 获取类型(all,page),all会忽略页数直接返回所有
		IncludeNode string `url:"include_node,omitempty"` // 是否获取区域下所有节点信息(1获取0不获取)

		Page    string `url:"page,omitempty"`
		PerPage string `url:"per_page,omitempty"`
		Orderby string `url:"orderby,omitempty"` // 排序(id,name,ip,node_num,store_num,cloud_num,ip_num,status,ip_used_num,ip_free_num)
		Sort    string `url:"sort,omitempty"`
	}
	ListResp struct {
		Data []struct {
			Id           int    `json:"id"`   // 区域ID
			Name         string `json:"name"` // 区域名
			ShortName    string `json:"short_name"`
			Ip           string `json:"ip"`
			Country      int    `json:"country"`
			Username     string `json:"username"`
			Password     string `json:"password"`
			Status       int    `json:"status"`
			VncIp        string `json:"vnc_ip"`
			VncPrivateIp string `json:"vnc_private_ip"`
			VncDomain    string `json:"vnc_domain"`
			VncSslPem    string `json:"vnc_ssl_pem"`
			VncSslKey    string `json:"vnc_ssl_key"`
			VncSwitch    int    `json:"vnc_switch"`
			AlternateIp  string `json:"alternate_ip"`
			NodeNum      int    `json:"node_num"`
			StoreNum     int    `json:"store_num"`
			CloudNum     int    `json:"cloud_num"`
			CountryName  string `json:"country_name"`
			CountryCode  string `json:"country_code"`
			Removable    bool   `json:"removable"`
			HostSumOn    int    `json:"host_sum_on"`
			HostSumOff   int    `json:"host_sum_off"`
			IpUsedNum    int    `json:"ip_used_num"`
			IpFreeNum    int    `json:"ip_free_num"`
			IpNum        int    `json:"ip_num"`
			HostSumOther int    `json:"host_sum_other"`
		} `json:"data"`
	}
)

func (r *ListReq) Url() string                 { return "/v1/areas" }
func (r *ListReq) Method() string              { return http.MethodGet }
func (r *ListReq) Header() http.Header         { return nil }
func (r *ListReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *ListReq) Form() (form url.Values)     { return }
func (r *ListReq) Body() any                   { return nil }

// List 区域列表
func List(ctx *cputil.Context, req *ListReq) (*ListResp, error) {
	if resp, err := cputil.Execute[*ListReq, *ListResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
