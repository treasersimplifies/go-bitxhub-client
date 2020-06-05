package rpcx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/meshplus/bitxhub-kit/crypto/asym"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAll(t *testing.T) {
	// Initialize client:
	privKey, err := asym.GenerateKey(asym.ECDSASecp256r1)
	require.Nil(t, err)
	var cfg = &config{
		addrs: []string{
			"localhost:60011",
			"localhost:60012",
			"localhost:60013",
			"localhost:60014",
		},
		logger:     logrus.New(),
		privateKey: privKey,
	}
	cli, err := New(
		WithAddrs(cfg.addrs),
		WithLogger(cfg.logger),
		WithPrivateKey(cfg.privateKey),
	)
	require.Nil(t, err)
	pubKey, err := privKey.PublicKey().Address()
	require.Nil(t, err)
	fmt.Println("pubKey:", pubKey)
	fmt.Println("cli:", cli)
	// Contract:
	// testContract(t, cli)
	/*
		Problems:
		Not Implemented: Freeze/Unfreeze/Update/Removed
	*/

	// Register AppChain:
	// testAppChain(t, cli)
	/*
		Problems:
		The operations was executed on blockchain in chaincode which was not readable,
		so we can't know meanings of operations and params clearly.
	*/

	// Blockchain:
	// testBlockchain(t, cli)

	// Others:
	testOthers(t, cli, "0x8266ba37c72f58e885234341202f62f2d8f94f41")
}

func testOthers(t *testing.T, cli *ChainClient, pubKey string) {
	meta, err := cli.GetNetworkMeta()
	require.Nil(t, err)
	fmt.Println("[TestAll] network meta:", meta)
	balance, err := cli.GetAccountBalance(pubKey)
	require.Nil(t, err)
	fmt.Println("[TestAll] ", pubKey, " balance:", balance)
}

func testBlockchain(t *testing.T, cli *ChainClient) {
	status, err := cli.GetChainStatus()
	require.Nil(t, err)
	fmt.Println("[TestAll] status:", status)
	meta, err := cli.GetChainMeta()
	require.Nil(t, err)
	fmt.Println("[TestAll] chain meta:", meta)
	block, err := cli.GetBlock("50", pb.GetBlockRequest_HEIGHT)
	require.Nil(t, err)
	require.Equal(t, block.BlockHeader.Number, uint64(50))
	fmt.Println("[TestAll] GetBlock:", block)
	blocks, err := cli.GetBlocks(2, 3)
	require.Nil(t, err)
	fmt.Println("[TestAll] GetBlocks:", len(blocks.Blocks))
}

func testAppChain(t *testing.T, cli *ChainClient) {
	args := []*pb.Arg{
		String(""),           //validators
		Int32(0),             //consensus_type
		String("hyperchain"), //chain_type
		String("税务链"),        //name
		String("趣链税务链"),      //desc
		String("1.8"),        //version
	}
	res, err := cli.InvokeBVMContract(InterchainContractAddr, "Register", args...)
	require.Nil(t, err)
	fmt.Println("[TestAll] AppChain Register:", string(res.Ret))
	appChain := &Appchain{}
	err = json.Unmarshal(res.Ret, appChain)
	require.Nil(t, err)
	fmt.Print("[TestAll] appChain id:", appChain.ID)
	appchainID := appChain.ID
	args = []*pb.Arg{
		String(appchainID),
		Int32(1),   //审核通过
		String(""), //desc
	}
	res, err = cli.InvokeBVMContract(InterchainContractAddr, "Aduit", args...)
	require.Nil(t, err)
	fmt.Println("[TestAll] AppChain Aduit:", string(res.Ret))
	res, err = cli.InvokeBVMContract(InterchainContractAddr, "DeleteAppchain", String(appchainID))
	require.Nil(t, err)
	fmt.Println("[TestAll] AppChain DeleteAppchain:", string(res.Ret))
}

func testContract(t *testing.T, cli *ChainClient) {
	//获取wasm合约字节数组
	contract, _ := ioutil.ReadFile("./testdata/example.wasm")
	//部署合约，获取合约地址
	addr, _ := cli.DeployContract(contract)
	fmt.Println("[TestAll] addr:", addr)
	//调用合约，获取交易回执
	result, _ := cli.InvokeXVMContract(addr, "a", Int32(1), Int32(2))
	//判断合约调用交易成功与否，打印合约调用数据
	assert.True(t, CheckReceipt(result))
	fmt.Println("[TestAll] XVM result.Ret:", string(result.Ret))
	result, err := cli.InvokeBVMContract(StoreContractAddr, "Set", String("a"), String("10"))
	require.Nil(t, err)
	res, err := cli.InvokeBVMContract(StoreContractAddr, "Get", String("a"))
	require.Nil(t, err)
	fmt.Println("[TestAll] BVM result.Ret:", string(res.Ret))
	assert.Equal(t, string(res.Ret), "10")
}
