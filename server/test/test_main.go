package main

import (
	"test/common"
)

func main() {
	token := common.TestLogin()
	common.TestManualEntryQues(token)
}
