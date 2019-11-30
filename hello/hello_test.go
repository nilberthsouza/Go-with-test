package main

import "testing"

func TestHello(t *testing.T){
    
    assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want{
            t.Error("got '%s' want '%s'", got, want)
        }
    
    }
    t.Run("saying hello to people", func(t *testing.T){
        got := Hello("Nilberth")
        want := "Hello, Nilberth"
        assertCorrectMessage(t, got, want)
    })

    T.Run("empty string defaults to 'Wold'", func(t *testing.T){
        got := Hello("")
        want := "Hello, World"
        assertCorrectMessage(t,got,want)
    })
}

