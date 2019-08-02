package encodings

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/tigrannajaryan/exp-otelproto/encodings/traceprotobuf"
)

// Encode request body into a message of continuous bytes. The message starts with one by
// specifying the length of the RequestHeader, followed by RequestHeader encoded in
// Protobuf format, followed by RequestBody encoded in Protobuf format.
// +--------------------+-------------------------------------------+-----------------------------------------+
// + Header Length Byte | Variable length Header (Protobuf-encoded) | Variable length Body (Protobuf-encoded) |
// +--------------------+-------------------------------------------+-----------------------------------------+
func Encode(
	requestBody *traceprotobuf.RequestBody,
	compression traceprotobuf.CompressionMethod,
) []byte {
	bodyBytes, err := proto.Marshal(requestBody)
	if err != nil {
		log.Fatal("cannot encode:", err)
	}

	switch compression {
	case traceprotobuf.CompressionMethod_NONE:
		break
	case traceprotobuf.CompressionMethod_ZLIB:
		var b bytes.Buffer
		w := zlib.NewWriter(&b)
		w.Write(bodyBytes)
		w.Close()
		bodyBytes = b.Bytes()
	}

	header := &traceprotobuf.RequestHeader{
		Compression: compression,
	}
	headerBytes, err := proto.Marshal(header)
	if err != nil {
		log.Fatal("cannot encode:", err)
	}

	b := bytes.NewBuffer(make([]byte, 0, 1+len(headerBytes)+len(bodyBytes)))
	b.WriteByte(byte(len(headerBytes)))
	b.Write(headerBytes)
	b.Write(bodyBytes)

	return b.Bytes()
}

// Decode a continuous message of bytes into a RequestBody. This function perform the
// reverse of Encode operation.
func Decode(messageBytes []byte) *traceprotobuf.RequestBody {
	headerLen := messageBytes[0]
	headerBytes := messageBytes[1 : headerLen+1]
	bodyBytes := messageBytes[headerLen+1:]

	var header traceprotobuf.RequestHeader
	err := proto.Unmarshal(headerBytes, &header)
	if err != nil {
		log.Fatal("cannot decode:", err)
	}

	switch header.Compression {
	case traceprotobuf.CompressionMethod_NONE:
		break

	case traceprotobuf.CompressionMethod_ZLIB:
		b := bytes.NewBuffer(bodyBytes)
		r, err := zlib.NewReader(b)
		if err != nil {
			log.Fatal("cannot decode:", err)
		}

		bodyBytes, err = ioutil.ReadAll(r)
		if err != nil {
			log.Fatal("cannot decode:", err)
		}
	}

	var body traceprotobuf.RequestBody
	err = proto.Unmarshal(bodyBytes, &body)
	if err != nil {
		log.Fatal("cannot decode:", err)
	}

	return &body
}