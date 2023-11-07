package services

import (
	"context"
	"io"

	"github.com/buildbeaver/buildbeaver/common/certificates"
	"github.com/buildbeaver/buildbeaver/common/models"
	"github.com/buildbeaver/buildbeaver/common/models/search"
	"github.com/buildbeaver/buildbeaver/server/dto"
	"github.com/buildbeaver/buildbeaver/server/store"
)

type QueueService interface {
	// EnqueueBuildFromCommit parses the commit's build definition and enqueues a new build from it.
	// If there is a problem with the commit's build definition, a skeleton build is enqueued that is immediately
	// set to failed with an error describing the problem, and no error will be returned from this function.
	// Returns an error only if there was a transient issue that could be retried.
	EnqueueBuildFromCommit(ctx context.Context, txOrNil *store.Tx, commit *models.Commit, ref string, opts *models.BuildOptions) (*dto.BuildGraph, error)
	// EnqueueBuildFromBuildDefinition enqueues a new build based on the specified build definition, which is assumed
	// to have come from the specified commit. Unlike EnqueueBuildFromCommit this function will return an error
	// if there is a problem with the build definition (as well as any transient errors).
	EnqueueBuildFromBuildDefinition(ctx context.Context, txOrNil *store.Tx, repoID models.RepoID, commitID models.CommitID, buildDef *models.BuildDefinition, ref string, opts *models.BuildOptions) (*dto.BuildGraph, error)
	// AddConfigToBuild enqueues new jobs for an existing build, taken from the supplied build configuration.
	// Returns the full build graph containing both existing and new jobs, as well as an array containing just the new jobs.
	// This function will return an error if there is a problem with the jobs, as well as any transient errors.
	AddConfigToBuild(ctx context.Context, txOrNil *store.Tx, buildID models.BuildID, config []byte, configType models.ConfigType) (*dto.BuildGraph, []*dto.JobGraph, error)
	// CheckBuildConfigLength returns an error if the supplied length (in bytes) is too long for a build configuration,
	// or if the configuration is empty.
	CheckBuildConfigLength(buildDefinitionLength int) error
	// Dequeue returns the next queued job that is ready for execution and that the specified
	// runner is capable of running, or a ErrCodeNotFound if no jobs are ready for execution.
	Dequeue(ctx context.Context, runnerID models.RunnerID) (*dto.RunnableJob, error)
	// UpdateJobStatus updates the status of a job that was previously dequeued. If the new status is
	// WorkflowStatusFailed then an error should be provided to indicate what happened.
	// This function will maintain the status of the build containing this job, to reflect the overall
	// status of the build each time the status of a job is changed.
	UpdateJobStatus(ctx context.Context, txOrNil *store.Tx, jobID models.JobID, update dto.UpdateJobStatus) (*models.Job, error)
	// UpdateJobFingerprint sets the fingerprint that has been calculated for a job. If the build is not configured
	// with the force option (e.g. force=false), the server will attempt to locate previously a successful job with a
	// matching fingerprint and indirect this job to it. If an indirection has been set, the agent must skip the job.
	UpdateJobFingerprint(ctx context.Context, jobID models.JobID, update dto.UpdateJobFingerprint) (*models.Job, error)
	// UpdateStepStatus updates the status of a step that is executing under a job that was previously dequeued.
	// If the new status is WorkflowStatusFailed then an error should be provided to indicate what happened.
	UpdateStepStatus(ctx context.Context, txOrNil *store.Tx, stepID models.StepID, update dto.UpdateStepStatus) (*models.Step, error)
	// ReadQueuedBuild makes a queued build DTO including all child jobs and steps.
	ReadQueuedBuild(ctx context.Context, txOrNil *store.Tx, buildID models.BuildID) (*dto.QueuedBuild, error)
	// ReadJobGraph makes and returns a JobGraph for the specified job.
	ReadJobGraph(ctx context.Context, txOrNil *store.Tx, jobID models.JobID) (*dto.JobGraph, error)
}

