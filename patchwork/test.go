package main

import "fmt"

func main() {
  // Define Tetris S shape:
  //
  //   X
  //   XX
  //    X
  //
  var mat = [][]bool {
    {true, false},
    {true, true},
    {false, true}}

  var patch = Patch{mat}
  patch.print()
  fmt.Println()

  patch.rotate()
  patch.print()
  fmt.Println()

  patch.mirror()
  patch.print()
  fmt.Println()

  patch.rotate()
  patch.print()
  fmt.Println()
}
