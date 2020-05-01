package internal

import "strings"

type MapValData struct {
	Length int
	Vals   []string
	Join   string
}

type MapData struct {
	MapType string
	KV      MapKeyValue
	Data    map[string]MapValData
}

func NewMapValData(data []string) MapValData {
	return MapValData{
		Length: len(data),
		Vals:   data,
		Join:   strings.Join(data, ", "),
	}
}

func NewMapData(mapType string, keyVal MapKeyValue, data map[string]MapValData) MapData {
	return MapData{
		MapType: mapType,
		KV:      keyVal,
		Data:    data,
	}
}
