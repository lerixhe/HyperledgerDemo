/*
这里处理与区块链的交互操作：
1.update
2.query
输入：方法名称，参数组（包含key，value）
*/
package models

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)


const (
    channelID      = "default"
    orgName        = "bcs06071508"
    orgAdmin       = "Admin"
    ordererOrgName = "bcs06071508"
    ccID           = "obcs-example02"
)
/* 获取client对象  */
func getClient(configOpt core.ConfigProvider, sdkOpts ...fabsdk.Option)*channel.Client{
	  //Init the sdk config
	  sdk, err := fabsdk.New(configOpt, sdkOpts...)
	  if err != nil {
		  fmt.Println("Failed to create new SDK: %s", err)
	  }
	  defer sdk.Close()
	  // ************ setup complete ************** //
	      //prepare channel client context using client context
		  clientChannelContext := sdk.ChannelContext(channelID, fabsdk.WithUser("Admin"), fabsdk.WithOrg(orgName))
	  	    // Channel client is used to query and execute transactions (Org1 is default org)
			  client, err := channel.New(clientChannelContext)

			  if err != nil {
				  fmt.Println("Failed to create new channel client: %s", err)
			  }
			  return client
}


/*  操作智能合约的类型定义: */
type ChainCodeSpec struct {
	// 需要New创建client,而参数为sdk对象，也需要new
	client      channel.Client
	chaincodeId string
}

/* 查询 */
func (this *ChainCodeSpec) ChaincodeQuery(function string, chaincodeArgs [][]byte) (response []byte, err error) {
	res :=
	return nil, nil
}

/* 更新*/
func (this *ChainCodeSpec) ChaincodeUpdate(function string, chaincodeArgs [][]byte) (response []byte, err error) {
	return nil, nil
}
