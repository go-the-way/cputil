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
	"github.com/rwscode/cputil"
	"net/http"
	"net/url"

	q "github.com/google/go-querystring/query"
)

type (
	ListReq struct {
		SearchType   string `url:"searchtype,omitempty"`    // 搜索类型，0默认搜索主机名和IP,1区域,2节点
		Node         string `url:"node,omitempty"`          // 节点ID
		Status       string `url:"status,omitempty"`        // 状态(init=创建中,on=开机,off=关机,suspend=暂停,paused挂起,rescue=救援系统,unknown=未知)
		User         string `url:"user,omitempty"`          // 用户ID,后台有效
		ListType     string `url:"list_type,omitempty"`     // 获取类型(all,page),all会忽略页数直接返回所有
		ImageVersion string `url:"image_version,omitempty"` // 镜像版本id(v2.4.3+)
		Security     string `url:"security,omitempty"`      // 安全组ID(传0可获取没有绑定安全组实例,v2.4.3+)
		Vpc          string `url:"vpc,omitempty"`           // VPCID(v2.4.4+)
		HostType     string `url:"host_type,omitempty"`     // 实例类型(host=专业魔方云实例,lightHost=轻量版魔方云实例,hyperv=Hyper-V,v2.5.2+)
		Id           string `url:"id,omitempty"`            // 按ID搜索
		KvmId        string `url:"kvmid,omitempty"`         // 按KVMID搜索
		AdvancedCpu  string `url:"advanced_cpu,omitempty"`  // 搜索智能CPUID
		AdvancedBw   string `url:"advanced_bw,omitempty"`   // 搜索智能带宽ID

		Page    string `url:"page,omitempty"`
		PerPage string `url:"per_page,omitempty"`
		// 排序(id,cpu,memory,hostname,node,uid,mainip,os,in_bw,out_bw,delete_time,recycle_time,cpu_usage,memory_usage,username,node_name,current_in_bw,current_out_bw,current_read_byte,current_write_byte)
		Orderby string `url:"orderby,omitempty"`
		Sort    string `url:"sort,omitempty"`
	}
	ListResp struct {
		Data []struct {
			Id               int         `json:"id"`
			Cpu              int         `json:"cpu"`
			Memory           float64     `json:"memory"`
			Kvm              string      `json:"kvm"`
			Hostname         string      `json:"hostname"`
			Node             int         `json:"node"`
			Status           string      `json:"status"`
			NetworkType      string      `json:"network_type"`
			DeleteTime       string      `json:"delete_time"`
			RecycleTime      string      `json:"recycle_time"`
			Rescue           int         `json:"rescue"`
			CpuUsage         string      `json:"cpu_usage"`
			CurrentInBw      string      `json:"current_in_bw"`
			CurrentOutBw     string      `json:"current_out_bw"`
			MemoryTotal      string      `json:"memory_total"`
			MemoryUsable     string      `json:"memory_usable"`
			MemoryUsage      string      `json:"memory_usage"`
			Lock             int         `json:"lock"`
			SuspendType      string      `json:"suspend_type"`
			CurrentReadByte  float64     `json:"current_read_byte"`
			CurrentWriteByte float64     `json:"current_write_byte"`
			NodeName         string      `json:"node_name"`
			Type             string      `json:"type"`
			Mainip           string      `json:"mainip"`
			Os               string      `json:"os"`
			Parent           int         `json:"parent"`
			ParentName       interface{} `json:"parent_name"`
			Svg              int         `json:"svg"`
			InBw             int         `json:"in_bw"`
			OutBw            int         `json:"out_bw"`
			Ipaddress        string      `json:"ipaddress"`
			Remark           interface{} `json:"remark"`
			Rid              int         `json:"rid"`
			ResourcePackage  interface{} `json:"resource_package"`
			IpNum            int         `json:"ip_num"`
			User             struct {
				Id         int    `json:"id"`
				Email      string `json:"email"`
				Username   string `json:"username"`
				Phone      string `json:"phone"`
				Parent     int    `json:"parent"`
				ParentName string `json:"parent_name"`
				RCount     int    `json:"r_count"`
			} `json:"user"`
			Area struct {
				Id        int    `json:"id"`
				Name      string `json:"name"`
				ShortName string `json:"short_name"`
			} `json:"area"`
			Disk []struct {
				Id      int    `json:"id"`
				Type    string `json:"type"`
				Size    int    `json:"size"`
				Name    string `json:"name"`
				Remarks string `json:"remarks"`
				Kvm     int    `json:"kvm"`
			} `json:"disk"`
			IsNatHost        bool   `json:"is_nat_host"`
			OtherInterfaceIp string `json:"other_interface_ip"`
		} `json:"data"`
		Meta struct {
			Total     int `json:"total"`
			TotalPage int `json:"total_page"`
			Page      int `json:"page"`
			PerPage   int `json:"per_page"`
		} `json:"meta"`
	}
)

func (r *ListReq) Url() string                 { return "/v1/clouds" }
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
