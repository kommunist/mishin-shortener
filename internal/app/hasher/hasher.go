// Модуль hasher предоставляет функционал хеширования для укорачивания ссылок.
package hasher

import (
	"crypto/md5"
	"encoding/hex"
)

// Функция хеширования переданных данных.
func GetMD5Hash(text []byte) string {
	hash := md5.Sum(text)
	return hex.EncodeToString(hash[:])
}
