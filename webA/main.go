package main

import (
	"fmt"
	"github.com/jianshao/sentinel-demo/store"
	"io"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

type SentinelServe struct {
	
}

func (s *SentinelServe)ServeHTTP(wr http.ResponseWriter, req *http.Request)  {
	
}

var (
	getCount  = 0
	postCount = 0
)

func main()  {

	if err := http.ListenAndServe(":6060", nil); err != nil {
		fmt.Println(err)
	}

	store.S = new(store.Store)
	if err := store.S.SetUp(); err != nil {
		fmt.Println("setup failed, err:", err)
	}
	http.HandleFunc("/sentinel/test", handle)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("listen failed")
	}
}

type RequestStruct struct {
	ip          string
	userId      string
	requestId   int
	requestType int
}

func parseRequest(req *http.Request) *RequestStruct {
	request := &RequestStruct{}

	request.ip = req.FormValue("ip")
	request.userId = req.FormValue("userid")
	requestId := req.FormValue("requestid")
	if requestId == "" {
		request.requestId = 0
	} else if value, err := strconv.Atoi(requestId); err != nil {
		request.requestId = value
	}
	requestType := req.FormValue("requesttype")
	if requestType == "" {
		request.requestType = 0
	} else if value, err := strconv.Atoi(requestType); err == nil{
		request.requestType = value
	}
	return request
}

func handle(wr http.ResponseWriter, req *http.Request)  {
	if err := req.ParseForm(); err != nil {
		return
	}

	_ = parseRequest(req)

	if req.Method == "GET" {
		getCount += 1
		io.WriteString(wr, fmt.Sprintf("get count = %d", getCount))
	} else if req.Method == "POST" {
		postCount += 1
		io.WriteString(wr, fmt.Sprintf("post count = %d", postCount))
	} else {
		fmt.Println("unknown method")
	}

	return
}
