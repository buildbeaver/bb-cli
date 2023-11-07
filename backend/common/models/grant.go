package models

import (
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

const GrantResourceKind ResourceKind = "grant"

type GrantID struct {
	ResourceID
}

func NewGrantID() GrantID {
	return GrantID{ResourceID: NewResourceID(GrantResourceKind)}
}

func GrantIDFromResourceID(id ResourceID) GrantID {
	return GrantID{ResourceID: id}
}

type GrantMetadata struct {
	// ID uniquely identifies the grant
	ID        GrantID `json:"id" goqu:"skipupdate" db:"access_control_grant_id"`
	CreatedAt Time    `json:"created_at" goqu:"skipupdate" db:"access_control_grant_created_at"`
	UpdatedAt Time    `json:"updated_at" db:"access_control_grant_updated_at"`
}

type Grant struct {
	// This is only a stub implementation; almost all fields missing
	GrantMetadata
}

func (m *Grant) GetCreatedAt() Time {
	return m.CreatedAt
}

func (m *Grant) GetID() ResourceID {
	return m.ID.ResourceID
}

func (m *Grant) GetKind() ResourceKind {
	return GrantResourceKind
}

func (m *Grant) Validate() error {
	var result *multierror.Error
	if !m.ID.Valid() {
		result = multierror.Append(result, errors.New("error id must be set"))
	}
	if m.CreatedAt.IsZero() {
		result = multierror.Append(result, errors.New("error created at must be set"))
	}
	return result.ErrorOrNil()
}
