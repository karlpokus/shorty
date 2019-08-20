package main

import (
  "fmt"
  "flag"

  "github.com/karlpokus/shorty"
)

var (
  version = flag.Bool("version", false, "print version and exit")
  host = flag.String("host", "localhost", "listening host")
  port = flag.String("port", "9012", "listening port")
)

func main() {
  flag.Parse()
  if *version {
    fmt.Println(shorty.Version)
    return
  }
  server, err := shorty.New(*host, *port)
  if err != nil {
    fmt.Printf("%s", err)
    return
  }
  err = server.Start()
  if err != nil {
    fmt.Printf("%s", err)
  }
}
