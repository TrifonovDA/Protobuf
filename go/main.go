package main

import (
	"fmt"
	"github.com/TrifonovDA/Protobuf/protos/my"
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

	params := &my.SellerParams{}
	err = proto.Unmarshal(result, params)

	fmt.Printf("result: \n%+v\n\n\n", params)
	fmt.Printf("first item:\n%+v\n\n\n", params.GetResult()[0])

	//fmt.Print(my)

}
