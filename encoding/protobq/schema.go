package protobq

import (
	"cloud.google.com/go/bigquery"
	"go.einride.tech/protobuf-bigquery/internal/wkt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// InferSchema infers a BigQuery schema for the given proto.Message using default options.
func InferSchema(msg proto.Message) bigquery.Schema {
	return SchemaOptions{}.InferSchema(msg)
}

// SchemaOptions contains configuration options for BigQuery schema inference.
type SchemaOptions struct {
	// UseEnumNumbers converts enum values to INTEGER types.
	UseEnumNumbers bool
}

// InferSchema infers a BigQuery schema for the given proto.Message using options in
// MarshalOptions.
func (o SchemaOptions) InferSchema(msg proto.Message) bigquery.Schema {
	return o.inferMessageSchema(msg.ProtoReflect().Descriptor())
}

// inferMessageSchema infers the BigQuery schema for the given protoreflect.MessageDescriptor.
func (o SchemaOptions) inferMessageSchema(msg protoreflect.MessageDescriptor) bigquery.Schema {
	schema := make(bigquery.Schema, 0, msg.Fields().Len())
	for i := 0; i < msg.Fields().Len(); i++ {
		schema = append(schema, o.inferFieldSchema(msg.Fields().Get(i)))
	}
	return schema
}

func (o SchemaOptions) inferFieldSchema(field protoreflect.FieldDescriptor) *bigquery.FieldSchema {
	if field.IsMap() {
		return o.inferMapFieldSchema(field)
	}
	fieldSchema := &bigquery.FieldSchema{
		Name:     string(field.Name()),
		Type:     o.inferFieldSchemaType(field),
		Repeated: field.IsList(),
	}
	if fieldSchema.Type == bigquery.RecordFieldType && fieldSchema.Schema == nil {
		fieldSchema.Schema = o.inferMessageSchema(field.Message())
	}
	return fieldSchema
}

func (o SchemaOptions) inferFieldSchemaType(field protoreflect.FieldDescriptor) bigquery.FieldType {
	switch field.Kind() {
	case protoreflect.DoubleKind, protoreflect.FloatKind:
		return bigquery.FloatFieldType
	case protoreflect.Int64Kind,
		protoreflect.Uint64Kind,
		protoreflect.Int32Kind,
		protoreflect.Fixed64Kind,
		protoreflect.Fixed32Kind,
		protoreflect.Uint32Kind,
		protoreflect.Sfixed32Kind,
		protoreflect.Sfixed64Kind,
		protoreflect.Sint32Kind,
		protoreflect.Sint64Kind:
		return bigquery.IntegerFieldType
	case protoreflect.BoolKind:
		return bigquery.BooleanFieldType
	case protoreflect.StringKind:
		return bigquery.StringFieldType
	case protoreflect.BytesKind:
		return bigquery.BytesFieldType
	case protoreflect.EnumKind:
		return o.inferEnumFieldType(field)
	case protoreflect.MessageKind, protoreflect.GroupKind:
		switch field.Message().FullName() {
		case wkt.Timestamp:
			return bigquery.TimestampFieldType
		case wkt.Duration:
			return bigquery.FloatFieldType
		case "google.protobuf.DoubleValue",
			"google.protobuf.FloatValue":
			return bigquery.FloatFieldType
		case "google.protobuf.Int32Value",
			"google.protobuf.Int64Value",
			"google.protobuf.UInt32Value",
			"google.protobuf.UInt64Value":
			return bigquery.IntegerFieldType
		case "google.protobuf.BoolValue":
			return bigquery.BooleanFieldType
		case "google.protobuf.StringValue":
			return bigquery.StringFieldType
		case "google.protobuf.BytesValue":
			return bigquery.BytesFieldType
		case wkt.Struct:
			return bigquery.StringFieldType // JSON string
		case wkt.Date:
			return bigquery.DateFieldType
		case "google.type.DateTime":
			return bigquery.TimestampFieldType
		case wkt.LatLng:
			return bigquery.GeographyFieldType
		case wkt.TimeOfDay:
			return bigquery.TimeFieldType
		}
	}
	return bigquery.RecordFieldType
}

func (o SchemaOptions) inferEnumFieldType(field protoreflect.FieldDescriptor) bigquery.FieldType {
	if o.UseEnumNumbers {
		return bigquery.IntegerFieldType
	}
	return bigquery.StringFieldType
}

func (o SchemaOptions) inferMapFieldSchema(field protoreflect.FieldDescriptor) *bigquery.FieldSchema {
	return &bigquery.FieldSchema{
		Name:     string(field.Name()),
		Repeated: true,
		Type:     bigquery.RecordFieldType,
		Schema: bigquery.Schema{
			o.inferFieldSchema(field.MapKey()),
			o.inferFieldSchema(field.MapValue()),
		},
	}
}
