package interfaces

type Router interface {

	// PreHandle 处理业务之前
	PreHandle(request Request)

	// Handle 处理中
	Handle(request Request)
	// PostHandle 处理之后
	PostHandle(request Request)
}
