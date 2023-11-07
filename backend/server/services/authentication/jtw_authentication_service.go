package authentication

import (
	"context"
	"fmt"

	"github.com/buildbeaver/buildbeaver/common/logger"
	"github.com/buildbeaver/buildbeaver/common/models"
	"github.com/buildbeaver/buildbeaver/server/services"
	"github.com/buildbeaver/buildbeaver/server/store"
)

// JWTAuthenticationService is an implementation of AuthenticationService that supports authentication
// only using JWT tokens.
type JWTAuthenticationService struct {
	db                *store.DB
	identityStore     store.IdentityStore
	credentialService services.CredentialService
	logger.Log
}

func NewJWTAuthenticationService(
	db *store.DB,
	identityStore store.IdentityStore,
	credentialService services.CredentialService,
	logFactory logger.LogFactory) *JWTAuthenticationService {

	return &JWTAuthenticationService{
		db:                db,
		identityStore:     identityStore,
		credentialService: credentialService,
		Log:               logFactory("AuthenticationService"),
	}
}

// AuthenticateJWT authenticates an identity using a JWT.
func (s *JWTAuthenticationService) AuthenticateJWT(ctx context.Context, jwt string) (*models.Identity, error) {
	// Verify the JWT and extract an Identity ID
	identityID, err := s.credentialService.VerifyIdentityJWT(jwt)
	if err != nil {
		return nil, err
	}

	// Check the identity is in the database
	identity, err := s.identityStore.Read(ctx, nil, identityID)
	if err != nil {
		return nil, fmt.Errorf("error reading legal entity for identity ID specified in JWT: %w", err)
	}

	return identity, nil
}
