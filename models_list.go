package page_generator

import (
	"fmt"
)

type InitFunction func() error

var pgModels map[string]*PageModel // page models list

var serviceName string

func SetDefinitions(init InitFunction, service string) error {
	serviceName = service
	if err := startPaging(); err != nil {
		return err
	}
	creatorsInit()
	return init()
}

func startPaging() error {
	if serviceName == "" {
		return fmt.Errorf("service name is empty ")
	}
	if pgModels == nil {
		pgModels = make(map[string]*PageModel)
	}
	return nil
}

func AddPageModel(key string, models *PageModel) {
	if pgModels == nil {
		pgModels = make(map[string]*PageModel)
	}
	pgModels[key] = models
}
