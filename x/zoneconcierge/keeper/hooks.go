package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	checkpointingtypes "github.com/babylonchain/babylon/x/checkpointing/types"
	epochingtypes "github.com/babylonchain/babylon/x/epoching/types"
	"github.com/babylonchain/babylon/x/zoneconcierge/types"
)

type Hooks struct {
	k Keeper
}

// ensures Hooks implements ClientHooks interfaces
var _ checkpointingtypes.CheckpointingHooks = Hooks{}
var _ epochingtypes.EpochingHooks = Hooks{}

func (k Keeper) Hooks() Hooks { return Hooks{k} }

// AfterEpochEnds is triggered upon an epoch has ended
func (h Hooks) AfterEpochEnds(ctx sdk.Context, epoch uint64) {
	// upon an epoch has ended, index the current chain info for each CZ
	for _, chainID := range h.k.GetAllChainIDs(ctx) {
		h.k.recordEpochChainInfo(ctx, chainID, epoch)
	}
}

// AfterRawCheckpointFinalized is triggered upon an epoch has been finalised
func (h Hooks) AfterRawCheckpointFinalized(ctx sdk.Context, epoch uint64) error {
	// upon an epoch has been finalised, update the last finalised epoch
	h.k.setFinalizedEpoch(ctx, epoch)

	headersToBroadcast := h.k.getHeadersToBroadcast(ctx)
	// send BTC timestamp to all open channels with ZoneConcierge
	h.k.BroadcastBTCTimestamps(ctx, epoch, headersToBroadcast)

	// Update the last broadcasted segment
	h.k.setLastSentSegment(ctx, &types.BTCChainSegment{
		BtcHeaders: headersToBroadcast,
	})
	return nil
}

// Other unused hooks

func (h Hooks) AfterBlsKeyRegistered(ctx sdk.Context, valAddr sdk.ValAddress) error { return nil }
func (h Hooks) AfterRawCheckpointConfirmed(ctx sdk.Context, epoch uint64) error     { return nil }

func (h Hooks) AfterRawCheckpointForgotten(ctx sdk.Context, ckpt *checkpointingtypes.RawCheckpoint) error {
	return nil
}
func (h Hooks) AfterRawCheckpointBlsSigVerified(ctx sdk.Context, ckpt *checkpointingtypes.RawCheckpoint) error {
	return nil
}
func (h Hooks) AfterEpochBegins(ctx sdk.Context, epoch uint64)                          {}
func (h Hooks) BeforeSlashThreshold(ctx sdk.Context, valSet epochingtypes.ValidatorSet) {}
