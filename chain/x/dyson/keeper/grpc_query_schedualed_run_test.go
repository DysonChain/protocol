package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/org/dyson/testutil/keeper"
	"github.com/org/dyson/testutil/nullify"
	"github.com/org/dyson/x/dyson/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestScheduledRunQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DysonKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNScheduledRun(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetScheduledRunRequest
		response *types.QueryGetScheduledRunResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetScheduledRunRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetScheduledRunResponse{ScheduledRun: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetScheduledRunRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetScheduledRunResponse{ScheduledRun: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetScheduledRunRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ScheduledRun(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestScheduledRunQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DysonKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNScheduledRun(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllScheduledRunRequest {
		return &types.QueryAllScheduledRunRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ScheduledRunAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ScheduledRun), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ScheduledRun),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ScheduledRunAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ScheduledRun), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ScheduledRun),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ScheduledRunAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ScheduledRun),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ScheduledRunAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
