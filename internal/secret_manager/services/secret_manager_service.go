package services

import (
	"context"
)

// NOTE: only use this at the startup of the application!
func (s *SecretManagerService) SetupCache(ctx context.Context) error {

	entrys, err := s.repo.GetAllClientSecrets(ctx)

	if err != nil {
		return err
	}

	for i := range entrys {
		s.cacheManager.SetOrUpdateItem(entrys[i].Key, entrys[i])
	}

	return nil
}