type LogService interface {
	// Create a new log descriptor.
	// Returns store.ErrAlreadyExists if a log descriptor with matching unique properties already exists.
	Create(ctx context.Context, txOrNil *store.Tx, log *models.LogDescriptor) (*models.LogDescriptor, error)
	// Read an existing log descriptor, looking it up by ID.
	// Returns models.ErrNotFound if the log descriptor does not exist.
	Read(ctx context.Context, txOrNil *store.Tx, id models.LogDescriptorID) (*models.LogDescriptor, error)
	// Seal a log descriptor and its data, making it immutable going forward.
	Seal(ctx context.Context, txOrNil *store.Tx, id models.LogDescriptorID) error
	// Search all log descriptors. If searcher is set, the results will be limited to log descriptors the searcher
	// is authorized to see (via the read:build permission). Use cursor to page through results, if any.
	Search(ctx context.Context, txOrNil *store.Tx, searcher models.IdentityID, search models.LogDescriptorSearch) ([]*models.LogDescriptor, *models.Cursor, error)
	// WriteData pipes data from reader and writes it to the log descriptor's data.
	WriteData(ctx context.Context, logDescriptorID models.LogDescriptorID, reader io.Reader) error
	// ReadData opens a read stream to a log descriptor's data.
	ReadData(ctx context.Context, logID models.LogDescriptorID, search *models.LogSearch) (io.ReadCloser, error)
}

// BlobStore is an interface for storing and retrieving flat files.
type BlobStore interface {
	// PutBlob writes all data in the source reader to a blob identified by key.
	// The caller is responsible for closing the reader.
	PutBlob(ctx context.Context, key string, source io.Reader) error
	// GetBlob returns a reader positioned at the beginning of the blob identified by key.
	// The caller is responsible for closing the reader.
	GetBlob(ctx context.Context, key string) (io.ReadCloser, error)
	// GetBlobRange returns a reader positioned at the specified offset of the blob identified
	// by key, which will read up to length bytes. The caller is responsible for closing the reader.
	GetBlobRange(ctx context.Context, key string, offset, length int64) (io.ReadCloser, error)
	// DeleteBlob deletes a blob. Returns nil if the blob does not exist.
	DeleteBlob(ctx context.Context, key string) error
	// ListBlobs lists blobs matching prefix, starting at marker. Use cursor to page through results, if any.
	ListBlobs(ctx context.Context, prefix string, marker string, pagination models.Pagination) ([]*models.BlobDescriptor, *models.Cursor, error)
}

type RepoService interface {
	// Read an existing repo, looking it up by ID.
	Read(ctx context.Context, txOrNil *store.Tx, id models.RepoID) (*models.Repo, error)
	// ReadByExternalID reads an existing repo, looking it up by its external id.
	// Returns models.ErrNotFound if the repo does not exist.
	ReadByExternalID(ctx context.Context, txOrNil *store.Tx, externalID models.ExternalResourceID) (*models.Repo, error)
	// Upsert creates a repo if it does not exist, otherwise it updates its mutable properties
	// if they differ from the in-memory instance. Returns true,false if the resource was created
	// and false,true if the resource was updated. false,false if neither a create or update was necessary.
	// Repo Metadata and selected fields will not be updated (including Enabled and SSHKeySecretID fields).
	Upsert(ctx context.Context, txOrNil *store.Tx, repo *models.Repo) (bool, bool, error)
	// Search all repos. If searcher is set, the results will be limited to repos the searcher is authorized to
	// see (via the read:repo permission). Use cursor to page through results, if any.
	Search(ctx context.Context, txOrNil *store.Tx, searcher models.IdentityID, query search.Query) ([]*models.Repo, *models.Cursor, error)
	// UpdateRepoEnabled enables or disables builds for a repo.
	UpdateRepoEnabled(ctx context.Context, repoID models.RepoID, update dto.UpdateRepoEnabled) (*models.Repo, error)
	// SoftDelete soft deletes an existing repo.
	// Returns store.ErrOptimisticLockFailed if there is an optimistic lock mismatch, i.e. if the repo has changed in
	// the database since the supplied object was read.
	SoftDelete(ctx context.Context, txOrNil *store.Tx, repo *models.Repo) error
	// AllocateBuildNumber increments and returns the build counter for the specified repo. The returned counter
	// is safe to assign to a build and will never be returned again by this function.
	AllocateBuildNumber(ctx context.Context, txOrNil *store.Tx, repoID models.RepoID) (models.BuildNumber, error)
}

