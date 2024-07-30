package keeper_test

import (
	"math/rand"
	"testing"

	"github.com/babylonlabs-io/babylon/app"
	"github.com/babylonlabs-io/babylon/testutil/datagen"
	"github.com/stretchr/testify/require"
)

func FuzzConsumerRegistry(f *testing.F) {
	datagen.AddRandomSeedsToFuzzer(f, 10)

	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))

		babylonApp := app.Setup(t, false)
		zcKeeper := babylonApp.BTCStkConsumerKeeper
		ctx := babylonApp.NewContext(false)

		// generate a random consumer register
		consumerRegister := datagen.GenRandomConsumerRegister(r)

		// check that the consumer is not registered
		isRegistered := zcKeeper.IsConsumerRegistered(ctx, consumerRegister.ConsumerId)
		require.False(t, isRegistered)

		// Check that the consumer is not registered
		consumerRegister2, err := zcKeeper.GetConsumerRegister(ctx, consumerRegister.ConsumerId)
		require.Error(t, err)
		require.Nil(t, consumerRegister2)

		// Register the consumer
		zcKeeper.SetConsumerRegister(ctx, consumerRegister)

		// check that the consumer is registered
		consumerRegister2, err = zcKeeper.GetConsumerRegister(ctx, consumerRegister.ConsumerId)
		require.NoError(t, err)
		require.Equal(t, consumerRegister.ConsumerId, consumerRegister2.ConsumerId)
		require.Equal(t, consumerRegister.ConsumerName, consumerRegister2.ConsumerName)
		require.Equal(t, consumerRegister.ConsumerDescription, consumerRegister2.ConsumerDescription)
	})
}
