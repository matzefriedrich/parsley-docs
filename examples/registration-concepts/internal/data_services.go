package internal

import "fmt"

type DataService interface {
	FetchData() string
}

type remoteDataService struct {
	tenantId string
}

func (s *remoteDataService) FetchData() string {
	return fmt.Sprintf("Data from remote server for tenant %s", s.tenantId)
}

type localDataService struct{}

func (s *localDataService) FetchData() string {
	return "Data from local database"
}

func NewRemoteDataService() DataService {
	return &remoteDataService{}
}

func NewRemoteDataServiceFactory() func(string) DataService {
	return func(tenantId string) DataService {
		return &remoteDataService{
			tenantId: tenantId,
		}
	}
}

func NewLocalDataService() DataService {
	return &localDataService{}
}

func NewLocalDataServiceFactory() func(string) DataService {
	return func(_ string) DataService {
		return &localDataService{}
	}
}
