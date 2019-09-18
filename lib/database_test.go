// Copyright 2019 Migrant
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

import "testing"

func TestInitDB(t *testing.T) {
	err := InitDB("mysql", "root@tcp(127.0.0.1:3306)/migrant?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Error(err)
	}
}

func TestClose(t *testing.T) {
	err := CloseDB()
	if err != nil {
		t.Error(err)
	}
}
