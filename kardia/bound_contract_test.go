// Package kardia
package kardia

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/kardiachain/go-kardia/lib/abi"
	"github.com/kardiachain/go-kardia/lib/common"
	"github.com/kardiachain/go-kardia/lib/crypto"
	"github.com/stretchr/testify/assert"
)

func TestBoundContract_Deploy(t *testing.T) {
	node, err := SetupNodeClient()
	assert.Nil(t, err)
	r := strings.NewReader(wheelABI)
	abiData, err := abi.JSON(r)
	assert.Nil(t, err)
	pubKey, privateKey, err := SetupTestAccount()
	assert.Nil(t, err)
	fromAddress := crypto.PubkeyToAddress(*pubKey)
	nonce, err := node.NonceAt(context.Background(), fromAddress.String())
	assert.Nil(t, err)
	balance, err := node.Balance(context.Background(), fromAddress.String())
	assert.Nil(t, err)
	fmt.Println("Balance", balance)
	gasLimit := uint64(30000000)
	gasPrice := big.NewInt(1)
	auth := NewKeyedTransactor(privateKey)
	auth.Nonce = nonce
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = gasLimit   // in units
	auth.GasPrice = gasPrice

	smc := NewBoundContract(node, &abiData, common.HexToAddress(WheelSMCAddr))
	smcAddress, txHash, err := smc.DeployKRC20(auth)
	assert.Nil(t, err)
	fmt.Println("SMC Addr", smcAddress.String())
	fmt.Println("TxHash", txHash.String())
}