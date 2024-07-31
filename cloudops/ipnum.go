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

package cloudops

import (
	"errors"
	"fmt"
	"github.com/go-the-way/cputil"
	"net/http"
	"net/url"
)

type (
	IpNumReq struct {
		ID           uint `json:"-"`              // 实例ID
		Num          uint `json:"num"`            // 目标IP数量
		IpGroup      uint `json:"ip_group"`       // IP分组ID(增加IP时生效分配指定分组的IP,v2.3.0+)
		IsForceGroup uint `json:"is_force_group"` // 是否强制使用IP分组,不强制指定分组不够会使用其他的组(0不强制,1强制,v2.3.0+)
	}
	IpNumResp struct{}
)

func (r *IpNumReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/ip", r.ID) }
func (r *IpNumReq) Method() string              { return http.MethodPut }
func (r *IpNumReq) Header() http.Header         { return nil }
func (r *IpNumReq) Values() (values url.Values) { return }
func (r *IpNumReq) Form() (form url.Values)     { return }
func (r *IpNumReq) Body() any                   { return r }

// IpNum 修改实例IP数量
// 修改实例IP数量,第一个网卡
func IpNum(ctx *cputil.Context, req *IpNumReq) (*IpNumResp, error) {
	if resp, err := cputil.Execute[*IpNumReq, *IpNumResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
