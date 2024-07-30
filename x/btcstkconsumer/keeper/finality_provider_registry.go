package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/runtime"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"

	bbn "github.com/babylonlabs-io/babylon/types"
	btcstaking "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	btcstktypes "github.com/babylonlabs-io/babylon/x/btcstkconsumer/types"
)

// SetConsumerFinalityProvider adds the given finality provider to CZ chains KVStore
func (k Keeper) SetConsumerFinalityProvider(ctx context.Context, fp *btcstaking.FinalityProvider) {
	store := k.finalityProviderStore(ctx, fp.ConsumerId)
	fpBytes := k.cdc.MustMarshal(fp)
	store.Set(fp.BtcPk.MustMarshal(), fpBytes)
	// Update the finality provider chain store
	fpChainStore := k.finalityProviderChainStore(ctx)
	fpChainStore.Set(fp.BtcPk.MustMarshal(), []byte(fp.ConsumerId))
}

// GetConsumerFinalityProvider gets the finality provider in the chain id with the given finality provider Bitcoin PK
func (k Keeper) GetConsumerFinalityProvider(ctx context.Context, consumerId string, fpBTCPK *bbn.BIP340PubKey) (*btcstaking.FinalityProvider, error) {
	if !k.HasConsumerFinalityProvider(ctx, fpBTCPK) {
		return nil, btcstaking.ErrFpNotFound
	}
	store := k.finalityProviderStore(ctx, consumerId)
	fpBytes := store.Get(*fpBTCPK)
	if fpBytes == nil {
		// FP exists, but not for this chain id
		return nil, btcstaking.ErrFpNotFound
	}
	var fp btcstaking.FinalityProvider
	k.cdc.MustUnmarshal(fpBytes, &fp)
	return &fp, nil
}

// HasConsumerFinalityProvider checks if the finality provider already exists, across / independently of all chain ids
func (k Keeper) HasConsumerFinalityProvider(ctx context.Context, fpBTCPK *bbn.BIP340PubKey) bool {
	store := k.finalityProviderChainStore(ctx)
	return store.Has(*fpBTCPK)
}

// GetConsumerOfFinalityProvider gets the finality provider chain id for the given finality provider Bitcoin PK
func (k Keeper) GetConsumerOfFinalityProvider(ctx context.Context, fpBTCPK *bbn.BIP340PubKey) (string, error) {
	if !k.HasConsumerFinalityProvider(ctx, fpBTCPK) {
		return "", btcstaking.ErrFpNotFound
	}
	store := k.finalityProviderChainStore(ctx)
	chainID := store.Get(*fpBTCPK)
	return string(chainID), nil
}

// IterateFPs iterates over all finality providers for a given chain
func (k Keeper) IterateFPs(ctx context.Context, chainId string, handler func(fp *btcstaking.FinalityProvider) bool) {
	// filter out all finality providers with positive voting power
	fpIter := k.finalityProviderStore(ctx, chainId).Iterator(nil, nil)
	defer fpIter.Close()
	for ; fpIter.Valid(); fpIter.Next() {
		var fp btcstaking.FinalityProvider
		k.cdc.MustUnmarshal(fpIter.Value(), &fp)

		shouldContinue := handler(&fp)
		if !shouldContinue {
			return
		}
	}
}

// HasFPs checks if the FPs list exists for a given chain id
func (k Keeper) HasFPs(ctx context.Context, chainID string) bool {
	store := k.finalityProviderStore(ctx, chainID)
	iter := store.Iterator(nil, nil)
	defer iter.Close()
	return iter.Valid()
}

// GetFPs gets the list of FPs, i.e., the finality provider set for a given chain id
func (k Keeper) GetFPs(ctx context.Context, chainID string) map[string]btcstaking.FinalityProvider {
	store := k.finalityProviderStore(ctx, chainID)
	iter := store.Iterator(nil, nil)
	defer iter.Close()

	// if no finality provider for this chain id, return nil
	if !iter.Valid() {
		return nil
	}

	// get all finality providers for this chain id
	fpSet := map[string]btcstaking.FinalityProvider{}
	for ; iter.Valid(); iter.Next() {
		fpBTCPK, err := bbn.NewBIP340PubKey(iter.Key())
		if err != nil {
			// failing to unmarshal finality provider BTC PK in KVStore is a programming error
			panic(errorsmod.Wrapf(bbn.ErrUnmarshal, "Failed to unmarshall FP BTC PK %s: %v", iter.Key(), err))
		}
		var fp btcstaking.FinalityProvider
		k.cdc.MustUnmarshal(iter.Value(), &fp)
		fpSet[fpBTCPK.MarshalHex()] = fp
	}

	return fpSet
}

// finalityProviderStore returns the KVStore of the finality provider set per chain
// prefix: ConsumerFinalityProviderKey || chain id
// key: Bitcoin PubKey
// value: FinalityProvider
func (k Keeper) finalityProviderStore(ctx context.Context, chainID string) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	fpStore := prefix.NewStore(storeAdapter, btcstktypes.ConsumerFinalityProviderKey)
	return prefix.NewStore(fpStore, []byte(chainID))
}

// finalityProviderChainStore returns the KVStore of the finality provider chain per FP BTC PubKey
// prefix: ConsumerFinalityProviderChainKey
// key: Bitcoin PubKey
// value: chain id
func (k Keeper) finalityProviderChainStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, btcstktypes.ConsumerFinalityProviderChainKey)
}
