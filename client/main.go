package main

import (
  "fmt"
  "bufio"
  "log"
  "log/slog"
  "net"
  "os"
)

func main() {
  fmt.Println(Prompt("My Tiny Chat"))
  conn, err := net.Dial("tcp", "localhost:8080")
  if err != nil {
    log.Fatal(err)
  }
  defer conn.Close()

  go receiveMessage(conn)

  reader := bufio.NewReader(os.Stdin)

  for{
    fmt.Print("> ")
    text, _ := reader.ReadString('\n')
    conn.Write([]byte(text))
  }
}

func receiveMessage(conn net.Conn) {
  for {
    message, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
      slog.Error(err.Error())
      return
    }
    fmt.Print(message)
  }
}
