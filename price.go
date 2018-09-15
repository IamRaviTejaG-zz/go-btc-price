package main

import (
  "fmt"
  "flag"
  "github.com/levigross/grequests"
)

func main() {
  fromPtr := flag.String("from", "BTC", "Convert from")
  toPtr := flag.String("to", "USD", "Convert to")
  valuePtr := flag.Int("value", 1, "Value")
  flag.Parse()
  resp, err := grequests.Get("https://blockchain.info/ticker", nil)
  fmt.Println("from: ", *fromPtr)
  fmt.Println("to: ", *toPtr)
  fmt.Println("value: ", *valuePtr)
  // if err != nil {
	//    fmt.Println("Unable to make request!")
  //    fmt.Println(err)
  //  }
  fmt.Println(err)
  fmt.Println(resp)
}
