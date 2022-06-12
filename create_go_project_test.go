package main

import (
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	if err := create_go_project("hogehoge", "hogehoge", "."); err != nil {
		if _, err := os.Stat("hogehoge"); err != nil {
			os.RemoveAll("hogehoge")
		}
		t.Errorf("create error, %s", err)
	}
	if _, err := os.Stat("./hogehoge/go.mod"); err != nil {
		if _, err := os.Stat("hogehoge"); err != nil {
			os.RemoveAll("hogehoge")
		}
	}


}
