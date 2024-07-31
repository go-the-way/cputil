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
	UpdateReq struct {
		ID     uint   `json:"id"`     // 磁盘ID
		Size   uint   `json:"size"`   // 磁盘大小（G）
		Driver string `json:"driver"` // 驱动(ide,virtio,sata,scsi)
		Cache  string `json:"cache"`  // 缓存(none,writethrough,writeback,directsync)
		Io     string `json:"io"`     // IO(native,threads)

		// ReadBytesSec     int `json:"read_bytes_sec"`
		// WriteBytesSec    int `json:"write_bytes_sec"`
		// ReadIopsSec      int `json:"read_iops_sec"`
		// WriteIopsSec     int `json:"write_iops_sec"`
		// ReadBytesSecMax  int `json:"read_bytes_sec_max"`
		// WriteBytesSecMax int `json:"write_bytes_sec_max"`
		// ReadIopsSecMax   int `json:"read_iops_sec_max"`
		// WriteIopsSecMax  int `json:"write_iops_sec_max"`
		// IopsMax          int `json:"iops_max"`
		// IopsMin          int `json:"iops_min"`

		// size	int	-		磁盘大小
		// driver	string	-		驱动(ide,virtio,sata,scsi)
		// cache	string	-		缓存(none,writethrough,writeback,directsync)
		// io	string	-		IO(native,threads)
		// read_bytes_sec	int	-		读取限制(MB/s)
		// write_bytes_sec	int	-		写入限制(MB/s)
		// read_iops_sec	int	-		读取限制(ops/s)
		// write_iops_sec	int	-		写入限制(ops/s)
		// read_bytes_sec_max	int	-		读取最大突发(MB)
		// write_bytes_sec_max	int	-		写入最大突发(MB)
		// read_iops_sec_max	int	-		读取最大突发(ops)
		// write_iops_sec_max	int	-		写入最大突发(ops)
		// iops_min	int	-		最小iops(0-100,Hyper-V,v2.5.2+)
		// iops_max	int	-		最大iops(0-100,Hyper-V,v2.5.2+)

	}
	UpdateResp struct{}
)

func (r *UpdateReq) Url() string                 { return fmt.Sprintf("/v1/disks/%d", r.ID) }
func (r *UpdateReq) Method() string              { return http.MethodPut }
func (r *UpdateReq) Header() http.Header         { return nil }
func (r *UpdateReq) Values() (values url.Values) { return }
func (r *UpdateReq) Form() (form url.Values)     { return }
func (r *UpdateReq) Body() any                   { return r }

func Update(ctx *cputil.Context, req *UpdateReq) (*UpdateResp, error) {
	if resp, err := cputil.Execute[*UpdateReq, *UpdateResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
