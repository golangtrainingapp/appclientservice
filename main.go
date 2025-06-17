package main

import (
	"fmt"
	"github.com/golangtrainingapp/upstreamservice"
)

func main() {
	resp, err := upstreamservice.GetSingleResponseByLatAndLong(53.1900, -112.2500)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
