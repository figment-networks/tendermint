package deepmind

import (
	"github.com/figment-networks/tendermint-protobuf-def/codec"
	"github.com/golang/protobuf/proto"
	abci "github.com/tendermint/tendermint/abci/types"
	tmtypes "github.com/tendermint/tendermint/proto/tendermint/types"
	"github.com/tendermint/tendermint/types"
)

func encodeBlock(bh types.EventDataNewBlock) ([]byte, error) {
	ecBlock, err := bh.Block.ToProto()
	if err != nil {
		return nil, err
	}

	nb := &codec.EventBlock{
		Block: ecBlock,
	}

	// Adds the Block ID into the New Block Data
	nb.BlockId = &tmtypes.BlockID{
		Hash: bh.Block.Header.Hash(),
		PartSetHeader: tmtypes.PartSetHeader{
			Total: bh.Block.LastBlockID.PartSetHeader.Total,
			Hash:  bh.Block.LastBlockID.PartSetHeader.Hash,
		},
	}

	return proto.Marshal(nb)
}

func encodeTx(result *abci.TxResult) ([]byte, error) {

	return proto.Marshal(result)
}

func encodeValidatorSetUpdates(updates *types.EventDataValidatorSetUpdates) ([]byte, error) {

	result := &codec.EventValidatorSetUpdates{}

	for _, update := range updates.ValidatorUpdates {
		valUpdate, err := update.ToProto()
		if err != nil {
			return nil, err
		}

		result.ValidatorUpdates = append(result.ValidatorUpdates, valUpdate)

	}

	return proto.Marshal(result)
}
