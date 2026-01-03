package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 日期
func main() {
	// 1.  当前时间
	now := time.Now()
	// now time : 2020-10-11 14:08:19.687646516 +0800 CST m=+0.000066408, type is time.Time
	fmt.Printf("now time : %v, type is %T\n", now, now)

	// 2. 当前时间的年月日
	fmt.Printf("%v年%v月%v日 %v时%v分%v秒\n", now.Year(), int(now.Month()), now.Day(), now.Hour(),
		now.Minute(), now.Second())

	// 3. 格式化
	format := now.Format("2006-01-02 15:04:05")
	fmt.Println(format)

	// 4. 休眠xx
	time.Sleep(time.Second)

	// 5. 时间戳
	unix := now.Unix()
	fmt.Println(unix)           //秒级时间戳
	fmt.Println(now.UnixNano()) // 纳秒时间戳

	// 拼接字符串用 10万次 + 号耗时 12秒
	appendStr()

	// 拼接字符串 strings.builder 耗时 10101552 纳秒
	appendBuilderStr()
}

func appendStr() {

	now := time.Now().Unix()
	str := ""
	for i := 0; i < 100000; i++ {
		str += strconv.Itoa(i) + "hello"
	}
	end := time.Now().Unix()
	fmt.Println("耗时: ", end-now)
}

func appendBuilderStr() {
	now := time.Now().UnixNano()
	builder := strings.Builder{}
	for i := 0; i < 100000; i++ {
		builder.WriteString(strconv.Itoa(i) + "hello")
	}

	builder.String()
	end := time.Now().UnixNano()
	fmt.Println("耗时: ", end-now)
}
