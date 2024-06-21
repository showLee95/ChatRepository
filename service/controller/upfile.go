package controller

import (
	"chatim/dao/sqlites"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// func Updata(ctx *gin.Context) {
// 	w := ctx.Writer
// 	req := ctx.Request
// 	//获取文件
// 	srcFile, head, err := req.FormFile("file")
// 	if err != nil {
// 		fmt.Println("srcFile err:", err)
// 		return
// 	}

// 	//检查文件后缀
// 	suffix := ".png"
// 	ofilName := head.Filename
// 	tem := strings.Split(ofilName, ".")
// 	if len(tem) > 1 {
// 		suffix = "." + tem[len(tem)-1]
// 	}

// 	//保存文件
// 	fileName := fmt.Sprintf("%d%04d%s", time.Now().Unix(), rand.Int31(), suffix)
// 	dstFile, err := os.Create("./asset/upload/" + fileName)
// 	if err != nil {
// 		fmt.Println("dstFile err:", err)
// 		return
// 	}
// 	_, err = io.Copy(dstFile, srcFile)
// 	if err != nil {
// 		fmt.Println("dstFile err2:", err)
// 	}
// 	url := "./asset/upload/" + fileName
// 	fmt.Println(w, url, "发送成功")
// }

// Updata handles the file upload and saves it to the /data directory
func UpFile(c *gin.Context) {

	// Single file
	file, err := c.FormFile("file")
	// zap.L().Debug(za)
	// fmt.Println(file)
	if err != nil {
		zap.L().Error("File is required:", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	// Save the file to the /data directory
	dst := fmt.Sprintf("%s/%s", viper.GetString("app.path"), filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = sqlites.UpFile(file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"db file error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "filename": file.Filename})
}