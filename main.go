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
	// defer meta.Logout()
	fmt.Println(meta)

	sessions, err := meta.SessionList()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("Sessions :")
	fmt.Println(len(sessions))
	for _, session := range sessions {
		fmt.Printf("%+v\n", session)
		fmt.Println("=====")
	}
	meta.Version()
}
