package legal_entity

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/buildbeaver/buildbeaver/server/store"

	"github.com/buildbeaver/buildbeaver/server/services"

	"github.com/buildbeaver/buildbeaver/common/logger"
	"github.com/buildbeaver/buildbeaver/common/models"
)

type BasicLegalEntityService struct {
	db                   *store.DB
	legalEntityStore     store.LegalEntityStore
	ownershipStore       store.OwnershipStore
	resourceLinkStore    store.ResourceLinkStore
	identityStore        store.IdentityStore
	authorizationService services.AuthorizationService
	groupService         services.GroupService
	logger.Log
}

func NewBasicLegalEntityService(
	db *store.DB,
	legalEntityStore store.LegalEntityStore,
	ownershipStore store.OwnershipStore,
	resourceLinkStore store.ResourceLinkStore,
	identityStore store.IdentityStore,
	authorizationService services.AuthorizationService,
	groupService services.GroupService,
	logFactory logger.LogFactory,
) *BasicLegalEntityService {
	return &BasicLegalEntityService{
		db:                   db,
		legalEntityStore:     legalEntityStore,
		ownershipStore:       ownershipStore,
		resourceLinkStore:    resourceLinkStore,
		identityStore:        identityStore,
		authorizationService: authorizationService,
		groupService:         groupService,
		Log:                  logFactory("LegalEntityService"),
	}
}

// Create creates a new legal entity and configures default access control rules.
func (s *BasicLegalEntityService) Create(ctx context.Context, txOrNil *store.Tx, legalEntityData *models.LegalEntityData) (*models.LegalEntity, error) {
	return nil, fmt.Errorf("error: Not implemented")
}

// Read an existing legal entity, looking it up by ID.
func (s *BasicLegalEntityService) Read(ctx context.Context, txOrNil *store.Tx, id models.LegalEntityID) (*models.LegalEntity, error) {
	return s.legalEntityStore.Read(ctx, txOrNil, id)
}

func (s *BasicLegalEntityService) ReadByExternalID(ctx context.Context, txOrNil *store.Tx, externalID models.ExternalResourceID) (*models.LegalEntity, error) {
	return nil, fmt.Errorf("error: Not implemented")
}

func (s *BasicLegalEntityService) ReadByIdentityID(ctx context.Context, txOrNil *store.Tx, identityID models.IdentityID) (*models.LegalEntity, error) {
	return nil, fmt.Errorf("error: Not implemented")
}

func (s *BasicLegalEntityService) ReadIdentity(ctx context.Context, txOrNil *store.Tx, id models.LegalEntityID) (*models.Identity, error) {
	return nil, fmt.Errorf("error: Not implemented")
}

func (s *BasicLegalEntityService) FindOrCreate(ctx context.Context, txOrNil *store.Tx, legalEntityData *models.LegalEntityData) (legalEntity *models.LegalEntity, created bool, err error) {
	return nil, false, fmt.Errorf("error: Not implemented")
}

// Upsert creates a legal entity if no legal entity with the same External ID already exists, otherwise it updates
// the existing legal entity's data if it differs from the supplied data.
// Returns the LegalEntity as it exists in the database after the create or update, and
// true,false if the resource was created, false,true if the resource was updated, or false,false if
// neither create nor update was necessary.
func (s *BasicLegalEntityService) Upsert(
	ctx context.Context,
	txOrNil *store.Tx,
	legalEntityData *models.LegalEntityData,
) (*models.LegalEntity, bool, bool, error) {
	err := legalEntityData.Validate()
	if err != nil {
		return nil, false, false, errors.Wrap(err, "error validating legal entity")
	}
	var (
		legalEntity *models.LegalEntity
		created     bool
		updated     bool
	)
	err = s.db.WithTx(ctx, txOrNil, func(tx *store.Tx) error {
		legalEntity, created, updated, err = s.legalEntityStore.Upsert(ctx, tx, legalEntityData)
		if err != nil {
			return fmt.Errorf("error upserting legal entity: %w", err)
		}

		if created || updated {
			// Ensure the legal entity has an up-to-date resource link matching its name
			_, _, err = s.resourceLinkStore.Upsert(ctx, tx, legalEntity)
			if err != nil {
				return fmt.Errorf("error upserting resource link: %w", err)
			}
		}
		if created {
			s.Infof("Created legal entity %q", legalEntity.ID)
		}
		return nil
	})
	if err != nil {
		return nil, false, false, err
	}
	return legalEntity, created, updated, err
}

func (s *BasicLegalEntityService) Update(ctx context.Context, txOrNil *store.Tx, legalEntity *models.LegalEntity) error {
	return fmt.Errorf("error: Not implemented")
}

func (s *BasicLegalEntityService) ListParentLegalEntities(ctx context.Context, txOrNil *store.Tx, legalEntityID models.LegalEntityID, pagination models.Pagination) ([]*models.LegalEntity, *models.Cursor, error) {
	return nil, nil, fmt.Errorf("error: Not implemented")
}

func (s *BasicLegalEntityService) ListMemberLegalEntities(ctx context.Context, txOrNil *store.Tx, parentLegalEntityID models.LegalEntityID, pagination models.Pagination) ([]*models.LegalEntity, *models.Cursor, error) {
	return nil, nil, fmt.Errorf("error: Not implemented")
}

func (s *BasicLegalEntityService) ListAllLegalEntities(ctx context.Context, txOrNil *store.Tx, pagination models.Pagination) ([]*models.LegalEntity, *models.Cursor, error) {
	return nil, nil, fmt.Errorf("error: Not implemented")
}

func (s *BasicLegalEntityService) AddCompanyMember(ctx context.Context, txOrNil *store.Tx, companyID models.LegalEntityID, memberID models.LegalEntityID) error {
	return fmt.Errorf("error: Not implemented")
}

func (s *BasicLegalEntityService) RemoveCompanyMember(ctx context.Context, txOrNil *store.Tx, companyID models.LegalEntityID, memberID models.LegalEntityID) error {
	return fmt.Errorf("error: Not implemented")
}
