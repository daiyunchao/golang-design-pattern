package tmpl

import "github.com/marmotedu/iam/pkg/log"

// Cooker 定义接口
type Cooker interface {
	fire()
	cook()
	outFire()
}

// CookSomething 定义模板,定义方法的顺序
func CookSomething(cook Cooker) {
	cook.fire()
	cook.cook()
	cook.outFire()
}

// CookBase 类似父类,实现接口
type CookBase struct {
}

func (cookBase *CookBase) fire() {
	log.Infof("fire cookBase")
}

func (cookBase *CookBase) cook() {
	log.Infof("cook cookBase")
}

func (cookBase *CookBase) outFire() {
	log.Infof("out fire cookBase")
}

// XiHongShi 通过组合的方式,也相当于实现了接口
type XiHongShi struct {
	CookBase
}

// 只改动自己感兴趣的地方,至于在什么地方用,其实是不关心的
func (x *XiHongShi) cook() {
	log.Infof("cook xihong shi")
}

type ChaoJiDan struct {
	CookBase
}

func (chaoJiDan *ChaoJiDan) cook() {
	log.Infof("cook chaoJiDan")
}
