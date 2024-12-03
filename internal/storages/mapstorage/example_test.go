package mapstorage

import (
	"context"
	"fmt"
	"mishin-shortener/internal/delasync"
)

func ExampleStorage_Push() {
	stor, err := Make()

	err = stor.Push(context.Background(), "short", "original", "userId")
	if err != nil {
		fmt.Println(err)
	}
}

func ExampleStorage_PushBatch() {
	stor, err := Make()

	err = stor.PushBatch(context.Background(), &map[string]string{"short": "original"}, "userId")
	if err != nil {
		fmt.Println(err)
	}
}

func ExampleStorage_Get() {
	stor, err := Make()

	value, err := stor.Get(context.Background(), "short")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
}

func ExampleStorage_GetByUserID() {
	stor, err := Make()

	value, err := stor.GetByUserID(context.Background(), "userId")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
}

func ExampleStorage_DeleteByUserID() {
	stor, err := Make()

	err = stor.DeleteByUserID(context.Background(), []delasync.DelPair{{UserID: "UserID", Item: "Short"}})
	if err != nil {
		fmt.Println(err)
	}
}

func ExampleStorage_Ping() {
	stor, err := Make()

	err = stor.Ping(context.Background())
	if err != nil {
		fmt.Println(err)
	}
}

func ExampleStorage_Finish() {
	stor, err := Make()

	err = stor.Finish()
	if err != nil {
		fmt.Println(err)
	}
}
