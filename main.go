package main

import (
	"crypto/ecdsa"
	"flag"
	"fmt"
	"math/rand"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
)

func getAddress(devContractAddress common.Address, initCodeHash []byte, p string, s string) {
	for {
		i := rand.Uint64()
		salt := uint256.NewInt(i).Bytes32()
		address := crypto.CreateAddress2(devContractAddress, salt, initCodeHash).String()

		if strings.HasPrefix(address, p) && strings.HasSuffix(address, s) {
			fmt.Println(i)
			fmt.Println(address)
			fmt.Println()
		}
	}
}

func getAddressEoa(p string, s string) {
	for {
		privateKey, _ := crypto.GenerateKey()
		publicKey := privateKey.Public()
		publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
		address := crypto.PubkeyToAddress(*publicKeyECDSA).String()
		if strings.HasPrefix(address, p) && strings.HasSuffix(address, s) {
			fmt.Println(address)
			privateKeyBytes := crypto.FromECDSA(privateKey)
			fmt.Println(hexutil.Encode(privateKeyBytes))
			fmt.Println()
	}
	}
}

var (
	address = flag.String("a", "", "dev contract address")
	byteCode = flag.String("i", "", "init byte code")
	prefix = flag.String("p", "", "prefix of address")
	suffix = flag.String("s", "", "suffix of address")
)

func main() {
	flag.Parse()
	p := "0x" + *prefix
	s := *suffix
	fmt.Println("[+] Start!")

	if *address == "" {
		go getAddressEoa(p, s)
		go getAddressEoa(p, s)
		go getAddressEoa(p, s)
		go getAddressEoa(p, s)
		getAddressEoa(p, s)
	} else {
		devContractAddress := common.BytesToAddress(common.FromHex(*address))
		initByteCode := common.FromHex(*byteCode)
		initCodeHash := crypto.Keccak256(initByteCode)
		go getAddress(devContractAddress, initCodeHash, p ,s)
		go getAddress(devContractAddress, initCodeHash, p ,s)
		go getAddress(devContractAddress, initCodeHash, p ,s)
		go getAddress(devContractAddress, initCodeHash, p ,s)
		getAddress(devContractAddress, initCodeHash, p ,s)
	}
}
