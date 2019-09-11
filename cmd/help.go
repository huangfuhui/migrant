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
	"errors"
	"fmt"
)

// show the usage of migrant.
func ShowUsage() string {
	return `
A migrant tool for MySQL written by golang.

Usage:
	command [options] [arguments]

Available commands:
	help		Displays help for a command.
	migrate		Migrate the migration file.
	make		Create a new migration file.
	rollback	Rollback the last database migration.
	reset		Rollback all database migrations.
	fresh		Drop all tables and re-run all migrations.
`
}

type HelpCmd struct {
	Command *Cmd
}

func (h *HelpCmd) Exec() (err error) {
	args := h.Command.Args
	if len(args) == 0 {
		fmt.Println(ShowUsage())
		return
	} else if len(args) > 1 {
		return errors.New("command 'help' just need one argument.")
	} else {
		helper := helperFactory(args[0])
		if helper == nil {
			fmt.Println("could not find usage of '" + args[0] + "'.")
			return
		}
		fmt.Println(helper.Usage())
	}

	return
}

func (h *HelpCmd) Usage() string {
	return `
Displays help for a command.

Usage:
	help [command]
`
}

type helper interface {
	Usage() string
}

// This factory would return an instance of Command,
// and suppose be use on help command only.
func helperFactory(c string) helper {
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
