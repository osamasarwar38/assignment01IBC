// Pass arguments to Satoshi.go (portNumber address of Others , portNumber address of Satoshi)

package main

import (

  a1 "github.com/osamasarwar38/assignment01IBC"
	"fmt"
  "net"
 "encoding/gob"
	"log"
	"os"
//	"strconv"
)


type Transaction struct{
  Sender string
  Reciever string
  Coins int
}


func listen(othersAddress string){
   net.Listen("tcp", ":"+othersAddress)
  fmt.Println("Listening")
}



func printChain(conn net.Conn){
  fmt.Println("Printing the whole chain...")
  var recvdBlock a1.Block
  dec := gob.NewDecoder(conn)
  err := dec.Decode(&recvdBlock)
  if err != nil {
    //handle error
  }

  var trans Transaction
  dec = gob.NewDecoder(conn)
  err = dec.Decode(&trans)
  if err != nil {
    //handle error
    fmt.Println(err)
  }

  fmt.Println(trans)




  a1.ListBlocks (&recvdBlock)
}

func connectToAllPeers(conn net.Conn){
  fmt.Println("Connecting to all Peers...")
  //var connectedPeers string


}


func main() {
    // Pass arguments to Satoshi.go (portNumber address of Others , portNumber address of Satoshi)
othersAddress := os.Args[1]
satoshiAddress := os.Args[2]
conn, err := net.Dial("tcp", "localhost:"+satoshiAddress)
if err != nil {
  //handle error
}

conn.Write([]byte(""+othersAddress))

ln,err := net.Listen("tcp", ":"+othersAddress)
if err!=nil{
  log.Fatal(err)
  fmt.Println("Error")
}
conn, err2 := ln.Accept()
if err2 != nil {
  log.Println(err2)
  }else{
  //Satoshi sends chain and addresses of connectedPeers
  printChain(conn)
//  printChain(conn)
  connectToAllPeers(conn)

}


//
// conn, err3 := ln.Accept()
// if err3 != nil {
//   log.Println(err3)
//   }else{
//
//
// }

}
