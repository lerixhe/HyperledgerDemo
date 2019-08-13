package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	_ "github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// chaincode类型定义
type Test struct {
	// 简单的读写实现
}

func main() {
	// 主函数中启动chaincode
	shim.Start(new(Test))
}

// Test需要实现Chaincode接口,必须实现这两个方法：init invoke
func (test *Test) Init(shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (test *Test) Invoke(shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}
