package parser

import (
	"os"
	"testing"
)

func TestParseJSONFile(t *testing.T) {
	// тестовый JSON
	testJSON := `{"name": "Batan", "age": 23}`

	// + временный файл
	err := os.WriteFile("test.json", []byte(testJSON), 0644)
	if err != nil {
		t.Fatalf("Ошибка при создании временного файла: %v", err)
	}
	defer os.Remove("test.json")

	// парсим JSON
	data, err := ParseJSONFile("test.json")
	if err != nil {
		t.Fatalf("Ошибка при парсинге JSON: %v", err)
	}

	// проверим значения
	if data["name"] != "Batan" {
		t.Errorf("Ожидалось name='Batan', получено: %v", data["name"])
	}
	if data["age"] != float64(23) {
		t.Errorf("Ожидалось age=23, получено: %v", data["age"])
	}
}
