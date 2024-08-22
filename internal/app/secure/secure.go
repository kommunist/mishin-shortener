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

type UserIdKeyType int

const UserIdKey UserIdKeyType = 0

func setKeyRandom(size int) {
	if len(keyRandom) > 0 {
		return
	} else {
		keyRandom, _ = generateRandom(size) // добавить обработку ошибки
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

func Encrypt(data string) (string, error) {
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

	nonce, err := generateRandom(aesgcm.NonceSize())
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", err
	}

	encrypted := aesgcm.Seal(nil, nonce, []byte(data), nil) // зашифровываем
	encrypted = append(encrypted, nonce...)                 // подложим вектор в конце

	slog.Info("Encrypted data + nonce", "encrypted", encrypted)

	result := base64.StdEncoding.EncodeToString(encrypted) // в base64

	slog.Info("Encrypted in base64", "base64", result)

	return result, nil
}

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

	input, _ := base64.StdEncoding.DecodeString(data) // из base64 // обработать ошибку

	slog.Info("Decoded from base64", "input", input)

	token := []byte(input[:len(input)-aesgcm.NonceSize()])
	nonce := []byte(input[len(input)-aesgcm.NonceSize():])

	slog.Info("decrypt data", "token", token, "nonce", nonce)

	result, _ := aesgcm.Open(nil, nonce, token, nil) // расшифровываем // обработать ошибку

	return string(result), nil
}
