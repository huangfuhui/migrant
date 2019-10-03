package cmd

import (
	"migrant/lib"
	"os"
)

type ConfCmd struct {
	Command *Cmd
	File    string
}

func NewConfCmd(cmd *Cmd) (cc *ConfCmd, err error) {
	cc = &ConfCmd{}
	cc.Command = cmd

	param, err := lib.NewParam(cmd.Args)
	if err != nil {
		return
	}
	cc.File = param.Get("file")

	return
}

func (c *ConfCmd) Usage() string {
	return `
Create and init a config file.

Usage:
	conf [--argument=value]...

Available commands:
	file		Config file path, create './migrant_conf.json' by default if there ara not this arg.
	
`
}

func (c *ConfCmd) Exec() (err error) {
	if c.File == "" {
		c.File = "./migrant.conf"
	}

	f, err := os.OpenFile(c.File, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return
	}
	defer func() {
		_ = f.Close()
	}()
	_, err = f.Write(c.conf())
	return
}

func (c *ConfCmd) conf() []byte {
	return []byte(`{
    "driver":"mysql",
    "dns":"",
    "table_name":"migrant",
    "table_prefix":""
}`)
}
