package gomake

const (
	FILE_TYPE = iota
)

type Target interface {
	Name() string
	Type() int
}

type Command func() error

type Rule interface {
	Target() Target
	Dependents() []Target
	Commands() []Command
}

func Build(target Target) error {
	return nil
}
