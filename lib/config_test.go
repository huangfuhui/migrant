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
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	var err error
	if _, err := os.Stat("../test"); err != nil {
		err = os.Mkdir("../test", os.ModePerm)
		if err != nil {
			t.Error(err)
		}
	}

	confStr := `{
  "dns": "root:@tcp(127.0.0.1:3306)?charset=utf8&parseTime=True&loc=Local",
  "driver": "mysql",
  "database_name": "migrant",
  "table_prefix": ""
}`
	file, err := os.OpenFile("../test/migrant.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		t.Error(err)
	}

	_, err = file.WriteString(confStr)
	if err != nil {
		t.Error(err)
	}

	_, err = NewConfig("../test/migrant.json")
	if err != nil {
		t.Error(err)
	}
}
