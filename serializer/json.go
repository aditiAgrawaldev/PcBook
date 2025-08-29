package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) (string, error) {
	marshaler := protojson.MarshalOptions{
		UseEnumNumbers:  false, // Enums as strings
		EmitUnpopulated: true,  // Show fields even if empty
		Indent:          "  ",  // Pretty print
		UseProtoNames:   true,  // Use original field names (snake_case)
	}

	data, err := marshaler.Marshal(message)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
