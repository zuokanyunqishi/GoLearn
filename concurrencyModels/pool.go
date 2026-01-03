package concurrencyModels

import "fmt"

// 应用场景，需要下载上千万数量级图片，用一定数据量携程处理，处理完后，工作携程退出

// 使用 close(chan)，所有接收该管道的协程会收到 nil 和 false，false 代表管道已经关闭

// WorkPoolsByChan  工作池模型
// urls []string
func WorkPoolsByChan(urls []string, workNums int) {

	// 投递管道
	taskChan := make(chan string, 30)
	exit := make(chan struct{})

	// 开启工作携程

	for i := 0; i < workNums; i++ {
		go WorkByChan(taskChan, i, exit)
	}

	// 投递任务
	for _, url := range urls {
		taskChan <- url
	}

	// 关闭管道，
	close(taskChan)

	// 等待全部携程退出
	for i := 0; i < workNums; i++ {
		<-exit
	}

}

func WorkByChan(taskChan chan string, i int, exit chan struct{}) {

	// 在生产数据端关闭携程，for range 会自动退出循环
	for s := range taskChan {
		// do something
		_ = s
	}
	fmt.Println(i, "号携程退出")
	exit <- struct{}{}

}
