package task

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type TaskMgr struct {
	host string
}

func NewTaskMgr() *TaskMgr {
	return &TaskMgr{
		host: "10.202.13.23",
	}
}

func (m *TaskMgr) StartJob(uri, domain string, cnt int) {
	var wg sync.WaitGroup
	requrl := "http://" + m.host + uri + "?num="
	for i := 0; i < cnt; i += 1 {
		wg.Add(1)
		r := requrl + strconv.Itoa(i)
		go m.callAndCheck(r, domain, &wg)
	}
	wg.Wait()

	time.Sleep(5 * time.Second)

	var wgA sync.WaitGroup
	for i := 0; i < 1; i += 1 {
		wgA.Add(1)
		go func(wgA *sync.WaitGroup) {
			var wgB sync.WaitGroup
			for j := 0; j < 1; j += 1 {
				wgB.Add(1)
				num := rand.Intn(cnt)
				r := requrl + strconv.Itoa(num)
				go m.callAndCheck(r, domain, &wgB)
			}
			wgB.Wait()
			wgA.Done()
		}(&wgA)
	}
	wgA.Wait()

	return
}

func (m *TaskMgr) callAndCheck(requrl, domain string, wg *sync.WaitGroup) {
	defer wg.Done()

	req, err := http.NewRequest("GET", requrl, nil)
	if err != nil {
		log.Println(err)
	}

	req.Host = domain

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("req: ", req.URL.String())
	//fmt.Println("resp code: ", resp.StatusCode, " header: ")
	//for k, v := range resp.Header {
	//fmt.Println(k, ", ", v)
	//}

	expectMd5 := resp.Header.Get("Content-Md5")

	defer resp.Body.Close()

	hc := NewHashCheck(resp.Body)

	//io.Discard
	buf := make([]byte, 1024)
	io.CopyBuffer(io.Discard, hc, buf)

	if hc.hashStr != expectMd5 {
		//log.Println("hash check err, expectMd5: ", expectMd5, ", but actul: ", hc.hashStr)
	}
}
