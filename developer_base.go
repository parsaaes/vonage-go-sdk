package nexmo

import "github.com/judy2k/nexmo/sling"

type DeveloperService struct {
	sling   *sling.Sling
	authSet *AuthSet
}

func newDeveloperService(base *sling.Sling, authSet *AuthSet) *DeveloperService {
	sling := base.Base("https://rest.nexmo.com/")
	return &DeveloperService{
		sling:   sling,
		authSet: authSet,
	}
}

func (c *DeveloperService) SetBaseURL(baseURL string) {
	c.sling.Base(baseURL)
}
