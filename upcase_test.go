package main

import (
  "testing"
)

func TestUpcasing(t *testing.T) {
  ret := UpCaseStrings("Hi my name is mikko")
  if (ret != "Hi my Name is Mikko") {
    t.Fatalf("Bad upcasing %s", ret)
  }

  ret = UpCaseStrings("please try   with   spaces")
  if (ret != "Please try   With   Spaces") {
    t.Fatalf("Bad upcasing %s", ret)
  }
}
