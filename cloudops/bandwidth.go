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
	BandwidthReq struct {
		ID    uint `json:"-"`      // 实例ID
		InBw  int  `json:"in_bw"`  // 进带宽
		OutBw int  `json:"out_bw"` // 出带宽
	}
	BandwidthResp struct{}
)

func (r *BandwidthReq) Url() string                 { return fmt.Sprintf("/v1/clouds/%d/bw", r.ID) }
func (r *BandwidthReq) Method() string              { return http.MethodPut }
func (r *BandwidthReq) Header() http.Header         { return nil }
func (r *BandwidthReq) Values() (values url.Values) { return }
func (r *BandwidthReq) Form() (form url.Values)     { return }
func (r *BandwidthReq) Body() any                   { return r }

// Bandwidth 修改实例带宽
func Bandwidth(ctx *cputil.Context, req *BandwidthReq) (*BandwidthResp, error) {
	if resp, err := cputil.Execute[*BandwidthReq, *BandwidthResp](ctx, req); err != nil {
		return nil, err
	} else if e := resp.Error; e != "" {
		return nil, errors.New(e)
	} else {
		return resp.Result, err
	}
}
