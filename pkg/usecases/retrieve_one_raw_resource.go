package usecases

import "github.com/falcosecurity/cloud-native-security-hub/pkg/resource"

type RetrieveOneRawResource struct {
	ResourceRepository resource.Repository
	ResourceID         string
}

func (useCase *RetrieveOneRawResource) Execute() (resource *resource.Resource, err error) {
	res, err := useCase.ResourceRepository.FindById(useCase.ResourceID)
	if err != nil {
		return
	}
	return res, nil
}
