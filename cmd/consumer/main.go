package main

import (
	"fmt"

	"github.com/fanchann/redis-pubsub-example-impl/db"
	"github.com/fanchann/redis-pubsub-example-impl/types"
)

func main() {
	subcriber := db.RedisConnection().Subscribe(types.Ctx, types.ChannelName)
	defer subcriber.Close()

	msg := subcriber.Channel()

	for ms := range msg {

		fmt.Printf("Message %s | From %s \n", ms.Payload, ms.Channel)
	}
}
