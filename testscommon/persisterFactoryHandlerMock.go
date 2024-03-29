package testscommon

import (
	"github.com/Dharitri-org/drtg-storage/common"
	"github.com/Dharitri-org/drtg-storage/leveldb"
	"github.com/Dharitri-org/drtg-storage/memorydb"
	"github.com/Dharitri-org/drtg-storage/storageUnit"
	"github.com/Dharitri-org/drtg-storage/types"
)

type persisterFactoryHandlerMock struct {
	dbType            storageUnit.DBType
	batchDelaySeconds int
	maxBatchSize      int
	maxOpenFiles      int
}

// NewPersisterFactoryHandlerMock -
func NewPersisterFactoryHandlerMock(dbType storageUnit.DBType, batchDelaySeconds int, maxBatchSize int, maxOpenFiles int) *persisterFactoryHandlerMock {
	return &persisterFactoryHandlerMock{
		dbType:            dbType,
		batchDelaySeconds: batchDelaySeconds,
		maxBatchSize:      maxBatchSize,
		maxOpenFiles:      maxOpenFiles,
	}
}

// Create -
func (mock *persisterFactoryHandlerMock) Create(path string) (types.Persister, error) {
	switch mock.dbType {
	case storageUnit.LvlDB:
		return leveldb.NewDB(path, mock.batchDelaySeconds, mock.maxBatchSize, mock.maxOpenFiles)
	case storageUnit.LvlDBSerial:
		return leveldb.NewSerialDB(path, mock.batchDelaySeconds, mock.maxBatchSize, mock.maxOpenFiles)
	case storageUnit.MemoryDB:
		return memorydb.New(), nil
	default:
		return nil, common.ErrNotSupportedDBType
	}
}

// IsInterfaceNil -
func (mock *persisterFactoryHandlerMock) IsInterfaceNil() bool {
	return mock == nil
}
