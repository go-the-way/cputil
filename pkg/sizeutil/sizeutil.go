// Copyright 2023 panelbase Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sizeutil

import (
	"fmt"
	"strings"
)

const (
	_byte     = 1024
	SizeByte  = _byte
	SizeKByte = SizeByte * SizeByte
	SizeMByte = SizeKByte * SizeByte
	SizeGByte = SizeMByte * SizeByte
	SizeTByte = SizeGByte * SizeByte
	_         = SizeTByte
)

// FormatSize returns auto calculate size format string
//
// 0	     ~ SizeByte   : in Bytes  => $N B
//
// SizeByte  ~ SizeKByte  : in KBytes => $N B
//
// SizeKByte ~ SizeMByte  : in MBytes => $N KB
//
// SizeMByte ~ SizeGByte  : in GBytes => $N MB
//
// SizeGByte ~ SizeTByte  : in GBytes => $N GB
//
// SizeTByte ~   ~        : in TBytes => $N TB
func FormatSize(bytes uint64, withSpace ...bool) (str string) {
	spaceStr := ""
	if len(withSpace) > 0 {
		if withSpace[0] {
			spaceStr = " "
		}
	}
	scaleStr := ""
	switch {
	case bytes < SizeByte:
		str = fmt.Sprintf("%d%sB", bytes, spaceStr)
	case bytes < SizeKByte:
		str = fmt.Sprintf("%.1f%sKB", float64(bytes)/SizeByte, spaceStr)
		scaleStr = ".0"
	case bytes < SizeMByte:
		str = fmt.Sprintf("%.2f%sMB", float64(bytes)/SizeKByte, spaceStr)
		scaleStr = ".00"
	case bytes < SizeGByte:
		str = fmt.Sprintf("%.3f%sGB", float64(bytes)/SizeMByte, spaceStr)
		scaleStr = ".000"
	default:
		str = fmt.Sprintf("%.4f%sTB", float64(bytes)/SizeGByte, spaceStr)
		scaleStr = ".0000"
	}

	if scaleStr != "" {
		// replace like 120.0KB    => 120KB
		// replace like 120.00MB   => 120MB
		// replace like 120.000GB  => 120GB
		// replace like 120.0000TB => 120TB
		str = strings.ReplaceAll(str, scaleStr, "")
	}

	return
}
