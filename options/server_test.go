package options

import (
	"github.com/marmotedu/iam/pkg/log"
	"testing"
)

func TestServer(t *testing.T) {
	server := NewServer(WithHost("127.0.0.1"), WithPort(9000), WithTLS(true))
	log.Infof("server: %v", server)
}
