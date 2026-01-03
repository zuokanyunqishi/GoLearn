package main

import (
	p "GoLearn/proto/pb"
	"fmt"
	"google.golang.org/protobuf/proto"
)

func main() {

	person := &p.Person{
		Base: &p.PersonBase{
			Name:   "小明",
			Age:    32,
			Height: 175,
			Sex:    p.Sex_man,
			Address: &p.Address{
				Province: "北京",
				City:     "北京",
				Region:   "海淀",
			},
			Email: "111@qq.com",
		},
		Friends: nil,
		Father:  nil,
		School: &p.Person_School{
			Name:    "北京第八中学",
			Address: "海淀区",
		},
		Company:   nil,
		IsMarried: false,
		Skill: map[string]*p.Skill{
			"学习": {
				Desc:  "学习诗人快乐",
				Grade: p.Skill_high,
			},
		},
	}

	bytes, _ := proto.Marshal(person)

	var pclone p.Person

	proto.Unmarshal(bytes, &pclone)

	fmt.Println(pclone.Base, pclone.GetSkill())

}
