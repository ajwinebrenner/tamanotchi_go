package main

import (
	"github.com/ajwinebrenner/tamanotchi_go/api/internal/config"
	"github.com/ajwinebrenner/tamanotchi_go/api/internal/server"
)

func main() {
	server.Start(config.Load())
}
