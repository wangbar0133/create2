# CREATE2

A create2 salt generator that gets a specific address.

## compile

```bash
go bulid main.go
```

## usage

```bash
./main -a <address> -i <init byte code> -p <prefix> -s <suffix>
```

address: Factory contract Deployer's address
init byte code: contract's init code
prefix: prefix for generating address
suffix: suffix for generating address

## example

1. Changing the Test contract.
2. Deploy Factory contract Deployer, this code is in deployer.sol.
3. Get init byte code by Deployer contract.
4. Using ./main to get salt

```bash
% ./main -a 0x26b989b9525Bb775C8DEDf70FeE40C36B397CE67 -i 0x608060405234801561001057600080fd5b506101f6806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80638da5cb5b14610046578063901b94c014610064578063f96339301461006e575b600080fd5b61004e610078565b60405161005b91906101a5565b60405180910390f35b61006c61009c565b005b6100766100de565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461013657600080fd5b7f23ddb4dbb8577d03ebf1139a17a5c016963c43761e8ccd21eaa68e9b8ce6a68e60405160405180910390a1565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061018f82610164565b9050919050565b61019f81610184565b82525050565b60006020820190506101ba6000830184610196565b9291505056fea26469706673582212202ea2184ab1262a50cc7664bc90ba9589e49538a91682af89b0e9f2fc79f31e1564736f6c63430008110033 -p 123 -s abc
[+] Start!
5229343943791775475
0x123Bc50381F5F6a0D4931C614FE5b38978ed5abc
```

5. call deploy function in Deployer contract.
