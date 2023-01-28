package gas

import (
	"context"

	"github.com/ajhodges/stackup-bundler/pkg/errors"
	"github.com/ajhodges/stackup-bundler/pkg/userop"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CallGasEstimate(
	eth *ethclient.Client,
	from common.Address,
	op *userop.UserOperation,
) (uint64, error) {
	est, err := eth.EstimateGas(context.Background(), ethereum.CallMsg{
		From: from,
		To:   &op.Sender,
		Data: op.CallData,
	})
	if err != nil {
		return 0, errors.NewRPCError(errors.EXECUTION_REVERTED, err.Error(), err.Error())
	}

	return est, nil
}
