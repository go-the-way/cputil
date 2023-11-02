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
	"github.com/rwscode/cputil"
	"net/http"
	"net/url"
)

type (
	ReinstallReq struct {
		ID             uint   `json:"-"`                   // 实例ID
		Os             uint   `json:"os"`                  // 镜像ID/模板ID必须有一个,都有为镜像ID镜像ID
		Template       uint   `json:"template"`            // 镜像ID/模板ID必须有一个,都有为镜像ID模板ID
		Password       string `json:"password"`            // 自动生成新密码
		Port           uint   `json:"port"`                // 端口(可选,专业版支持)
		SystemDiskSize uint   `json:"system_disk_size"`    // 系统盘大小(可以传入数组如[20,30],第一个为windows,第二个为其他,v2.3.9+,v2.4.1改为不传不修改磁盘大小)
		FormatDataDisk uint   `json:"format_data_disk"`    // 是否格式化数据盘(0=不格式,1=格式化)
		Key            string `json:"key"`                 // 镜像设置的密钥密钥(重装为Windows可选)
		PasswordType   uint   `json:"image_client_hidden"` // query参数是否验证镜像是否客户端隐藏(0=客户端可见,1=客户端隐藏)
	}
	ReinstallResp struct {
		TaskID         string `json:"taskid"`
		User           string `json:"user"`             // 操作系统用户
		Password       string `json:"password"`         // 操作系统密码
		Port           uint   `json:"port"`             // 重装随机端口(v2.2.2+)
		SystemDiskSize uint   `json:"system_disk_size"` // 重装后系统盘大小(0表示未变更,v2.3.9+)
	}
)

func (r *ReinstallReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/reinstall", r.ID) }
func (r *ReinstallReq) Method() string              { return http.MethodPut }
func (r *ReinstallReq) Header() http.Header         { return nil }
func (r *ReinstallReq) Values() (values url.Values) { return }
func (r *ReinstallReq) Form() (form url.Values)     { return }
func (r *ReinstallReq) Body() any                   { return r }

// Reinstall 重装
func Reinstall(ctx *cloudplatform.Context, req *ReinstallReq) (*ReinstallResp, error) {
	if resp, err := cloudplatform.Execute[*ReinstallReq, *ReinstallResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
