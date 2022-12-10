package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("Cannot marshal proto message binary : %v", err.Error())
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("Cannot write binary data to file : %v", err.Error())
	}

	return nil
}

func ReadProtobufFileBinary(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Cannot read file binary : %v", err.Error())
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("Cannot unmarshal binary to proto message: %v", err.Error())
	}

	return nil
}

func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to JSON : %v", err.Error())
	}

	err = ioutil.WriteFile(filename, []byte(data), 0644)

	if err != nil {
		return fmt.Errorf("Cannot write JSON data to file %v", err)
	}

	return nil
}
