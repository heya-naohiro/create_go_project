package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func create_go_project(reponame string, directory string, base string) error {
	os.Chdir(base)
	if _, err := os.Stat(directory); err == nil {
		return fmt.Errorf("error : already exist dir %s", directory)
	}
	if err := os.Mkdir(directory, 0777); err != nil {
		return fmt.Errorf("error : directory create error %s, %s", directory, err)
	}
	cmd := exec.Command("go", "mod", "init", reponame)
	cmd.Dir = directory
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		return fmt.Errorf("error : go mod init error %s, %s", reponame, err)
	}
	return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Println("args error")
		fmt.Println("Usage: creatego hogehoge")
		os.Exit(1)
	}
	reponame := flag.Arg(0)
	sl := strings.Split(reponame, "/")
	dirname := sl[len(sl)-1]
	usr, _ := user.Current()
	homedir := usr.HomeDir
	if err := create_go_project(reponame, dirname, homedir); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
