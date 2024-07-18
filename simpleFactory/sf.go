package simpleFactory

type SF struct {
	Name string
}

func NewSF(name string) *SF {
	return &SF{
		Name: name,
	}
}
