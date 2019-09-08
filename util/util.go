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

package util

// tips for wrong arguments.
func ArgsErrMsg(args string, maybe []string) string {
	msg := "the argument '" + args + "' is not valid."
	if len(maybe) > 0 {
		msg += " maybe try with '"
		for _, v := range maybe {
			msg += v + " "
		}
		msg += "' again."
	}

	return msg
}
