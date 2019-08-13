package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
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

func (test *Test) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// 这里是交互入口，在这里调用链码中的方法进行增删改查
	// 获取要调用的的方法和所需参数
	function, parameters := stub.GetFunctionAndParameters()
	// 约定根据方法名来调用方法，第一个参数为key，第二个参数为value（若存在）
	if function == "get" {
		return test.get(stub, parameters[0])
	} else if function == "set" {
		return test.set(stub, parameters[0], []byte(parameters[1]))
	}
	// 其他所有情况
	return shim.Error("Invalid Smart Contract function  name")
}

/*
设计读写方法
set get
*/
//  读数据
func (test *Test) get(stub shim.ChaincodeStubInterface, key string) peer.Response {
	// 根据传入的key读出其值
	// 未读取到，返回空和错误

	// 处理异常
	data, err := stub.GetState(key)
	if err != nil {
		return shim.Error(err.Error())
	}
	// 处理data
	if len(data) == 0 {
		return shim.Error("Data not Available!")
	}
	return shim.Success(data)
}

//  写数据
func (test *Test) set(stub shim.ChaincodeStubInterface, key string, value []byte) peer.Response {
	// 将传入的值写入区块链
	err := stub.PutState(key, value)
	// 处理异常
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}
