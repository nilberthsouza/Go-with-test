package main

import "testing"

func TestSum(t *testing.T){
    got := Sum()
    want := 3

    if got != want{
    t.Errorf("got '%q' want '%q'",got, want)
    }

}
