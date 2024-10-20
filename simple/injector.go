//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializedService(ErrBool bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

func InitializedDatabase() (*DatabaseRepository) {
	wire.Build(NewDatabaseMongo, NewDatabaseMysql,NewDatabaseRepository)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBar() (*FooBarService) {
	wire.Build(NewFooBarService, fooSet, barSet)
	return nil
}

var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializedHello() *SayHelloService{
	wire.Build(NewSayHelloService, helloSet)
	return nil
}

func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}