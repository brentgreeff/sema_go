package main

import "testing"

func TestSum(t *testing.T) {
  want := 7
  if got := add(2, 5); got != want {
    t.Errorf("add(2, 5) = %d, want %d", got, want)
  }

  want2 := 102
  if got := add(2, 100); got != want2 {
    t.Errorf("add(2, 100) = %d, want %d", got, want2)
  }

  want3 := 322
  if got := add(222, 100); got != want3 {
    t.Errorf("add(2, 5) = %d, want %d", got, want3)
  }
}

func TestProduct(t *testing.T) {
  if multiply(2, 5) != 10 {
    t.Fail()
  }
  if multiply(2, 100) != 200 {
    t.Fail()
  }
  if multiply(222, 3) != 666 {
    t.Fail()
  }
}
