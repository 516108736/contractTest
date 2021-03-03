package main

import (
	cc "./contractGo"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main()  {
	fmt.Println("???")
	client, err := ethclient.Dial("wss://kovan.infura.io/ws/v3/5f85acad140a4286858886f080177bc9")
	if err != nil {
		panic(err)
	}
	address:=common.HexToAddress("0xb043097ba18f6b17022F994A3dC2DB88054b00b1")
	instance,err:=cc.NewScf(address,client)
	if err!=nil{
		panic(err)
	}
	data,err:=instance.GetLatestPrice(nil)
	fmt.Println("errr",data,err)
}
