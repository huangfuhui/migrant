package cmd

import "testing"

var c *ConfCmd

func TestNewConfCmd(t *testing.T) {
	var err error
	c, err = NewConfCmd(&Cmd{Args: []string{"--file=../debug/migrant.conf"}})
	if err != nil {
		t.Error(err)
	}
}

func TestConfCmd_Usage(t *testing.T) {
	if len(c.Usage()) == 0 {
		t.Error("usage for config command should not be nil")
	}
}

func TestConfCmd_Exec(t *testing.T) {
	err := c.Exec()
	if err != nil {
		t.Error(err)
	}
}

func TestConfCmd_conf(t *testing.T) {
	if len(c.conf()) == 0 {
		t.Error("the standard config string should not be nil")
	}
}
