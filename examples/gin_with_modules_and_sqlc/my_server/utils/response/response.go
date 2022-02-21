package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// all of the correct status codes can be found here
// https://pkg.go.dev/net/http?utm_source=gopls#StatusOK
func Response(c *gin.Context, status int, data interface{}, message string) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

const debug = true

// return the acutual error message during debug only
func DbConnErrorMessage(err_msg string) string {

	if debug {
		return fmt.Sprintln("Could not connect to the database: ", err_msg)
	}

	return fmt.Sprintln("Could not connect to the database!")
}
