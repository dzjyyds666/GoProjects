package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/labstack/echo/v4"
)

type reszieParam struct {
	Width     int `json:"width"`
	Height    int `json:"height"`
	Mode      int `json:"mode"`
	Quality   int `json:"quality"`
	X         int `json:"x"`
	Y         int `json:"y"`
	Interlace int `json:"interlace"`
}

func DealWhitImage(c echo.Context) error {

	param := &reszieParam{}

	param.Width, _ = strconv.Atoi(c.QueryParam("width"))
	param.Height, _ = strconv.Atoi(c.QueryParam("height"))
	param.Quality, _ = strconv.Atoi(c.QueryParam("quality"))
	param.Mode, _ = strconv.Atoi(c.QueryParam("mode"))
	param.X, _ = strconv.Atoi(c.QueryParam("x"))
	param.Y, _ = strconv.Atoi(c.QueryParam("y"))
	param.Interlace, _ = strconv.Atoi(c.QueryParam("interlace"))

	if param.Quality == 0 {
		param.Quality = 100
	}

	fmt.Println(param)
	// 获取前端请求中的图片信息
	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, "请求参数错误")
	}

	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusBadRequest, "文件打开失败")
	}
	defer src.Close()

	//创建一个缓冲区存储文件数据
	inputBuffer := new(bytes.Buffer)
	if _, err := io.Copy(inputBuffer, src); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to read file"})
	}

	//创建一个缓冲区来接收文件缩放后的数据
	outputBuffer := new(bytes.Buffer)

	switch param.Mode {
	case 0:
		// 缩放图片
		resizeImage(inputBuffer, outputBuffer, param.Width, param.Height, param.Quality)
	case 1:
		// 裁剪图片
		cropImage(inputBuffer, outputBuffer, param.Width, param.Height, param.X, param.Y, param.Quality)
	case 2:
		if param.Interlace == 1 {
			interlace(inputBuffer, outputBuffer)
		} else {
			return c.String(http.StatusBadRequest, "参数错误")
		}
	}
	return c.Stream(http.StatusOK, "image/jpeg", outputBuffer)
}

// 裁切图片
func cropImage(inputBuffer, outputBuffer *bytes.Buffer, width, height, x, y, quality int) (string, error) {
	if quality == 0 {
		quality = 100
	}
	// 构造 ImageMagick 的 `convert` 命令
	// `-crop {width}x{height}+{x}+{y}` 用于定义裁切区域
	cmd := exec.Command("convert", "-", "-crop", fmt.Sprintf("%dx%d+%d+%d", width, height, x, y), "-quality", fmt.Sprintf("%d", quality), "-")
	cmd.Stdin = inputBuffer
	cmd.Stdout = outputBuffer
	//执行命令并捕捉可能的错误
	err := cmd.Run()

	if err != nil {
		return "失败", err
	}
	return "成功", nil
}

// 缩放图片
func resizeImage(inputBUffer, outputBuffer *bytes.Buffer, width, height, quality int) (string, error) {

	fmt.Println("width:", width, "height:", height, "quality:", quality)
	if quality == 0 {
		quality = 100
	}
	cmd := exec.Command("convert", "-", "-resize", fmt.Sprintf("%dx%d", width, height), "-quality", fmt.Sprintf("%d", quality), "-")
	cmd.Stdin = inputBUffer
	cmd.Stdout = outputBuffer

	err := cmd.Run()

	if err != nil {
		return "失败", err
	}
	return "成功", nil
}

// 渐进显示
func interlace(inputBuffer, outputBuffer *bytes.Buffer) (string, error) {
	cmd := exec.Command("convert", "-", "-interlace", "plane", "jpeg:-")
	cmd.Stdin = inputBuffer
	cmd.Stdout = outputBuffer

	err := cmd.Run()
	if err != nil {
		fmt.Println("err:", err)
		return "失败", err
	}
	return "成功", nil
}
