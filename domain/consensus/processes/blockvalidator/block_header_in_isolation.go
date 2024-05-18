package blockvalidator

import (
	"github.com/Pugdag/pugdagd/domain/consensus/model"
	"github.com/Pugdag/pugdagd/domain/consensus/model/externalapi"
	"github.com/Pugdag/pugdagd/domain/consensus/ruleerrors"
	"github.com/Pugdag/pugdagd/domain/consensus/utils/consensushashing"
	"github.com/Pugdag/pugdagd/domain/consensus/utils/constants"
	"github.com/Pugdag/pugdagd/infrastructure/logger"
	"github.com/Pugdag/pugdagd/util/mstime"
	"github.com/pkg/errors"
)

// ValidateHeaderInIsolation validates block headers in isolation from the current
// consensus state
func (v *blockValidator) ValidateHeaderInIsolation(stagingArea *model.StagingArea, blockHash *externalapi.DomainHash) error {
	onEnd := logger.LogAndMeasureExecutionTime(log, "ValidateHeaderInIsolation")
	defer onEnd()

	header, err := v.blockHeaderStore.BlockHeader(v.databaseContext, stagingArea, blockHash)
	if err != nil {
		return err
	}

	//todo : drop this
	//log.Info("blockHash %s - genesisHash %s", blockHash, v.genesisHash)

	if !blockHash.Equal(v.genesisHash) {
		err = v.checkBlockVersion(header)
		if err != nil {
			return err
		}
	}

	err = v.checkBlockTimestampInIsolation(header)
	if err != nil {
		return err
	}

	err = v.checkParentsLimit(header)
	if err != nil {
		return err
	}

	return nil
}

func (v *blockValidator) checkParentsLimit(header externalapi.BlockHeader) error {
	hash := consensushashing.HeaderHash(header)
	if len(header.DirectParents()) == 0 && !hash.Equal(v.genesisHash) {
		return errors.Wrapf(ruleerrors.ErrNoParents, "block has no parents")
	}

	if uint64(len(header.DirectParents())) > uint64(v.maxBlockParents) {
		return errors.Wrapf(ruleerrors.ErrTooManyParents, "block header has %d parents, but the maximum allowed amount "+
			"is %d", len(header.DirectParents()), v.maxBlockParents)
	}
	return nil
}

func (v *blockValidator) checkBlockVersion(header externalapi.BlockHeader) error {
	if header.Version() != constants.BlockVersion {
		return errors.Wrapf(
			ruleerrors.ErrWrongBlockVersion, "The block version should be %d", constants.BlockVersion)
	}
	return nil
}

func (v *blockValidator) checkBlockTimestampInIsolation(header externalapi.BlockHeader) error {
	blockTimestamp := header.TimeInMilliseconds()
	now := mstime.Now().UnixMilliseconds()
	maxCurrentTime := now + int64(v.timestampDeviationTolerance)*v.targetTimePerBlock.Milliseconds()
	if blockTimestamp > maxCurrentTime {
		return errors.Wrapf(
			ruleerrors.ErrTimeTooMuchInTheFuture, "The block timestamp is in the future.")
	}
	return nil
}
