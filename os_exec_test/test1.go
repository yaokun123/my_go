package main

import (
	"os/exec"
	"fmt"
	"strings"
	"bytes"
	"time"
	"io/ioutil"
	"context"
)

func main()  {
	//1、使用LookPath函数查找系统命令所在位置
	path,err :=exec.LookPath("ls")
	if err != nil{
		fmt.Println(1,err)
	}else{
		fmt.Println(1,path) //   /bin/ls
	}

	//2、使用Command返回cmd结构体
	cmd := exec.Command("tr","a-z","A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err2 := cmd.Run()
	if err != nil{
		fmt.Println(2,err2)
	}else{
		fmt.Println(2,out.String())
	}

	//3、使用combinedOutput运行命令，并返回标准输出和标准错误
	arg := []string{"nofile","hasfile"}
	cmd3 := exec.Command("mv",arg...)
	out3,_ := cmd3.CombinedOutput()
	fmt.Println(3,string(out3))//mv: rename nofile to hasfile: No such file or directory

	//4、使用Output运行命令并返回其标准输出,不包含错误信息
	cmd4 := exec.Command("mv",arg...)
	out4,_ := cmd4.Output()
	fmt.Println(4,string(out4))//空，没有返回错误信息

	//5、使用Run执行并等待（捕捉不到输出）
	t5_start := time.Now()
	cmd5 := exec.Command("sleep","5")
	err5 := cmd5.Run()//执行并等待
	t5_end := time.Now()
	fmt.Println(5,"执行并等待",t5_end.Sub(t5_start))
	if err5 != nil{
		fmt.Println(5,err5)
	}

	//6、使用Run执行但是不等待（捕捉不到输出）
	t6_start := time.Now()
	cmd6 := exec.Command("sleep","5")
	err6 := cmd6.Start()//执行但是不等待
	t6_end := time.Now()
	fmt.Println(6,"执行但是不等待",t6_end.Sub(t6_start))
	if err6 != nil{
		fmt.Println(6,err6)
	}

	//7、使用Wait等待执行的命令
	t7_start := time.Now()
	cmd7 := exec.Command("sleep","5")
	err7 := cmd7.Start()//执行但是不等待
	t7_end1 := time.Now()
	fmt.Println(7,"执行但是不等待",t7_end1.Sub(t7_start))

	cmd7.Wait()//等待命令执行完毕
	t7_end2 := time.Now()
	fmt.Println(7,"加上等待",t7_end2.Sub(t7_start))

	if err7 != nil{
		fmt.Println(1,err7)
	}


	//8、使用StderrPipe连接到command的标准错误
	//这里记得要保证在cmd命令结束前来读取内容，不然会读不到（一般结合Start与Wait来保证）
	cmd8 := exec.Command("mv","nofile","hasfile")
	stderr,_ := cmd8.StderrPipe()//获取标准错误
	cmd8.Start()
	output,_ := ioutil.ReadAll(stderr)
	cmd8.Wait()
	fmt.Println(8,string(output))


	//9、使用StdinPipe连接到command标准输入的管道pipe
	cmd9 := exec.Command("cat")
	stdin,_ := cmd9.StdinPipe()
	stdin.Write([]byte("/tmp/tmp.txt"))
	stdin.Close()
	output9,_ := cmd9.CombinedOutput()
	fmt.Println(9,string(output9))

	//10、使用StdoutPipe连接到命令启动时标准输出的管道
	cmd10 := exec.Command("ls")
	stdout10,_ := cmd10.StdoutPipe()
	cmd10.Start()
	content,_ := ioutil.ReadAll(stdout10)
	fmt.Println(10,string(content))


	//11、CommandContext的使用示例
	tt1 := time.Now()
	ctx,cancel := context.WithTimeout(context.Background(),1*time.Second)
	defer cancel()

	cmd11 := exec.CommandContext(ctx,"sleep","5")
	cmd11.CombinedOutput()
	tt2 := time.Now()//只用了一秒

	fmt.Println(11,tt2.Sub(tt1))
}
