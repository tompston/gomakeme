package response



{{- if ( eq .ProjectInfo.Framework "fiber") }}
import (
	"github.com/gofiber/fiber/v2"
	"fmt"
)

// all of the correct status codes can be found here
// https://pkg.go.dev/net/http?utm_source=gopls#StatusOK
func Response(c *fiber.Ctx, status int,data interface{}, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  status,
		"message": message,
		"data":    data,
	})
}
{{- end }}





{{- if ( eq .ProjectInfo.Framework "gin") }}
import (
	"github.com/gin-gonic/gin"
	"fmt"
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
{{- end }}




const debug = true

// return the acutual error message during debug only
func DbConnErrorMessage(err_msg string) string {

	if debug {
		return fmt.Sprintln("Could not connect to the database: ", err_msg)
	}

	return fmt.Sprintln("Could not connect to the database!")
}