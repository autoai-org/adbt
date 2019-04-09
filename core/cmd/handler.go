package main

import (
	mongo "../internal/drivers/mongo"
	storages "../internal/storages"
	"github.com/tucnak/climax"
)

func backupHandler() climax.CmdHandler {
	return func(ctx climax.Context) int {
		var connString string
		if conn, ok := ctx.Get("host"); ok {
			connString = conn
		}
		dbService := mongo.NewDefaultMongoService(connString)
		storageService := storages.NewLocalStorageService()
		return 0
	}
}
