package main

import "github.com/proxy/socks5"

func main(){
 conf := &socks5.Config{}
 server, err := socks5.New(conf)
 if err != nil {
   panic(err)
 }

// Create SOCKS5 proxy on localhost port 8000
 if err := server.ListenAndServe("tcp", "0.0.0.0:2080"); err != nil {
   panic(err)
 }
}
