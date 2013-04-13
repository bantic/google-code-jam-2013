package main

import (
    "os"
    "bufio"
    "bytes"
    "io"
    "fmt"
    "strings"
    "strconv"
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

func getTestCases(lines []string) (caseCount int, testLines [][]string) {
  caseCount, err := strconv.Atoi(lines[0])
  if err != nil { panic(err) }

  // pop the first line off (it was the number of test cases)
  lines = lines[1:]
  linesPerTest := 4

  for i := 0; i < caseCount; i++ {
    individualTestLines := make([]string, linesPerTest)
    startLine := (linesPerTest * i)

    for j := startLine; j < (startLine + linesPerTest); j++ {
      testLineIdx := j - startLine
      individualTestLines[ testLineIdx ] = lines[j]
    }
    testLines = append(testLines, individualTestLines)
  }


  return
}

func processTest(testNumber int, testLines []string) {
  fmt.Printf("Case #%d: \n", testNumber + 1)
  fmt.Printf("game lines: %s\n", testLines)

  var grid [][]byte

  for i, line := range testLines {
    fmt.Printf("line #%d: %s %s\n", i, line, []byte(line))
    grid = append(grid, []byte(line))
  }

  game := TicTacToeGame{grid:grid}
  fmt.Printf("game is: %s\n",game)
  fmt.Printf("game row 1 is: %s\n",game.row(1))
  fmt.Printf("game col 2 is: %s\n",game.column(2))
  fmt.Printf("game diag left is: %s\n",game.diagonal(DIAG_LEFT))
  fmt.Printf("game diag right is: %s\n",game.diagonal(DIAG_RIGHT))
}

func processInputFile(path string) {
  lines, err := readLines(path)
  if err != nil { panic(err) }

  caseCount, testLinesArray := getTestCases(lines)
  fmt.Println(testLinesArray)
  for i := 0; i < caseCount; i++ {
    processTest(i, testLinesArray[i])
  }
}

type TicTacToeGame struct {
  grid [][]byte
}

func (game *TicTacToeGame) getWinner() byte {
  //# check each column
  //# check both diagonals

  //# check each row
  for i := 0; i < 4; i++ {
    rowResult := game.checkRow(i)
    if rowResult == '.' {
      break
    } else {
      return rowResult
    }
  }

  for j:= 0; j < 4; j++ {
    colResult := game.checkCol(j)
    if colResult == '.' {
      break
    } else {
      return colResult
    }
  }

  return '.'
}

func (game *TicTacToeGame) hasWinner() bool {
  if game.getWinner() == '.' { return false }
  return true
}

func (game *TicTacToeGame) column(colNum int) []byte {

  colBytes := make([]byte, 4)

  for i := 0; i < 4; i++ {
    colBytes = append(colBytes, game.grid[i][colNum])
  }

  return colBytes
}

func (game *TicTacToeGame) row(rowNum int) []byte {
  return game.grid[rowNum]
}

const DIAG_LEFT int  = 1
const DIAG_RIGHT int= 2

func (game *TicTacToeGame) diagonal(diagDir int) []byte {
  diagBytes := make([]byte, 4)

  for i := 0; i < 4; i = i + 1 {
    if diagDir == DIAG_LEFT {
      diagBytes = append(diagBytes, game.grid[i][i])
    }

    if diagDir == DIAG_RIGHT {
      diagBytes = append(diagBytes, game.grid[i][3-i])
    }
  }

  return diagBytes
}

func (game *TicTacToeGame) checkCol(colNum int) byte {
  var prevChar byte
  var currentChar byte

  for i := 0; i < 4; i++ {
    currentChar = game.grid[i][colNum]
    if currentChar == '.' { return '.' }
    if currentChar == 'T' { continue }

    if i == 0 {
      prevChar = currentChar
    }

    if currentChar != prevChar { return '.' }
    prevChar = currentChar
  }
  return currentChar
}

func (game *TicTacToeGame) checkRow(rowNum int) byte {
  var prevChar byte
  var currentChar byte

  for j := 0; j < 4; j++ {
    currentChar = game.grid[rowNum][j]
    if currentChar == '.' { return '.' }
    if currentChar == 'T' { continue }

    if j == 0 {
      prevChar = currentChar
    }

    if currentChar != prevChar { return '.' }
    prevChar = currentChar
  }
  return currentChar
}

func (game *TicTacToeGame) isFilled() bool {
  for i := 0; i < 4; i++ {
    for j := 0; j < 4; j++ {
      if game.grid[i][j] == '.' { return false }
    }
  }
  return true
}

func (game *TicTacToeGame) isDraw() bool {
  return game.isFilled() && !game.hasWinner()
}

func main() {
  argsWithoutProg := os.Args[1:]
  processInputFile( argsWithoutProg[0] )
}
