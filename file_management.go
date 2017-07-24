package frango

import (
	"os"

	"github.com/gin-gonic/gin"
)

func CreateFile(filePath string, content string) {
	file, err := os.Create(filePath)
	PrintErr(err)
	file.WriteString(content)
}

func ServeFile(c *gin.Context, filePath string, fileName string) {
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.File(filePath)
}
