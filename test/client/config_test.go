// Copyright (c) 2008-2019, Hazelcast, Inc. All Rights Reserved.
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

package client

import (
	"testing"

	"github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/config"
	"github.com/hazelcast/hazelcast-go-client/test/testutil"
	"github.com/stretchr/testify/assert"
)

func TestSetGroupConfig(t *testing.T) {
	shutdownFunc := testutil.CreateCluster(remoteController)
	defer shutdownFunc()
	cfg := hazelcast.NewConfig()
	groupCfg := config.NewGroupConfig()
	groupCfg.SetName("wrongName")
	groupCfg.SetPassword("wrongPassword")
	cfg.SetGroupConfig(groupCfg)
	client, err := hazelcast.NewClientWithConfig(cfg)
	assert.Error(t, err)
	client.Shutdown()
}

func TestSetNetworkConfig(t *testing.T) {
	var expected int32 = 10
	cfg := hazelcast.NewConfig()
	nCfg := config.NewNetworkConfig()
	nCfg.SetConnectionAttemptLimit(10)
	cfg.SetNetworkConfig(nCfg)
	actual := cfg.NetworkConfig().ConnectionAttemptLimit()
	assert.Equal(t, expected, actual)
}
