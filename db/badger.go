package db

import (
	"github.com/danskeren/database/kv"
	"github.com/dgraph-io/badger/v2"
	"github.com/dgraph-io/badger/v2/options"
)

var BadgerDB kv.KV

func init() {
	var err error
	opts := badger.DefaultOptions("./badger.db")
	opts.ValueLogLoadingMode = options.FileIO
	opts.TableLoadingMode = options.FileIO
	BadgerDB, err = kv.Open(opts)
	if err != nil {
		panic(err)
	}
}
