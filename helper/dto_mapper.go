// Package helper wrapper for mapper https://github.com/dranikpg/dto-mapper

package helper

import "github.com/dranikpg/dto-mapper"

func Mapper(to interface{}, from interface{}) error {
	return dto.Map(to, from)
}
