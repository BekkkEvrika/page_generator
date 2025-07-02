package page_generator

import "github.com/gin-gonic/gin"

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func success(c *gin.Context, mess string) {
	e := Error{Code: "200", Message: mess}
	c.JSON(200, e)
}

func internalError(c *gin.Context, mess string) {
	e := Error{Code: "500", Message: mess}
	c.JSON(500, e)
}

func notFound(c *gin.Context, mess string) {
	e := Error{Code: "404", Message: mess}
	c.JSON(404, e)
}

func methodNotAllowed(c *gin.Context, mess string) {
	e := Error{Code: "405", Message: mess}
	c.JSON(405, e)
}

func created(c *gin.Context, mess string) {
	e := Error{Code: "201", Message: mess}
	c.JSON(201, e)
}

func badRequest(c *gin.Context, mess string) {
	e := Error{Code: "400", Message: mess}
	c.JSON(400, e)
}
