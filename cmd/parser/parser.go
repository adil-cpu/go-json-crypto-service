package parser

import (
	"encoding/json"
	"os"
)

// читаем файл и возвращаем его содержимое как map[string]interface{}
func ParseJSONFile(filename string) (map[string]interface{}, error) {
	// читаем...
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// парсим JSON
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
