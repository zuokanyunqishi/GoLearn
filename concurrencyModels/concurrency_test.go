package concurrencyModels

import (
	"testing"
	"time"
)

func TestWorkPoolsByChan(t *testing.T) {
	urls := make([]string, 10000)
	// 输出
	WorkPoolsByChan(urls, 100)
}

func TestRunQuit(t *testing.T) {
	RunQuit1()
}

func TestWhispers(t *testing.T) {
	Whispers()
}

func TestFastOne(t *testing.T) {
	if result, err := FastOne(); err != nil {
		t.Error(err)
	} else {
		t.Log(result)
	}

	start := time.Now()

	result, err := FastOne(time.Millisecond*200, time.Second*2, time.Millisecond*1000)
	duration := time.Since(start)
	t.Log("耗时", duration)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(result)
	}
}