type AuthorizationService interface {
	// IsAuthorized returns true if the identity is authorized to perform operation on resource.
	IsAuthorized(ctx context.Context, identityID models.IdentityID, operation *models.Operation, resourceID models.ResourceID) (bool, error)
	// CreateGrantsForIdentity grants the specified identity a set of permissions.
	// For each operation in the supplied list, the identity will be allowed to perform the
	// specified operation on the specified resource or on any resource it owns (directly or indirectly),
	// as long as the resource kind matches the kind specified in the operation.
	CreateGrantsForIdentity(
		ctx context.Context,
		txOrNil *store.Tx,
		grantedByLegalEntityID models.LegalEntityID,
		authorizedIdentityID models.IdentityID,
		operations []*models.Operation,
		resourceID models.ResourceID,
	) error
	// CreateGrantsForGroup grants the specified access control Group a set of permissions.
	// For each operation in the supplied list, identities in the specified group will be allowed to perform the
	// specified operation on the specified resource or on any resource it owns (directly or indirectly),
	// as long as the resource kind matches the kind specified in the operation.
	CreateGrantsForGroup(
		ctx context.Context,
		txOrNil *store.Tx,
		grantedByLegalEntityID models.LegalEntityID,
		authorizedGroupID models.GroupID,
		operations []*models.Operation,
		resourceID models.ResourceID,
	) error
	// DeleteGrant permanently and idempotently deletes a grant, identifying it by id.
	DeleteGrant(ctx context.Context, txOrNil *store.Tx, id models.GrantID) error
	// DeleteAllGrantsForIdentity permanently and idempotently deletes all grants for the specified identity.
	DeleteAllGrantsForIdentity(ctx context.Context, txOrNil *store.Tx, identityID models.IdentityID) error
	// FindOrCreateGrant finds and returns a grant with the data specified in the supplied grant data.
	// The GrantStore.ReadByAuthorizedOperation function is used to find matching grants.
	// If no such grant exists then a new one is created and returned, and true is returned for 'created'.
	FindOrCreateGrant(ctx context.Context, txOrNil *store.Tx, grantData *models.Grant) (grant *models.Grant, created bool, err error)
	// ListGrantsForGroup finds and returns all grants that give permissions to the specified group.
	ListGrantsForGroup(ctx context.Context, txOrNil *store.Tx, groupID models.GroupID, pagination models.Pagination) ([]*models.Grant, *models.Cursor, error)
}

