package cache

import (
	"encoding/json"
	"myapp/ent"
)

type CacheValue struct {
	//register model / entity here for parsing in cache
	SystemParameter *ent.SystemParameter
	Pet             *ent.Pet
	User            *ent.User
}

func (m *CacheValue) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m *CacheValue) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}
