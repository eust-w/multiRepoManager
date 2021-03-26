package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"time"
)

type conf struct {
	PullRepoPath []string `yaml:"pullRepoPath,flow"`
	PushRepoPath []string `yaml:"pushRepoPath,flow"`
	Proxy        string   `yaml:"proxy"`
	ProxyName    string   `yaml:"proxyName"`
}

func getConf(c *conf) {
	ex, _ := os.Executable()
	config, err := ioutil.ReadFile(path.Join(filepath.Dir(ex), "./config.yaml"))
	if err != nil {
		fmt.Print(err)
	}
	err1 := yaml.Unmarshal(config, &c)
	if err1 != nil {
		fmt.Println("error")
	}
}

func run(name string, dir string, cmd ...string){
	var _tem = []string{"/c"}
	cmds := append(_tem, cmd...)
	println(cmds)
	cmdRun := exec.Command(name, cmds...)
	cmdRun.Dir = dir
	cmdRun.Run()
}

func Push() {
	var c conf
	getConf(&c)
	cmd1 := []string{"/C","start", c.Proxy}
	cmd2 := []string{"/f", "/im", c.ProxyName}
	k:= exec.Command("taskkill", cmd2...).Run()
	exec.Command("cmd", cmd1...).Start()
	time.Sleep(time.Duration(1)*time.Second)
	for _, dir := range c.PushRepoPath {
		run("cmd",dir,  "git", "add", "*")
		run("cmd",dir,  "git","commit", "-m", "\"auto commit\"")
		run("cmd",dir,  "git", "push")
		time.Sleep(time.Duration(1)*time.Second)
	}
	if k != nil{
		exec.Command("taskkill", cmd2...).Start()
	}
}

func Pull() {
	var c conf
	getConf(&c)
	cmd1 := []string{"/C","start", c.Proxy}
	cmd2 := []string{"/f", "/im", c.ProxyName}
	k:= exec.Command("taskkill", cmd2...).Run()
	exec.Command("cmd", cmd1...).Start()
	time.Sleep(time.Duration(1)*time.Second)
	for _, dir := range c.PushRepoPath {
		run("cmd",dir,  "git", "pull")
		run("cmd",dir,  "git", "merge")
		time.Sleep(time.Duration(1)*time.Second)
	}
	if k != nil{
		exec.Command("taskkill", cmd2...).Start()
	}
}

func main() {
	defer func() {

		if r := recover(); r != nil {
			fmt.Println("please with arg: push or pull")
		}
	}()
	cmd := os.Args[1]
	if cmd == "push" {
		Push()
	} else if cmd == "pull" {
		Pull()
	} else {
		fmt.Println("please with arg: push or pull ")
	}

}
