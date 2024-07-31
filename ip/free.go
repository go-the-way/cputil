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

package ip

import (
	"errors"
	"github.com/go-the-way/cputil"
	q "github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

type (
	FreeReq struct {
		HostId      string `url:"hostid,omitempty"`        // 实例ID
		IpSegmentId string `url:"ip_segment_id,omitempty"` // 所属ip段id
	}
	FreeResp struct {
		Id     int    `json:"id"`
		IpName string `json:"ip_name"`
		Ip     []struct {
			Id int    `json:"id"`
			Ip string `json:"ip"`
		}
	}
)

func (r *FreeReq) Url() string                 { return "/v1/ip/free?type=cloud" }
func (r *FreeReq) Method() string              { return http.MethodGet }
func (r *FreeReq) Header() http.Header         { return nil }
func (r *FreeReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *FreeReq) Form() (form url.Values)     { return }
func (r *FreeReq) Body() any                   { return nil }

func Free(ctx *cputil.Context, req *FreeReq) ([]*FreeResp, error) {
	if resp, err := cputil.Execute[*FreeReq, []*FreeResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
