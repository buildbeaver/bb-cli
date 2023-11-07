package models

import (
	"errors"

	"github.com/hashicorp/go-multierror"
)

const GroupResourceKind ResourceKind = "group"

type GroupID struct {
	ResourceID
}

func NewGroupID() GroupID {
	return GroupID{ResourceID: NewResourceID(GroupResourceKind)}
}

func GroupIDFromResourceID(id ResourceID) GroupID {
	return GroupID{ResourceID: id}
}

type GroupMetadata struct {
	ID        GroupID `json:"id" goqu:"skipupdate" db:"access_control_group_id"`
	CreatedAt Time    `json:"created_at" goqu:"skipupdate" db:"access_control_group_created_at"`
	UpdatedAt Time    `json:"updated_at" db:"access_control_group_updated_at"`
	DeletedAt *Time   `json:"deleted_at,omitempty" db:"access_control_group_deleted_at"`
	ETag      ETag    `json:"etag" db:"access_control_group_etag" hash:"ignore"`
}

type Group struct {
	// This is only a stub implementation; almost all fields missing
	GroupMetadata
	// Name of the group, unique within the groups owned by the owner legal entity
	Name ResourceName `json:"name" db:"access_control_group_name"`
}

func (m *Group) GetID() ResourceID {
	return m.ID.ResourceID
}

func (m *Group) GetKind() ResourceKind {
	return GroupResourceKind
}

func (m *Group) GetCreatedAt() Time {
	return m.CreatedAt
}

func (m *Group) GetUpdatedAt() Time {
	return m.UpdatedAt
}

func (m *Group) SetUpdatedAt(t Time) {
	m.UpdatedAt = t
}

func (m *Group) GetETag() ETag {
	return m.ETag
}

func (m *Group) SetETag(eTag ETag) {
	m.ETag = eTag
}

func (m *Group) Validate() error {
	var result *multierror.Error
	if !m.ID.Valid() {
		result = multierror.Append(result, errors.New("error id must be set"))
	}
	if m.CreatedAt.IsZero() {
		result = multierror.Append(result, errors.New("error created at must be set"))
	}
	if m.UpdatedAt.IsZero() {
		result = multierror.Append(result, errors.New("error updated at must be set"))
	}
	if m.DeletedAt != nil && m.DeletedAt.IsZero() {
		result = multierror.Append(result, errors.New("error deleted at must be non-zero when set"))
	}
	return result.ErrorOrNil()
}

// StandardGroupDefinition defines a standard access control group that can easily be created for a
// legal entity (normally for a company or other organization).
type StandardGroupDefinition struct {
	Name        ResourceName
	Description string
	Operations  []*Operation
}

var RunnerStandardGroup = &StandardGroupDefinition{}
