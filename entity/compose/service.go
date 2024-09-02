package compose

import "github.com/AmadlaOrg/hery/storage"

// NewComposeService to set up the compose service
func NewComposeService() *SComposer {
	return &SComposer{
		Storage: storage.NewStorageService(),
	}
}
