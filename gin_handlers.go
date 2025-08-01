package page_generator

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
)

func getListPageHandler(pg *PageModel) func(c *gin.Context) {
	return func(c *gin.Context) {
		params := QueryParams{
			Claims: ExtractClaims(c),
			QData:  c.Request.URL.Query(),
			Token:  c.GetHeader("Authorization"),
		}
		c.JSON(200, pg.getDataPage(&params))
	}
}

func getTablePageHandler(pg *PageModel) func(c *gin.Context) {
	return func(c *gin.Context) {
		params := QueryParams{
			Claims: ExtractClaims(c),
			QData:  c.Request.URL.Query(),
			Token:  c.GetHeader("Authorization"),
		}
		c.JSON(200, pg.getOnlyTable(&params))
	}
}

func getCreatePageHandler(pg *PageModel) func(c *gin.Context) {
	return func(c *gin.Context) {
		params := QueryParams{
			Claims: ExtractClaims(c),
			QData:  c.Request.URL.Query(),
			Token:  c.GetHeader("Authorization"),
		}
		c.JSON(200, pg.model.getCreatePage(&params))
	}
}

func getUpdatePageHandler(pg *PageModel) func(c *gin.Context) {
	return func(c *gin.Context) {
		params := QueryParams{
			Claims: ExtractClaims(c),
			QData:  c.Request.URL.Query(),
			Token:  c.GetHeader("Authorization"),
		}
		c.JSON(200, pg.model.getUpdatePage(&params))
	}
}

func getCountItemsHandler(pg *PageModel) func(c *gin.Context) {
	return func(c *gin.Context) {
		params := QueryParams{
			Claims: ExtractClaims(c),
			QData:  c.Request.URL.Query(),
			Token:  c.GetHeader("Authorization"),
		}
		kol, err := pg.pagination.GetCount(&params)
		if err != nil {
			notFound(c, "Not found: "+err.Error())
			return
		}
		mp := make(map[string]int)
		mp["id"] = kol
		c.JSON(200, mp)
	}
}

func postFilterDataHandler(pg *PageModel) func(c *gin.Context) {
	return func(c *gin.Context) {
		filter := reflect.New(pg.filterType.Elem()).Interface()
		list := reflect.New(pg.listType.Elem()).Interface()
		params := QueryParams{
			Claims: ExtractClaims(c),
			QData:  c.Request.URL.Query(),
			Token:  c.GetHeader("Authorization"),
		}
		if err := c.ShouldBind(filter); err != nil {
			badRequest(c, "Bad request: "+err.Error())
			return
		}
		listInt, ok := list.(IGetList)
		if !ok {
			badRequest(c, "internal error: list does not implement Filterable")
			return
		}
		if err := listInt.Filter(filter, &params); err != nil {
			notFound(c, "Not found: "+err.Error())
			return
		}
		c.JSON(200, list)
	}
}

func postCreateDataHandler(pg *PageModel) func(c *gin.Context) {
	return func(c *gin.Context) {
		crPtr := reflect.New(pg.modelType.Elem()).Interface()
		params := QueryParams{
			Claims: ExtractClaims(c),
			QData:  c.Request.URL.Query(),
			Token:  c.GetHeader("Authorization"),
		}
		if err := c.ShouldBind(crPtr); err != nil {
			badRequest(c, "Bad request: "+err.Error())
			return
		}
		crObj, ok := crPtr.(ICreate)
		if !ok {
			badRequest(c, "internal error: list does not implement ICreate")
			return
		}
		if err := crObj.Create(&params); err != nil {
			internalError(c, "Internal error: "+err.Error())
			return
		}
		success(c, "success")
	}
}

func putUpdateDataHandler(pg *PageModel) func(c *gin.Context) {
	return func(c *gin.Context) {
		params := QueryParams{
			Claims: ExtractClaims(c),
			QData:  c.Request.URL.Query(),
			Token:  c.GetHeader("Authorization"),
		}
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
		if err := crObj.Update(&params); err != nil {
			internalError(c, "Internal error: "+err.Error())
			return
		}
		success(c, "success")
	}
}

func deleteDataHandler(pg *PageModel) func(c *gin.Context) {
	return func(c *gin.Context) {
		params := QueryParams{
			Claims: ExtractClaims(c),
			QData:  c.Request.URL.Query(),
			Token:  c.GetHeader("Authorization"),
		}
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
		if err := crObj.Delete(&params); err != nil {
			internalError(c, "Internal error: "+err.Error())
			return
		}
		success(c, "success")
	}
}

func getDefaultListHandler(pg *PageModel) func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error
		var limit, offset int
		lim := c.Query("limit")
		off := c.Query("offset")
		if lim != "" {
			limit, err = strconv.Atoi(lim)
		}
		if off != "" {
			offset, err = strconv.Atoi(off)
		}
		if err != nil {
			badRequest(c, err.Error())
			return
		}
		params := QueryParams{
			Claims: ExtractClaims(c),
			QData:  c.Request.URL.Query(),
			Token:  c.GetHeader("Authorization"),
			Limit:  limit,
			Offset: offset,
		}
		listPtr := reflect.New(pg.listType.Elem()).Interface()
		listInt, ok := listPtr.(IGetList)
		if !ok {
			badRequest(c, "internal error: list does not implement Filterable")
			return
		}
		if err := listInt.GetList(&params); err != nil {
			notFound(c, "Not found: "+err.Error())
			return
		}
		c.JSON(200, listPtr)
	}
}
