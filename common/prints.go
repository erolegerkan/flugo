package common

import "fmt"

func VerbosePrint(sentence string){
	fmt.Println("🟠 Verbose Mode : " + sentence)
}

func WarningPrint(sentence string){
	fmt.Println("⚠️ Warning : " + sentence)
}

func ErrorPrint(sentence string){
	fmt.Println("❗️ Error : " + sentence)
}

func SuccessPrint(sentence string){
	fmt.Println("✅ Success : " + sentence)
}