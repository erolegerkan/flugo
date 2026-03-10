package common

import (
	"github.com/erolegerkan/flugo/styles"
)

func VerbosePrint(sentence string){
	styles.RenderWithStyle("🟠 Verbose Mode : " + sentence, "verbose")
}

func WarningPrint(sentence string, isPrefixNeeeded bool){
	if isPrefixNeeeded {
		styles.RenderWithStyle("⚠️ Warning : " + sentence,"warning")
	} else {
		styles.RenderWithStyle(sentence,"warning")
	}
}

func ErrorPrint(sentence string){
	styles.RenderWithStyle("❗️ Error : " + sentence, "error")
}

func SuccessPrint(sentence string){
	styles.RenderWithStyle("✅ Success : " + sentence, "success")
}