package automation

type Command struct {
	ModuleName string
	MethodName string

	Parameters  string
	Description string

	Action interface{}
}

func (c Command) FullMethodName() string {
	if c.ModuleName != "" {
		return c.ModuleName + "." + c.MethodName
	} else {
		return c.MethodName
	}
}
