package task

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type TaskMgr struct {
}

func NewTaskMgr() *TaskMgr {
	return &TaskMgr{}
}

func (m *TaskMgr) StartJob(cnt int) {
	var wg sync.WaitGroup
	for i := 0; i < cnt; i += 1 {
		wg.Add(1)
		go m.callAndCheck(i, &wg)
	}

	wg.Wait()

	return
}

func (m *TaskMgr) callAndCheck(index int, wg *sync.WaitGroup) {
	defer wg.Done()

	requrl := "http://127.0.0.1:8080/abc.file?index=" + strconv.Itoa(index)

	req, err := http.NewRequest("GET", requrl, nil)
	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	expectMd5 := resp.Header.Get("X-Md5")

	defer resp.Body.Close()

	hc := NewHashCheck(resp.Body)

	//io.Discard
	buf := make([]byte, 1024)
	io.CopyBuffer(io.Discard, hc, buf)

	if hc.hashStr != expectMd5 {
		log.Println("hash check err")
	}
}