type GroupService interface {
	// ReadByName reads an existing access control Group, looking it up by group name and the ID of the
	// legal entity that owns the group. Returns models.ErrNotFound if the group does not exist.
	ReadByName(ctx context.Context, txOrNil *store.Tx, ownerLegalEntityID models.LegalEntityID, groupName models.ResourceName) (*models.Group, error)
	// ReadByExternalID reads an existing group, looking it up by its unique external id.
	// Returns models.ErrNotFound if the group does not exist.
	ReadByExternalID(ctx context.Context, txOrNil *store.Tx, externalID models.ExternalResourceID) (*models.Group, error)
	// FindOrCreateStandardGroup finds or creates a new group for a legal entity and sets up permissions for
	// any new group, based on the supplied standard group definition.
	FindOrCreateStandardGroup(ctx context.Context, tx *store.Tx, legalEntity *models.LegalEntity, groupDefinition *models.StandardGroupDefinition) (*models.Group, error)
	// FindOrCreateByName finds and returns the access control Group with the name and legal entity specified in
	// the supplied group data.
	// If no such group exists then a new group is created and returned, and true is returned for 'created'.
	FindOrCreateByName(ctx context.Context, txOrNil *store.Tx, groupData *models.Group) (*models.Group, bool, error)
	// UpsertByExternalID creates a group if no group with the same External ID already exists, otherwise it updates
	// the existing group's mutable properties if they differ from the in-memory instance.
	// Returns true,false if the resource was created, false,true if the resource was updated, or false,false if
	// neither create nor update was necessary.
	// Returns an error if no External ID is filled out in the supplied Group.
	// In all cases group.ID will be filled out in the supplied group object.
	UpsertByExternalID(ctx context.Context, txOrNil *store.Tx, group *models.Group) (bool, bool, error)
	// Delete permanently and idempotently deletes an access control group.
	// All memberships and grants for this group will also be permanently deleted.
	Delete(ctx context.Context, txOrNil *store.Tx, id models.GroupID) error
	// ListGroups returns a list of groups. Use cursor to page through results, if any.
	// If groupParent is provided then only groups owned by the supplied parent legal entity will be returned.
	// If memberID is provided then only groups that include the provided identity as a member will be returned.
	ListGroups(ctx context.Context, txOrNil *store.Tx, groupParent *models.LegalEntityID, memberID *models.IdentityID, pagination models.Pagination) ([]*models.Group, *models.Cursor, error)
	// FindOrCreateMembership adds the specified identity to an access control Group by adding a group membership
	// for a specific source system.
	// This method is idempotent, and returns true if a new membership was created or false if there was already
	// a membership for this identity for the group with the specified source system
	FindOrCreateMembership(ctx context.Context, txOrNil *store.Tx, membershipData *models.GroupMembershipData) (membership *models.GroupMembership, created bool, err error)
	// RemoveMembership removes a membership for the specified identity from an access control group.
	// If sourceSystem is not nil then only the membership record matching the source system will be deleted;
	// otherwise membership records from all source systems for the member will be deleted.
	// This method is idempotent.
	RemoveMembership(ctx context.Context, txOrNil *store.Tx, groupID models.GroupID, memberID models.IdentityID, sourceSystem *models.SystemName) error
	// ReadMembership reads an existing access control group membership, looking it up by group member, identity and
	// source system. Returns models.ErrNotFound if the group membership does not exist.
	ReadMembership(ctx context.Context, txOrNil *store.Tx, groupID models.GroupID, memberID models.IdentityID, sourceSystem models.SystemName) (*models.GroupMembership, error)
	// ListGroupMemberships returns a list of group memberships. Use cursor to page through results, if any.
	// If groupID is provided then only memberships of the specified group will be returned.
	// If memberID is provided then only groups that include the provided identity as a member will be returned.
	// If sourceSystem is provided then only memberships with matching source system values will be returned.
	ListGroupMemberships(
		ctx context.Context,
		txOrNil *store.Tx,
		groupID *models.GroupID,
		memberID *models.IdentityID,
		sourceSystem *models.SystemName,
		pagination models.Pagination,
	) ([]*models.GroupMembership, *models.Cursor, error)
}

type AuthenticationService interface {
	// AuthenticateJWT authenticates an identity using a JWT.
	AuthenticateJWT(ctx context.Context, jwt string) (*models.Identity, error)
}

