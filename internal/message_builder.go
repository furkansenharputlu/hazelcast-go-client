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

package internal

import (
	"github.com/hazelcast/hazelcast-go-client/internal/common"
	"github.com/hazelcast/hazelcast-go-client/internal/protocol"
)

type clientMessageBuilder struct {
	incompleteMessages map[int64]*protocol.ClientMessage
	responseChannel    chan *protocol.ClientMessage
}

func (mb *clientMessageBuilder) onMessage(msg *protocol.ClientMessage) {
	if msg.HasFlags(common.BeginEndFlag) > 0 {
		mb.responseChannel <- msg
	} else if msg.HasFlags(common.BeginFlag) > 0 {
		mb.incompleteMessages[msg.CorrelationId()] = msg
	} else {
		message, found := mb.incompleteMessages[msg.CorrelationId()]
		if !found {
			return
		}
		message.Accumulate(msg)
		if msg.HasFlags(common.EndFlag) > 0 {
			message.AddFlags(common.BeginEndFlag)
			mb.responseChannel <- message
			delete(mb.incompleteMessages, msg.CorrelationId())
		}
	}
}
