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
	"github.com/go-the-way/cputil"
	"net/http"
	"net/url"
)

type (
	UpdateReq struct {
		ID             uint   `json:"-"`                          // 实例ID
		CPU            uint   `json:"cpu,omitempty"`              // CPU
		Memory         uint   `json:"memory,omitempty"`           // 内存(MB)
		Vncport        uint   `json:"vncport,omitempty"`          // vnc端口
		Vncpass        uint   `json:"vncpass,omitempty"`          // vnc密码
		Vnc            uint   `json:"vnc,omitempty"`              // 是否启用vnc(0禁用,1启用)
		Bootorder      string `json:"bootorder,omitempty"`        // bootorder(dc,cd,d,c)
		Cpumodel       uint   `json:"cpumodel,omitempty"`         // cpumodel(1host-passthrough,2host-model,3custom)
		Pae            uint   `json:"pae,omitempty"`              // pae是否启用(0禁用,1启用)
		Clock          uint   `json:"clock,omitempty"`            // 时钟(0Default,1UTC,2Localtime)
		Acpi           uint   `json:"acpi,omitempty"`             // acpi是否启用(0禁用,1启用)
		TrafficQuota   uint   `json:"traffic_quota,omitempty"`    // 限额流量
		TmpTraffic     uint   `json:"tmp_traffic,omitempty"`      // 临时流量
		TrafficType    uint   `json:"traffic_type,omitempty"`     // 统计方向1=进2=出3=总计
		SnapNum        int    `json:"snap_num,omitempty"`         // 快照数量限制(0不限-1不允许)
		BackupNum      int    `json:"backup_num,omitempty"`       // 备份数量限制(0不限-1不允许)
		NatAclLimit    int    `json:"nat_acl_limit,omitempty"`    // nat转发限制(0不限-1不允许)
		NatWebLimit    int    `json:"nat_web_limit,omitempty"`    // nat建站限制(0不限-1不允许)
		CpuTopCheck    uint   `json:"cputopcheck,omitempty"`      // 是否启用CPUTopology(0关闭1开启)
		CpuTuneCheck   uint   `json:"cputunecheck,omitempty"`     // 是否启用CPUTune(0关闭1开启)
		CpuNumaCheck   uint   `json:"cpunumacheck,omitempty"`     // 是否启用NUMATopology(0关闭1开启)
		CpuSockets     uint   `json:"cpusockets,omitempty"`       // cpusockets
		CpuCores       uint   `json:"cpucores,omitempty"`         // cpucores
		CpuThreads     uint   `json:"cputhreads,omitempty"`       // cputhreads
		CpuTune        string `json:"cputune,omitempty"`          // cputune
		CpuNumaTop     string `json:"cpunumatop,omitempty"`       // cpunumatop
		Console        uint   `json:"console,omitempty"`          // console(0关闭,1开启,v2.2.7+)
		ResetFlowDay   uint   `json:"reset_flow_day,omitempty"`   // 流量重置日(只能是1-31,v2.3.9+)
		Queues         uint   `json:"queues,omitempty"`           // 网卡多队列(v2.5.2+)
		DynamicMemory  uint   `json:"dynamic_memory,omitempty"`   // 动态内存(0=关闭,1开启,v2.5.2+)
		Gpu            uint   `json:"gpu,omitempty"`              // 虚拟显存开关(0=关闭,1=开启)
		VRam           uint   `json:"vram,omitempty"`             // 虚拟显存大小(只有gpu=1时生效,单位MB,只能是16,32,64,128,256,512,1024,2048)
		AdvancedCpu    uint   `json:"advanced_cpu,omitempty"`     // 智能CPU规则
		AdvancedBw     uint   `json:"advanced_bw,omitempty"`      // 智能带宽规则
		SkipColdBackup uint   `json:"skip_cold_backup,omitempty"` // 是否跳过备份(0=不跳过,1=跳过冷备份)
	}
	UpdateResp struct{}
)

func (r *UpdateReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d", r.ID) }
func (r *UpdateReq) Method() string              { return http.MethodPut }
func (r *UpdateReq) Header() http.Header         { return nil }
func (r *UpdateReq) Values() (values url.Values) { return }
func (r *UpdateReq) Form() (form url.Values)     { return }
func (r *UpdateReq) Body() any                   { return r }

// Update 修改实例
func Update(ctx *cputil.Context, req *UpdateReq) (*UpdateResp, error) {
	if resp, err := cputil.Execute[*UpdateReq, *UpdateResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
