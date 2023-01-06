package schema

import (
	"encoding/json"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"myapp/helper"
)

// SystemParameter holds the schema definition for the SystemParameter entity.
type SystemParameter struct {
	ent.Schema
}

// Fields of the SystemParameter.
func (SystemParameter) Fields() []ent.Field {
	schema := []ent.Field{
		field.String("key").NotEmpty().Unique(),
		field.String("value").NotEmpty(),
	}

	return helper.InitBaseSchema(schema)
}

// Edges of the SystemParameter.
func (SystemParameter) Edges() []ent.Edge {
	return nil
}

func (SystemParameter) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("key"),
		index.Fields("value"),
	}
}

func (s *SystemParameter) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
	//return msgpack.Marshal(s)
	/*r, err := msgpack.Marshal((*[]CacheValue)(s))
	if err != nil {
		log.Fatal("MarshalBinary", err)
	}

	return r, nil*/
}

//func (s *SystemParameter) UnmarshalBinary(data []byte) error {
//	return json.Unmarshal(data, &s)
//	//return msgpack.Unmarshal(data, s)
//	/*err := msgpack.Unmarshal(data, (*[]CacheValue)(s))
//	if err != nil {
//		log.Fatal("UnmarshalBinary", err)
//	}
//
//	return err*/
//}
