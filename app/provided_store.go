package app

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Fantom-foundation/go-lachesis/inter/idx"
	"github.com/Fantom-foundation/go-lachesis/inter/sfctype"
)

/*
 * NOTE: all the methods are temporary and will be refactored during Tendermint implementation.
 */

// GetEpochValidators provides store's method.
func (a App) GetEpochValidators(epoch idx.Epoch) []sfctype.SfcStakerAndID {
	return a.store.GetEpochValidators(epoch)
}

// GetSfcConstants provides store's method.
func (a App) GetSfcConstants(epoch idx.Epoch) SfcConstants {
	return a.store.GetSfcConstants(epoch)
}

// SetReceipts provides store's method.
func (a App) SetReceipts(n idx.Block, receipts types.Receipts) {
	a.store.SetReceipts(n, receipts)
}

//  provides store's method.
func (a App) GetReceipts(n idx.Block) types.Receipts {
	return a.store.GetReceipts(n)
}

// StateDB provides store's method.
func (a App) StateDB(from common.Hash) *state.StateDB {
	return a.store.StateDB(from)
}

// HasSfcStaker provides store's method.
func (a App) HasSfcStaker(stakerID idx.StakerID) bool {
	return a.store.HasSfcStaker(stakerID)
}

// HasEpochValidator provides store's method.
func (a App) HasEpochValidator(epoch idx.Epoch, stakerID idx.StakerID) bool {
	return a.store.HasEpochValidator(epoch, stakerID)
}

// GetSfcStaker provides store's method.
func (a App) GetSfcStaker(stakerID idx.StakerID) *sfctype.SfcStaker {
	return a.store.GetSfcStaker(stakerID)
}

// ForEachSfcStaker provides store's method.
func (a App) ForEachSfcStaker(do func(sfctype.SfcStakerAndID)) {
	a.store.ForEachSfcStaker(do)
}

// ForEachSfcDelegator provides store's method.
func (a App) ForEachSfcDelegator(do func(sfctype.SfcDelegatorAndAddr)) {
	a.store.ForEachSfcDelegator(do)
}

// GetSfcDelegator provides store's method.
func (a App) GetSfcDelegator(addr common.Address) *sfctype.SfcDelegator {
	return a.store.GetSfcDelegator(addr)
}

// GetDelegatorClaimedRewards provides store's method.
func (a App) GetDelegatorClaimedRewards(addr common.Address) *big.Int {
	return a.store.GetDelegatorClaimedRewards(addr)
}

// GetStakerClaimedRewards provides store's method.
func (a App) GetStakerClaimedRewards(stakerID idx.StakerID) *big.Int {
	return a.store.GetStakerClaimedRewards(stakerID)
}

// GetStakerDelegatorsClaimedRewards provides store's method.
func (a App) GetStakerDelegatorsClaimedRewards(stakerID idx.StakerID) *big.Int {
	return a.store.GetStakerDelegatorsClaimedRewards(stakerID)
}
