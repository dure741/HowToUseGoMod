# HowToUseGoMod
#大概介绍一下go mod的使用方法
#使用go mod 管理项目，就不需要非得把项目放到GOPATH指定目录下，你可以在你磁盘的任何位置新建一个项目，#比如：
#  现在我在centos7系统中root用户下创建新目录（该目录不是GOPATH目录）
mkdir ~/gomod

#并进入目录
cd ~/gomod

#再次我们需要在centos中使用命令：
#go mod init [一个自定义的初始化名字，不过建议和当前目录或项目名一致]
go mod init test

#这时候会出现一个提示：
go: creating new go.mod: module test

#该目录下会多出一个目录:
go.mod
#包含go.mod文件的目录也被称为模块根，也就是说，go.mod 文件的出现定义了它所在的目录为一个模块。

#在目录下测试，创建main.go和min所需的包目录
touch main.go

#main.go:
package main                                                 
import (
   "fmt"
   "gomod/route" //这里的包路径式相对于模块根目录的路径，在包前还要加上模块根目录
 
   "github.com/jinzhu/configor"
 )
 
 func main() {
   fmt.Println("this is a test form main.main\n使用外部包测试>    ：", configor.Config{})
   fmt.Println("使用项目内包测试：", route.Name{})
}


#在模块根目录创建包目录route
makedir route

#在目录内编辑 route包文件test.go
cd route

#test.go:
package route                                                
 
type Name struct {
  i int
}

#返回项目根目录gomod
go run main.go
#输出：
使用外部包测试： {  false false false false 0s <nil> false}
使用项目内包测试： {0}
