package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/org/dyson/testutil/keeper"
	"github.com/org/dyson/testutil/nullify"
	"github.com/org/dyson/x/dyson/keeper"
	"github.com/org/dyson/x/dyson/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNScheduledRun(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ScheduledRun {
	items := make([]types.ScheduledRun, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetScheduledRun(ctx, items[i])
	}
	return items
}

func TestScheduledRunGet(t *testing.T) {
	keeper, ctx := keepertest.DysonKeeper(t)
	items := createNScheduledRun(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetScheduledRun(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestScheduledRunRemove(t *testing.T) {
	keeper, ctx := keepertest.DysonKeeper(t)
	items := createNScheduledRun(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveScheduledRun(ctx,
			item.Index,
		)
		_, found := keeper.GetScheduledRun(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestScheduledRunGetAll(t *testing.T) {
	keeper, ctx := keepertest.DysonKeeper(t)
	items := createNScheduledRun(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllScheduledRun(ctx)),
	)
}
