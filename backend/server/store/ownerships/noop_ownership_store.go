package ownerships

import (
	"context"
	"fmt"

	"github.com/buildbeaver/buildbeaver/common/logger"
	"github.com/buildbeaver/buildbeaver/common/models"
	"github.com/buildbeaver/buildbeaver/server/store"
)

func init() {
	store.MustDBModel(&models.Ownership{})
}

type NoOpOwnershipStore struct {
	table *store.ResourceTable
}

func NewNoOpOwnershipStore(db *store.DB, logFactory logger.LogFactory) *NoOpOwnershipStore {
	return &NoOpOwnershipStore{
		table: store.NewResourceTable(db, logFactory, &models.Ownership{}),
	}
}

// Create a new ownership.
// This no-op implementation just returns success.
func (d *NoOpOwnershipStore) Create(ctx context.Context, txOrNil *store.Tx, ownership *models.Ownership) error {
	return nil
}

// Read an existing ownership, looking it up by ResourceID.
// Returns models.ErrNotFound if the ownership does not exist.
func (d *NoOpOwnershipStore) Read(ctx context.Context, txOrNil *store.Tx, id models.OwnershipID) (*models.Ownership, error) {
	return nil, fmt.Errorf("error: Not implemented")
}

// readByOwnedResource reads an existing ownership, looking it up by the owned resource.
// Returns models.ErrNotFound if the ownership does not exist.
func (d *NoOpOwnershipStore) readByOwnedResource(ctx context.Context, txOrNil *store.Tx, ownedResourceID models.ResourceID) (*models.Ownership, error) {
	return nil, fmt.Errorf("error: Not implemented")
}

// Update an existing ownership with optimistic locking. Overrides all previous values using the supplied model.
// Returns store.ErrOptimisticLockFailed if there is an optimistic lock mismatch.
func (d *NoOpOwnershipStore) Update(ctx context.Context, txOrNil *store.Tx, ownership *models.Ownership) error {
	return nil
}

// Upsert creates an ownership if it does not exist, otherwise it updates its mutable properties
// if they differ from the in-memory instance. Returns true,false if the resource was created
// and false,true if the resource was updated. false,false if neither a create or update was necessary.
// This no-op implementation just returns success, and true, false (as if the record was created).
func (d *NoOpOwnershipStore) Upsert(ctx context.Context, txOrNil *store.Tx, ownership *models.Ownership) (bool, bool, error) {
	return true, false, nil
}

// Delete permanently and idempotently deletes an ownership, identifying it by owned resource id.
func (d *NoOpOwnershipStore) Delete(ctx context.Context, txOrNil *store.Tx, ownedResourceID models.ResourceID) error {
	return nil
}
