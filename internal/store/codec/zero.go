package codec

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	zeroBool    bool
	zeroInt32   int32
	zeroInt64   int64
	zeroUint32  uint32
	zeroUint64  uint64
	zeroFloat32 float32
	zeroFloat64 float64
	zeroString  string
	zeroEnum    protoreflect.EnumNumber
	zeroBytes   = make([]byte, 0)
	zeroSlice   = make([]interface{}, 0)
	zeroMap     = make(map[string]interface{}, 0)
)

func zeroValue(fd protoreflect.FieldDescriptor, v protoreflect.Value) any {
	if fd.Cardinality() == protoreflect.Repeated {
		if fd.IsList() {
			return zeroSlice
		}
		if fd.IsMap() {
			return zeroMap
		}
	}
	switch fd.Kind() {
	case protoreflect.BoolKind:
		return zeroBool
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return zeroInt32
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return zeroInt64
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return zeroUint32
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return zeroUint64
	case protoreflect.FloatKind:
		return zeroFloat32
	case protoreflect.DoubleKind:
		return zeroFloat64
	case protoreflect.StringKind:
		return zeroString
	case protoreflect.BytesKind:
		return zeroBytes
	case protoreflect.EnumKind:
		return zeroEnum
	default:
		panic(fmt.Sprintf("unsupported zero value: kind=%v", fd.Kind()))
	}
}
