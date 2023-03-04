// Code generated by ent, DO NOT EDIT.

package ent

import (
	"myapp/ent/pet"
	"myapp/ent/post"
	"myapp/ent/schema"
	"myapp/ent/system_parameter"
	"myapp/ent/user"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	petFields := schema.Pet{}.Fields()
	_ = petFields
	// petDescName is the schema descriptor for name field.
	petDescName := petFields[1].Descriptor()
	// pet.NameValidator is a validator for the "name" field. It is called by the builders before save.
	pet.NameValidator = func() func(string) error {
		validators := petDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// petDescCode is the schema descriptor for code field.
	petDescCode := petFields[3].Descriptor()
	// pet.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	pet.CodeValidator = petDescCode.Validators[0].(func(string) error)
	// petDescAgeMonth is the schema descriptor for age_month field.
	petDescAgeMonth := petFields[4].Descriptor()
	// pet.AgeMonthValidator is a validator for the "age_month" field. It is called by the builders before save.
	pet.AgeMonthValidator = petDescAgeMonth.Validators[0].(func(int) error)
	// petDescIsDeleted is the schema descriptor for is_deleted field.
	petDescIsDeleted := petFields[5].Descriptor()
	// pet.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	pet.DefaultIsDeleted = petDescIsDeleted.Default.(bool)
	// petDescCreatedBy is the schema descriptor for created_by field.
	petDescCreatedBy := petFields[6].Descriptor()
	// pet.CreatedByValidator is a validator for the "created_by" field. It is called by the builders before save.
	pet.CreatedByValidator = petDescCreatedBy.Validators[0].(func(string) error)
	// petDescCreatedAt is the schema descriptor for created_at field.
	petDescCreatedAt := petFields[7].Descriptor()
	// pet.DefaultCreatedAt holds the default value on creation for the created_at field.
	pet.DefaultCreatedAt = petDescCreatedAt.Default.(time.Time)
	// petDescUpdatedAt is the schema descriptor for updated_at field.
	petDescUpdatedAt := petFields[9].Descriptor()
	// pet.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	pet.DefaultUpdatedAt = petDescUpdatedAt.Default.(time.Time)
	// petDescID is the schema descriptor for id field.
	petDescID := petFields[0].Descriptor()
	// pet.DefaultID holds the default value on creation for the id field.
	pet.DefaultID = petDescID.Default.(func() uuid.UUID)
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescTitle is the schema descriptor for title field.
	postDescTitle := postFields[0].Descriptor()
	// post.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	post.TitleValidator = postDescTitle.Validators[0].(func(string) error)
	// postDescContent is the schema descriptor for content field.
	postDescContent := postFields[1].Descriptor()
	// post.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	post.ContentValidator = postDescContent.Validators[0].(func(string) error)
	// postDescSlug is the schema descriptor for slug field.
	postDescSlug := postFields[2].Descriptor()
	// post.SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	post.SlugValidator = postDescSlug.Validators[0].(func(string) error)
	// postDescStatus is the schema descriptor for status field.
	postDescStatus := postFields[3].Descriptor()
	// post.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	post.StatusValidator = postDescStatus.Validators[0].(func(int) error)
	// postDescCreatedAt is the schema descriptor for created_at field.
	postDescCreatedAt := postFields[4].Descriptor()
	// post.DefaultCreatedAt holds the default value on creation for the created_at field.
	post.DefaultCreatedAt = postDescCreatedAt.Default.(func() time.Time)
	system_parameterFields := schema.System_parameter{}.Fields()
	_ = system_parameterFields
	// system_parameterDescKey is the schema descriptor for key field.
	system_parameterDescKey := system_parameterFields[0].Descriptor()
	// system_parameter.KeyValidator is a validator for the "key" field. It is called by the builders before save.
	system_parameter.KeyValidator = system_parameterDescKey.Validators[0].(func(string) error)
	// system_parameterDescValue is the schema descriptor for value field.
	system_parameterDescValue := system_parameterFields[1].Descriptor()
	// system_parameter.ValueValidator is a validator for the "value" field. It is called by the builders before save.
	system_parameter.ValueValidator = system_parameterDescValue.Validators[0].(func(string) error)
	// system_parameterDescIsDeleted is the schema descriptor for is_deleted field.
	system_parameterDescIsDeleted := system_parameterFields[2].Descriptor()
	// system_parameter.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	system_parameter.DefaultIsDeleted = system_parameterDescIsDeleted.Default.(bool)
	// system_parameterDescCreatedBy is the schema descriptor for created_by field.
	system_parameterDescCreatedBy := system_parameterFields[3].Descriptor()
	// system_parameter.CreatedByValidator is a validator for the "created_by" field. It is called by the builders before save.
	system_parameter.CreatedByValidator = system_parameterDescCreatedBy.Validators[0].(func(string) error)
	// system_parameterDescCreatedAt is the schema descriptor for created_at field.
	system_parameterDescCreatedAt := system_parameterFields[4].Descriptor()
	// system_parameter.DefaultCreatedAt holds the default value on creation for the created_at field.
	system_parameter.DefaultCreatedAt = system_parameterDescCreatedAt.Default.(time.Time)
	// system_parameterDescUpdatedAt is the schema descriptor for updated_at field.
	system_parameterDescUpdatedAt := system_parameterFields[6].Descriptor()
	// system_parameter.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	system_parameter.DefaultUpdatedAt = system_parameterDescUpdatedAt.Default.(time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = func() func(string) error {
		validators := userDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPhone is the schema descriptor for phone field.
	userDescPhone := userFields[3].Descriptor()
	// user.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	user.PhoneValidator = userDescPhone.Validators[0].(func(string) error)
	// userDescIsDeleted is the schema descriptor for is_deleted field.
	userDescIsDeleted := userFields[4].Descriptor()
	// user.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	user.DefaultIsDeleted = userDescIsDeleted.Default.(bool)
	// userDescCreatedBy is the schema descriptor for created_by field.
	userDescCreatedBy := userFields[5].Descriptor()
	// user.CreatedByValidator is a validator for the "created_by" field. It is called by the builders before save.
	user.CreatedByValidator = userDescCreatedBy.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[6].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[8].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
