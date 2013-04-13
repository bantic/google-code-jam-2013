package main

import (
    "os"
    "bufio"
    "bytes"
    "io"
    "fmt"
    "strings"
)

// Read a whole file into the memory and store it as array of lines
func readLines(path string) (lines []string, err error) {
    fmt.Println("READING",path)
    var (
        file *os.File
        part []byte
        prefix bool
    )
    if file, err = os.Open(path); err != nil {
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    buffer := bytes.NewBuffer(make([]byte, 0))
    for {
        if part, prefix, err = reader.ReadLine(); err != nil {
            break
        }
        buffer.Write(part)
        if !prefix {
            str := strings.TrimSpace(buffer.String())
            if len(str) > 0 { lines = append(lines, str) }
            buffer.Reset()
        }
    }
    if err == io.EOF {
        err = nil
    }
    return
}

func main() {
  fmt.Println("HELLO, starting")
  lines, err := readLines( "input.in" )
  if err != nil { panic(err) }

  for i, line := range lines {
    fmt.Printf("Line: #%d %s (len %d)\n", i, line, len(line))
  }
}
