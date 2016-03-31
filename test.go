package main

import(
	"fmt"
	"os"
	"github.com/cooljiansir/doge/buse"
)



func main(){
	file,err := os.Create("test.iso")
	if err != nil{
		panic(err)
	}
	bd := buse.Create(file,1024*1024*100)
	dev,err := bd.Connect()
	if err != nil{
		panic(err)
	}
	fmt.Println("blk:",dev)
}
