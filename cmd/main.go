package main

import (
	"github.com/DiegoAndresMarmota/Local-Storage/api/models"
	"github.com/DiegoAndresMarmota/Local-Storage/internal"
	"github.com/DiegoAndresMarmota/Local-Storage/pkg"
)

func main() {
	models.InitConnect()

	internal.ExistTables("suppliers")

	models.CloseConnect()
	pkg.Ping()
}
