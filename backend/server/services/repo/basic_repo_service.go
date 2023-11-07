package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/buildbeaver/buildbeaver/common/logger"
	"github.com/buildbeaver/buildbeaver/common/models"
	"github.com/buildbeaver/buildbeaver/common/models/search"
	"github.com/buildbeaver/buildbeaver/server/dto"
	"github.com/buildbeaver/buildbeaver/server/store"
)

type BasicRepoService struct {
	db                *store.DB
	ownershipStore    store.OwnershipStore
	repoStore         store.RepoStore
	resourceLinkStore store.ResourceLinkStore
	logger.Log
}

func NewBasicRepoService(
	db *store.DB,
	ownershipStore store.OwnershipStore,
	repoStore store.RepoStore,
	resourceLinkStore store.ResourceLinkStore,
	logFactory logger.LogFactory) *BasicRepoService {

	return &BasicRepoService{
		db:                db,
		ownershipStore:    ownershipStore,
		repoStore:         repoStore,
		resourceLinkStore: resourceLinkStore,
		Log:               logFactory("RepoService"),
	}
}

// Read an existing repo, looking it up by ID.
// Returns models.ErrNotFound if the repo does not exist.
func (s *BasicRepoService) Read(ctx context.Context, txOrNil *store.Tx, id models.RepoID) (*models.Repo, error) {
	return s.repoStore.Read(ctx, txOrNil, id)
}

func (s *BasicRepoService) ReadByExternalID(ctx context.Context, txOrNil *store.Tx, externalID models.ExternalResourceID) (*models.Repo, error) {
	return nil, fmt.Errorf("error: Not implemented")
}

// Upsert creates a repo if it does not exist, otherwise it updates its mutable properties
// if they differ from the in-memory instance. Returns true,false if the resource was created
// and false,true if the resource was updated. false,false if neither a create or update was necessary.
// Repo Metadata and selected fields will not be updated (including Enabled and SSHKeySecretID fields).
func (s *BasicRepoService) Upsert(ctx context.Context, txOrNil *store.Tx, repo *models.Repo) (created bool, updated bool, err error) {
	err = repo.Validate()
	if err != nil {
		return false, false, errors.Wrap(err, "error validating repo")
	}
	err = s.db.WithTx(ctx, txOrNil, func(tx *store.Tx) error {
		created, updated, err = s.repoStore.Upsert(ctx, tx, repo)
		if err != nil {
			return fmt.Errorf("error upserting repo: %w", err)
		}
		ownership := models.NewOwnership(models.NewTime(time.Now()), repo.LegalEntityID.ResourceID, repo.GetID())
		_, _, err = s.ownershipStore.Upsert(ctx, tx, ownership)
		if err != nil {
			return fmt.Errorf("error upserting ownership: %w", err)
		}
		if created || updated {
			_, _, err = s.resourceLinkStore.Upsert(ctx, tx, repo)
			if err != nil {
				return fmt.Errorf("error upserting resource link: %w", err)
			}
		}
		if created {
			err := s.repoStore.InitializeBuildCounter(ctx, tx, repo.ID)
			if err != nil {
				return fmt.Errorf("error initializing repo build counter: %w", err)
			}
			s.Infof("Created repo %q", repo.ID)
		}
		return nil
	})
	return created, updated, err
}

func (s *BasicRepoService) UpdateRepoEnabled(ctx context.Context, repoID models.RepoID, update dto.UpdateRepoEnabled) (*models.Repo, error) {
	return nil, fmt.Errorf("error: Not implemented")
}

func (s *BasicRepoService) SoftDelete(ctx context.Context, txOrNil *store.Tx, repo *models.Repo) error {
	return fmt.Errorf("error: Not implemented")
}

func (s *BasicRepoService) Search(ctx context.Context, txOrNil *store.Tx, searcher models.IdentityID, query search.Query) ([]*models.Repo, *models.Cursor, error) {
	return nil, nil, fmt.Errorf("error: Not implemented")
}

func (s *BasicRepoService) AllocateBuildNumber(ctx context.Context, txOrNil *store.Tx, repoID models.RepoID) (models.BuildNumber, error) {
	return 0, fmt.Errorf("error: Not implemented")
}
