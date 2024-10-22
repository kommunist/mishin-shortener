package secure

import (
	"fmt"
)

func Example() {
	original := "Lenin is a live!"

	encrypted, err := Encrypt(original)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(encrypted)

	decrypted, err := Decrypt(encrypted)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(decrypted)
}
