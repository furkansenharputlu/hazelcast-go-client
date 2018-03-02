// Copyright (c) 2008-2018, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
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
package protocol

import ()

type MultiMapSizeResponseParameters struct {
	Response int32
}

func MultiMapSizeCalculateSize(name *string) int {
	// Calculates the request payload size
	dataSize := 0
	dataSize += StringCalculateSize(name)
	return dataSize
}

func MultiMapSizeEncodeRequest(name *string) *ClientMessage {
	// Encode request into clientMessage
	clientMessage := NewClientMessage(nil, MultiMapSizeCalculateSize(name))
	clientMessage.SetMessageType(MULTIMAP_SIZE)
	clientMessage.IsRetryable = true
	clientMessage.AppendString(name)
	clientMessage.UpdateFrameLength()
	return clientMessage
}

func MultiMapSizeDecodeResponse(clientMessage *ClientMessage) *MultiMapSizeResponseParameters {
	// Decode response from client message
	parameters := new(MultiMapSizeResponseParameters)
	parameters.Response = clientMessage.ReadInt32()
	return parameters
}
