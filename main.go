package main 

import (
	"fmt"
)

func hashLetter(key int, letter string)(result string){
	
}

func encrypt(key int, text string)(result string){

}

func decrypt(key int, encrypted string)(result string){

}

func main(){
	plainText := "Hello World"
	fmt.Println("Plain text : ",plainText)
	encrypted := encrypt(5, plainText)
	fmt.Println("Encrypted : ",encrypted)
	decrypted := decrypt(5, encrypted)
	fmt.Println("Decrypted : ",decrypted)
}