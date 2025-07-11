package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"go-json-crypto-service/cmd/crypto"
	"go-json-crypto-service/cmd/db"
	"go-json-crypto-service/cmd/generator"
	"go-json-crypto-service/cmd/parser"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла:", err)
	}
}

func main() {
	mode := os.Getenv("MODE")
	key := os.Getenv("FERNET_KEY")

	switch mode {
	case "ENCRYPTION":
		fmt.Println("...шифруем JSON...")

		// + случайный JSON
		data := generator.GenerateJSON()

		// 2. сохраняем JSON
		err := generator.SaveJSONToFile(data, "data.json")
		if err != nil {
			log.Fatal("JSON не сохранился: ", err)
		}

		// 3. шифруем
		err = crypto.EncryptFile("data.json", "data.encrypted", key)
		if err != nil {
			log.Fatal("шифровка не сработала: ", err)
		}

		fmt.Println("УРА, JSON сгенерирован, зашифрован и сохранен → data.encrypted")

	case "DECRYPTION":
		fmt.Println("...расшифровываем...")

		key := os.Getenv("FERNET_KEY")
		err := crypto.DecryptFile("data.encrypted", "data_decrypted.json", key)
		if err != nil {
			log.Fatal("неудачная расшифровка: ", err)
		}

		fmt.Println("расшифровали → data_decrypted.json")

		// Парсим JSON
		parsed, err := parser.ParseJSONFile("data_decrypted.json")
		if err != nil {
			log.Fatal("проблемы с парсингом JSON", err)
		}

		fmt.Println("расшифрованные данные: ", parsed)

		// Подключение к БД
		database, err := db.ConnectDB()
		if err != nil {
			log.Fatal("грохнули БД... ", err)
		}
		defer database.Close()

		fmt.Println("есть соединение с БД")

		// получаем список нужных ключей
		wantedKeys, err := db.GetWantedKeys(database)
		if err != nil {
			log.Fatal("ошибка при получении ключей из key_list: ", err)
		}

		if len(wantedKeys) == 0 {
			fmt.Println("внимание: key_list пуст, нечего фильтровать( ")
			return
		}

		fmt.Println("ключи из key_list:", wantedKeys)

		// фильтруем JSON по нужным ключам и сохраняем
		for _, key := range wantedKeys {
			if val, ok := parsed[key]; ok {
				strVal := fmt.Sprintf("%v", val)
				err := db.InsertFilteredData(database, key, strVal)
				if err != nil {
					log.Printf("не сохранили %s: %v\n", key, err)
				} else {
					fmt.Printf("сохранено: %s = %s\n", key, strVal)
				}
			} else {
				fmt.Printf("ключ %s не найден в JSON\n", key)
			}
		}

	default:
		log.Fatal("Некорректный MODE. Укажи ENCRYPTION или DECRYPTION в .env")
	}
}
