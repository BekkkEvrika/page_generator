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
				val.pageListUrl = "/" + key + "/list/page"
				val.defaultUrl = "/" + key + "/list/data"
				g.GET("/"+key+"/list/page", getListPageHandler(val))
				g.GET("/"+key+"/list/table", getTablePageHandler(val))
				g.GET("/"+key+"/list/data", getDefaultListHandler(val))
			}
			if val.pagination != nil {
				val.countUrl = "/" + key + "/list/count"
				g.GET("/"+key+"/list/count", getCountItemsHandler(val))
			}
			if val.filterModel != nil {
				val.filterModel.filterUrl = "/" + key + "/list/filter"
				g.POST("/"+key+"/list/filter", postFilterDataHandler(val))
			}
			if val.model.delete != nil {
				val.deleteUrl = "/" + key + "/delete/data"
				g.DELETE("/"+key+"/delete/data", deleteDataHandler(val))
			}
			if val.model.create != nil {
				val.addUrl = "/" + key + "/create/page"
				val.model.createUrl = "/" + key + "/create/data"
				g.GET("/"+key+"/create/page", getCreatePageHandler(val))
				g.POST("/"+key+"/create/page", postCreatePageHandler(val))
				g.POST("/"+key+"/create/data", postCreateDataHandler(val))
			}
			if val.model.update != nil {
				val.editUrl = "/" + key + "/update/page"
				val.model.updateUrl = "/" + key + "/update/data"
				g.GET("/"+key+"/update/page", getUpdatePageHandler(val))
				g.POST("/"+key+"/update/page", postUpdatePageHandler(val))
				g.PUT("/"+key+"/update/data", putUpdateDataHandler(val))
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
				val.pageListUrl = "/" + key + "/list/page"
				val.defaultUrl = "/" + key + "/list/data"
				rg.GET("/"+key+"/list/page", getListPageHandler(val))
				rg.GET("/"+key+"/list/table", getTablePageHandler(val))
				rg.GET("/"+key+"/list/data", getDefaultListHandler(val))
			}
			if val.pagination != nil {
				val.countUrl = "/" + key + "/list/count"
				rg.GET("/"+key+"/list/count", getCountItemsHandler(val))
			}
			if val.filterModel != nil {
				val.filterModel.filterUrl = "/" + key + "/list/filter"
				rg.POST("/"+key+"/list/filter", postFilterDataHandler(val))
			}
			if val.model.delete != nil {
				val.deleteUrl = "/" + key + "/delete/data"
				rg.DELETE("/"+key+"/delete/data", deleteDataHandler(val))
			}
			if val.model.create != nil {
				val.addUrl = "/" + key + "/create/page"
				val.model.createUrl = "/" + key + "/create/data"
				rg.GET("/"+key+"/create/page", getCreatePageHandler(val))
				rg.POST("/"+key+"/create/page", postCreatePageHandler(val))
				rg.POST("/"+key+"/create/data", postCreateDataHandler(val))
			}
			if val.model.update != nil {
				val.editUrl = "/" + key + "/update/page"
				val.model.updateUrl = "/" + key + "/update/data"
				rg.GET("/"+key+"/update/page", getUpdatePageHandler(val))
				rg.POST("/"+key+"/update/page", postUpdatePageHandler(val))
				rg.PUT("/"+key+"/update/data", putUpdateDataHandler(val))
			}
		}
	} else {
		return fmt.Errorf("not defined ")
	}
	return nil
}
