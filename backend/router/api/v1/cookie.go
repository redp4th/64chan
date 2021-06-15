package v1

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"nmb/config"
	"nmb/pkg/util"
	"time"

	"github.com/gin-gonic/gin"
)

func quickHash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])[:8]
}

func AssignCookie(c *gin.Context) {
	ip := c.ClientIP()
	cookie := quickHash(fmt.Sprintf("%s%d", ip, time.Now().Unix()))
	token, err := util.GenerateToken(cookie, config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"token":   nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "successfully assigned cookie",
			"token":   token,
		})
	}
}
