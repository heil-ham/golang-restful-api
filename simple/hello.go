package simple

type SayHello interface {
	Hello() string
}

type SayHelloService struct {
	SayHello
}

type SayHelloImpl struct {

}

func (sayHelloImpl *SayHelloImpl) Hello() string {
	return "Helloooo"
}

func NewSayHelloImpl() *SayHelloImpl{
	return &SayHelloImpl{

	}
}

func NewSayHelloService(sayhello SayHello) *SayHelloService{
	return &SayHelloService{
		SayHello: sayhello,
	}
}