package options

// Server 结构体
type Server struct {
	Host string
	Port int
	TLS  bool
}
type ServerOption func(server *Server)

func NewServer(opts ...ServerOption) *Server {
	server := &Server{
		Host: "",
		Port: 8888,  //设置默认中
		TLS:  false, //设置默认值
	}
	for _, opt := range opts {
		opt(server)
	}
	return server
}

func WithHost(host string) ServerOption {
	return func(server *Server) {
		server.Host = host
	}
}

func WithPort(port int) ServerOption {
	return func(server *Server) {
		server.Port = port
	}
}

func WithTLS(tls bool) ServerOption {
	return func(server *Server) {
		server.TLS = tls
	}
}
