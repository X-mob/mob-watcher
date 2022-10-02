package task

import (
	"context"
	"fmt"

	"github.com/X-mob/mob-watcher/lib"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

func WaitTx(tx *types.Transaction, txErr error) {
	if txErr != nil && tx == nil {
		fmt.Printf("tx failed to send, %s\n", txErr.Error())
		return
	}

	txData := TxData{
		To:    *tx.To(),
		Data:  tx.Data(),
		Value: tx.Value(),
	}

	isExits := GlobalTxManager.Check(txData.Hash().Hex())
	if isExits == NotFound {
		GlobalTxManager.Add(txData)
	}

	// check
	if txErr != nil {
		fmt.Printf("tx failed: %s\n", txErr.Error())
		// update failed status
		GlobalTxManager.Update(txData.Hash().Hex(), Failed)
		if tx != nil {
			failedTx := GlobalTxManager.GetTx(txData.Hash().Hex())
			failedTx.Hash = tx.Hash().Hex()
			GlobalTxManager.UpdateTx(failedTx)
		}
		return
	} else {
		// update sent status
		sentTx := GlobalTxManager.GetTx(txData.Hash().Hex())
		if tx != nil {
			sentTx.Hash = tx.Hash().Hex()
		}
		sentTx.Status = Sent
		GlobalTxManager.UpdateTx(sentTx)
	}

	fmt.Printf("tx Hash: %s, wait to be mined..\n", tx.Hash())
	receipt, err := bind.WaitMined(context.Background(), lib.EthClient, tx)
	if err != nil {
		fmt.Printf("get tx receipt failed: %s\n", err.Error())
		// update failed status
		GlobalTxManager.Update(txData.Hash().Hex(), Failed)
		return
	}

	// update success status
	successTx := GlobalTxManager.GetTx(txData.Hash().Hex())
	successTx.Receipt = receipt
	successTx.Status = OnChain
	GlobalTxManager.UpdateTx(successTx)

	fmt.Printf("tx mined at block %d\n", receipt.BlockNumber.Int64())
}
