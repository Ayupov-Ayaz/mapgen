package main

import (
	"log"

	"github.com/ayupov-ayaz/mapgen/analysis"
	"github.com/ayupov-ayaz/mapgen/services"
)

func mapByType(writer services.FileRecorder) error {
	const (
		_package   = "models"
		filepath   = "./example/models/payout"
		structName = "LinePayout"
		mapType    = "SymbolPayouts"
	)

	mapData := analysis.NewMapParams(_package, filepath, mapType, structName)

	return analysis.GenerateMapByString(writer, mapData)
}

func main() {
	writer := services.NewRecorder()

	if err := mapByType(writer); err != nil {
		log.Fatal(err)
	}
}
