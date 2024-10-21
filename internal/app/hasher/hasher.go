// Модуль hasher предоставляет функционал хеширования для укорачивания ссылок.
package hasher

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5Hash(text []byte) string {
	hash := md5.Sum(text)
	return hex.EncodeToString(hash[:])
}
