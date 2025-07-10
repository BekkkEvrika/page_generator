package page_generator

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

func getListPageHandler(pg *PageModel) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, pg.getDataPage())
	}
}

func getTablePageHandler(pg *PageModel) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, pg.getOnlyTable())
	}
}

func getCreatePageHandler(pg *PageModel) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, pg.getCreatePage())
	}
}

func getUpdatePageHandler(pg *PageModel) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, pg.getUpdatePage())
	}
}

func postFilterDataHandler(pg *PageModel) func(c *gin.Context) {
	return func(c *gin.Context) {
		filter := reflect.New(pg.filterType).Interface()
		list := reflect.New(pg.listType.Elem()).Interface()
		if err := c.ShouldBind(filter); err != nil {
			badRequest(c, "Bad request: "+err.Error())
			return
		}
		listInt, ok := list.(IGetList)
		if !ok {
			badRequest(c, "internal error: list does not implement Filterable")
			return
		}
		if err := listInt.Filter(filter); err != nil {
			notFound(c, "Not found: "+err.Error())
			return
		}
		c.JSON(200, list)
	}
}

func postCreateDataHandler(pg *PageModel) func(c *gin.Context) {
	return func(c *gin.Context) {
		claims := ExtractClaims(c)
		crPtr := reflect.New(pg.modelType.Elem()).Interface()
		if err := c.ShouldBind(crPtr); err != nil {
			badRequest(c, "Bad request: "+err.Error())
			return
		}
		crObj, ok := crPtr.(ICreate)
		if !ok {
			badRequest(c, "internal error: list does not implement ICreate")
			return
		}
		if err := crObj.Create(claims); err != nil {
			internalError(c, "Internal error: "+err.Error())
			return
		}
		success(c, "success")
	}
}

func putUpdateDataHandler(pg *PageModel) func(c *gin.Context) {
	return func(c *gin.Context) {
		claims := ExtractClaims(c)
		crPtr := reflect.New(pg.modelType.Elem()).Interface()
		if err := c.ShouldBind(crPtr); err != nil {
			badRequest(c, "Bad request: "+err.Error())
			return
		}
		crObj, ok := crPtr.(IUpdate)
		if !ok {
			badRequest(c, "internal error: list does not implement IUpdate")
			return
		}
		if err := crObj.Update(claims); err != nil {
			internalError(c, "Internal error: "+err.Error())
			return
		}
		success(c, "success")
	}
}

func deleteDataHandler(pg *PageModel) func(c *gin.Context) {
	return func(c *gin.Context) {
		claims := ExtractClaims(c)
		crPtr := reflect.New(pg.modelType.Elem()).Interface()
		if err := c.ShouldBind(crPtr); err != nil {
			badRequest(c, "Bad request: "+err.Error())
			return
		}
		crObj, ok := crPtr.(IDelete)
		if !ok {
			badRequest(c, "internal error: list does not implement IDelete")
			return
		}
		if err := crObj.Delete(claims); err != nil {
			internalError(c, "Internal error: "+err.Error())
			return
		}
		success(c, "success")
	}
}

func getDefaultListHandler(pg *PageModel) func(c *gin.Context) {
	return func(c *gin.Context) {
		claims := ExtractClaims(c)
		listPtr := reflect.New(pg.listType.Elem()).Interface()
		listInt, ok := listPtr.(IGetList)
		if !ok {
			badRequest(c, "internal error: list does not implement Filterable")
			return
		}
		if err := listInt.GetList(claims); err != nil {
			notFound(c, "Not found: "+err.Error())
			return
		}
		c.JSON(200, listPtr)
	}
}