type CredentialService interface {
	// Create a new credential.
	Create(ctx context.Context, txOrNil *store.Tx, credential *models.Credential) error
	// CreateSharedSecretCredential creates a new shared secret credential for the specified identity.
	// The plaintext shared secret value is returned.
	// The plaintext can never be reconstructed so do something useful with it now.
	CreateSharedSecretCredential(
		ctx context.Context,
		txOrNil *store.Tx,
		identityID models.IdentityID,
		enabled bool,
	) (models.PublicSharedSecretToken, *models.Credential, error)
	// CreateClientCertificateCredential creates a new client certificate credential for the specified identity.
	// clientCert is an X.509 certificate used to identify the client.
	CreateClientCertificateCredential(
		ctx context.Context,
		txOrNil *store.Tx,
		identityID models.IdentityID,
		enabled bool,
		clientCert certificates.CertificateData,
	) (*models.Credential, error)
	// CreateIdentityJWT creates a new JWT (JSON Web Token) credential that can be used to authenticate as
	// the specified identity.
	// The JWT will not be stored but can be verified by calling VerifyJWTIdentityCredential, or by any third
	// party using the server's JTW public key.
	CreateIdentityJWT(identityID models.IdentityID) (string, error)
	// VerifyIdentityJWT verifies the signature on the supplied JWT (JSON Web Token) and returns the identity ID
	// specified in the subject field. The identity ID is NOT checked against the database.
	VerifyIdentityJWT(token string) (models.IdentityID, error)
	// Delete permanently and idempotently deletes a credential.
	Delete(ctx context.Context, txOrNil *store.Tx, id models.CredentialID) error
	// ListCredentialsForIdentity returns a list of all credentials for the specified identity ID.
	// Use cursor to page through results, if any.
	ListCredentialsForIdentity(
		ctx context.Context,
		txOrNil *store.Tx,
		identityID models.IdentityID,
		pagination models.Pagination,
	) ([]*models.Credential, *models.Cursor, error)
}

type ArtifactService interface {
	// Create a new artifact with its contents provided by reader. It is the caller's responsibility to close reader.
	// Optionally specify expectedMD5 to verify the file contents matches the expected MD5.
	// If storeData is true then the artifact data obtained from the reader will be stored in the blob store.
	Create(
		ctx context.Context,
		jobID models.JobID,
		groupName models.ResourceName,
		relativePath string,
		expectedMD5 string,
		reader io.Reader,
		storeData bool,
	) (*models.Artifact, error)
	// Read an existing artifact, looking it up by ID.
	Read(ctx context.Context, txOrNil *store.Tx, id models.ArtifactID) (*models.Artifact, error)
	// Search all artifacts. If searcher is set, the results will be limited to artifacts the searcher is authorized to
	// see (via the read:artifact permission). Use cursor to page through results, if any.
	Search(ctx context.Context, txOrNil *store.Tx, searcher models.IdentityID, search models.ArtifactSearch) ([]*models.Artifact, *models.Cursor, error)
	// GetArtifactData returns a reader to the data of an artifact.
	// It is the callers responsibility to close reader.
	GetArtifactData(ctx context.Context, artifactID models.ArtifactID) (io.ReadCloser, error)
}

