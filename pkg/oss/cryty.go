package oss

import (
	"encoding/base64"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

func (c *AliConfig) EnCrypt(file multipart.File) string {
	buffer := make([]byte, 500000)
	n, _ := file.Read(buffer)
	//base64压缩图片
	base64string := base64.StdEncoding.EncodeToString(buffer[:n])
	// 指定值 + 反转图片base64值
	base64string = c.Salt + Reverse(base64string)
	return base64string
}

func (c *AliConfig) EnCryptByBytes(file []byte) string {
	//base64压缩图片
	base64string := base64.StdEncoding.EncodeToString(file)
	// 指定值 + 反转图片base64值
	base64string = c.Salt + Reverse(base64string)
	return base64string
}

func (c *AliConfig) DeCrypt(url string) (string, error) {
	//获取远端图片
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	// 读取获取的[]byte数据
	data, _ := io.ReadAll(res.Body)
	imageBase64 := base64.StdEncoding.EncodeToString(data)
	decodeString, err := base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		return "", err
	}
	trim := strings.TrimLeft(string(decodeString), c.Salt)
	reverse := Reverse(trim)
	return "data:image/jpeg;base64," + reverse, nil
}

func (c *AliConfig) Base64(url string) (string, error) {
	//获取远端图片
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	// 读取获取的[]byte数据
	data, _ := io.ReadAll(res.Body)
	imageBase64 := base64.StdEncoding.EncodeToString(data)
	return "data:image/jpeg;base64," + imageBase64, nil
}

func (c *AliConfig) DeCryptToByte(url string) ([]byte, error) {
	//获取远端图片
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	// 读取获取的[]byte数据
	data, _ := io.ReadAll(res.Body)
	trim := strings.TrimLeft(string(data), c.Salt)
	reverse := Reverse(trim)
	decodeString, err := base64.StdEncoding.DecodeString(reverse)
	if err != nil {
		return nil, err
	}
	return decodeString, nil
}

// Reverse 字符串反转
func Reverse(str string) string {
	var result []byte
	tmp := []byte(str)
	length := len(str)
	for i := 0; i < length; i++ {
		result = append(result, tmp[length-i-1])
	}
	return string(result)
}

// TODO eryue 客户端不需要缩略图也可以显示，暂时注释
//func (c *AliConfig) Scale(w http.ResponseWriter, imageData []byte) {
//	// 读取获取的[]byte数据
//
//	_ = scale(strings.NewReader(string(imageData)), w, 120, 120, 100)
//}
//
//func scale(in io.Reader, out io.Writer, width, height, quality int) error {
//	origin, fm, err := image.Decode(in)
//	if err != nil {
//		return err
//	}
//	if width == 0 || height == 0 {
//		width = origin.Bounds().Max.X
//		height = origin.Bounds().Max.Y
//	}
//	if quality == 0 {
//		quality = 100
//	}
//	canvas := resize.Thumbnail(uint(width), uint(height), origin, resize.Lanczos3)
//	//return jpeg.Encode(out, canvas, &jpeg.Options{quality})
//	switch fm {
//	case "jpeg":
//		return jpeg.Encode(out, canvas, &jpeg.Options{quality})
//	case "png":
//		return png.Encode(out, canvas)
//	case "gif":
//		return gif.Encode(out, canvas, &gif.Options{})
//	case "bmp":
//		return bmp.Encode(out, canvas)
//	}
//	return nil
//}
