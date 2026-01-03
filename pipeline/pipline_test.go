package pipeline

import (
	context2 "context"
	"fmt"
	"golang.org/x/net/context"
	"testing"
)

func TestPipeline_Run(t *testing.T) {
	type fields struct {
		pipes []pipeFun
	}
	type args struct {
		data any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.

		{
			name: "test1",
			fields: fields{
				pipes: []pipeFun{
					func(next HandleFunc) HandleFunc {
						return func(c context.Context) {
							fmt.Println("f1-before")
							next(c)
							fmt.Println("f1-after")
						}
					},
					func(next HandleFunc) HandleFunc {
						return func(c context.Context) {
							fmt.Println("f2-before")
							next(c)
							fmt.Println("f2-after")
						}
					},
					func(next HandleFunc) HandleFunc {
						return func(c context.Context) {
							fmt.Println("f3-before")
							next(c)
							fmt.Println("f3-after")
						}
					},
				},
			},
			args: args{
				data: "pipeline",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pipeline{
				pipes: tt.fields.pipes,
			}
			p.Run(tt.args.data)
		})
	}
}

func TestNew(t *testing.T) {
	New(func(next HandleFunc) HandleFunc {
		return func(c context2.Context) {
			fmt.Println("1:before")
			next(c)
			fmt.Println("1:after")
		}
	}, func(next HandleFunc) HandleFunc {
		return func(c context2.Context) {
			fmt.Println("2:before")
			next(c)
			fmt.Println("2:after")

		}
	}, func(next HandleFunc) HandleFunc {
		return func(c context2.Context) {
			fmt.Println("3:before")
			next(c)
			fmt.Println("3:after")

		}
	}).Run("hello,testNew")
}
