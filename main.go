package main

import (
	"fmt"

	"./protobuf"
	"github.com/golang/protobuf/proto"
)

func main() {
	// Create a User format, and put data into it.
	data := protobuf.User{
		Id:       12345,
		Username: "Yu Wang",
		Password: "Hello",
	}

	// Make data encode to Protocol buffer format (Be aware of using Pointer)
	dataBuffer, _ := proto.Marshal(&data)

	// Make the data which is already encoded, decode to protobuf.User format
	var user protobuf.User
	proto.Unmarshal(dataBuffer, &user)

	// Output the result
	fmt.Println(user.Id, " ", user.Username, " ", user.Password)
}
