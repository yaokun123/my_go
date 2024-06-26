package main

import (
	"errors"
	"path/filepath"
	"log"
	"context"
	"os/exec"
	"fmt"
	"io/ioutil"
)

var (
	argsError = errors.New("no input file or out path")
	fileTypeError = errors.New("this file must be in format")
)

type HtmlToPdf struct {
	Commond string
	in string
	out string
	argsMap map[string]string
	prams []string
}

func NewPdf() *HtmlToPdf {
	args := map[string]string{
		"--load-error-handling":"ignore",
		"--footer-center":"第[page]页/共[topage]页",
		"--footer-font-size":"8",
		"-B":"31",
		"-T":"32",
	}
	return &HtmlToPdf{
		Commond:"wkhtmltopdf",
		argsMap:args,
	}
}

func (this *HtmlToPdf) OutFile(input string,outPath string) (string,error) {
	var pdfPath string

	//输入、输出参数不能为空
	if input=="" || outPath==""{
		return pdfPath,argsError
	}

	//判断是否生成pdf文件
	ext := filepath.Ext(outPath)
	if ext != ".pdf"{
		return pdfPath,fileTypeError
	}

	this.in = input
	this.out = outPath

	//构建参数
	this.buildPrams()

	//执行命令
	bytes,err := this.doExec()
	if err != nil{
		return pdfPath,err
	}
	log.Printf("[wkhtmltopdf-stdout]:%s",string(bytes))
	return pdfPath,nil
}

func (this *HtmlToPdf) buildPrams()  {
	for key,val := range this.argsMap{
		this.prams = append(this.prams,key,val)
	}

	//添加输入输出参数
	this.prams = append(this.prams,this.in,this.out)
}

func (this *HtmlToPdf) doExec() ([]byte,error) {
	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd := exec.CommandContext(ctx,this.Commond,this.prams...)
	stdout,err := cmd.StderrPipe()
	if err != nil{
		return nil,err
	}
	defer stdout.Close()
	
	//运行命令
	err = cmd.Start()
	if err != nil{
		return nil,err
	}
	
	bytes,err := ioutil.ReadAll(stdout)
	if err != nil{
		return nil,err
	}
	
	fmt.Println(string(bytes))
	fmt.Println("htmltopdf退出程序中=",cmd.Process.Pid)
	cmd.Wait()
	return bytes,err
}

func main()  {
	pdf := NewPdf()
	url,err := pdf.OutFile("https://shuidi.cn/company-0381a1a7be0c246b71c09a6ca6722a93.html?from_search=1","/tmp/test.pdf")

	if err != nil{
		log.Println(err)
	}
	fmt.Println(url)
}
