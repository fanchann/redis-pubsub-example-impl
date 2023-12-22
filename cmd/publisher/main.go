package main

import (
	"belajar-cli/db"
	"belajar-cli/types"
	"fmt"
)

func main() {

	for {
		var topic string

		fmt.Print("\nEnter Topic :")
		fmt.Scanln(&topic)

		err := db.RedisConnection().Publish(types.Ctx, types.ChannelName, topic).Err()
		if err != nil {
			panic(err)
		}

		fmt.Printf("\nTopic %s was delivered \n", topic)
	}
}
