package wiki

type WikiService struct{}

func NewWikiService() *WikiService {
	return &WikiService{}
}

func (s *WikiService) GetPages(projectID string) ([]interface{}, error) {
	return []interface{}{}, nil
}

func (s *WikiService) GetPage(projectID, pageID string) (interface{}, error) {
	return nil, nil
}

func (s *WikiService) CreatePage(projectID string, data interface{}) (interface{}, error) {
	return nil, nil
}

func (s *WikiService) UpdatePage(projectID, pageID string, data interface{}) (interface{}, error) {
	return nil, nil
}

func (s *WikiService) DeletePage(projectID, pageID string) error {
	return nil
}
