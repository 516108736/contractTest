package main

import (
	"fmt"
	cc "github.com/contractTest/gcode/contractGo"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main()  {
	fmt.Println("???")
	client, err := ethclient.Dial("wss://kovan.infura.io/ws/v3/5f85acad140a4286858886f080177bc9")
	if err != nil {
		panic(err)
	}
	address:=common.HexToAddress("0xb155628574d067cc89fab428a430eaa9621abda65118f468ed7254d156c0418f")
	instance,err:=cc.NewScf(address,client)
	if err!=nil{
		panic(err)
	}
	data,err:=instance.GetETHUSDTLatestPrice(nil)
	fmt.Println("GetETHUSDTLatestPrice errr",data,err)
	data,err=instance.GetBTCUSDTLatestPrice(nil)
	fmt.Println("GetBTCUSDTLatestPrice err",data,err)
}
