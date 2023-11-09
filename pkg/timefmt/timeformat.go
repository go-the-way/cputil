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

package timefmt

import "time"

func Utc2Gmt8(utcTime string) string {
	if utcTime == "" {
		return ""
	}
	parsed, err := time.Parse("2006-01-02T15:04:05Z", utcTime)
	if err != nil {
		return ""
	}
	return parsed.Add(time.Hour * 8).Format("2006-01-02 15:04:05")
}
