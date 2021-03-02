package main

import (
	"fmt"
	"log"
	"metaspoilt-client/rpc"
	"os"
)

func main() {
	host := os.Getenv("MSFHOST")
	pass := os.Getenv("MSFPASS")
	user := "msf"
	meta, err := rpc.New(host, user, pass)
	if err != nil {
		log.Fatalln(err)
	}
	defer meta.Logout()
	fmt.Println(meta)
}
