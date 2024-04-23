package template

import (
	"github.com/apibrew/apibrew/pkg/api"
	model2 "github.com/apibrew/sso/pkg/model"
)

type testResourceProcessor struct {
	api                    api.Interface
	oauth2ConfigRepository api.Repository[*model2.TestResource]
}

func (t testResourceProcessor) Mapper() Mapper[*model2.TestResource] {
	//TODO implement me
	panic("implement me")
}

func (t testResourceProcessor) Register(entity *model2.TestResource) error {
	//TODO implement me
	panic("implement me")
}

func (t testResourceProcessor) Update(entity *model2.TestResource) error {
	//TODO implement me
	panic("implement me")
}

func (t testResourceProcessor) UnRegister(entity *model2.TestResource) error {
	//TODO implement me
	panic("implement me")
}
