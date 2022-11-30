package main

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
	"io"
)

type HashCheck struct {
	h       hash.Hash
	body    io.ReadCloser
	hashStr string
}

// clt 读
func NewHashCheck(body io.ReadCloser) *HashCheck {
	return &HashCheck{
		h:    md5.New(),
		body: body,
	}
}

func (hc *HashCheck) Read(bs []byte) (n int, err error) {
	n, err = hc.body.Read(bs)
	if n > 0 {
		if n, err := hc.h.Write(bs[:n]); err != nil {
			return n, err
		}
	}
	if err == io.EOF {
		hc.hashStr = hex.EncodeToString(hc.h.Sum(nil))
	}
	return
}

/*
  hash是一个通用接口
  具体的 crypto/xxx 实现hash接口
  通常，我们会封装一个 md5 函数，对字符串进行处理，有时候也会将文件读取成字节进行处理。
  这里不建议将文件读取到 []byte 后，进行 hash，
  而是直接使用 writer 接口进行处理，这样占用的内存更少，速度也更快。
*/

/*
f, err := os.Open("./test.file")
if err != nil {
	return
}
hasher := md5.New()
io.Copy(hasher, f) //将数据写入hasher  cp(dst, src)
hash.Sum(nil) //计算hash值
*/
