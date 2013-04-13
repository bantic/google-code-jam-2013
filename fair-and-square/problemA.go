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

func integerIsFairAndSquare(x int, seenPtr *map[int]bool, valuePtr *map[int]bool) bool {

  seenMap  := *seenPtr
  valueMap := *valuePtr

  if seenMap[x] {
    // fmt.Printf("returning from cache: %d -> %t", x, valueMap[x])
    return valueMap[x]
  } else {
    seenMap[x] = true
  }

  if !integerIsPalindrome(x) {
    valueMap[x] = false
    return false
  }

  possibleSqrRoot := PerfectSqrRoot(x)
  if possibleSqrRoot == 0 {
    valueMap[x] = false
    return false
  }

  if integerIsPalindrome( possibleSqrRoot ) {
    valueMap[x] = true
    return true
  }

  valueMap[x] = false
  return false
}

func PerfectSqrRoot( x int ) int {
  sqrt := Sqrt( float64(x) )
  sqrtInt := int(sqrt)
  if sqrtInt * sqrtInt == x {
    return sqrtInt
  }

  return 0
}

func Sqrt(x float64) (sqrt float64) {
    z := x / 2.0
    for i := 0; i < 5; i++ {
       prevZ := z
       z = z - ( (z*z - x) / (2*z) )
       if math.Abs(z - prevZ) < 0.5 {
           break
       }
    }
    return z
}

func reverseString(str string) string {
  bytes := make([]byte, len(str))
  var j int = len(bytes) - 1
  for i := 0; i <= j; i++ {
    bytes[j-i] = str[i]
  }
  return string(bytes)
}

func processInputFile(path string, seenPtr *map[int]bool, valuePtr *map[int]bool) {
  lines, err := readLines(path)
  if err != nil { panic(err) }

  nextLine, lines := lines[0], lines[1:]
  testCases, err := strconv.Atoi( nextLine )
  if err != nil { panic(err) }

  // fmt.Printf("Test cases: %d\n", testCases)

  for i := 0; i < testCases; i++ {
    nextLine := lines[i]
    fmt.Printf("Case #%d: %s\n", (i+1), lineResult(nextLine, seenPtr, valuePtr))
  }
}

func lineResult(line string, seenPtr *map[int]bool, valuePtr *map[int]bool) string {
  bounds := integersFromString( line )
  return fmt.Sprintf("%d", countFairAndSquareIntegersInBounds( bounds[0], bounds[1], seenPtr, valuePtr ) )
}

func countFairAndSquareIntegersInBounds(startInt, endInt int, seenPtr *map[int]bool, valuePtr *map[int]bool) int {
  count := 0
  for i := startInt; i <= endInt; i++ {
    if integerIsFairAndSquare(i, seenPtr, valuePtr) { count++ }
  }
  return count
}

func main() {
  seenMap := make(map[int]bool)
  valueMap := make(map[int]bool)

  argsWithoutProg := os.Args[1:]
  processInputFile( argsWithoutProg[0], &seenMap, &valueMap)
}
