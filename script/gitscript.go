package main
import (
	"fmt"
	"os"
	"os/exec"
)

func main()  {
	//合并分支
	gitMerge("dev","dev1")
	gitMerge("master","dev")
	gitCheckout("master")
}

//参数初始化

//执行git命令
//切换分支
func gitCheckout(branchName string){
	//指定的命令必须是 git checkout k 依次传入，不能分开传入
	c := exec.Command("git","checkout",branchName)
	c.Stdout = os.Stdout
	err := c.Run()
	if err !=nil{
		fmt.Print("切换分支"+branchName+"错误",err)
	}else{
		fmt.Printf("切换到分支:%s\n",branchName)
	}
} 

//合并分支
func gitMerge(masterName string,branchName string){
	//切换分支
	gitCheckout(masterName)
	//合并分支
	c := exec.Command("git","merge",branchName)
	c.Stdout = os.Stdout
	err := c.Run()
	if err !=nil{
		fmt.Print("合并分支"+branchName+"错误",err)
	}else{
		fmt.Printf("%s分支合并%s分支成功\n",masterName,branchName)
	}
	//切回原分支
	// gitCheckout(branchName)
}