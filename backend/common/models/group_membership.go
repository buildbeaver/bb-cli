package models

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"
)

const GroupMembershipResourceKind ResourceKind = "group-membership"

type GroupMembershipID struct {
	ResourceID
}

func NewGroupMembershipID() GroupMembershipID {
	return GroupMembershipID{ResourceID: NewResourceID(GroupMembershipResourceKind)}
}

func GroupMembershipIDFromResourceID(id ResourceID) GroupMembershipID {
	return GroupMembershipID{ResourceID: id}
}

type GroupMembershipMetadata struct {
	ID        GroupMembershipID `json:"id" goqu:"skipupdate" db:"access_control_group_membership_id"`
	CreatedAt Time              `json:"created_at"  db:"access_control_group_membership_created_at"`
}

type GroupMembershipData struct {
	// This is only a stub implementation; all fields missing
}

type GroupMembership struct {
	GroupMembershipMetadata
	GroupMembershipData
}

func NewGroupMembershipData(
	groupID GroupID,
	memberIdentityID IdentityID,
	sourceSystem SystemName,
	addedByLegalEntityID LegalEntityID,
) *GroupMembershipData {
	return &GroupMembershipData{}
}

func (m *GroupMembership) GetCreatedAt() Time {
	return m.CreatedAt
}

func (m *GroupMembership) GetID() ResourceID {
	return m.ID.ResourceID
}

func (m *GroupMembership) GetKind() ResourceKind {
	return GroupMembershipResourceKind
}

func (m *GroupMembership) Validate() error {
	var result *multierror.Error
	if !m.ID.Valid() {
		result = multierror.Append(result, errors.New("error id must be set"))
	}
	if m.CreatedAt.IsZero() {
		result = multierror.Append(result, errors.New("error created at must be set"))
	}
	err := m.GroupMembershipData.Validate()
	if err != nil {
		result = multierror.Append(result, fmt.Errorf("data is invalid: %s", err))
	}
	return result.ErrorOrNil()
}

func (m *GroupMembershipData) Validate() error {
	return nil
}
