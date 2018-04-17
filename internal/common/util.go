// Copyright (c) 2008-2018, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License")
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"crypto/rand"
	"fmt"
	"io"
	"math"
	"net"
	"strconv"
	"strings"
	"time"
)

func IsValidIpAddress(addr string) bool {
	return net.ParseIP(addr) != nil
}

func GetIpAndPort(addr string) (string, int32) {
	var port int
	var err error
	parts := strings.Split(addr, ":")
	if len(parts) == 2 {
		port, err = strconv.Atoi(parts[1])
		if err != nil {
			port = 5701 // Default port
		}
	} else {
		port = -1
	}
	addr = parts[0]
	return addr, int32(port)
}

func GetTimeInMilliSeconds(duration time.Duration) int64 {
	if duration == -1 {
		return -1
	}
	if duration > 0 && duration < time.Millisecond {
		return int64(time.Millisecond)
	}
	return duration.Nanoseconds() / int64(time.Millisecond)
}

func ConvertMillisToDuration(timeInMillis int64) time.Duration {
	if timeInMillis == math.MaxInt64 {
		return time.Duration(timeInMillis)
	}
	return time.Duration(timeInMillis) * time.Millisecond
}

func ConvertMillisToUnixTime(timeInMillis int64) time.Time {
	if timeInMillis == 0 {
		return time.Time{}
	} else if timeInMillis == math.MaxInt64 {
		return time.Unix(0, timeInMillis)
	}
	return time.Unix(0, timeInMillis*int64(time.Millisecond))
}

// NewUUID generates a random UUID according to RFC 4122
func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
