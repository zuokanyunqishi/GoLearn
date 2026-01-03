package concurrencyModels

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"time"
)

// 场景 并发请求n个接口，取最快返回的那一个结果，取到后退出所有携程

// 三个模拟器的“气象数据服务中心”，然后将这三个“气象数据服务中心”的实例传入 first 函数。
// 后者创建了三个 goroutine，每个 goroutine 对应向一个“气象数据服务中心”发起查询请求。
// 三个发起查询的 goroutine 都会将应答结果写入同一个 channel 中，
// first 获取第一个结果数据后就返回了。并关闭其他携程

type Result struct {
	Value string
}

func First(servers ...*httptest.Server) (Result, error) {

	ctx, cancel := context.WithCancel(context.Background())

	rchan := make(chan Result)
	defer cancel()
	queryFun := func(i int, server *httptest.Server) {
		url := server.URL

		request, _ := http.NewRequest("GET", url, nil)
		request.WithContext(ctx)
		response, err := http.DefaultClient.Do(request)

		if err != nil {
			fmt.Println(i, " 号server 失败 ", err.Error())
			return
		}

		defer response.Body.Close()
		res, _ := io.ReadAll(response.Body)
		rchan <- Result{
			Value: string(res),
		}
		return
	}

	for i, server := range servers {
		go queryFun(i, server)
	}

	select {
	case result := <-rchan:
		return result, nil
	case <-time.After(time.Second * 3):
		return Result{}, errors.New("time out")

	}
}

func fakeHttpServer(name string, timeout time.Duration) *httptest.Server {

	hf := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(name, " server 接收到请求....")
		time.Sleep(timeout)

		writer.Write([]byte(name + " server预报\r\n 北京： 晴朗 ，0° "))
	})
	return httptest.NewServer(hf)
}

func FastOne(times ...time.Duration) (Result, error) {
	if len(times) > 0 {
		servers := make([]*httptest.Server, len(times))
		for i, duration := range times {
			servers[i] = fakeHttpServer(fmt.Sprintf("%d", i), duration)
		}

		return First(servers...)
	}
	return First(
		fakeHttpServer("北京电视台", time.Second*5),
		fakeHttpServer("中央电视台", time.Second*10),
		fakeHttpServer("河南电视台", time.Second*4))

}