type LegalEntityService interface {
	// Create creates a new legal entity and configures default access control rules.
	Create(ctx context.Context, txOrNil *store.Tx, legalEntityData *models.LegalEntityData) (*models.LegalEntity, error)
	// Read an existing legal entity, looking it up by ID.
	Read(ctx context.Context, txOrNil *store.Tx, id models.LegalEntityID) (*models.LegalEntity, error)
	// ReadByExternalID reads an existing legal entity, looking it up by its external id.
	// Returns models.ErrNotFound if the legal entity does not exist.
	ReadByExternalID(ctx context.Context, txOrNil *store.Tx, externalID models.ExternalResourceID) (*models.LegalEntity, error)
	// ReadByIdentityID reads an existing legal entity, looking it up by the ID of its associated Identity.
	ReadByIdentityID(ctx context.Context, txOrNil *store.Tx, identityID models.IdentityID) (*models.LegalEntity, error)
	// ReadIdentity reads and returns the Identity for the specified Legal Entity.
	ReadIdentity(ctx context.Context, txOrNil *store.Tx, id models.LegalEntityID) (*models.Identity, error)
	// FindOrCreate creates a legal entity if no legal entity with the same External ID already exists,
	// otherwise it reads and returns the existing legal entity.
	// Returns the legal entity as it is in the database, and true iff a new legal entity was created.
	FindOrCreate(ctx context.Context, txOrNil *store.Tx, legalEntityData *models.LegalEntityData) (legalEntity *models.LegalEntity, created bool, err error)
	// Upsert creates a legal entity if no legal entity with the same External ID already exists, otherwise it updates
	// the existing legal entity's data if it differs from the supplied data.
	// Returns the LegalEntity as it exists in the database after the create or update, and
	// true,false if the resource was created, false,true if the resource was updated, or false,false if
	// neither create nor update was necessary.
	Upsert(ctx context.Context, txOrNil *store.Tx, legalEntityData *models.LegalEntityData) (*models.LegalEntity, bool, bool, error)
	// Update an existing legal entity with optimistic locking. Overrides all previous values using the supplied model.
	// Returns store.ErrOptimisticLockFailed if there is an optimistic lock mismatch.
	Update(ctx context.Context, txOrNil *store.Tx, legalEntity *models.LegalEntity) error
	// ListParentLegalEntities lists all legal entities a legal entity is a member of. Use cursor to page through results, if any.
	ListParentLegalEntities(ctx context.Context, txOrNil *store.Tx, legalEntityID models.LegalEntityID, pagination models.Pagination) ([]*models.LegalEntity, *models.Cursor, error)
	// ListMemberLegalEntities lists all legal entities that are members of a parent legal entity. Use cursor to page through results, if any.
	ListMemberLegalEntities(ctx context.Context, txOrNil *store.Tx, parentLegalEntityID models.LegalEntityID, pagination models.Pagination) ([]*models.LegalEntity, *models.Cursor, error)
	// ListAllLegalEntities lists all legal entities in the system. Use cursor to page through results, if any.
	ListAllLegalEntities(ctx context.Context, txOrNil *store.Tx, pagination models.Pagination) ([]*models.LegalEntity, *models.Cursor, error)
	// AddCompanyMember adds a user as a member of a particular company. The legal entity for the user and the company
	// must already exist in the database. This method is idempotent.
	AddCompanyMember(ctx context.Context, txOrNil *store.Tx, companyID models.LegalEntityID, memberID models.LegalEntityID) error
	// RemoveCompanyMember removes records for a user who is no longer a member of a particular company.
	// The user will be removed from all groups owned by the company, and removed as a member of the company's legal entity.
	// The legal entity for the user and the company must already exist in the database.
	RemoveCompanyMember(ctx context.Context, txOrNil *store.Tx, companyID models.LegalEntityID, memberID models.LegalEntityID) error
}

