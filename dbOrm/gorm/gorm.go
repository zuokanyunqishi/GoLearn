package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BaseCreditInfo struct {
	Id         int    `gorm:"column:id"`
	BankName   string `gorm:"column:bankName"`
	BillDate   string `gorm:"column:billDate"`
	DueDate    string `gorm:"column:bankName"`
	CardNo     string `gorm:"column:dueDate"`
	Score      int    `gorm:"column:score"`
	CardType   int8   `gorm:"column:cardType"`
	CreateTime string `gorm:"column:createTime"`
	UpdateTIme string `gorm:"column:updateTIme"`

	attr []map[string]interface{}
}

func (BaseCreditInfo) TableName() string {
	return `baseCreditInfo`
}

func (this BaseCreditInfo) ToMap() map[string]interface{} {
	return map[string]interface{}{
		`id`:         this.Id,
		`bankName`:   this.BankName,
		`billDate`:   this.BillDate,
		`dueDate`:    this.DueDate,
		`cardNo`:     this.CardNo,
		`score`:      this.Score,
		`cardType`:   this.CardType,
		`createTime`: this.CreateTime,
		`updateTIme`: this.UpdateTIme,
	}

}

func (this BaseCreditInfo) GetAttr(attr []BaseCreditInfo) []map[string]interface{} {
	for _, attr := range attr {
		this.attr = append(this.attr, attr.ToMap())
	}

	return this.attr

}

func main() {

	db, err := gorm.Open(mysql.Open("root:pass@tcp(127.0.0.1:3306)/creditCard?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var model []BaseCreditInfo

	for i := 0; i < 500; i++ {
		db.Where("id >= ?", 2).Find(&model)

		fmt.Println(new(BaseCreditInfo).GetAttr(model), i)
	}

}
