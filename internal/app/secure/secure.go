// Модуль secure предоставляет функции шифрования и дешифрования для аутентификации.
package secure

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log/slog"
)

var keyRandom []byte // использовать глобальные переменные не хорошо. Но пока не стал загонять в файл

// Используется для прокидывания id пользователя через контекст
type UserIDKeyType int

// Используется для прокидывания id пользователя через контекст
const UserIDKey UserIDKeyType = 0

func setKeyRandom(size int) {
	if len(keyRandom) > 0 {
		return
	} else {
		keyRandom, _ = generateRandom(size) // TODO сделать обработку ошибки
	}
}

func generateRandom(size int) ([]byte, error) {
	// генерируем криптостойкие случайные байты в b
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Шифрование исходной строки. Используется для аутентификации
func Encrypt(data string) (string, error) {
	setKeyRandom(aes.BlockSize)

	slog.Info("Encrypt data", "data", data)

	aesblock, err := aes.NewCipher(keyRandom)
	if err != nil {
		slog.Error("Error when create encrypter", "err", err)
		return "", err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", err
	}

	nonce, err := generateRandom(aesgcm.NonceSize())
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", err
	}

	encrypted := aesgcm.Seal(nil, nonce, []byte(data), nil) // зашифровываем
	encrypted = append(encrypted, nonce...)                 // подложим вектор в конце

	result := base64.StdEncoding.EncodeToString(encrypted) // в base64

	return result, nil
}

// Расшифровка исходной строки. Используется для аутентификации
func Decrypt(data string) (string, error) {
	setKeyRandom(aes.BlockSize)
	aesblock, err := aes.NewCipher(keyRandom)
	if err != nil {
		slog.Error("Error when create encrypter", "err", err)
		return "", err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", err
	}

	input, _ := base64.StdEncoding.DecodeString(data) // из base64 // TODO сделать обработку ошибки

	slog.Info("Decoded from base64", "input", input)

	token := []byte(input[:len(input)-aesgcm.NonceSize()])
	nonce := []byte(input[len(input)-aesgcm.NonceSize():])

	result, _ := aesgcm.Open(nil, nonce, token, nil) // расшифровываем // TODO сделать обработку ошибки

	return string(result), nil
}