type BuildService interface {
	// Create a new build.
	// Returns store.ErrAlreadyExists if a build with matching unique properties already exists.
	Create(ctx context.Context, txOrNil *store.Tx, build *models.Build) error
	// Update an existing build with optimistic locking. Overrides all previous values using the supplied model.
	// Returns store.ErrOptimisticLockFailed if there is an optimistic lock mismatch.
	Update(ctx context.Context, txOrNil *store.Tx, build *models.Build) error
	// Read an existing build, looking it up by ID.
	// Returns models.ErrNotFound if the build does not exist.
	Read(ctx context.Context, txOrNil *store.Tx, id models.BuildID) (*models.Build, error)
	// ReadByIdentityID looks up the build that corresponds to the specified identity, or returns a not found error
	// if the identity doesn't correspond to a build.
	ReadByIdentityID(ctx context.Context, txOrNil *store.Tx, identityID models.IdentityID) (*models.Build, error)
	// FindOrCreateIdentity returns an Identity that has permission to read and add jobs for a specific build only,
	// for use by dynamic jobs running as part of that build.
	// If no identity exists for the build then a new identity is created and returned.
	FindOrCreateIdentity(ctx context.Context, txOrNil *store.Tx, buildID models.BuildID) (*models.Identity, error)
	// DeleteIdentity deletes any existing Identity associated with a build.
	DeleteIdentity(ctx context.Context, txOrNil *store.Tx, buildID models.BuildID) error
	// LockRowForUpdate takes out an exclusive row lock on the build table row for the specified build.
	// This function must be called within a transaction, and will block other transactions from locking, updating
	// or deleting the row until this transaction ends.
	LockRowForUpdate(ctx context.Context, tx *store.Tx, id models.BuildID) error
	// Search all builds. If a searcher identity is provided then the search will be constrained to include only
	// results that the identity has access to. Use cursor to page through results, if any.
	Search(ctx context.Context, txOrNil *store.Tx, searcher models.IdentityID, search *models.BuildSearch) ([]*models.BuildSearchResult, *models.Cursor, error)
	// Summary returns a summary of builds for the given legalEntityId. If searcher is set, the results will be limited
	// to build(s) the searcher is authorized to see (via the read:build permission).
	Summary(ctx context.Context, txOrNil *store.Tx, legalEntityId models.LegalEntityID, searcher models.IdentityID) (*models.BuildSummaryResult, error)
	// UniversalSearch searches all builds. If searcher is set, the results will be limited to build(s) the searcher is authorized to
	// see (via the read:build permission). Use cursor to page through results, if any.
	UniversalSearch(ctx context.Context, txOrNil *store.Tx, searcher models.IdentityID, search search.Query) ([]*models.BuildSearchResult, *models.Cursor, error)
}

type JobService interface {
	// Create a new job.
	// Returns store.ErrAlreadyExists if a job with matching unique properties already exists.
	Create(ctx context.Context, txOrNil *store.Tx, create *dto.CreateJob) error
	// Read an existing job, looking it up by ID.
	// Returns models.ErrNotFound if the job does not exist.
	Read(ctx context.Context, txOrNil *store.Tx, id models.JobID) (*models.Job, error)
	// ReadByFingerprint reads the most recent successful job inside a repo with a matching workflow, name
	// and fingerprint. Returns models.ErrNotFound if the job does not exist.
	ReadByFingerprint(
		ctx context.Context,
		txOrNil *store.Tx,
		repoID models.RepoID,
		workflow models.ResourceName,
		jobName models.ResourceName,
		jobFingerprint string,
		jobFingerprintHashType *models.HashType) (*models.Job, error)
	// Update an existing job with optimistic locking. Overrides all previous values using the supplied model.
	// Returns store.ErrOptimisticLockFailed if there is an optimistic lock mismatch.
	Update(ctx context.Context, txOrNil *store.Tx, job *models.Job) error
	// ListDependencies lists all jobs that the specified job depends on.
	ListDependencies(ctx context.Context, txOrNil *store.Tx, jobID models.JobID) ([]*models.Job, error)
	// FindQueuedJob locates a queued job that the runner is capable of running, and which is ready for
	// execution (e.g all dependencies are completed).
	FindQueuedJob(ctx context.Context, txOrNil *store.Tx, runner *models.Runner) (*models.Job, error)
	// ListByBuildID gets all jobs that are associated with the specified build id.
	ListByBuildID(ctx context.Context, txOrNil *store.Tx, id models.BuildID) ([]*models.Job, error)
	// ListByStatus returns all jobs that have the specified status, regardless of who owns the jobs or which build
	// they are part of. Use cursor to page through results, if any.
	ListByStatus(ctx context.Context, txOrNil *store.Tx, status models.WorkflowStatus, pagination models.Pagination) ([]*models.Job, *models.Cursor, error)
}

