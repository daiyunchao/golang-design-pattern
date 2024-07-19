package strategy

// Doer 为了保证,不同策略的行为统一,所以这样定义一个接口,不同的策略都实现这个接口
type Doer interface {
	DoSomething(int, int) int
}

// Sa 策略 A
type Sa struct {
}

func NewSa() *Sa {
	return &Sa{}
}

func (s *Sa) DoSomething(a int, b int) int {
	return a + b
}

type Sb struct {
}

func NewSb() *Sb {
	return &Sb{}
}
func (s *Sb) DoSomething(a int, b int) int {
	return a - b
}

type Op struct {
	doer Doer
}

func NewOp() *Op {
	return &Op{}
}
func (o *Op) SetDoer(doer Doer) {
	o.doer = doer
}

func (o *Op) Do(a, b int) int {
	return o.doer.DoSomething(a, b)
}
