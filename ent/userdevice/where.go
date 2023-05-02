// Code generated by ent, DO NOT EDIT.

package userdevice

import (
	"myapp/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uint64) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldEQ(FieldUserID, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldEQ(FieldVersion, v))
}

// Platform applies equality check predicate on the "platform" field. It's identical to PlatformEQ.
func Platform(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldEQ(FieldPlatform, v))
}

// LatestSkipUpdate applies equality check predicate on the "latest_skip_update" field. It's identical to LatestSkipUpdateEQ.
func LatestSkipUpdate(v time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldEQ(FieldLatestSkipUpdate, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeviceID applies equality check predicate on the "device_id" field. It's identical to DeviceIDEQ.
func DeviceID(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldEQ(FieldDeviceID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uint64) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uint64) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uint64) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uint64) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uint64) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uint64) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uint64) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uint64) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldLTE(FieldUserID, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldLTE(FieldVersion, v))
}

// VersionContains applies the Contains predicate on the "version" field.
func VersionContains(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldContains(FieldVersion, v))
}

// VersionHasPrefix applies the HasPrefix predicate on the "version" field.
func VersionHasPrefix(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldHasPrefix(FieldVersion, v))
}

// VersionHasSuffix applies the HasSuffix predicate on the "version" field.
func VersionHasSuffix(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldHasSuffix(FieldVersion, v))
}

// VersionEqualFold applies the EqualFold predicate on the "version" field.
func VersionEqualFold(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldEqualFold(FieldVersion, v))
}

// VersionContainsFold applies the ContainsFold predicate on the "version" field.
func VersionContainsFold(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldContainsFold(FieldVersion, v))
}

// PlatformEQ applies the EQ predicate on the "platform" field.
func PlatformEQ(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldEQ(FieldPlatform, v))
}

// PlatformNEQ applies the NEQ predicate on the "platform" field.
func PlatformNEQ(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldNEQ(FieldPlatform, v))
}

// PlatformIn applies the In predicate on the "platform" field.
func PlatformIn(vs ...string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldIn(FieldPlatform, vs...))
}

// PlatformNotIn applies the NotIn predicate on the "platform" field.
func PlatformNotIn(vs ...string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldNotIn(FieldPlatform, vs...))
}

// PlatformGT applies the GT predicate on the "platform" field.
func PlatformGT(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldGT(FieldPlatform, v))
}

// PlatformGTE applies the GTE predicate on the "platform" field.
func PlatformGTE(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldGTE(FieldPlatform, v))
}

// PlatformLT applies the LT predicate on the "platform" field.
func PlatformLT(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldLT(FieldPlatform, v))
}

// PlatformLTE applies the LTE predicate on the "platform" field.
func PlatformLTE(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldLTE(FieldPlatform, v))
}

// PlatformContains applies the Contains predicate on the "platform" field.
func PlatformContains(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldContains(FieldPlatform, v))
}

// PlatformHasPrefix applies the HasPrefix predicate on the "platform" field.
func PlatformHasPrefix(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldHasPrefix(FieldPlatform, v))
}

// PlatformHasSuffix applies the HasSuffix predicate on the "platform" field.
func PlatformHasSuffix(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldHasSuffix(FieldPlatform, v))
}

// PlatformEqualFold applies the EqualFold predicate on the "platform" field.
func PlatformEqualFold(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldEqualFold(FieldPlatform, v))
}

// PlatformContainsFold applies the ContainsFold predicate on the "platform" field.
func PlatformContainsFold(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldContainsFold(FieldPlatform, v))
}

// LatestSkipUpdateEQ applies the EQ predicate on the "latest_skip_update" field.
func LatestSkipUpdateEQ(v time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldEQ(FieldLatestSkipUpdate, v))
}

// LatestSkipUpdateNEQ applies the NEQ predicate on the "latest_skip_update" field.
func LatestSkipUpdateNEQ(v time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldNEQ(FieldLatestSkipUpdate, v))
}

// LatestSkipUpdateIn applies the In predicate on the "latest_skip_update" field.
func LatestSkipUpdateIn(vs ...time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldIn(FieldLatestSkipUpdate, vs...))
}

// LatestSkipUpdateNotIn applies the NotIn predicate on the "latest_skip_update" field.
func LatestSkipUpdateNotIn(vs ...time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldNotIn(FieldLatestSkipUpdate, vs...))
}

