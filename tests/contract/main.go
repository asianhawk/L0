package main

import (
	"bytes"
	"encoding/hex"
	//"encoding/json"
	//"fmt"

	"github.com/bocheninc/L0/components/crypto"
	//"github.com/bocheninc/L0/components/utils"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"

	"fmt"
	"os"

	"github.com/bocheninc/L0/components/utils"
	"github.com/bocheninc/L0/core/accounts"
	"github.com/bocheninc/L0/core/coordinate"
	"github.com/bocheninc/L0/core/types"
)

var (
	fromChain      = []byte{0}
	toChain        = []byte{0}
	issuePriKeyHex = "496c663b994c3f6a8e99373c3308ee43031d7ea5120baf044168c95c45fbcf83"
	sender         = "6ce1bb0858e71b50d603ebe4bec95b11d8833e6d"
)

func main() {
	issueTX()
	time.Sleep(40 * time.Second)
	DeploySmartContractTX()
	time.Sleep(40 * time.Second)
	ExecSmartContractTX()
	time.Sleep(40 * time.Second)

}

func sendTransaction(tx chan string) {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
		},
		Timeout: time.Second * 500,
	}
	URL := "http://" + "localhost" + ":" + "8881"
	for {
		select {
		case txHex := <-tx:
			//tx := new(types.Transaction)
			//tx.Deserialize(utils.HexToBytes(txHex))
			//fmt.Println(" hash: ", tx.Hash().String(), " nonce: ", tx.Nonce(), " amount: ", tx.Amount())
			req, _ := http.NewRequest("POST", URL, bytes.NewBufferString(
				`{"id":1,"method":"Transaction.Broadcast","params":["`+txHex+`"]}`,
			))
			req.Header.Set("Content-Type", "application/json")
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			io.Copy(ioutil.Discard, resp.Body)
			defer resp.Body.Close()
		}
	}
}

func issueTX() {
	issueKey, _ := crypto.HexToECDSA(issuePriKeyHex)
	nonce := 1
	txChan := make(chan string, 5)
	go sendTransaction(txChan)
	contractSpec := new(types.ContractSpec)
	f, _ := os.Open("/home/rocky/pro/Go_Pro/src/github.com/bocheninc/L0/tests/contract/I0vote.lua")
	buf, _ := ioutil.ReadAll(f)

	var a accounts.Address
	pubBytes := []byte(sender + string(buf))
	a.SetBytes(crypto.Keccak256(pubBytes[1:])[12:])

	contractSpec.ContractCode = []byte("")
	contractSpec.ContractAddr = a.Bytes()
	contractSpec.ContractParams = []string{"send", "8ce1bb0858e71b50d603ebe4bec95b11d8833e68", "100"}
	privateKey, _ := crypto.GenerateKey()
	accounts.PublicKeyToAddress(*privateKey.Public())
	tx := types.NewTransaction(
		coordinate.NewChainCoordinate(fromChain),
		coordinate.NewChainCoordinate(toChain),
		types.TypeIssue,
		uint32(nonce),
		accounts.HexToAddress(sender),
		accounts.NewAddress(contractSpec.ContractAddr),
		big.NewInt(10e11),
		big.NewInt(1),
		uint32(time.Now().Unix()),
	)
	tx.Payload = utils.Serialize(contractSpec)
	sig, _ := issueKey.Sign(tx.SignHash().Bytes())
	tx.WithSignature(sig)
	txChan <- hex.EncodeToString(tx.Serialize())
}

func DeploySmartContractTX() {
	issueKey, _ := crypto.HexToECDSA(issuePriKeyHex)
	nonce := 2
	txChan := make(chan string, 5)
	go sendTransaction(txChan)
	contractSpec := new(types.ContractSpec)
	f, _ := os.Open("/home/rocky/pro/Go_Pro/src/github.com/bocheninc/L0/tests/contract/I0vote.lua")
	buf, _ := ioutil.ReadAll(f)
	var a accounts.Address
	pubBytes := []byte(sender + string(buf))
	a.SetBytes(crypto.Keccak256(pubBytes[1:])[12:])

	contractSpec.ContractCode = buf
	contractSpec.ContractAddr = a.Bytes()
	contractSpec.ContractParams = []string{"init", "100"}
	privateKey, _ := crypto.GenerateKey()
	accounts.PublicKeyToAddress(*privateKey.Public())
	fmt.Println(len(string(contractSpec.ContractAddr)))
	tx := types.NewTransaction(
		coordinate.NewChainCoordinate(fromChain),
		coordinate.NewChainCoordinate(toChain),
		types.TypeContractInit,
		uint32(nonce),
		accounts.HexToAddress(sender),
		accounts.Address{}, //HexToAddress("0x00000000000000000000"),
		big.NewInt(10000),
		big.NewInt(0),
		uint32(time.Now().Unix()),
	)
	tx.Payload = utils.Serialize(contractSpec)
	sig, _ := issueKey.Sign(tx.SignHash().Bytes())
	tx.WithSignature(sig)
	txChan <- hex.EncodeToString(tx.Serialize())
}

func ExecSmartContractTX() {
	issueKey, _ := crypto.HexToECDSA(issuePriKeyHex)
	nonce := 3
	txChan := make(chan string, 5)
	go sendTransaction(txChan)
	contractSpec := new(types.ContractSpec)
	f, _ := os.Open("/home/rocky/pro/Go_Pro/src/github.com/bocheninc/L0/tests/contract/I0vote.lua")
	buf, _ := ioutil.ReadAll(f)

	var a accounts.Address
	pubBytes := []byte(sender + string(buf))
	a.SetBytes(crypto.Keccak256(pubBytes[1:])[12:])

	contractSpec.ContractCode = []byte("")
	contractSpec.ContractAddr = a.Bytes()
	contractSpec.ContractParams = []string{"transfer", "8ce1bb0858e71b50d603ebe4bec95b11d8833e68", "100"}
	privateKey, _ := crypto.GenerateKey()
	accounts.PublicKeyToAddress(*privateKey.Public())
	tx := types.NewTransaction(
		coordinate.NewChainCoordinate(fromChain),
		coordinate.NewChainCoordinate(toChain),
		types.TypeContractInvoke,
		uint32(nonce),
		accounts.HexToAddress(sender),
		accounts.NewAddress(contractSpec.ContractAddr),
		big.NewInt(1000),
		big.NewInt(0),
		uint32(time.Now().Unix()),
	)
	tx.Payload = utils.Serialize(contractSpec)
	sig, _ := issueKey.Sign(tx.SignHash().Bytes())
	tx.WithSignature(sig)
	txChan <- hex.EncodeToString(tx.Serialize())
}
