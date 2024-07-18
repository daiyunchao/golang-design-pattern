package simpleFactory

type SF struct {
	Name string
	Age  int
}

func NewSF(name string) *SF {
	return &SF{
		Name: name,
		Age:  10,
	}
}
