package keeper

import (
	"github.com/auagho/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
