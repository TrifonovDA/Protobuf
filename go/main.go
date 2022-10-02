package main

import (
	"/mnt/C/Go/JSON_AND_HTTP/protos/my"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("/mnt/C/Go/JSON_AND_HTTP/output.bin", os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}

	result, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	//my := '/mnt/C/Go/JSON_AND_HTTP/protos/my'
	params := &my.SellerParams{}
	err = proto.Unmarshal(result, params)

	fmt.Printf("result: \n%+v\n\n\n", params)
	fmt.Printf("first item:\n%+v\n\n\n", params.GetResult()[0])
}
