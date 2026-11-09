package startup

import (
	"context"
	"os"

	"github.com/hashicorp/vault/api"
	vault "github.com/hashicorp/vault/api"
	auth "github.com/hashicorp/vault/api/auth/approle"
	"github.com/hashicorp/vault/api/auth/userpass"
	"grapevine.boobles.cloud/internal/secret_manager/repositorys"
	"grapevine.boobles.cloud/internal/secret_manager/services"
)

// Deprecated: you cant create app-roles at the moment. So this has to wait..
func ConfigureSecretManager(ctx context.Context) (*services.SecretManagerService, error) {

	conf := vault.DefaultConfig()
	conf.Address = os.Getenv("secret_default_address")

	secretId := auth.SecretID{FromString: os.Getenv("secret_secretId")}
	appRoleAuth, err := auth.NewAppRoleAuth(os.Getenv("secret_role_id"), &secretId)

	if err != nil {
		return &services.SecretManagerService{}, err
	}

	client, err := vault.NewClient(conf)

	if err != nil {
		return &services.SecretManagerService{}, err
	}

	// Authenticate our client
	if _, err := client.Auth().Login(ctx, appRoleAuth); err != nil {
		return &services.SecretManagerService{}, err
	}

	// Our current provider
	repository := repositorys.CreateNewStackitRepository(client, os.Getenv("secret_tenant_path"), "secret_client_path", os.Getenv("secret_instance_id"))

	return services.CreateNewSecretManagerService(repository), nil
}

func ConfigureSecretManagerWitPw(ctx context.Context) (*services.SecretManagerService, error) {

	conf := vault.DefaultConfig()
	conf.Address = os.Getenv("secret_default_address")

	client, err := api.NewClient(conf)

	if err != nil {
		return &services.SecretManagerService{}, err
	}

	userpassAuth, err := userpass.NewUserpassAuth(
		os.Getenv("secret_user"),
		&userpass.Password{FromString: os.Getenv("secret_pw")},
	)

	if err != nil {
		return &services.SecretManagerService{}, err
	}

	if _, err := client.Auth().Login(ctx, userpassAuth); err != nil {
		return &services.SecretManagerService{}, err
	}

	repository := repositorys.CreateNewStackitRepository(client, os.Getenv("secret_tenant_path"), "secret_client_path", os.Getenv("secret_instance_id"))

	return services.CreateNewSecretManagerService(repository), nil
}
