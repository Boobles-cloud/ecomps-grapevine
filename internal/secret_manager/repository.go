package secretmanager

import (
	"context"
	"uuid"

	smstructs "grapevine.boobles.cloud/internal/secret_manager/sm_structs"
)

type SecretManagerRepository interface {
	GetTenantPw(ctx context.Context, tenantId uuid.UUID) (smstructs.SecretEntry, error)
	GetClientSecret(ctx context.Context, appId uint) (smstructs.SecretEntry, error)
	GetAllClientSecrets(ctx context.Context) ([]smstructs.SecretEntry, error)
}
