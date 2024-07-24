# golang设计模式整理

### 单例模式

##### 定义

​	全局只有一个实例,各种地方用的都是这一个实例,是一种节省内存和系统开销的方法

##### 场景

1. 配置管理, 当应用程序需要一个集中管理配置的类，且这些配置在整个应用程序生命周期内保持唯一并且可全局访问时
2. 各种连接池, 如果是需要做成连接池的,说明资源是有限的,需要进行管理,如果做成单例来全局管理这些连接就比较合理
3. 资源管理, 比如全局缓存,数据库资源入口等

##### golang如何实现

​	golang 可以使用`sync.Once`来实现,`sync.Once`的作用就是确保某段代码只执行一次的机制,而且也是一种线程安全的方法来执行一次性的初始化任务

实现代码:

[golang 单例模式,代码库链接地址](https://github.com/daiyunchao/golang-design-pattern/tree/main/singleton)

核心代码如下:

```go
func GetConfigManager() *ConfigManager {
	once.Do(func() {
		count := 100
		configManager = &ConfigManager{
			ConfigCount: &(count),
		}
	})
	return configManager
}
```

### 简单工厂

##### 定义

​	不直接通过构造函数来创建实例的一种方式,作用在于传递必要的参数

##### 好处

1. 调用者只需要知道必要参数
2. 调用者不在乎你需要的这个实例是否是单例的
3. 如果都是通过工厂创建函数,那么对象的扩展就更容易了

##### golang如何实现

​	简单工厂还是很简单,可以直接看代码

​	[golang 简单工厂模式,代码链接地址](https://github.com/daiyunchao/golang-design-pattern/tree/main/simpleFactory)

核心代码如下:

```go
type SF struct {
	Name string
	Age  int
}
//通过工厂去实例化结构体,通过这个函数让调用者知道应该关心哪些参数
func NewSF(name string) *SF {
	return &SF{
		Name: name,
		Age:  10,
	}
}

```



### 抽象工厂

##### 定义

​	通过工厂创建出来的实例是一个接口的具体实现

##### 好处

​	通过返回接口，可以**在你不公开内部实现的情况下，让调用者使用你提供的各种功能**,这个在做一些抽象的时候特别好用,比如我通过 if/else 或是 swatch/case 的方式创建了很多实例,但这些实例都有一个共同点,都实现了一个接口,那么我们就实现了对具体实现的抽象

##### golang如何实现

​	[golang 抽象工厂模式, 代码链接地址](https://github.com/daiyunchao/golang-design-pattern/tree/main/abstractFactory)

```go
type Opener interface {
	OpenFile() error
}

type Closer interface {
	CloseFile() error
}

type OpenAndCloser interface {
	Opener
	Closer
}
//核心点在于: 创建的并不是实例,而是接口
func NewFileA() OpenAndCloser {
	return &FileA{}
}

func NewFileB() OpenAndCloser {
	return &FileB{}
}
```



### 策略模式

##### 定义

​	定义一组算法，将每个算法都封装起来，并且使它们之间可以互换

##### 好处

​	少了很多if/else 的硬编码

##### golang 如何实现

​	为了每一个策略都有相同的行为,所以在需要定义一个接口,并且每一个策略都需要实现这个接口

一共可以分为 一个接口,N 个具体的策略, 一个操作者

[golang 策略模式](https://github.com/daiyunchao/golang-design-pattern/tree/main/strategy)

核心代码:

```go
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
```



### 模板模式

##### 定义

​	模板模式 (Template Pattern)定义一个操作中算法的骨架，而将一些步骤延迟到子类中。这种方法让子类在不改变一个算法结构的情况下，就能重新定义该算法的某些特定步骤。简单来说，模板模式就是将一个类中能够公共使用的方法放置在抽象类中实现，将不能公共使用的方法作为抽象方法，强制子类去实现，这样就做到了将一个类作为一个模板，让开发者去填充需要填充的地方

##### golang如何实现

​	因为在 golang 中是没有父子关系的,但是我们可以通过包含的关系来实现类似父子的关系

```go
type Fa struct {}
function (f *Fa)doSomething()

//通过这种包含的关系,让 Son 也拥有了 Fa的全部属性和方法
type Son struct {
  Fa
}

```

[golang 模板模式](https://github.com/daiyunchao/golang-design-pattern/tree/main/tmpl)

### 代理模式

##### 定义

​	简单理解就是,A 君去买房,但是流程冗长而且麻烦,A 君不是很了解里面的门门道道,于是就找了专业机构协助买房的工作人员 B 姐,B 姐带着 A 君去走这些流程,B 姐去准备资料,A 君只要关键性给钱和签字就行了, 故事里的 B 姐就是代理



##### 使用场景

访问控制

附加功能,其实访问控制也算是附加功能的一种,只是附加的是权限判断的逻辑而已

##### golang如何实现

[golang 代理模式,权限控制](https://github.com/daiyunchao/golang-design-pattern/tree/main/proxy)

核心代码

```go
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


//测试用例:
func TestAccess(t *testing.T) {
	access := Proxy{
		PersonA{Name: "B"},
	}
	access.Visit()
	access = Proxy{
		PersonA{Name: "A"},
	}
	access.Visit()
}

//输出结果:
INFO	proxy/access.go:28	Proxy visit
INFO	proxy/access.go:33	B 无权限访问
INFO	proxy/access.go:28	Proxy visit
INFO	proxy/access.go:16	A visit
```



### 选项模式

##### 定义

​	用于构建具有许多可选参数的结构体或函数。这个模式用一种简介而灵活的方式来处理各种选项，使代码更易于扩展和维护

##### 使用场景

​	在构建 Struct 时有这样的需求 传了就用传入的参数当值,如果没传就用一个默认值,但是 golang 中并没有这种设定,所以一般有两种解决方法

方法 1: 我们要分别开发两个用来创建实例的函数，一个可以创建带默认值的实例，一个可以定制化创建实例

```go
package options

import (
	"time"
)

const (
	defaultTimeout = 10
	defaultCaching = false
)

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

// NewConnect creates a connection.
func NewConnect(addr string) (*Connection, error) {
	return &Connection{
		addr:    addr,
		cache:   defaultCaching,//设置默认值
		timeout: defaultTimeout,//设置默认值
	}, nil
}

// 使用参数的形式传递
func NewConnectWithOptions(addr string, cache bool, timeout time.Duration) (*Connection, error) {
	return &Connection{
		addr:    addr,
		cache:   cache,
		timeout: timeout,
	}, nil
}
```

调用者通过自己的需求,要么用默认值,要么就用传递的参数,但这种太极端了,而且不优雅,就像上面的例子一样,如果现在需要`cache`默认值,但`timeout`非默认值,那是不是还需要创建一个新的函数叫 `NewConnectWithCacheOptions`呢

还有另外一种方式:

我们需要创建一个带默认值的选项，并用该选项创建实例

```go
package options

import (
	"time"
)

const (
	defaultTimeout = 10
	defaultCaching = false
)

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

type ConnectionOptions struct {
	Caching bool
	Timeout time.Duration
}

func NewDefaultOptions() *ConnectionOptions {
	return &ConnectionOptions{
		Caching: defaultCaching,
		Timeout: defaultTimeout,
	}
}

// NewConnect creates a connection with options.
func NewConnect(addr string, opts *ConnectionOptions) (*Connection, error) {
	return &Connection{
		addr:    addr,
		cache:   opts.Caching,
		timeout: opts.Timeout,
	}, nil
}
```

这种方式本质上和第一种没什么区别,只是将需要传递的参数封装到了 Struct 里,并且提供一个设置默认值的方法,还是不够灵活

于是就有了`选项模式`



##### golang如何实现

```go
// Server 结构体
type Server struct {
	Host string
	Port int
	TLS  bool
}
type ServerOption func(server *Server)

func NewServer(opts ...ServerOption) *Server {
  //先设置默认值
	server := &Server{
		Host: "",
		Port: 8888,  //设置默认中
		TLS:  false, //设置默认值
	}
  //通过传入参数,再设置具体的值
	for _, opt := range opts {
		opt(server)
	}
	return server
}

//设置 Host 参数
func WithHost(host string) ServerOption {
	return func(server *Server) {
		server.Host = host
	}
}

//设置 Port 参数
func WithPort(port int) ServerOption {
	return func(server *Server) {
		server.Port = port
	}
}

//设置 TLS 参数
func WithTLS(tls bool) ServerOption {
	return func(server *Server) {
		server.TLS = tls
	}
}
```

[golang 实现选项模式](https://github.com/daiyunchao/golang-design-pattern/tree/main/options)

