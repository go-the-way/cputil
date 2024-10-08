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
	"github.com/go-the-way/cputil"
)

// AreaName 区域名称
func AreaName(ctx *cputil.Context, cloudId uint) (string, error) {
	detailResp, err := Detail(ctx, &DetailReq{cloudId})
	if err != nil {
		return "", err
	}
	return detailResp.AreaName, err
}
