package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"strconv"
	"strings"
)

var InputIntError = errors.New("输入参数转换数字错误。。。")

// CustomerView
type CustomerView struct {
	input   string
	reader  *bufio.Reader
	writer  *bufio.Writer
	service *CustomerService
}

func NewCustomerView() *CustomerView {
	return &CustomerView{
		reader:  bufio.NewReader(os.Stdin),
		writer:  bufio.NewWriter(os.Stdout),
		service: NewCustomerService(),
	}
}

// 读取输入
func (v *CustomerView) read() string {
	readString, _ := v.reader.ReadString('\n')
	return strings.Trim(readString, "\n")
}

// 写入终端
func (v *CustomerView) writeln(contents ...string) {
	for _, content := range contents {
		v.writer.WriteString(content + "\n")
	}
	v.writer.Flush()
}

// 数据转换
func (v *CustomerView) formatInt(input string) (int, error) {
	inputInt, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err, "转换错误。。。")
		return 0, InputIntError
	}
	return inputInt, nil
}

// 数据转换
func (v *CustomerView) formatUnit8(input string) (uint8, error) {
	inputUnit, err := strconv.ParseUint(input, 10, 8)
	if err != nil {
		return 0, InputIntError
	}
	return uint8(inputUnit), nil
}

func (v *CustomerView) formatUnit(input string) (uint, error) {
	inputUnit, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		return 0, InputIntError
	}

	return uint(inputUnit), nil
}

// 主目录
func (v *CustomerView) MainMenu() {
	for {
		v.writeln(
			"----------------客户信息管理系统----------",
			"              1. 添加客户          ",
			"              2. 修改客户          ",
			"              3. 删除客户          ",
			"              4. 客户列表          ",
			"              5. 退出系统          ",
		)
		v.handleInput()
	}
}

// 处理输入
func (v *CustomerView) handleInput() {
	switch v.read() {
	case "1":
		// 添加客户
		v.Add()
		break
	case "2":
		//
		v.Modify()
		break
	case "3":
		// 删除客户
		v.Delete()
		break
	case "4":
		// 客户列表
		v.list()
		break
	default:
		v.writeln("退出中...")
		os.Exit(0)
	}
}

// 添加客户
func (v *CustomerView) Add() {

RetryAdd:
	v.writeln("-------------- 新增客户---------")
	name, email, phone, gender, age := v.ReadForCustomer()
	result := v.service.Add(NewCustomerModel(name, email, phone, gender, age))

	if result {
		v.writeln("添加成功,返回主菜单中...")
		return
	}

	v.writeln("添加失败,是否重新添加 ? (1:确认,2:放弃返回主界面)")
	if v.read() == "1" {
		goto RetryAdd
	}

}

func (v *CustomerView) write(contents ...string) {
	for _, content := range contents {
		v.writer.WriteString(content)
	}
	v.writer.Flush()
}

func (v *CustomerView) list() {

	tableModel := table.NewWriter()

	tableModel.AppendHeader([]interface{}{"编号", "姓名", "性别", "年纪", "手机", "电邮"})
	for _, row := range v.service.List() {
		tableModel.AppendRow(table.Row{row.Id, row.Name, row.Gender, row.Age, row.Phone, row.Email}, table.RowConfig{AutoMerge: true})
	}

	tableModel.SetColumnConfigs([]table.ColumnConfig{
		{Align: text.AlignCenter, Number: 1},
		{Align: text.AlignCenter, Number: 2},
		{Align: text.AlignCenter, Number: 3},
		{Align: text.AlignCenter, Number: 4},
		{Align: text.AlignCenter, Number: 5},
		{Align: text.AlignCenter, Number: 6},
	})

	v.writeln("--------------客户信息列表-----------------")
	v.writeln("")
	v.writeln(tableModel.Render())
	v.writeln("")
	v.writeln("--------------客户信息列表渲染完毕-----------------")

}

// 修改客户
func (v *CustomerView) Modify() {
	// 请输入要更改的客户编号
RetryM:
	v.write("请输入要更改的客户编号")
	var no uint
	var err error
	if no, err = v.formatUnit(v.read()); err != nil || no < 1 {
		v.writeln("输入非法..请重新输入")
		goto RetryM
	}
	index := v.FindById(no)

	if index < 0 {
		v.writeln(fmt.Sprintf("%d 客户不存在", no))
		return
	}
	// 渲染 存在的数据
	customer := v.service.List()[index]
	v.writeln(
		fmt.Sprintf("客户编号: %d", customer.Id),
		fmt.Sprintf("姓名: %s", customer.Name),
		fmt.Sprintf("性别: %s", customer.Gender),
		fmt.Sprintf("年龄: %d", customer.Age),
		fmt.Sprintf("手机: %s", customer.Phone),
		fmt.Sprintf("电邮: %s", customer.Email),
	)
	v.writeln("")
	v.writeln("开始修改-----")
	v.service.Update(index, NewCustomerModel(v.ReadForCustomer()))
	v.writeln("", "修改完成....")

}

func (v *CustomerView) FindById(id uint) int {
	for index, customer := range v.service.List() {
		if customer.Id == id {
			fmt.Println(customer.Phone)
			return index
		}
	}
	return -1
}

func (v *CustomerView) ReadForCustomer() (name, email, phone, gender string, age uint8) {
	var err error
	v.write("姓名: ")
	name = v.read()

	v.write("性别: ")
	gender = v.read()

Age:
	v.write("年龄: ")
	age, err = v.formatUnit8(v.read())
	if err != nil {
		v.writeln(err.Error())
		v.writeln("请重新输入....")
		goto Age
	}

	v.write("电话: ")
	phone = v.read()

	v.write("电邮: ")
	email = v.read()
	return

}

func (v *CustomerView) Delete() {
	v.writeln("---------------删除客户----------")
RetryInputCustomerId:
	v.write("请输入要删除的客户编号： ")
	customerId, err := v.formatUnit(v.read())

	if err != nil {
		v.writeln("输入参数错误，请重新输入")
		goto RetryInputCustomerId
	}
	// 查找是 customerId 是否存在
	var index = v.FindById(customerId)
	if index < 0 {
		v.writeln(fmt.Sprintf("客户 %d 不存在。。。", index))
		v.write("继续删除请输入 : 1, 返回主菜单输入： 0--")

		inputMap := map[string]int{"1": 0, "0": 0}
		v.input = v.read()
		if v.input == "1" {
			goto RetryInputCustomerId
		}

		if v.input == "0" {
			return
		}

		if _, ok := inputMap[v.input]; !ok {
			v.writeln("输入错误。。返回主菜单中。。")
			return
		}
	}

	// 存在
	v.write("确认删除 ? (y/n)")
	isDelStr := v.read()
	if strings.ToLower(isDelStr) == "y" {
		// 删除数据
		v.service.DeleteCustomer(index)
		v.writeln("删除成功。。。")
		return
	}

	if strings.ToLower(isDelStr) == "n" {
		v.writeln("放弃删除。。返回主菜单")
	}

}
