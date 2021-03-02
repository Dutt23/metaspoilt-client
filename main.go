package main

import (
	"fmt"
	"metaspoilt-minimal/rpc"
	"os"
)

func main() {
	host := os.Getenv("MSFHOST")
	pass := os.Getenv("MSFPASS")
	user := "msf"
	meta, err := rpc.New(host, user, pass)
	fmt.Println(meta)
	fmt.Println(err)
}
