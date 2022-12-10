package serializer_test

import (
	"gocpu/sample"
	"log"
	"testing"

	pb "gocpu/pb"
	"gocpu/serializer"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile1 := "../tmp/laptop1.json"
	jsonFile2 := "../tmp/laptop2.json"

	laptop1 := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFileBinary(binaryFile, laptop2)
	require.NoError(t, err)

	log.Println(laptop1)

	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile1)
	require.NoError(t, err)

	err = serializer.WriteProtobufToJSONFile(laptop2, jsonFile2)
	require.NoError(t, err)
}
