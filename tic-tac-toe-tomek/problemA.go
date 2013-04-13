package main

import (
  "fmt"
  "os"
  "io"
  "bufio"
  "strings"
)


func main() {
  fmt.Println("HELLO, starting")

  file, err := os.Open("input.in")
  if err != nil { panic(err) }

  reader := bufio.NewReader(file)

  for {
    line, err := reader.ReadString('\n')
    if err == io.EOF {
      fmt.Println("Done reading")
      break
    } else if err != nil { panic(err) }

    line = strings.TrimSpace(line)

    if len(line) == 0 { continue }

    fmt.Printf("Read line %s (%d)\n", line, len(line))
  }
}
