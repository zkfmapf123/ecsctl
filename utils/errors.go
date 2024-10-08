package utils

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	ERR_NOT_EXISTS_FILE = "file not exists"
)

func SuccessGreen(msg string) {
	fmt.Println(color.GreenString("[success] %s", msg))
	os.Exit(0)
}

func PanicRed(err error) {
	fmt.Println(color.RedString("[err] %s", err.Error()))
	os.Exit(1)
}

func WarningYellow(err error) {
	fmt.Println(color.YellowString("[warning] %s", err.Error()))
	os.Exit(0)
}

func NoticeGreen(err error) {
	fmt.Println(color.GreenString("[notice] %s", err.Error()))
}

func MustCheckError(err error) {
	if err != nil {
		PanicRed(err)
	}
}
