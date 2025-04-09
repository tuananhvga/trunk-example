package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-feature/go-sdk/openfeature"
	"github.com/tuananhvga/trunk-example/internal/featureflag"
)

func getRequestContext(c *gin.Context) openfeature.EvaluationContext {
	userId := c.Query("userId")
	if userId == "" {
		userId = "anonymous"
	}
	attributes := map[string]any{
		"foo":    "bar",
		"ip":     c.ClientIP(),
		"userId": userId,
	}
	fmt.Println("attributes:", attributes)
	return openfeature.NewEvaluationContext("default", attributes)
}

func main() {
	r := gin.Default()
	client := featureflag.New(featureflag.Flagsmith)
	r.GET("/boolean", func(c *gin.Context) {
		details, err := client.BooleanValue(c, "boolean-flag", false, getRequestContext(c))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.String(http.StatusOK, fmt.Sprintln("boolean:", details))
	})
	r.GET("/features", func(c *gin.Context) {
		details, err := client.ObjectValue(c, "features", "", getRequestContext(c))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Printf("features: %T\n", details)
		s, ok := details.([]any)
		if ok {
			for _, v := range s {
				fmt.Printf("features: %v\n", v)
			}
		}
		c.String(http.StatusOK, fmt.Sprintln("features:", details))
	})
	r.Run("0.0.0.0:1977") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
