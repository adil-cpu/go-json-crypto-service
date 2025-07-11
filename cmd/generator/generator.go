package generator

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
)

func GenerateJSON() map[string]interface{} {
	rand.Seed(time.Now().UnixNano())

	names := []string{"Alice", "Bob", "Charlie", "Dana", "Erlan"}
	domains := []string{"example.com", "gmail.com", "mail.kz"}

	jsonData := map[string]interface{}{
		"id":        rand.Intn(1000),
		"name":      names[rand.Intn(len(names))],
		"email":     names[rand.Intn(len(names))] + "@" + domains[rand.Intn(len(domains))],
		"age":       rand.Intn(40) + 18,
		"is_active": rand.Intn(2) == 1,
	}

	return jsonData
}

func SaveJSONToFile(data map[string]interface{}, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}
