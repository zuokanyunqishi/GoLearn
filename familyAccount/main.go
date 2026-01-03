package main

import (
	"GoLearn/chat/util/color"
	"bufio"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
)

var writer = bufio.NewWriter(os.Stdin)
var reader = bufio.NewReader(os.Stdin)

var db *gorm.DB

func init() {
	db1, err := gorm.Open(sqlite.Open("account.db"), &gorm.Config{})
	if err != nil {
		panic("database open fail" + err.Error())
	}
	db = db1
}

type account struct {
	ID       int
	Name     string
	Amount   float64
	CreateAt string
}

func (a account) TableName() string {
	return "account"
}

type detail struct {
	ID         int
	DetailType int
	Amount     float64
	CreateAt   string
}

func (d detail) TableName() string {
	return "details"
}

// 家庭记账终端软件
func main() {

	for {
		drawMainUI()
		notice("请输入 1~4 选择菜单")
		input := readInput()
		handleInput(input)

	}

}

// 渲染主界面

func drawMainUI() {
	writer.WriteString("------------ 家庭记账系统 --------\n")
	writer.WriteString("------------ 1. 收支明细 --------\n")
	writer.WriteString("------------ 2. 记录收入 --------\n")
	writer.WriteString("------------ 3. 记录支出 --------\n")
	writer.WriteString("------------ 3. 退出系统 --------\n")
	writer.Flush()
}

func readInput() string {
	readString, _ := reader.ReadString('\n')
	return strings.Trim(readString, "\n")
}

func handleInput(input string) {
	switch input {
	case "1":
		drawPayments()
		break
	case "2":
		handleRecord()
		break
	case "3":
		break
	case "4":
		notice("你选择了退出系统。。。")
		os.Exit(0)
	default:
		return

	}
}

func handleRecord() {
	drawRecordIncomeUi()
	var model = account{}
	amount, _ := strconv.ParseFloat(readInput(), 64)
	model.Amount = amount
	drawUi("请输入姓名: ")
	model.Name = readInput()
	db.Create(&model)
	db.Create(&detail{
		DetailType: 1,
		Amount:     model.Amount,
	})
	notice("记录成功")

}
func notice(output string) {
	writer.WriteString(output + "\n")
	writer.Flush()
}

func drawPayments() {
	var account []account
	db.Find(&account)

	tr := table.NewWriter()
	tr.AppendHeader(table.Row{"序号", "姓名", "金额"})

	for key, a := range account {
		tr.AppendRow(table.Row{key + 1, color.Red(a.Name), color.Yellow(fmt.Sprintf("%.2f", a.Amount))})
	}

	writer.WriteString(tr.Render() + "\n")
}

func drawRecordIncomeUi() {
	writer.WriteString("本次收入金额:\n")
	writer.Flush()
}

func drawUi(str ...string) {

	for _, value := range str {
		writer.WriteString(value + "\n")
	}
	writer.Flush()

}
