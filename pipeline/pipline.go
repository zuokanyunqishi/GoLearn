package pipeline

import (
	"context"
	"fmt"
)

type HandleFunc func(c context.Context)

type pipeFun func(next HandleFunc) HandleFunc
type Pipeline struct {
	pipes []pipeFun
}

func New(handles ...pipeFun) *Pipeline {
	return &Pipeline{
		pipes: handles,
	}
}

func (p *Pipeline) Run(data any) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "data", data)

	stack := HandleFunc(func(c context.Context) {
		fmt.Println("my is mid :", c.Value("data"))
	})
	for i := len(p.pipes) - 1; i >= 0; i-- {
		stack = p.pipes[i](stack)
	}

	stack(ctx)
}
