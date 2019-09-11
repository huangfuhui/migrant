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

package cmd

import (
	"fmt"
	"migrant/util"
	"strings"
)

const (
	Help     = "help"
	Migrate  = "migrate"
	Make     = "make"
	Rollback = "rollback"
	Reset    = "reset"
	Fresh    = "fresh"
)

type Cmd struct {
	Command string
	Args    []string
}

var cmd *Cmd

// init Cmd.
func NewCmd(args []string) (command *Cmd, err error) {
	cmd = &Cmd{}

	args = args[1:]
	if len(args) == 0 {
		fmt.Println(ShowUsage())
		return
	}

	c := strings.ToLower(args[0])
	cmd.Command = c
	cmd.Args = args[1:]

	commander := CommandFactory(c)
	if commander == nil {
		// args error
		fmt.Println(util.ArgsErrMsg(c, []string{}))
		return
	}
	err = commander.Exec()
	if err != nil {
		fmt.Println(err.Error())
	}

	return
}

// execute migration.
func (c *Cmd) Exec() error {
	return nil
}

// Any Commands should inherit this interface.
type Command interface {
	Usage() string
	Exec() error
}

// A factory which produce CMD instance.
func CommandFactory(c string) Command {
	switch c {
	case Help:
		return &HelpCmd{cmd}
	case Migrate:

	case Make:

	case Rollback:

	case Reset:

	case Fresh:
	}

	return nil
}
