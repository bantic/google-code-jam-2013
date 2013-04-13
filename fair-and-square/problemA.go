package main

import (
    "os"
    "bufio"
    "bytes"
    "io"
    "fmt"
    "strings"
    "strconv"
    "math"
)

// Read a whole file into the memory and store it as array of lines
func readLines(path string) (lines []string, err error) {
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

func integersFromString(str string) (ints []int) {
  intStrings := strings.Split(str, " ")
  for _, intString := range intStrings {
    intVal, err := strconv.Atoi(intString)
    if err != nil { panic(err) }

    ints = append(ints, intVal)
  }
  return
}

func integerIsPalindrome(x int) bool {
  str := strconv.Itoa(x)
  return str == reverseString(str)
}

func integerIsPerfectSquare(x int) bool {
  sqrt := math.Sqrt( float64(x) )
  return sqrt == math.Floor(sqrt)
}

func integerIsFairAndSquare(x int) bool {
  if !integerIsPalindrome(x) { return false }

  if !integerIsPerfectSquare(x) { return false }

  sqrt := math.Sqrt( float64(x) )
  if integerIsPalindrome( int(sqrt) ) { return true }

  return false
}

func reverseString(str string) string {
  bytes := make([]byte, len(str))
  var j int = len(bytes) - 1
  for i := 0; i <= j; i++ {
    bytes[j-i] = str[i]
  }
  return string(bytes)
}

func processInputFile(path string) {
  lines, err := readLines(path)
  if err != nil { panic(err) }

  nextLine, lines := lines[0], lines[1:]
  testCases, err := strconv.Atoi( nextLine )
  if err != nil { panic(err) }

  // fmt.Printf("Test cases: %d\n", testCases)

  for i := 0; i < testCases; i++ {
    nextLine := lines[i]
    fmt.Printf("Case #%d: %s\n", (i+1), lineResult(nextLine))
  }
}

func lineResult(line string) string {
  bounds := integersFromString( line )
  // fmt.Println("start %d, end %d", bounds[0], bounds[1])
  return fmt.Sprintf("%d", countFairAndSquareIntegersInBounds( bounds[0], bounds[1] ) )
}

func countFairAndSquareIntegersInBounds(startInt, endInt int) int {
  count := 0
  for i := startInt; i <= endInt; i++ {
    if integerIsFairAndSquare(i) { count++ }
  }
  return count
}

func main() {
  processInputFile( "input.in" )
}
