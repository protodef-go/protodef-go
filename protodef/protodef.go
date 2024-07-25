package protodef

import (
	"errors"
	"io"
	"os"

	"github.com/protodef-go/protodef-go/protocol"
	"github.com/tidwall/gjson"
)

func ReadProtocolFile(filename string) (*protocol.Protocol, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	data := gjson.ParseBytes(b)
	if !data.IsObject() {
		return nil, errors.New("protocol file is obviously wrong")
	}

	protocol := &protocol.Protocol{}
	err = protocol.ReadJSON(data)
	if err != nil {
		return nil, err
	}

	return protocol, nil
}
