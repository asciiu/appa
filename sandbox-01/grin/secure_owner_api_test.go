// Copyright 2020 BlockCypher
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http//www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grin_test

import (
	"os"
	"testing"

	"github.com/asciiu/appa/lib/config"
	"github.com/asciiu/appa/sandbox-01/grin"
	"github.com/blockcypher/libgrin/owner"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"github.com/tj/assert"
)

func checkErr(label string, err error) {
	if err != nil {
		log.Errorf("%s: %s\n", label, err.Error())
		os.Exit(1)
	}
}

func TestReal(t *testing.T) {
	config.LoadEnv("../config/dev.env")

	var cfg grin.GrinConfig
	err := envconfig.Process("", &cfg)
	checkErr("process config", err)

	ownerAPI := owner.NewSecureOwnerAPI(cfg.URL)
	err = ownerAPI.Init()
	assert.Nil(t, err, "init owner api err was not nil")

	err = ownerAPI.Open(nil, cfg.Password)
	assert.Nil(t, err, "open wallet err was not nil")

	nodeHeight, err := ownerAPI.NodeHeight()
	assert.Nil(t, err, "node height err was not nil")
	assert.NotNil(t, nodeHeight)

	//torConfig := libwallet.TorConfig{
	//	UseTorListener: true,
	//	SocksProxyAddr: "127.0.0.1:59050",
	//	SendConfigDir:  ".",
	//}
	//if err := ownerAPI.SetTorConfig(torConfig); err != nil {
	//	assert.Error(t, err)
	//}

	fresh, summary, err := ownerAPI.RetrieveSummaryInfo(true, 10)
	assert.Nil(t, err, "retrieve summary err was not nil")

	assert.True(t, fresh, "refresh from node should be true")

	log.Infof("%+v\n", summary)

	if err := ownerAPI.Close(nil); err != nil {
		assert.Error(t, err)
	}
}
