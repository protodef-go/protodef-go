package datatypes

import (
	"github.com/tidwall/gjson"
)

type Type struct {
	Name     string
	TypeName string

	Extras any
}

func GetTypeFromJSON(name string, option gjson.Result) *Type {
	if option.Type == gjson.String && option.String() == "native" {
		return GetNativeType(name)
	}
	t := GetType(name, option)
	if t != nil {
		t.Name = name
		return t
	}
	return nil
}

func GetType(name string, d gjson.Result) *Type {
	var t *Type
	if d.Type == gjson.String {
		t = GetNativeType(d.String())
		if t != nil {
			return t
		}
	}

	if d.IsArray() {
		t = &Type{}
		arr := d.Array()
		arr_len := len(arr)
		if arr_len == 2 {
			arr_type := arr[0]
			if arr_type.Type == gjson.String {
				t.TypeName = arr_type.String()
				switch t.TypeName {
				case "container":
					t.Extras = &Container{}

					t.Extras.(*Container).ReadJSON(arr[1])
				}
			}
		}
		return t
	}
	return nil
}
