package internal

import "strings"

type MapValData struct {
	Length int
	Vals   []string
	Join   string
}

type MapData struct {
	MapType string
	KeyType string
	ValType string
	Data    map[string]MapValData
}

func NewMapValData(data []string) MapValData {
	return MapValData{
		Length: len(data),
		Vals:   data,
		Join:   strings.Join(data, ", "),
	}
}

func NewMapData(mapType, keyType, valType string, data map[string]MapValData) MapData {
	return MapData{
		MapType: mapType,
		KeyType: keyType,
		ValType: valType,
		Data:    data,
	}
}
