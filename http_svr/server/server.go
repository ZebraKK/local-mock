package server

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"hash"
	"io"
	"net/http"
	//"os"
	"strconv"
	"time"
)

type Server struct {
	datas   []byte
	size    int
	hasher  hash.Hash
	hashStr string
}

func NewServer() *Server {
	size := 32 * 1024
	svr := &Server{
		size:   size,
		datas:  make([]byte, size),
		hasher: md5.New(),
	}

	rand.Read(svr.datas)
	svr.hasher.Write(svr.datas)
	svr.hashStr = hex.EncodeToString(svr.hasher.Sum(nil))

	return svr
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//1.req 处理信息
	/*
		filename := "myfile"
		fileSize := 100 * 1024

		f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return
		}
		hasher := md5.New()
		io.Copy(hasher, f) //将数据写入hasher  cp(dst, src)
		hashStr := hex.EncodeToString(hasher.Sum(nil))
	*/
	//2.写头
	s.setRespHeader(r, w)
	statusCode := s.getStatusCode(r)
	w.WriteHeader(statusCode)

	//3.写body
	buf := make([]byte, 1024)
	io.CopyBuffer(w, bytes.NewReader(s.datas), buf)

}

func (s *Server) dealReq(r *http.Request) (err error) {
	return
}

func (s *Server) setRespHeader(r *http.Request, w http.ResponseWriter) {
	w.Header().Set("X-Md5", s.hashStr)
	w.Header().Set("Content-Type", "txt")
	w.Header().Set("ContentLength", strconv.Itoa(s.size))
	path := r.URL.EscapedPath()
	etag := hex.EncodeToString(s.hasher.Sum([]byte(path)))
	w.Header().Set("ETag", etag)
	lm := time.Now().Format(time.UnixDate)
	w.Header().Set("Last-Modify", lm)
}

// 304/206/200等
func (s *Server) getStatusCode(r *http.Request) int {
	return http.StatusOK //200
}
