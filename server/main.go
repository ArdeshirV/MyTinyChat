package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"os"

	"log"
)

func main() {
	fmt.Println(Prompt("My Tiny Chat"))
  ln, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Fatal(err)
  }
  defer ln.Close()

  slog.Info("Server started. Listen on port 8080")

  conn, err := ln.Accept()
  if err != nil {
    log.Fatal(err)
  }
  defer conn.Close()

  go receiveMessages(conn)
  scanner := bufio.NewScanner(os.Stdin)
  for {
    fmt.Print("Server> ")
    scanner.Scan()
    message := scanner.Text()
    if message == "/quit" {
      break
    }
    sendMessage(conn, message)
  }
  slog.Info("Server closed")
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

func sendMessage(conn net.Conn, message string) {
  conn.Write([]byte(message + "\n"))
}

