package interfaces

type Error interface {
	error
	GetErrorCode() uint32
	GetTrace() string
	SetTrace(string)
}
