package datatypes

import (
	"fmt"

	"github.com/tidwall/gjson"
)

type ContainerField struct {
	Name string
	Type *Type
}

type Container struct {
	Name   string
	Fields []*ContainerField
}

func (cf *Container) ReadJSON(d gjson.Result) error {
	if !d.IsArray() {
		return fmt.Errorf("container %s data is not array", cf.Name)
	}

	for _, value := range d.Array() {
		name := value.Get("name").String()
		typeData := value.Get("type")
		cf.Fields = append(cf.Fields, &ContainerField{
			Name: name,
			Type: GetTypeFromJSON(name, typeData),
		})

	}
	return nil
}
