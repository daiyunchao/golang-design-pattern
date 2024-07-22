package tmpl

import "testing"

func TestTmpl(t *testing.T) {
	var cook Cooker
	cook = &XiHongShi{}
	CookSomething(cook)

	cook = &ChaoJiDan{}
	CookSomething(cook)
}
