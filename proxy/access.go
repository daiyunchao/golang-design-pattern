package proxy

import "github.com/marmotedu/iam/pkg/log"

// Accesses 定义了一个访问的接口
type Accesses interface {
	Visit()
}

// PersonA 实际操作人
type PersonA struct {
	Name string
}

func (p *PersonA) Visit() {
	log.Infof("%s visit", p.Name)
}

// Proxy 代理操作人
type Proxy struct {
	PersonA
}

func (p *Proxy) Visit() {
	if p.PersonA.Name == "" {
		p.PersonA = PersonA{Name: "Test"}
	}
	log.Infof("Proxy visit")
	//附加了权限的验证功能
	if p.PersonA.Name == "A" {
		p.PersonA.Visit()
	} else {
		log.Infof("%s 无权限访问", p.Name)
	}
}
