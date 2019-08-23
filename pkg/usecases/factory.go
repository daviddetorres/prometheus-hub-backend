package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
	"log"
	"os"
)

type Factory interface {
	NewRetrieveAllResourcesUseCase() *RetrieveAllResources
	NewRetrieveOneResourceUseCase(hash string) *RetrieveOneResource
	NewRetrieveAllVendorsUseCase() *RetrieveAllVendors
	NewRetrieveOneVendorUseCase(vendorID string) *RetrieveOneVendor

	NewResourcesRepository() resource.Repository
	NewVendorRepository() vendor.Repository
}

func NewFactory() Factory {
	return &factory{}
}

type factory struct {
}

func (f *factory) NewRetrieveAllResourcesUseCase() *RetrieveAllResources {
	return &RetrieveAllResources{
		ResourceRepository: f.NewResourcesRepository(),
	}
}

func (f *factory) NewRetrieveOneResourceUseCase(hash string) *RetrieveOneResource {
	return &RetrieveOneResource{
		ResourceRepository: f.NewResourcesRepository(),
		Hash:               hash,
	}
}

func (f *factory) NewRetrieveAllVendorsUseCase() *RetrieveAllVendors {
	return &RetrieveAllVendors{
		VendorRepository: f.NewVendorRepository(),
	}
}

func (f *factory) NewRetrieveOneVendorUseCase(vendorID string) *RetrieveOneVendor {
	return &RetrieveOneVendor{
		VendorRepository: f.NewVendorRepository(),
		VendorID:         vendorID,
	}
}

func (f *factory) NewResourcesRepository() resource.Repository {
	resourcesPath, ok := os.LookupEnv("RESOURCES_PATH")
	if !ok {
		log.Println("The RESOURCES_PATH env var is not set")
		os.Exit(1)
	}
	repo, err := resource.FromPath(resourcesPath)
	if err != nil {
		log.Println("the resource repository of type file does not exist")
		os.Exit(1)
	}
	return repo
}

func (f *factory) NewVendorRepository() vendor.Repository {
	vendorPath, ok := os.LookupEnv("VENDOR_PATH")
	if !ok {
		log.Println("The VENDOR_PATH env var is not set")
		os.Exit(1)
	}
	repo, err := vendor.FromPath(vendorPath)
	if err != nil {
		log.Println("the resource repository of type file does not exist")
		os.Exit(1)
	}
	return repo
}
