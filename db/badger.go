package db

import (
	"github.com/danskeren/database/kv"
	"github.com/dgraph-io/badger"
)

var BadgerDB kv.KV

func init() {
	var err error
	BadgerDB, err = kv.Open(badger.DefaultOptions("./badger.db"))
	if err != nil {
		panic(err)
	}
}
