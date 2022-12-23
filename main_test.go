package main

import (
	"crypto/ecdsa"
	"fmt"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
)

func TestCreate2(t *testing.T) {
	devContractAddress := common.BytesToAddress(common.FromHex("0x26b989b9525Bb775C8DEDf70FeE40C36B397CE67"))
	initByteCode := common.FromHex("0x608060405234801561001057600080fd5b506101f6806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80638da5cb5b14610046578063901b94c014610064578063f96339301461006e575b600080fd5b61004e610078565b60405161005b91906101a5565b60405180910390f35b61006c61009c565b005b6100766100de565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461013657600080fd5b7f23ddb4dbb8577d03ebf1139a17a5c016963c43761e8ccd21eaa68e9b8ce6a68e60405160405180910390a1565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061018f82610164565b9050919050565b61019f81610184565b82525050565b60006020820190506101ba6000830184610196565b9291505056fea26469706673582212202ea2184ab1262a50cc7664bc90ba9589e49538a91682af89b0e9f2fc79f31e1564736f6c63430008110033")
	initCodeHash := crypto.Keccak256(initByteCode)

	i := rand.Uint64()
	salt := uint256.NewInt(i).Bytes32()
	address := crypto.CreateAddress2(devContractAddress, salt, initCodeHash).String()
	fmt.Println(address)
}

func TestEoa(t *testing.T) {
	privateKey, _ := crypto.GenerateKey()
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).String()
	fmt.Println(address)
}
