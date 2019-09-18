// Copyright 2019 migrant
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lib

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	// database name source
	DNS string `json:"dns"`
	// database driver, only support MySQL now
	Driver       string `json:"driver"`
	DatabaseName string `json:"database_name"`
	TablePrefix  string `json:"table_prefix"`
}

// load config by config file
func NewConfig(configFile string) (conf *Config, err error) {
	conf = &Config{}

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, conf)
	if err != nil {
		return
	}

	return
}
