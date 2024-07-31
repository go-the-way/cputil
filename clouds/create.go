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
	"github.com/go-the-way/cputil"
	"net/http"
	"net/url"
)

type (
	CreateReq struct {
		CPU               uint   `json:"cpu"`                 // cpu
		Memory            uint   `json:"memory"`              // 内存(MB)
		OS                uint   `json:"os"`                  // 操作系统ID
		InBw              uint   `json:"in_bw"`               // 进带宽(可选)
		OutBw             uint   `json:"out_bw"`              // 出带宽(可选)
		IpNum             uint   `json:"ip_num"`              // IP数量 和IP地址必须传其中一个
		Ipaddress         []uint `json:"ipaddress"`           // IP地址ID 和IP数量必须传其中一个,节点通过IP地址分配,创建数量强制为1
		Node              uint   `json:"node"`                // 节点ID
		Area              uint   `json:"area"`                // 区域ID
		Store             uint   `json:"store"`               // 存储ID
		Client            uint   `json:"client"`              // 用户ID
		SystemDiskSize    uint   `json:"system_disk_size"`    // 系统盘大小
		DataDiskSize      uint   `json:"data_disk_size"`      // 数据盘大小(可选)
		Rootpass          string `json:"rootpass"`            // 操作系统密码
		Hostname          string `json:"hostname"`            // 主机名(可选)
		NetworkType       string `json:"network_type"`        // 网络类型(normal经典网络,vpc专有网络,轻量/Hyper-V/拨号只支持normal)
		Vpc               uint   `json:"vpc"`                 // vpc网络ID(当是vpc网络才有效)
		VpcName           string `json:"vpc_name"`            // vpc网络名称(当是vpc网络没选ID才有效)
		VpcIps            string `json:"vpc_ips"`             // vpc网络IP段(只能是192.168.0.0/16、172.16.0.0/12、10.0.0.0/8他们的子网,掩码在/16-/28,当是vpc网络没选ID才有效,v2.3.9+)
		Num               uint   `json:"num"`                 // 实例数量
		ClientName        string `json:"client_name"`         // 用户名(可选)
		TrafficQuota      uint   `json:"traffic_quota"`       // 流量限额
		Port              uint   `json:"port"`                // 端口(22,100-65535,auto=随机)
		SecurityGroup     uint   `json:"security_group"`      // 安全组ID(转发机没用,v2.3.0+)
		SecurityGroupName string `json:"security_group_name"` // 新增并关联安全组的名称(如果有安全组ID优先使用安全组ID,转发机没用,批量创建且端口随机时不生效,v2.3.0+)
		Type              string `json:"type"`                // 节点类型(host=KVM加强版,lightHost=KVM轻量,hyperv=Hyper-V,adsl=拨号云)
		PasswordType      uint   `json:"password_type"`       // 密码类型(0=密码,1=SSH密钥,windows镜像和Hyper-V节点不支持该参数)

		SystemReadBytesSec     int `json:"system_read_bytes_sec"`      // 系统盘读取限制(MB/s)
		SystemWriteBytesSec    int `json:"system_write_bytes_sec"`     // 系统盘写入限制(MB/s)
		SystemReadIopsSec      int `json:"system_read_iops_sec"`       // 系统盘读取限制(ops/s)
		SystemWriteIopsSec     int `json:"system_write_iops_sec"`      // 系统盘写入限制(ops/s)
		SystemReadBytesSecMax  int `json:"system_read_bytes_sec_max"`  // 系统盘读取最大突发(MB)
		SystemWriteBytesSecMax int `json:"system_write_bytes_sec_max"` // 系统盘写入最大突发(MB)
		SystemReadIopsSecMax   int `json:"system_read_iops_sec_max"`   // 系统盘读取最大突发(ops)
		SystemWriteIopsSecMax  int `json:"system_write_iops_sec_max"`  // 系统盘写入最大突发(ops)

		DataReadBytesSec     int `json:"data_read_bytes_sec"`      // 系统盘读取限制(MB/s)
		DataWriteBytesSec    int `json:"data_write_bytes_sec"`     // 系统盘写入限制(MB/s)
		DataReadIopsSec      int `json:"data_read_iops_sec"`       // 系统盘读取限制(ops/s)
		DataWriteIopsSec     int `json:"data_write_iops_sec"`      // 系统盘写入限制(ops/s)
		DataReadBytesSecMax  int `json:"data_read_bytes_sec_max"`  // 系统盘读取最大突发(MB)
		DataWriteBytesSecMax int `json:"data_write_bytes_sec_max"` // 系统盘写入最大突发(MB)
		DataReadIopsSecMax   int `json:"data_read_iops_sec_max"`   // 系统盘读取最大突发(ops)
		DataWriteIopsSecMax  int `json:"data_write_iops_sec_max"`  // 系统盘写入最大突发(ops)

		BackupNum    int `json:"backup_num"`    // 备份数量(-1=不支持备份,0=不限)
		SnapNum      int `json:"snap_num"`      // 快照数量(-1=不支持备份,0=不限)
		IsNatHost    int `json:"is_nat_host"`   // 是否是节点转发机(0=不是,1=是)
		NatAclLimit  int `json:"nat_acl_limit"` // NAT转发数量(-1=不支持,0=不限)
		NatWebLimit  int `json:"nat_web_limit"` // NAT建站数量(-1=不支持,0=不限)
		Template     int `json:"template"`      // 和操作系统2选1(都传优先使用操作系统)	模板ID
		NodePriority int `json:"node_priority"` // 节点选择优先级(1=实例数量平均,2=负载最低,3=内存最低)

		OtherDataDisk []DataDisk `json:"other_data_disk"` //	其他数据盘

		LinkClone     int    `json:"link_clone"`               // 本地系统盘创建方式(0=完整克隆,1=链接克隆,模板创建不支持该参数,v2.4.4+)
		BindMac       int    `json:"bind_mac"`                 // 是否启用IP-MAC绑定(0关闭,1开启,v2.3.4+)
		CpuLimit      int    `json:"cpu_limit"`                // 跟随系统设置		cpu限制(0-100,v2.3.9+)
		ResetFlowDay  int    `json:"reset_flow_day,omitempty"` // 流量重置日(1-31,v2.3.9+)
		CpuModel      int    `json:"cpu_model"`                // CPU模式(1=host-passthrough,2=host-model,3=custom,v2.4.4+)
		Rid           int    `json:"rid"`                      // 资源包ID(前台代理商创建实例用)
		DynamicMemory int    `json:"dynamic_memory"`           // 动态内存(0=关闭,1=开启,Hyper-V,v2.5.2+)
		AdvancedCpu   int    `json:"advanced_cpu"`             // 智能CPU规则ID
		AdvancedBw    int    `json:"advanced_bw"`              // 智能带宽规则ID
		Key           string `json:"key"`                      // 镜像设置的密钥		windows密钥
		TheSamePass   int    `json:"the_same_pass"`            // 批量创建时可以指定密码是否保持一致(0=随机,1=相同密码)
		Adsl          int    `json:"adsl"`                     // ADSL账号ID(0=自动,拨号云支持)
		SshKey        int    `json:"ssh_key"`                  // SSH密钥ID(有ID优先使用ID)
		SshKeyName    string `json:"ssh_key_name"`             //	SSH密钥名称(没有密钥ID时将使用密钥名称创建,password_type=1时需要ID或者名称的一种)
		TrafficType   int    `json:"traffic_type"`             //	流量统计方向(1=进,2=出,3=总计)

		FloatIp struct {
			Num   int `json:"num"`    // 浮动IP数量
			InBw  int `json:"in_bw"`  // 浮动IP进带宽(有浮动IP数量生效)
			OutBw int `json:"out_bw"` // 浮动IP出带宽(有浮动IP数量生效)
		} `json:"float_ip"` //
	}
	DataDisk struct {
		Size             uint `json:"size"`                // 其他磁盘大小
		ReadBytesSec     int  `json:"read_bytes_sec"`      // 读取限制(MB/s)
		WriteBytesSec    int  `json:"write_bytes_sec"`     // 写入限制(MB/s)
		ReadIopsSec      int  `json:"read_iops_sec"`       // 读取限制(ops/s)
		WriteIopsSec     int  `json:"write_iops_sec"`      // 写入限制(ops/s)
		ReadBytesSecMax  int  `json:"read_bytes_sec_max"`  // 读取最大突发(MB)
		WriteBytesSecMax int  `json:"write_bytes_sec_max"` // 写入最大突发(MB)
		ReadIopsSecMax   int  `json:"read_iops_sec_max"`   // 读取最大突发(ops)
		WriteIopsSecMax  int  `json:"write_iops_sec_max"`  // 写入最大突发(ops)
		Store            int  `json:"store"`               // 跟随系统盘存储ID
	}
	CreateResp struct {
		Num           uint   `json:"num"`            // 创建成功的数量
		ID            string `json:"id"`             // 创建的实例ID
		TaskID        uint   `json:"task_id"`        // 任务ID(只有当创建1台时返回)
		SecurityGroup uint   `json:"security_group"` // 关联的安全组ID(只有当创建1台时返回,v2.3.0+)
		Ssh           struct {
			ID  uint `json:"id"` // 新创建的SSH密钥ID
			Key uint `json:"id"` // 新创建的SSH密钥内容
		} `json:"ssh"` // SSH
	}
)

func (r *CreateReq) Url() string                 { return "/v1/clouds" }
func (r *CreateReq) Method() string              { return http.MethodPost }
func (r *CreateReq) Header() http.Header         { return nil }
func (r *CreateReq) Values() (values url.Values) { return }
func (r *CreateReq) Form() (form url.Values)     { return }
func (r *CreateReq) Body() any                   { return r }

// Create 创建实例
func Create(ctx *cputil.Context, req *CreateReq) (*CreateResp, error) {
	if resp, err := cputil.Execute[*CreateReq, *CreateResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