type StepService interface {
	// Create a new step.
	// Returns store.ErrAlreadyExists if a step with matching unique properties already exists.
	Create(ctx context.Context, txOrNil *store.Tx, create *dto.CreateStep) error
	// Read an existing step, looking it up by ID.
	// Returns models.ErrNotFound if the step does not exist.
	Read(ctx context.Context, txOrNil *store.Tx, id models.StepID) (*models.Step, error)
	// Update an existing step with optimistic locking. Overrides all previous values using the supplied model.
	// Returns store.ErrOptimisticLockFailed if there is an optimistic lock mismatch.
	Update(ctx context.Context, txOrNil *store.Tx, step *models.Step) error
	// ListByJobID gets all steps that are associated with the specified job id.
	ListByJobID(ctx context.Context, txOrNil *store.Tx, id models.JobID) ([]*models.Step, error)
}

type RunnerService interface {
	// Create a new runner. clientCert is an optional ASN.1 DER-encoded X.509 client certificate; if provided then
	// a client certificate credential will be created for authentication using TLS mutual authentication.
	// Returns store.ErrAlreadyExists if a runner with matching unique properties already exists.
	Create(ctx context.Context, txOrNil *store.Tx, runner *models.Runner, clientCert []byte) error
	// Read an existing runner, looking it up by ID.
	// Returns models.ErrNotFound if the runner does not exist.
	Read(ctx context.Context, txOrNil *store.Tx, id models.RunnerID) (*models.Runner, error)
	// ReadByName reads an existing runner, looking it up by name and the ID of the legal entity that owns the runner.
	// Returns models.ErrNotFound if the runner is not found.
	ReadByName(ctx context.Context, txOrNil *store.Tx, legalEntityID models.LegalEntityID, name models.ResourceName) (*models.Runner, error)
	// ReadByIdentityID reads an existing runner, looking it up by the ID of its associated Identity.
	ReadByIdentityID(ctx context.Context, txOrNil *store.Tx, identityID models.IdentityID) (*models.Runner, error)
	// ReadIdentity reads and returns the Identity for the specified runner.
	ReadIdentity(ctx context.Context, txOrNil *store.Tx, id models.RunnerID) (*models.Identity, error)
	// Update an existing runner.
	// Returns store.ErrOptimisticLockFailed if there is an optimistic lock mismatch.
	Update(ctx context.Context, txOrNil *store.Tx, runner *models.Runner) (*models.Runner, error)
	// RunnerCompatibleWithJob returns true if a runner exists that is capable of running job.
	RunnerCompatibleWithJob(ctx context.Context, txOrNil *store.Tx, job *models.Job) (bool, error)
	// SoftDelete soft deletes an existing runner.
	// Returns store.ErrOptimisticLockFailed if there is an optimistic lock mismatch.
	SoftDelete(ctx context.Context, txOrNil *store.Tx, runnerID models.RunnerID, delete dto.DeleteRunner) error
	// Search all runners. If searcher is set, the results will be limited to runners the searcher is authorized to
	// see (via the read:runner permission). Use cursor to page through results, if any.
	Search(ctx context.Context, txOrNil *store.Tx, searcher models.IdentityID, search models.RunnerSearch) ([]*models.Runner, *models.Cursor, error)
}

type EventService interface {
	// PublishEvent publishes a new event. Subscribers matching the event type and resource will be notified.
	PublishEvent(ctx context.Context, txOrNil *store.Tx, eventData *models.EventData) error
	// FetchEvents fetches new events for a given build, i.e. those with event numbers greater than lastEventNumber.
	// limit specifies the maximum number of events to return.
	// Events will be returned in order of event number; event numbers provide a unique ordering within a build.
	// If no new events are available then the function returns immediately.
	FetchEvents(ctx context.Context, txOrNil *store.Tx, buildID models.BuildID, lastEventNumber models.EventNumber, limit int) ([]*models.Event, error)
}
