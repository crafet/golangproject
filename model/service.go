package model

import (

)

type Service struct {
	Id string
	Name string
	Desc string
	Plans []Plan
	Tags []string
	MetaData map[string]interface{}
}

func NewService(id string,
				name string,
				desc string,
				plans []Plan,
				tags []string) *Service{

	return &Service {
					Id:id,
					Name:name,
					Desc:desc,
					Plans:plans,
					Tags:tags,
					}
}

