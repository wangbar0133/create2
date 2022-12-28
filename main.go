package main

import (
	"crypto/ecdsa"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
	"github.com/panjf2000/ants/v2"
	"math/rand"
	"os"
	"strings"
	"sync"
)

const ResultFile = "result.txt"

func getAddress(devContractAddress common.Address, initCodeHash []byte, p string, s string) {
	for {
		i := rand.Uint64()
		salt := uint256.NewInt(i).Bytes32()
		address := crypto.CreateAddress2(devContractAddress, salt, initCodeHash).String()

		if strings.HasPrefix(address, p) && strings.HasSuffix(address, s) {
			fmt.Println(i)
			fmt.Println(address)
			fmt.Println()

			f, err := os.OpenFile(ResultFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			_, err = fmt.Fprintf(f, "i: %d ; address: %s\n", i, address)
			if err != nil {
				fmt.Println(err)
				return
			}
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
			f, err := os.OpenFile(ResultFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			_, err = fmt.Fprintf(f, "address: %s ; privateKey: %s\n", address, hexutil.Encode(privateKeyBytes))
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
}

var (
	address  = flag.String("a", "", "dev contract address")
	byteCode = flag.String("i", "", "init byte code")
	prefix   = flag.String("p", "", "prefix of address")
	suffix   = flag.String("s", "", "suffix of address")
	goNum    = flag.Int("n", 1000, "goroutine number")
)

func main() {
	flag.Parse()
	p := "0x" + *prefix
	s := *suffix
	fmt.Println("[+] Start!")

	if *address == "" {
		fmt.Printf("[+] Generate EOA address with prefix %s and suffix %s\n", p, s)
		fmt.Printf("[+] Generate goroutine Num: %d \n", *goNum)
		defer ants.Release()

		runTimes := *goNum

		// Use the common pool.
		var wg sync.WaitGroup
		syncCalculateSum := func() {
			getAddressEoa(p, s)
			wg.Done()
		}
		for i := 0; i < runTimes; i++ {
			wg.Add(1)
			_ = ants.Submit(syncCalculateSum)
		}
		wg.Wait()
		fmt.Printf("running goroutines: %d\n", ants.Running())
		fmt.Printf("finish all tasks.\n")
	} else {
		fmt.Printf("[+] Generate contract address with prefix %s and suffix %s\n", p, s)
		fmt.Printf("[+] Generate goroutine Num: %d \n", *goNum)
		devContractAddress := common.BytesToAddress(common.FromHex(*address))
		initByteCode := common.FromHex(*byteCode)
		initCodeHash := crypto.Keccak256(initByteCode)

		defer ants.Release()

		runTimes := *goNum

		// Use the common pool.
		var wg sync.WaitGroup
		syncCalculateSum := func() {
			getAddress(devContractAddress, initCodeHash, p, s)
			wg.Done()
		}
		for i := 0; i < runTimes; i++ {
			wg.Add(1)
			_ = ants.Submit(syncCalculateSum)
		}
		wg.Wait()
		fmt.Printf("running goroutines: %d\n", ants.Running())
		fmt.Printf("finish all tasks.\n")
	}
}
