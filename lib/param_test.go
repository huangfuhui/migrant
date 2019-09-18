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
	"testing"
)

func TestNewParam(t *testing.T) {
	s := []string{
		`--name=hello`,
		`--Num0=98`,
		`--type==youKnown`,
		`--=`,
		`-ping=pong`,
	}

	param, err := NewParam(s)
	if err == nil || param == nil {
		t.Error("the regexp-parser doesn't work well, some params are not supposed.")
	} else if len(param.values) != 2 || param.values["name"] != "hello" || param.values["Num0"] != "98" {
		t.Error("the regexp-parser doesn't work well, and the parse result is:", param.values)
	}
}
