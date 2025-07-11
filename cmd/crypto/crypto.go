package crypto

import (
	"fmt"
	"os"

	"github.com/fernet/fernet-go"
)

func EncryptFile(inputPath string, outputPath string, keyString string) error {
	// читаем JSON-файл
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("ошибка чтения файла %s: %w", inputPath, err)
	}

	// декодируем строковый ключ в объект Fernet
	key, err := fernet.DecodeKey(keyString)
	if err != nil {
		return fmt.Errorf("ошибка декодирования ключа: %w", err)
	}

	// шифруем данные
	encrypted, err := fernet.EncryptAndSign(data, key)
	if err != nil {
		return fmt.Errorf("ошибка шифрования: %w", err)
	}

	// шифр -> в файл
	err = os.WriteFile(outputPath, encrypted, 0644)
	if err != nil {
		return fmt.Errorf("ошибка записи зашифрованного файла: %w", err)
	}

	return nil
}

func DecryptFile(inputPath string, outputPath string, keyString string) error {
	// читаем байты из зашифрованного файла
	encryptedData, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("ошибка чтения зашифрованного файла %s: %w", inputPath, err)
	}

	// декодируем ключ
	key, err := fernet.DecodeKey(keyString)
	if err != nil {
		return fmt.Errorf("ошибка декодирования ключа: %w", err)
	}

	// расшифровываем данные
	decrypted := fernet.VerifyAndDecrypt(encryptedData, 0, []*fernet.Key{key})
	if decrypted == nil {
		return fmt.Errorf("не удалось расшифровать данные: VerifyAndDecrypt вернул nil")
	}

	// сохраняем данные в файл
	err = os.WriteFile(outputPath, decrypted, 0644)
	if err != nil {
		return fmt.Errorf("ошибка записи расшифрованного файла: %w", err)
	}

	return nil
}
