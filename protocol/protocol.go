package protocol

import (
	"errors"

	"github.com/protodef-go/protodef-go/datatypes"
	"github.com/protodef-go/protodef-go/namespace"
	"github.com/tidwall/gjson"
)

type Protocol struct {
	Types      []*datatypes.Type
	Namespaces map[string]*namespace.Namespace
}

func (p *Protocol) ReadJSON(d gjson.Result) error {
	types := d.Get("types")

	if !types.Exists() {
		return errors.New("protocol types is missing")
	}

	if !types.IsObject() {
		return errors.New("protocol type is not object")
	}

	for name, option := range types.Map() {
		t := datatypes.GetTypeFromJSON(name, option)
		if t == nil {
			continue
		}
		p.Types = append(p.Types, t)
	}

	p.Namespaces = make(map[string]*namespace.Namespace)
	for name, data := range d.Map() {
		if name == "types" {
			continue
		}

		namespace := &namespace.Namespace{Name: name}
		err := namespace.ReadJSON(data)
		if err != nil {
			return err
		}

		p.Namespaces[name] = namespace
	}
	return nil
}
