package pty_interface

import (
  "fmt"
  "github.com/dapplebeforedawn/pty"
  "os/exec"
)

const READSIZE = 1024

func Pty(command string, rows uint16, cols uint16, in_chan chan []byte, out_chan chan []byte) {

  c := exec.Command(command)
  f, err := pty.Start(c)
  if err != nil { panic(err) }

  pty.Setsize( f, rows, cols )

  go func(){
    for bytes := range in_chan {
      // fmt.Print( string(bytes) )
      fmt.Print( bytes )
      f.Write(bytes)
    }
  }()

  go func(){
    for {
      bytes   := make([]byte, READSIZE)
      read, _ := f.Read(bytes)
      out_chan <- bytes[:read]
    }
  }()

  c.Wait()
}

