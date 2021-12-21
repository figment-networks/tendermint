package state

import (
	"github.com/tendermint/tendermint/types"
)

var (
	execHook ExecutionHook = NoopExecutionHook{}
)

func SetExecutionHook(hook ExecutionHook) {
	execHook = hook
}

type ExecutionHook interface {
	Start(int64) error
	End(int64) error
	Block(types.EventDataNewBlock) error
	BlockHeader(types.EventDataNewBlockHeader) error
	Evidence(types.EventDataNewEvidence) error
	Tx(types.EventDataTx) error
	ValidatorSetUpdates(types.EventDataValidatorSetUpdates) error
}

type NoopExecutionHook struct{}

func (NoopExecutionHook) Start(height int64) error {
	return nil
}

func (NoopExecutionHook) End(height int64) error {
	return nil
}

func (NoopExecutionHook) Block(data types.EventDataNewBlock) error {
	return nil
}

func (NoopExecutionHook) BlockHeader(data types.EventDataNewBlockHeader) error {
	return nil
}

func (NoopExecutionHook) Evidence(types.EventDataNewEvidence) error {
	return nil
}

func (NoopExecutionHook) Tx(data types.EventDataTx) error {
	return nil
}

func (NoopExecutionHook) ValidatorSetUpdates(data types.EventDataValidatorSetUpdates) error {
	return nil
}
