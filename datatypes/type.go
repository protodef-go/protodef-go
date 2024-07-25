package datatypes

import "github.com/tidwall/gjson"

type Type struct {
	Name string
	Type string
}

func GetTypeFromJSON(name string, option gjson.Result) *Type {
	if option.Type == gjson.String && option.String() == "native" {
		return GetNativeType(name)
	}
	t := GetType(option)
	if t != nil {
		t.Name = name
		return t
	}
	return nil
}

func GetType(d gjson.Result) *Type {
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
		if arr_len > 0 {
			arr_type := arr[0]
			if arr_type.Type == gjson.String {
				switch arr_type.String() {
				case "container":
					t.Type = "container"
				}
			}
		}
	}
	return nil
}
