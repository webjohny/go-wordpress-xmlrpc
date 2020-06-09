package wordpress

type BaseCall interface {
	GetMethod() string
	GetArgs(user string, pwd string) interface{}
}
