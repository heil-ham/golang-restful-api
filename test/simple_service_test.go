package test

import (
	"fmt"
	"golang-restful-api/simple"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService, errService := simple.InitializedService(false)
	fmt.Println(simpleService)
	fmt.Println(errService)
}

func TestHelloService(t *testing.T) {
	helloService := simple.InitializedHello()
	fmt.Println(helloService.Hello())
}

func TestConn(t *testing.T) {
	conn, closure := simple.InitializedConnection("file connectt")
	fmt.Println(conn.File.Name)
	closure()
	
}