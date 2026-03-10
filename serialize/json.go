package serialize

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ProtobufToJSON converts protocol buffer message to JSON string
func ProtobufToJSON(message proto.Message) (string, error) {
	marshaler := protojson.MarshalOptions{
		UseEnumNumbers:  true,  
		EmitUnpopulated: true,   
		Indent:          "  ",   
		UseProtoNames:   true,   
	}

	data, err := marshaler.Marshal(message)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// JSONToProtobuf converts JSON string to protocol buffer message
func JSONToProtobuf(data string, message proto.Message) error {
	return protojson.Unmarshal([]byte(data), message)
}