package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Age        int      `json:"age"`
	Name       string   `json:"name"`
	Sex        string   `json:"sex"`
	Skills     []string `json:"skills"`
	Num1, Num2 int
}

func (s Student) GetSub() {

	fmt.Println("完成计算 8-3 = ", s.Num1-s.Num2)
}

// 反射
func main() {
	var (
		xiaoming *Student
	)

	sType := reflect.TypeOf(xiaoming)

	name, ok := sType.Elem().FieldByName("Age")
	if !ok {
		panic(ok)
	}
	fmt.Println(name.Tag)
	trueType := sType.Elem()
	ptrStudent := reflect.New(trueType)
	xiaoming = ptrStudent.Interface().(*Student)
	xiaoming.GetSub()

	ptrStudent.Elem().FieldByName("Age").SetInt(13)
	ptrStudent.Elem().FieldByName("Name").SetString("张帅")
	ptrStudent.Elem().FieldByName("Sex").SetString("男")
	ptrStudent.Elem().FieldByName("Skills").Set(reflect.ValueOf([]string{"打酱油"}))
	ptrStudent.Elem().FieldByName("Num1").SetInt(3)
	ptrStudent.Elem().FieldByName("Num2").SetInt(10)
	ptrStudent.Elem().MethodByName("GetSub").Call(nil)
	ptrStudent.Elem().Method(0).Call(nil)
	fmt.Println(ptrStudent.Elem().FieldByName("Age"))

	xiaoming.GetSub()

}
