package main

import (
	"github.com/DiegoAndresMarmota/Local-Storage/api/models"
)

func main() {
	models.InitConnect()
	models.CloseConnect()
}
