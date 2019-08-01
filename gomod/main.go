package main

import (
	"fmt"
	"gomod/route"

	"github.com/jinzhu/configor"
)

func main() {
	fmt.Println("this is a test form main.main\n使用外部包测试：", configor.Config{})
	fmt.Println("使用项目内包测试：", route.Name{})
}
