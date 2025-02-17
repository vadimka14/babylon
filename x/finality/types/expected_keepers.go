package types

import (
	"context"

	bbn "github.com/babylonlabs-io/babylon/types"
	bstypes "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	etypes "github.com/babylonlabs-io/babylon/x/epoching/types"
)

type BTCStakingKeeper interface {
	GetParams(ctx context.Context) bstypes.Params
	GetFinalityProvider(ctx context.Context, fpBTCPK []byte) (*bstypes.FinalityProvider, error)
	HasFinalityProvider(ctx context.Context, fpBTCPK []byte) bool
	SlashFinalityProvider(ctx context.Context, fpBTCPK []byte) error
	GetVotingPower(ctx context.Context, fpBTCPK []byte, height uint64) uint64
	GetVotingPowerTable(ctx context.Context, height uint64) map[string]uint64
	GetBTCStakingActivatedHeight(ctx context.Context) (uint64, error)
	GetVotingPowerDistCache(ctx context.Context, height uint64) (*bstypes.VotingPowerDistCache, error)
	RemoveVotingPowerDistCache(ctx context.Context, height uint64)
	UnjailFinalityProvider(ctx context.Context, fpBTCPK []byte) error
}

type CheckpointingKeeper interface {
	GetEpoch(ctx context.Context) *etypes.Epoch
	GetLastFinalizedEpoch(ctx context.Context) uint64
}

// IncentiveKeeper defines the expected interface needed to distribute rewards.
type IncentiveKeeper interface {
	RewardBTCStaking(ctx context.Context, height uint64, filteredDc *bstypes.VotingPowerDistCache)
}

type BtcStakingHooks interface {
	AfterFinalityProviderActivated(ctx context.Context, btcPk *bbn.BIP340PubKey) error
}

type FinalityHooks interface {
	AfterSluggishFinalityProviderDetected(ctx context.Context, btcPk *bbn.BIP340PubKey) error
}
