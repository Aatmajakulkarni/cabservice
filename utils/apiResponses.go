package utils

import(
    "github.com/gin-gonic/gin"
    "net/http"
)
type Response struct {
  Status Status `json:"status"`
  Payload interface{} `json:"payload"`
}

type Status struct {
  StatusCode int `json:"status_code"`
  Message string `json:"message"`
}

func SendSuccess(c *gin.Context, data interface{}) {
  payload := Response {
    Payload: data,
    Status: Status {
      StatusCode: SUCCESS,
      Message: "Success",
    },
  }
  // fmt.Printf("response %+v", payload)
	c.JSON(http.StatusOK, payload)
}

func SendError(c *gin.Context, httpStatus int, errorCode int) {
	c.JSON(httpStatus, GetError(errorCode))
}
