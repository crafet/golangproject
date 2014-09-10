package controller

import (
	"net/http"
	"model"
	"encoding/json"
	"fmt"
)

type Services [] model.Service

type CatalogService struct {
	catalog_service Services
}


func(cs *CatalogService) marshalJson()([]byte, error) {
	return json.Marshal(cs.catalog_service)	
}

//相应CF 发起的catalog service 查询服务
func (cs *CatalogService) HandleCatalogService(w http.ResponseWriter, r *http.Request) {
	resp, err := cs.marshalJson()
	fmt.
	
}
