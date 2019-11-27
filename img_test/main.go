package main

import (
	"image"
	"os"
	"log"
	"code.google.com/p/graphics-go/graphics"
	"image/png"
	"fmt"
	"math"
	"strconv"
)

func main()  {
	path := "/tmp/yuan.png"
	src,err := LoadImage(path)
	if err != nil{
		log.Fatal(err)
	}

	dst := image.NewRGBA(image.Rect(0,0,250,400))

	angle := AngleToFloat(180)
	err = graphics.Rotate(dst,src,&graphics.RotateOptions{angle})
	if err != nil{
		log.Fatal(err)
	}

	//需要保存的文件
	imgcounter := 123
	saveImage(fmt.Sprintf("/tmp/%03d.png",imgcounter),dst)

}

//LoadImage decodes an image from a file
func LoadImage(path string) (img image.Image,err error) {
	file,err := os.Open(path)
	if err != nil{
		return
	}
	defer file.Close()
	img,_,err = image.Decode(file)
	return
}

func saveImage(path string,img image.Image) (err error) {
	imgfile,err := os.Create(path)
	defer  imgfile.Close()

	//以png的格式保存
	err = png.Encode(imgfile,img)
	if err != nil{
		log.Fatal(err)
	}
	return
}

func AngleToFloat(angle int) float64 {
	pi,_ :=  strconv.ParseFloat(fmt.Sprintf("%.2f",math.Pi),64)
	base := pi/180
	result := base*float64(angle)
	res,_ := strconv.ParseFloat(fmt.Sprintf("%.2f",result),64)
	return res
}

