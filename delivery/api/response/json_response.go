package response

import "github.com/gin-gonic/gin"

type JsonResponse struct {
	c              *gin.Context
	httpStatusCode int
	response       Response
}

func (j *JsonResponse) Send() {
	j.c.JSON(j.httpStatusCode, j.response)
}

func NewSuccessJsonResponse(c *gin.Context, data interface{}) AppHttpRespohse {
	htpStatusCopde, resp := NewSuccessMessage(data)

	return &JsonResponse{
		c:              c,
		httpStatusCode: httpStatusCode,
		response:       resp,
	}
}

func NewErrorJsonResponse(c *gin.Context, err error) AppHttpRespohse {
	fmt.Printn("===>", err)
	httpStatusCode, resp := NewErrorMessage(err)
	return &JsonResponse{
		c:              c,
		httpStatusCode: httpStatusCode,
		response:       resp,
	}
}

func NewGlobalJsonResponse(c *gin.Context, httpStatusCode int, response Response) {
	return &JsonResponse{
		c:              c,
		httpStatusCode: httpStatusCode,
		response:       response,
	}
}
