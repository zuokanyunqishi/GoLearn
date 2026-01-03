package main

import (
	"fmt"
	"reflect"
)

type Sharper interface {
	Area() float32
}

// Square //////////////////////
type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

// Rectangle /////////////////////////////
type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width

}

//////////////////////////

type valuable interface {
	getValue() float32
}

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

// 需要传入实现接口的结构体
func showValue(asset valuable) {
	fmt.Printf("value of asset is %f\n", asset.getValue())
}

// ---------------------
func main() {
	r := Rectangle{5, 3}

	q := &Square{5}

	//创建一个实现了 Sharper 接口的slice
	shapes := []Sharper{r, q}
	fmt.Println("looping through shapes for area")
	for n, _ := range shapes {
		fmt.Println("shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}

	//输出价格 car 和 stockPosition 分别实现了valuable 接口
	showValue(car{price: 34455.123})

	showValue(stockPosition{sharePrice: 45, count: 67})

	//判断数据类型
	var areaIntf Sharper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1
	// Is Square the type of areaIntf?
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := areaIntf.(Rectangle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Rectangle")
	}

	//类型关键字type 只能用在switch中
	var a interface{} = 7
	switch a.(type) {
	case string:
		fmt.Println(a)
	case int:
		fmt.Println(`是个整数`, 7)
	default:
		fmt.Println(a)

	}

	//使用if判断类型

	if k, Ok := a.(interface{}); Ok {
		fmt.Println(`a is ok`, k)
	}

	//还可以用反射判断类型
	fmt.Println(reflect.TypeOf(a))

	//判断数据类型
	classifier(213, -14.3, "BELGIUM", complex(1, 2), nil, false, []int{})

}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case float64:
			fmt.Printf("Param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)

		case []int:
			fmt.Printf("Param #%d is a []int\n", i)

		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}
