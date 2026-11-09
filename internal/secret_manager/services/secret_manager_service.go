package services

import (
	"context"
	"fmt"
	"uuid"

	"grapevine.boobles.cloud/global"
	smstructs "grapevine.boobles.cloud/internal/secret_manager/sm_structs"
)

// NOTE: only use this at the startup of the application!
func (s *SecretManagerService) SetupCache(ctx context.Context) error {

	entrys, err := s.repo.GetAllClientSecrets(ctx)

	if err != nil {
		return err
	}

	for i := range entrys {
		fmt.Println(entrys[i])
		s.cacheManager.SetOrUpdateItem(entrys[i].Key, entrys[i])
	}

	return nil
}

func (s *SecretManagerService) GetTenantPw(ctx context.Context, tenantId uuid.UUID) (smstructs.SecretEntry, error) {

	if tenantId.Compare(uuid.Nil()) == 0 || tenantId.Compare(uuid.Max()) == 0 {
		return smstructs.SecretEntry{}, global.ERR_TENANT_ID_INVALID
	}

	return s.repo.GetTenantPw(ctx, tenantId)
}

func (s *SecretManagerService) GetClientSecret(ctx context.Context, appKey string) (smstructs.SecretEntry, error) {

	if appKey == "" {
		return smstructs.SecretEntry{}, global.ERR_STRING_EMTPY
	}

	return s.repo.GetClientSecret(ctx, appKey)
}
