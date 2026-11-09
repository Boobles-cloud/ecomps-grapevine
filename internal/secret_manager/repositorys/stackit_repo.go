package repositorys

import (
	"context"
	"errors"
	"strconv"
	"uuid"

	vault "github.com/hashicorp/vault/api"
	"grapevine.boobles.cloud/global"
	smstructs "grapevine.boobles.cloud/internal/secret_manager/sm_structs"
)

type StackitRepository struct {
	client                  *vault.Client
	defaultTenantPath       string
	defaultClientSecretPath string
	instanceId              string
}

func CreateNewStackitRepository(c *vault.Client, dTP, dCP, instanceId string) *StackitRepository {
	return &StackitRepository{
		client:                  c,
		defaultTenantPath:       dTP,
		defaultClientSecretPath: dCP,
		instanceId:              instanceId,
	}
}

func (s *StackitRepository) GetTenantPw(ctx context.Context, tenantId uuid.UUID) (smstructs.SecretEntry, error) {

	secret, err := s.client.KVv2(s.instanceId).Get(ctx, tenantId.String())

	if err != nil {
		return smstructs.SecretEntry{}, err
	}

	dataMap, ok := secret.Data["tenantPw"].(string)

	if !ok {
		return smstructs.SecretEntry{}, global.ERR_FIELD_NOT_FOUND
	}

	return smstructs.SecretEntry{
		Key:   tenantId.String(),
		Value: dataMap,
	}, nil
}

func (s *StackitRepository) GetClientSecret(ctx context.Context, appId uint) (smstructs.SecretEntry, error) {

	secret, err := s.client.KVv2(s.instanceId).Get(ctx, "Test123")

	if err != nil {
		return smstructs.SecretEntry{}, err
	}

	dataMap, ok := secret.Data["client_secret"].(string)

	if !ok {
		return smstructs.SecretEntry{}, global.ERR_FIELD_NOT_FOUND
	}

	return smstructs.SecretEntry{
		Key:   strconv.Itoa(int(appId)),
		Value: dataMap,
	}, nil
}

func (s *StackitRepository) GetAllClientSecrets(ctx context.Context) ([]smstructs.SecretEntry, error) {
	return []smstructs.SecretEntry{}, errors.ErrUnsupported
}
