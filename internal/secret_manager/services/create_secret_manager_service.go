package services

import (
	secretmanager "grapevine.boobles.cloud/internal/secret_manager"
	smstructs "grapevine.boobles.cloud/internal/secret_manager/sm_structs"
	"grapevine.boobles.cloud/internal/utils"
)

type SecretManagerService struct {
	repo         secretmanager.SecretManagerRepository
	cacheManager *utils.CacheManager[smstructs.SecretEntry]
}

// Creates a new secret manager service
func CreateNewSecretManagerService(r secretmanager.SecretManagerRepository) *SecretManagerService {
	return &SecretManagerService{
		repo:         r,
		cacheManager: &utils.CacheManager[smstructs.SecretEntry]{},
	}
}
