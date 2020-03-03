package assignment01IBC


import (
  "crypto/sha256"
  "fmt"
)

type Transaction struct{
  sender string
  reciever string
  coins int
}

type Block struct {
//  transaction string
Transaction string
PrevPointer *Block
PrevHash [32] byte

}





func InsertBlock ( transaction string,
   chainHead *Block ) *Block{

     fmt.Println("InsertBlock Called")

     var newBlock Block
     newBlock.Transaction=transaction
     newBlock.PrevPointer=chainHead

     if (chainHead!=nil){
       newBlock.PrevHash =  sha256.Sum256([]byte(chainHead.Transaction))
      } else{ //if GenesisBlock is inserted

          //newBlock.prevHash = nil
      }
      fmt.Println(newBlock.Transaction," inserted! ")
      fmt.Println("")
     return &newBlock
   }




   func ListBlocks (chainHead *Block){

     fmt.Println("ListBlocks Called")

     fmt.Println("Transaction String    " , "   Previous Hash" )

     for chainHead!=nil{
       fmt.Println(chainHead.Transaction,"   ", (chainHead.PrevHash))
       chainHead=chainHead.PrevPointer
     }
     fmt.Println("")

   }




   func ChangeBlock(oldTrans string, newTrans string, chainHead *Block){

     fmt.Println("ChangeBlock Called")

     for chainHead!=nil {

       if (chainHead.Transaction==oldTrans){
         chainHead.Transaction=newTrans
         fmt.Println(oldTrans, " block changed to ",newTrans)
         break
       }

       chainHead=chainHead.PrevPointer
     }

     fmt.Println("")

   }




   func VerifyChain(chainHead *Block){

     fmt.Println("VerifyChain Called")

     var change = false

     for chainHead!=nil {  //verify chains till genesis block

       if (chainHead.PrevHash!=sha256.Sum256([]byte(chainHead.PrevPointer.Transaction))){
         change = true
          break
        }

        if (chainHead.PrevPointer.PrevPointer==nil){
          break
        }

      chainHead=chainHead.PrevPointer

     }

     if (change==true){
         fmt.Println("Hash Changed")
     }else {
         fmt.Println("Hash Not Changed")
     }

     fmt.Println("")

   }
