package page_generator

import (
	"fmt"
	"log"

	"github.com/BekkkEvrika/page_generator/keycloak"
)

type InitFunction func() error

var pgModels map[string]*PageModel // page models list

var serviceName string

var globalDateFormat string

var pageSize int

type PageSetting struct {
	Service          string
	DateFormat       string
	PageSize         int
	KeyCloakSettings *KeyCloakSettings
}

type KeyCloakSettings struct {
	BaseURL    string
	Realm      string
	ClientId   string
	ClientUUID string
	Secret     string
}

func SetDefinitions(init InitFunction, setting PageSetting) error {
	serviceName = setting.Service
	globalDateFormat = setting.DateFormat
	pageSize = setting.PageSize
	if err := startPaging(); err != nil {
		return err
	}
	creatorsInit()
	if err := init(); err != nil {
		return err
	}
	if setting.KeyCloakSettings == nil {
		setting.KeyCloakSettings = &KeyCloakSettings{
			BaseURL:    "https://admin.obukanal.ru/keycloak",
			Realm:      "billing-realm",
			ClientUUID: "d7e5b362-9794-42d2-b410-42a26341cdd6",
			Secret:     "SWcD68YKthmnTWe3itqGEwBon1eP6cla",
		}
	}
	return createKeycloakResources(setting.KeyCloakSettings)
}

func createKeycloakResources(cfg *KeyCloakSettings) error {

	kClient := keycloak.NewKeycloakClient(cfg.BaseURL, cfg.Realm, cfg.ClientId, cfg.ClientUUID, cfg.Secret)
	authorizer := keycloak.NewKeycloakAuthorizer(kClient)

	if err := authorizer.Register(pageModelMapping()); err != nil {
		log.Fatal("authz registration failed:", err)
	}
	return nil
}

func pageModelMapping() keycloak.Manifest {
	man := keycloak.Manifest{Service: serviceName}
	var models []keycloak.Model
	for key, val := range pgModels {
		mod := keycloak.Model{Name: key}
		mod.URIs = []string{
			"/" + serviceName + "/" + key + "/*",
		}
		mod.Scopes = append(mod.Scopes, keycloak.Scope{Name: "list"})
		if val.model.delete != nil {
			mod.Scopes = append(mod.Scopes, keycloak.Scope{Name: "delete"})
		}
		if val.model.create != nil {
			mod.Scopes = append(mod.Scopes, keycloak.Scope{Name: "create"})
		}
		if val.model.update != nil {
			mod.Scopes = append(mod.Scopes, keycloak.Scope{Name: "update"})
		}
		models = append(models, mod)
	}
	man.Models = models
	return man
}

func startPaging() error {
	if serviceName == "" {
		return fmt.Errorf("service name is empty ")
	}
	if globalDateFormat == "" {
		return fmt.Errorf("global date format is empty")
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
