package main

import (
  "testing"
)

func TestSuccess(t *testing.T) {
  a := 2
  b := 2
  if a != b {
    t.Error("Expected 2, got ", a)
  }

  c := a + b
  if c != 4 {
    t.Error("Expected 4, got ", c)
  }  
}

