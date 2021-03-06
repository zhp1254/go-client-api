package main

import (
	"fmt"
	"github.com/tronprotocol/go-client-api/common/hexutil"
	"github.com/tronprotocol/go-client-api/service"
)

const address = "47.91.216.69:50051"

func main() {
	client := service.NewGrpcClient(address)
	client.Start()
	defer client.Conn.Close()

	witnesses := client.ListWitnesses()

	for i, v := range witnesses.Witnesses {
		addr := hexutil.Encode(v.Address)
		u := v.Url
		totalProduced := v.TotalProduced
		totalMissed := v.TotalMissed
		latestBlockNum := v.LatestBlockNum
		latestSlotNum := v.LatestSlotNum
		isJobs := v.IsJobs
		fmt.Printf("index: %d, witness: address: %s, url: %s, "+
			"total produced: %d, total missed: %d, latest block num: %d, "+
			"latest slot num: %d, is jobs: %v\n", i,
			addr, u,
			totalProduced, totalMissed, latestBlockNum, latestSlotNum, isJobs)
	}

	nodes := client.ListNodes()

	for i, v := range nodes.Nodes {
		host := string(v.Address.Host)
		port := v.Address.Port
		fmt.Printf("index: %d, node: host: %v, port: %d\n", i, host, port)
	}

	account := client.GetAccount("a00a9309758508413039e4bc5a3d113f3ecc55031d")

	fmt.Printf("account: type: %s, address: %s, balance: %d\n", account.Type,
		hexutil.Encode(account.Address), account.Balance)
}
