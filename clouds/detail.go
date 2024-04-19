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
	"github.com/rwscode/cputil"
	"net/http"
	"net/url"
)

type (
	DetailReq struct {
		ID uint `json:"-"`
	}
	DetailRespDisk struct {
		Id               int      `json:"id"`
		Type             string   `json:"type"`
		Size             int      `json:"size"`
		Status           int      `json:"status"`
		Dev              string   `json:"dev"`
		Io               string   `json:"io"`
		Driver           string   `json:"driver"`
		Cache            string   `json:"cache"`
		CreateTime       string   `json:"create_time"`
		Name             string   `json:"name"`
		FsType           string   `json:"fs_type"`
		Remarks          string   `json:"remarks"`
		ReadBytesSec     int      `json:"read_bytes_sec"`
		WriteBytesSec    int      `json:"write_bytes_sec"`
		ReadIopsSec      int      `json:"read_iops_sec"`
		WriteIopsSec     int      `json:"write_iops_sec"`
		ReadBytesSecMax  int      `json:"read_bytes_sec_max"`
		WriteBytesSecMax int      `json:"write_bytes_sec_max"`
		ReadIopsSecMax   int      `json:"read_iops_sec_max"`
		WriteIopsSecMax  int      `json:"write_iops_sec_max"`
		StoreId          int      `json:"store_id"`
		IopsMin          int      `json:"iops_min"`
		IopsMax          int      `json:"iops_max"`
		DeleteWithHost   int      `json:"delete_with_host"`
		CreateWithHost   int      `json:"create_with_host"`
		EnableNode       []string `json:"enable_node"`
		StoreType        string   `json:"store_type"`
		StoreName        string   `json:"store_name"`
		RealDir          string   `json:"real_dir"`
		SnapNum          int      `json:"snap_num"`
	}
	DetailRespNetwork struct {
		Id         int    `json:"id"`
		Mac        string `json:"mac"`
		Bridge     string `json:"bridge"`
		Niccard    int    `json:"niccard"`
		InsideVlan string `json:"inside_vlan"`
		Ipaddress  []struct {
			Id         int    `json:"id"`
			Ipaddress  string `json:"ipaddress"`
			Bwid       int    `json:"bwid"`
			Remark     string `json:"remark"`
			Interface  int    `json:"interface"`
			Gateway    string `json:"gateway"`
			SubnetMask string `json:"subnet_mask"`
		} `json:"ipaddress"`
		NetId         int    `json:"net_id"`
		Vpc           int    `json:"vpc"`
		SubnetMask    string `json:"subnet_mask"`
		BindMac       int    `json:"bind_mac"`
		Bwid          int    `json:"bwid"`
		InBw          int    `json:"in_bw"`
		OutBw         int    `json:"out_bw"`
		NetworkName   string `json:"network_name"`
		InterfaceName string `json:"interface_name"`
		NetName       string `json:"net_name"`
		NiccardName   string `json:"niccard_name"`
		Name          string `json:"name"`
	}
	DetailResp struct {
		Acpi       int `json:"acpi"`
		AdvancedBw struct {
			Id int `json:"id"`
		} `json:"advanced_bw"`
		AdvancedCpu struct {
			Id int `json:"id"`
		} `json:"advanced_cpu"`
		Apic       int    `json:"apic"`
		AreaId     int    `json:"area_id"`
		AreaName   string `json:"area_name"`
		BackupNum  int    `json:"backup_num"`
		BackupUsed int    `json:"backup_used"`
		Bootorder  string `json:"bootorder"`
		BwGroup    []struct {
			Id    int `json:"id"`
			InBw  int `json:"in_bw"`
			OutBw int `json:"out_bw"`
		} `json:"bw_group"`
		ClientShowIpRemark int    `json:"client_show_ip_remark"`
		Clock              int    `json:"clock"`
		Console            int    `json:"console"`
		Cpu                int    `json:"cpu"`
		CpuLimit           string `json:"cpu_limit"`
		Cpucores           int    `json:"cpucores"`
		Cpumatch           int    `json:"cpumatch"`
		Cpumodel           int    `json:"cpumodel"`
		Cpunumacheck       int    `json:"cpunumacheck"`
		Cpunumatop         string `json:"cpunumatop"`
		Cpusockets         int    `json:"cpusockets"`
		Cputhreads         int    `json:"cputhreads"`
		Cputopcheck        int    `json:"cputopcheck"`
		Cputune            string `json:"cputune"`
		Cputunecheck       int    `json:"cputunecheck"`
		CreateTime         string `json:"create_time"`
		CrossAreaMigrate   bool   `json:"cross_area_migrate"`
		DefaultBwGroup     struct {
			Id    int `json:"id"`
			InBw  int `json:"in_bw"`
			OutBw int `json:"out_bw"`
			Ip    []struct {
				Id         int    `json:"id"`
				Ipaddress  string `json:"ipaddress"`
				SubnetMask string `json:"subnet_mask"`
				Gateway    string `json:"gateway"`
				Remark     string `json:"remark"`
			} `json:"ip"`
		} `json:"default_bw_group"`
		Disk          []DetailRespDisk `json:"disk"`
		DynamicMemory int              `json:"dynamic_memory"`
		Gpu           int              `json:"gpu"`
		Hostname      string           `json:"hostname"`
		Id            int              `json:"id"`
		ImageGroupId  int              `json:"image_group_id"`
		Internet      int              `json:"internet"`
		Ip            []struct {
			Id        uint   `json:"id"`
			Ipaddress string `json:"ipaddress"`
			Remark    string `json:"remark"`
			Interface int    `json:"interface"`
		} `json:"ip"`
		IpNum                  int                 `json:"ip_num"`
		IsNatHost              bool                `json:"is_nat_host"`
		Iso                    []interface{}       `json:"iso"`
		Kvmid                  string              `json:"kvmid"`
		LinkClone              int                 `json:"link_clone"`
		LinkFile               string              `json:"link_file"`
		LinkId                 int                 `json:"link_id"`
		LinkType               string              `json:"link_type"`
		Lock                   int                 `json:"lock"`
		Mainip                 string              `json:"mainip"`
		Memory                 float64             `json:"memory"`
		NatAclLimit            int                 `json:"nat_acl_limit"`
		NatAclUsed             int                 `json:"nat_acl_used"`
		NatWebLimit            int                 `json:"nat_web_limit"`
		NatWebUsed             int                 `json:"nat_web_used"`
		Network                []DetailRespNetwork `json:"network"`
		NetworkType            string              `json:"network_type"`
		NodeId                 int                 `json:"node_id"`
		NodeName               string              `json:"node_name"`
		NodeVlanRange          string              `json:"node_vlan_range"`
		OperateSystem          string              `json:"operate_system"`
		OsName                 string              `json:"os_name"`
		Osuser                 string              `json:"osuser"`
		Pae                    int                 `json:"pae"`
		PanelPass              string              `json:"panel_pass"`
		ParentName             interface{}         `json:"parent_name"`
		Port                   int                 `json:"port"`
		Queues                 int                 `json:"queues"`
		Remark                 string              `json:"remark"`
		Rescue                 int                 `json:"rescue"`
		RescueSystem           int                 `json:"rescue_system"`
		ResetFlowDay           int                 `json:"reset_flow_day"`
		ResourcePackage        []interface{}       `json:"resource_package"`
		Rid                    int                 `json:"rid"`
		Rootpassword           string              `json:"rootpassword"`
		Security               int                 `json:"security"`
		SecurityName           string              `json:"security_name"`
		SecurityRandPortDelete bool                `json:"security_rand_port_delete"`
		SingleIpNat            int                 `json:"single_ip_nat"`
		SkipColdBackup         int                 `json:"skip_cold_backup"`
		SnapNum                int                 `json:"snap_num"`
		SnapUsed               int                 `json:"snap_used"`
		SshKey                 struct {
			Id   uint   `json:"id"`   // SSH密钥ID
			Name string `json:"name"` // SSH密钥名称
		} `json:"ssh_key"`
		Status           string      `json:"status"`
		SupportNat       int         `json:"support_nat"`
		SuspendType      string      `json:"suspend_type"`
		Svg              int         `json:"svg"`
		System           int         `json:"system"`
		TmpTraffic       int         `json:"tmp_traffic"`
		TrafficQuota     int         `json:"traffic_quota"`
		TrafficStartTime string      `json:"traffic_start_time"`
		Trunk            int         `json:"trunk"`
		Type             string      `json:"type"`
		Uid              int         `json:"uid"`
		UserId           int         `json:"user_id"`
		Username         string      `json:"username"`
		Uuid             string      `json:"uuid"`
		Video            int         `json:"video"`
		Video2D          int         `json:"video2d"`
		Video3D          int         `json:"video3d"`
		Vlan             interface{} `json:"vlan"`
		Vnc              int         `json:"vnc"`
		Vncpass          string      `json:"vncpass"`
		Vncport          int         `json:"vncport"`
		VpcMac           string      `json:"vpc_mac"`
		VpcName          string      `json:"vpc_name"`
		Vram             int         `json:"vram"`
		WaitRebootReason string      `json:"wait_reboot_reason"`
	}
)

func (r *DetailReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d", r.ID) }
func (r *DetailReq) Method() string              { return http.MethodGet }
func (r *DetailReq) Header() http.Header         { return nil }
func (r *DetailReq) Values() (values url.Values) { return }
func (r *DetailReq) Form() (form url.Values)     { return }
func (r *DetailReq) Body() any                   { return nil }

// Detail 实例列表
func Detail(ctx *cputil.Context, req *DetailReq) (*DetailResp, error) {
	if resp, err := cputil.Execute[*DetailReq, *DetailResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
