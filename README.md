# local-mock
[![BuildStatus](https://github.com/ZebraKK/local-mock/actions/workflows/main.yml/badge.svg)]
[![Go Report Card](https://goreportcard.com/badge/github.com/zebrakk/local-mock)](https://goreportcard.com/report/github.com/zebrakk/local-mock)

## summary
简易的http client端和 server端，方便测试。
测试场景：
	local-client -> 目标缓存服务 -> local-server

## client
* hash校验响应内容

## server
* 不同path统一响应相同大小的随机内容
* 识别range/304
* 响应头：X-Md5/ETag/Last-Modify

## todo-list
* log
* 请求结果统计（clt，svr联动）
