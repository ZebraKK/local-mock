package main

import (
	"fmt"

	"http_clt/task"
)

type Cmd struct {
	name   string
	domain string
	uri    string
}

func main() {
	taskMgr := task.NewTaskMgr()

	// exchange in/out
	cmd := &Cmd{
		name: "job",
	}
	for {
		if exchangeCMD(cmd) == false {
			continue
		}
		switch cmd.name {
		case "test":
			fmt.Println("hello, test")
		case "job":
			domain := "yxw.iqy.qbox.net"
			uri := "/64K"
			cnt := 1
			go taskMgr.StartJob(uri, domain, cnt)
		default:
		}
	}
	return
}

func exchangeCMD(cmd *Cmd) bool {
	fmt.Println("请输入: <domain> <uri>")
	fmt.Scanln(&cmd.domain, &cmd.uri)
	fmt.Printf("domain: %v, uri: %v, 回车确认", cmd.domain, cmd.uri)
	var check string
	fmt.Scanln(&check)
	if check != "" {
		return false
	}

	return true
}
