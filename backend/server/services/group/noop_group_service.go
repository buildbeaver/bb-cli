package group

import (
	"context"
	"fmt"
	"time"

	"github.com/buildbeaver/buildbeaver/common/logger"
	"github.com/buildbeaver/buildbeaver/common/models"
	"github.com/buildbeaver/buildbeaver/server/store"
)

// dummyGroup is returned when reading groups (especially standard groups) that calling code expects
// to be there, to provide parameters to pass to subsequent calls to this service.
var dummyGroup = &models.Group{
	GroupMetadata: models.GroupMetadata{
		ID:        models.NewGroupID(),
		CreatedAt: models.NewTime(time.Now()),
		UpdatedAt: models.NewTime(time.Now()),
		DeletedAt: nil,
		ETag:      "",
	},
	Name: "Dummy group",
}

// dummyGroupMembership is returned when reading group memberships that calling code expects to be there.
var dummyGroupMembership = &models.GroupMembership{
	GroupMembershipMetadata: models.GroupMembershipMetadata{
		ID:        models.NewGroupMembershipID(),
		CreatedAt: models.NewTime(time.Now()),
	},
	GroupMembershipData: models.GroupMembershipData{},
}

type NoOpGroupService struct {
	db *store.DB
	logger.Log
}

func NewNoOpGroupService(db *store.DB, logFactory logger.LogFactory) *NoOpGroupService {
	return &NoOpGroupService{
		db:  db,
		Log: logFactory("NoOpGroupService"),
	}
}

// ReadByName reads an existing access control Group, looking it up by group name and the ID of the
// legal entity that owns the group. Returns models.ErrNotFound if the group does not exist.
// This no-op implementation just returns success, and an existing dummy group.
// This no-op implementation just returns a dummy group.
func (s *NoOpGroupService) ReadByName(
	ctx context.Context,
	txOrNil *store.Tx,
	ownerLegalEntityID models.LegalEntityID,
	groupName models.ResourceName,
) (*models.Group, error) {
	return dummyGroup, nil
}

// ReadByExternalID reads an existing group, looking it up by its unique external id.
// Returns models.ErrNotFound if the group does not exist.
func (s *NoOpGroupService) ReadByExternalID(ctx context.Context, txOrNil *store.Tx, externalID models.ExternalResourceID) (*models.Group, error) {
	return nil, fmt.Errorf("error: Not implemented")
}

// FindOrCreateStandardGroup finds or creates a new access control Group for a Legal Entity, and sets up
// permissions for any new group that was created, based on the supplied standard group definition.
func (s *NoOpGroupService) FindOrCreateStandardGroup(
	ctx context.Context,
	tx *store.Tx,
	legalEntity *models.LegalEntity,
	groupDefinition *models.StandardGroupDefinition,
) (*models.Group, error) {
	return nil, fmt.Errorf("error: Not implemented")
}

// FindOrCreateByName finds and returns the access control Group with the name and legal entity specified in
// the supplied group data.
// If no such group exists then a new group is created and returned, and true is returned for 'created'.
// This no-op implementation just returns success, and an existing dummy group.
func (s *NoOpGroupService) FindOrCreateByName(ctx context.Context, txOrNil *store.Tx, groupData *models.Group) (*models.Group, bool, error) {
	return dummyGroup, false, nil
}

// UpsertByExternalID creates a group if no group with the same External ID already exists, otherwise it updates
// the existing group's mutable properties if they differ from the in-memory instance.
// Returns true,false if the resource was created, false,true if the resource was updated, or false,false if
// neither create nor update was necessary.
// Returns an error if no External ID is filled out in the supplied Group.
// In all cases group.ID will be filled out in the supplied group object.
// This no-op implementation just returns success, and true, false as if created.
func (s *NoOpGroupService) UpsertByExternalID(ctx context.Context, txOrNil *store.Tx, group *models.Group) (created bool, updated bool, err error) {
	return true, false, nil
}

// Delete permanently and idempotently deletes an access control group.
// All memberships and grants for this group will also be permanently deleted.
// This no-op implementation just returns success.
func (s *NoOpGroupService) Delete(ctx context.Context, txOrNil *store.Tx, id models.GroupID) error {
	return nil
}

// ListGroups returns a list of groups. Use cursor to page through results, if any.
// If groupParent is provided then only groups owned by the supplied parent legal entity will be returned.
// If memberID is provided then only groups that include the provided identity as a member will be returned.
func (s *NoOpGroupService) ListGroups(
	ctx context.Context,
	txOrNil *store.Tx,
	groupParent *models.LegalEntityID,
	memberID *models.IdentityID,
	pagination models.Pagination,
) ([]*models.Group, *models.Cursor, error) {
	return nil, nil, fmt.Errorf("error: Not implemented")
}

// FindOrCreateMembership adds the specified identity to an access control Group by adding a group membership
// for a specific source system.
// This method is idempotent, and returns true if a new membership was created or false if there was already
// a membership for this identity for the group with the specified source system
// This no-op implementation just returns success and an existing dummy group.
func (s *NoOpGroupService) FindOrCreateMembership(
	ctx context.Context,
	txOrNil *store.Tx,
	membershipData *models.GroupMembershipData,
) (membership *models.GroupMembership, created bool, err error) {
	return dummyGroupMembership, false, nil
}

// RemoveMembership removes a membership for the specified identity from an access control group.
// If sourceSystem is not nil then only the membership record matching the source system will be deleted;
// otherwise membership records from all source systems for the member will be deleted.
// This no-op implementation just returns success.
func (s *NoOpGroupService) RemoveMembership(
	ctx context.Context,
	txOrNil *store.Tx,
	groupID models.GroupID,
	memberID models.IdentityID,
	sourceSystem *models.SystemName,
) error {
	return nil
}

// ReadMembership reads an existing access control group membership, looking it up by group member, identity and
// source system. Returns models.ErrNotFound if the group membership does not exist.
// This no-op implementation just returns a dummy group membership.
func (s *NoOpGroupService) ReadMembership(
	ctx context.Context,
	txOrNil *store.Tx,
	groupID models.GroupID,
	memberID models.IdentityID,
	sourceSystem models.SystemName,
) (*models.GroupMembership, error) {
	return dummyGroupMembership, nil
}

// ListGroupMemberships returns a list of group memberships. Use cursor to page through results, if any.
// If groupID is provided then only memberships of the specified group will be returned.
// If memberID is provided then only groups that include the provided identity as a member will be returned.
// If sourceSystem is provided then only memberships with matching source system values will be returned.
func (s *NoOpGroupService) ListGroupMemberships(
	ctx context.Context,
	txOrNil *store.Tx,
	groupID *models.GroupID,
	memberID *models.IdentityID,
	sourceSystem *models.SystemName,
	pagination models.Pagination,
) ([]*models.GroupMembership, *models.Cursor, error) {
	return nil, nil, fmt.Errorf("error: Not implemented")
}
