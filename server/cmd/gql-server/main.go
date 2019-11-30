package main

import (
	log "github.com/fedoratipper/bitkiosk/server/internal/logger"

	"github.com/fedoratipper/bitkiosk/server/internal/orm"
	"github.com/fedoratipper/bitkiosk/server/pkg/server"
)

func main() {
	orm, err := orm.Factory()
	if err != nil {
		log.Panic(err)
	}
	server.Run(orm)
}
