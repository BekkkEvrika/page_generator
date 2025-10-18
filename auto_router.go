package page_generator

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetModelsRoutes(g *gin.Engine) error {
	if pgModels != nil {
		for key, val := range pgModels {
			fmt.Println("key: ", key)
			if val.getList != nil {
				val.pageListUrl = "/page/" + key + "/list"
				val.defaultUrl = "/" + key + "s"
				g.GET("/page/"+key+"/list", getListPageHandler(val))
				g.GET("/page/"+key+"/table", getTablePageHandler(val))
				g.GET("/"+key+"s", getDefaultListHandler(val))
			}
			if val.pagination != nil {
				val.countUrl = "/" + key + "s/count"
				g.GET("/"+key+"s/count", getCountItemsHandler(val))
			}
			if val.filterModel != nil {
				val.filterModel.filterUrl = "/" + key + "s/filter"
				g.POST("/"+key+"s/filter", postFilterDataHandler(val))
			}
			if val.model.delete != nil {
				val.deleteUrl = "/" + key
				g.DELETE("/"+key, deleteDataHandler(val))
			}
			if val.model.create != nil {
				val.addUrl = "/page/" + key + "/create"
				val.model.createUrl = "/" + key
				g.GET("/page/"+key+"/create", getCreatePageHandler(val))
				g.POST("/page/"+key+"/create", postCreatePageHandler(val))
				g.POST("/"+key, postCreateDataHandler(val))
			}
			if val.model.update != nil {
				val.editUrl = "/page/" + key + "/update"
				val.model.updateUrl = "/" + key
				g.GET("/page/"+key+"/update", getUpdatePageHandler(val))
				g.POST("/page/"+key+"/update", postUpdatePageHandler(val))
				g.PUT("/"+key, putUpdateDataHandler(val))
			}
		}
	} else {
		return fmt.Errorf("not defined ")
	}
	return nil
}
