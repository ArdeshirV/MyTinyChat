package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println(Prompt("My Tiny Chat"))
}

func receiveMessages(conn net.Conn) {
  reader := bufio.NewReader(conn)
  for {
    message, err := reader.ReadString('\n')
    if err != nil {
      break
    }
    fmt.Printf("Client> %s", message)
  }
}


