package main

import "fmt"

type Patch struct {
  // TODO: Add button_price, time_cost, button_income
  mat [][]bool
}

func (this *Patch) print() {
  var i, j int
  for i = 0; i < len(this.mat); i++ {
    for j = 0; j < len(this.mat[i]); j++ {
      if this.mat[i][j] {
        fmt.Print("X")
      } else {
        fmt.Print(" ")
      }
    }
    fmt.Println()
  }
}


// Rotates clockwise = transpose + mirror
func (this *Patch) rotate() {
  this.transpose()
  this.mirror()
}

// Creates a new matrix and then replaces the old one.
func (this *Patch) transpose() {
  var old_h = len(this.mat)
  var old_w = len(this.mat[0])

  // init new matrix
  var new_h = old_w
  var new_w = old_h

  var new_mat = make([][]bool, new_h)
  for i := 0; i < len(new_mat); i++ {
    new_mat[i] = make([]bool, new_w)
  }

  // transpose
  for i := 0; i < old_h; i++ {
    for j := 0; j < old_w; j++ {
      new_mat[j][i] = this.mat[i][j]
    }
  }

  this.mat = new_mat
}

func (this *Patch) mirror() {
  var height = len(this.mat)
  var width = len(this.mat[0])

  for i := 0; i < height; i++ {
    for j := 0; j < width / 2; j++ {
      var temp = this.mat[i][j]
      this.mat[i][j] = this.mat[i][width - 1 - j]
      this.mat[i][width - 1 - j] = temp
    }
  }
}