// LatestSkipUpdateGT applies the GT predicate on the "latest_skip_update" field.
func LatestSkipUpdateGT(v time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldGT(FieldLatestSkipUpdate, v))
}

// LatestSkipUpdateGTE applies the GTE predicate on the "latest_skip_update" field.
func LatestSkipUpdateGTE(v time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldGTE(FieldLatestSkipUpdate, v))
}

// LatestSkipUpdateLT applies the LT predicate on the "latest_skip_update" field.
func LatestSkipUpdateLT(v time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldLT(FieldLatestSkipUpdate, v))
}

// LatestSkipUpdateLTE applies the LTE predicate on the "latest_skip_update" field.
func LatestSkipUpdateLTE(v time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldLTE(FieldLatestSkipUpdate, v))
}

// LatestSkipUpdateIsNil applies the IsNil predicate on the "latest_skip_update" field.
func LatestSkipUpdateIsNil() predicate.UserDevice {
	return predicate.UserDevice(sql.FieldIsNull(FieldLatestSkipUpdate))
}

// LatestSkipUpdateNotNil applies the NotNil predicate on the "latest_skip_update" field.
func LatestSkipUpdateNotNil() predicate.UserDevice {
	return predicate.UserDevice(sql.FieldNotNull(FieldLatestSkipUpdate))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.UserDevice {
	return predicate.UserDevice(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.UserDevice {
	return predicate.UserDevice(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.UserDevice {
	return predicate.UserDevice(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.UserDevice {
	return predicate.UserDevice(sql.FieldNotNull(FieldUpdatedAt))
}

// DeviceIDEQ applies the EQ predicate on the "device_id" field.
func DeviceIDEQ(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldEQ(FieldDeviceID, v))
}

// DeviceIDNEQ applies the NEQ predicate on the "device_id" field.
func DeviceIDNEQ(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldNEQ(FieldDeviceID, v))
}

// DeviceIDIn applies the In predicate on the "device_id" field.
func DeviceIDIn(vs ...string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldIn(FieldDeviceID, vs...))
}

// DeviceIDNotIn applies the NotIn predicate on the "device_id" field.
func DeviceIDNotIn(vs ...string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldNotIn(FieldDeviceID, vs...))
}

// DeviceIDGT applies the GT predicate on the "device_id" field.
func DeviceIDGT(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldGT(FieldDeviceID, v))
}

// DeviceIDGTE applies the GTE predicate on the "device_id" field.
func DeviceIDGTE(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldGTE(FieldDeviceID, v))
}

// DeviceIDLT applies the LT predicate on the "device_id" field.
func DeviceIDLT(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldLT(FieldDeviceID, v))
}

// DeviceIDLTE applies the LTE predicate on the "device_id" field.
func DeviceIDLTE(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldLTE(FieldDeviceID, v))
}

// DeviceIDContains applies the Contains predicate on the "device_id" field.
func DeviceIDContains(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldContains(FieldDeviceID, v))
}

// DeviceIDHasPrefix applies the HasPrefix predicate on the "device_id" field.
func DeviceIDHasPrefix(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldHasPrefix(FieldDeviceID, v))
}

// DeviceIDHasSuffix applies the HasSuffix predicate on the "device_id" field.
func DeviceIDHasSuffix(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldHasSuffix(FieldDeviceID, v))
}

// DeviceIDIsNil applies the IsNil predicate on the "device_id" field.
func DeviceIDIsNil() predicate.UserDevice {
	return predicate.UserDevice(sql.FieldIsNull(FieldDeviceID))
}

// DeviceIDNotNil applies the NotNil predicate on the "device_id" field.
func DeviceIDNotNil() predicate.UserDevice {
	return predicate.UserDevice(sql.FieldNotNull(FieldDeviceID))
}

// DeviceIDEqualFold applies the EqualFold predicate on the "device_id" field.
func DeviceIDEqualFold(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldEqualFold(FieldDeviceID, v))
}

// DeviceIDContainsFold applies the ContainsFold predicate on the "device_id" field.
func DeviceIDContainsFold(v string) predicate.UserDevice {
	return predicate.UserDevice(sql.FieldContainsFold(FieldDeviceID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserDevice) predicate.UserDevice {
	return predicate.UserDevice(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserDevice) predicate.UserDevice {
	return predicate.UserDevice(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserDevice) predicate.UserDevice {
	return predicate.UserDevice(func(s *sql.Selector) {
		p(s.Not())
	})
}