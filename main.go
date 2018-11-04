package main

import (
	"fmt"
	"github.com/ncw/swift"
	"os"
)

func main() {
	conn := swift.Connection{
		Domain:         "Default",
		UserName:       "ssDDDsssAAAAAaa",
		ApiKey:         "aqawsAQSAWSSsssSAWWSASSdddddSSSSS",
		AuthUrl:        "https://auth.cloud.ovh.net/v3",
		Region:         "SBG3",
	}
	err := conn.Authenticate()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = conn.ContainerCreate("containerName", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	content, err := os.Open("/tmp/intel.pdf")
	if err != nil {
		fmt.Println(err)
		return
	}

	headers, err := conn.ObjectPut("containerName", "book.pdf", content, false, "", "application/pdf", nil)
	if err!=nil{
		fmt.Println(err)
		return
	}
	
	fmt.Println(headers)
}