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

package disks

import (
	"errors"
	"fmt"
	"github.com/go-the-way/cputil"
	"net/http"
	"net/url"
)

// {
//    "id": 609,
//    "size": 500,
//    "driver": "virtio",
//    "cache": "writethrough",
//    "io": "native",
//    "read_bytes_sec": 0,
//    "write_bytes_sec": 0,
//    "read_iops_sec": 0,
//    "write_iops_sec": 0,
//    "read_bytes_sec_max": 0,
//    "write_bytes_sec_max": 0,
//    "read_iops_sec_max": 0,
//    "write_iops_sec_max": 0,
//    "iops_max": 0,
//    "iops_min": 0
// }

type (
	CreateReq struct {
		ID     uint   `json:"id"`     // 实例ID
		Size   uint   `json:"size"`   // 磁盘大小（G）
		Driver string `json:"driver"` // 驱动(ide,virtio,sata,scsi)
		Cache  string `json:"cache"`  // 缓存(none,writethrough,writeback,directsync)
		Io     string `json:"io"`     // IO(native,threads)
		Store  uint   `json:"store"`  // 存储ID
	}
	CreateResp struct{}
)

func (r *CreateReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/disks", r.ID) }
func (r *CreateReq) Method() string              { return http.MethodPost }
func (r *CreateReq) Header() http.Header         { return nil }
func (r *CreateReq) Values() (values url.Values) { return }
func (r *CreateReq) Form() (form url.Values)     { return }
func (r *CreateReq) Body() any                   { return r }

func Create(ctx *cputil.Context, req *CreateReq) (*CreateResp, error) {
	if resp, err := cputil.Execute[*CreateReq, *CreateResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
