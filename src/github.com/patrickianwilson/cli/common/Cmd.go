package common

type Cmd interface {
	GetCmdName() string

	GetCmdGroup() string

	Init(args []string, context *Context) error

	Validate() error

	PrintHelp()

	DoAction() error
}
