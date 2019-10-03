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
	"errors"
	"regexp"
)

type Param struct {
	values map[string]string
}

// Parse the command arguments and return Param.
func NewParam(s []string) (param *Param, err error) {
	param = &Param{}
	param.values = make(map[string]string)

	reg := regexp.MustCompile(`^--(\w+?)=(.+?)$`)
	for _, v := range s {
		res := reg.FindStringSubmatch(v)

		if len(res) != 3 {
			err = errors.New(ArgsErrMsg(v, []string{}))
			return
		}

		param.values[res[1]] = res[2]
	}

	return
}

func (p *Param) Get(k string) (v string) {
	return p.values[k]
}

func (p *Param) All() map[string]string {
	return p.values
}
