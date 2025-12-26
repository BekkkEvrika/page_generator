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
				val.pageListUrl = "/" + key + "/page/list"
				val.defaultUrl = "/" + key + "s"
				g.GET("/"+key+"/page/list", getListPageHandler(val))
				g.GET("/"+key+"/page/table", getTablePageHandler(val))
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
				val.addUrl = "/" + key + "/page/create"
				val.model.createUrl = "/" + key
				g.GET("/"+key+"/page/create", getCreatePageHandler(val))
				g.POST("/"+key+"/page/create", postCreatePageHandler(val))
				g.POST("/"+key, postCreateDataHandler(val))
			}
			if val.model.update != nil {
				val.editUrl = "/" + key + "/page/update"
				val.model.updateUrl = "/" + key
				g.GET("/"+key+"/page/update", getUpdatePageHandler(val))
				g.POST("/"+key+"/page/update", postUpdatePageHandler(val))
				g.PUT("/"+key, putUpdateDataHandler(val))
			}
		}
	} else {
		return fmt.Errorf("not defined ")
	}
	return nil
}

func GetModelsRoutesGroup(rg *gin.RouterGroup) error {
	if pgModels != nil {
		for key, val := range pgModels {
			fmt.Println("key: ", key)
			if val.getList != nil {
				val.pageListUrl = "/" + key + "/page/list"
				val.defaultUrl = "/" + key + "s"
				rg.GET("/"+key+"/page/list", getListPageHandler(val))
				rg.GET("/"+key+"/page/table", getTablePageHandler(val))
				rg.GET("/"+key+"s", getDefaultListHandler(val))
			}
			if val.pagination != nil {
				val.countUrl = "/" + key + "s/count"
				rg.GET("/"+key+"s/count", getCountItemsHandler(val))
			}
			if val.filterModel != nil {
				val.filterModel.filterUrl = "/" + key + "s/filter"
				rg.POST("/"+key+"s/filter", postFilterDataHandler(val))
			}
			if val.model.delete != nil {
				val.deleteUrl = "/" + key
				rg.DELETE("/"+key, deleteDataHandler(val))
			}
			if val.model.create != nil {
				val.addUrl = "/" + key + "/page/create"
				val.model.createUrl = "/" + key
				rg.GET("/"+key+"/page/create", getCreatePageHandler(val))
				rg.POST("/"+key+"/page/create", postCreatePageHandler(val))
				rg.POST("/"+key, postCreateDataHandler(val))
			}
			if val.model.update != nil {
				val.editUrl = "/" + key + "/page/update"
				val.model.updateUrl = "/" + key
				rg.GET("/"+key+"/page/update", getUpdatePageHandler(val))
				rg.POST("/"+key+"/page/update", postUpdatePageHandler(val))
				rg.PUT("/"+key, putUpdateDataHandler(val))
			}
		}
	} else {
		return fmt.Errorf("not defined ")
	}
	return nil
}
