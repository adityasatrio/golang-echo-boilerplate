package cache

import (
	"myapp/ent"
)

type CacheValue struct {
	//register model / entity here for parsing in cache
	//SystemParameter *ent.System_parameter
	Pet  *ent.Pet
	User *ent.User
}

type newType []CacheValue

/*func (m *CacheValue) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m *CacheValue) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}*/

/*
func (s *newType) MarshalBinary() ([]byte, error) {
	r, err := msgpack.Marshal((*[]CacheValue)(s))
	if err != nil {
		log.Fatal("MarshalBinary", err)
	}

	return r, nil
}

func (s *newType) UnmarshalBinary(data []byte) error {
	err := msgpack.Unmarshal(data, (*[]CacheValue)(s))
	if err != nil {
		log.Fatal("UnmarshalBinary", err)
	}

	return err
}*/

/*
import "github.com/vmihailenco/msgpack/v5"

type baseType struct {
	Tag string
}

type newType []baseType

func (s *newType) MarshalBinary() ([]byte, error) {
	return msgpack.Marshal((*[]baseType)(s))
}

func (s *newType) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, (*[]baseType)(s))
}

*/
