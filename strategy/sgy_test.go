package strategy

import (
	"github.com/marmotedu/iam/pkg/log"
	"testing"
)

func TestOp_Do(t *testing.T) {
	//策略A
	sa := NewSa()
	op := NewOp()
	op.SetDoer(sa)
	resA := op.Do(5, 3)
	if resA == 8 {
		log.Infof("SA result: %v", resA)
	} else {
		log.Errorf("SA result: %v", resA)
	}

	//策略 B
	sb := NewSb()
	op.SetDoer(sb)
	resB := op.Do(5, 3)
	if resB == 2 {
		log.Infof("SB result: %v", resB)
	} else {
		log.Errorf("SB result: %v", resB)
	}
}
