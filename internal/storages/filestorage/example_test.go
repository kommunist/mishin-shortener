package filestorage

import (
	"context"
	"fmt"
	"mishin-shortener/internal/app/delasync"
	"os"
)

func ExampleStorage_Push() {
	testFile, _ := os.CreateTemp("", "pattern")
	stor := Make(testFile.Name())

	err := stor.Push(context.Background(), "short", "original", "userId")
	if err != nil {
		fmt.Println(err)
	}
}

func ExampleStorage_PushBatch() {
	testFile, _ := os.CreateTemp("", "pattern")
	stor := Make(testFile.Name())

	err := stor.PushBatch(context.Background(), &map[string]string{"short": "original"}, "userId")
	if err != nil {
		fmt.Println(err)
	}
}

func ExampleStorage_Get() {
	testFile, _ := os.CreateTemp("", "pattern")
	stor := Make(testFile.Name())

	value, err := stor.Get(context.Background(), "short")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
}

func ExampleStorage_GetByUserID() {
	testFile, _ := os.CreateTemp("", "pattern")
	stor := Make(testFile.Name())

	value, err := stor.GetByUserID(context.Background(), "userId")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
}

func ExampleStorage_DeleteByUserID() {
	testFile, _ := os.CreateTemp("", "pattern")
	stor := Make(testFile.Name())

	err := stor.DeleteByUserID(context.Background(), []delasync.DelPair{{UserID: "UserID", Item: "Short"}})
	if err != nil {
		fmt.Println(err)
	}
}

func ExampleStorage_Ping() {
	testFile, _ := os.CreateTemp("", "pattern")
	stor := Make(testFile.Name())

	err := stor.Ping(context.Background())
	if err != nil {
		fmt.Println(err)
	}
}

func ExampleStorage_Finish() {
	testFile, _ := os.CreateTemp("", "pattern")
	stor := Make(testFile.Name())

	err := stor.Finish()
	if err != nil {
		fmt.Println(err)
	}
}
