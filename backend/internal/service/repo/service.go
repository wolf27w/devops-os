package repo

type RepoService struct{}

func NewRepoService() *RepoService {
	return &RepoService{}
}

func (s *RepoService) GetRepositories(projectID string) ([]interface{}, error) {
	return []interface{}{}, nil
}

func (s *RepoService) GetRepository(projectID, repoID string) (interface{}, error) {
	return nil, nil
}

func (s *RepoService) CreateRepository(projectID string, data interface{}) (interface{}, error) {
	return nil, nil
}

func (s *RepoService) UpdateRepository(projectID, repoID string, data interface{}) (interface{}, error) {
	return nil, nil
}

func (s *RepoService) DeleteRepository(projectID, repoID string) error {
	return nil
}
