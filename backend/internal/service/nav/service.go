package nav

type NavService struct{}

func NewNavService() *NavService {
	return &NavService{}
}

func (s *NavService) GetNavigation(projectID string) ([]interface{}, error) {
	return []interface{}{}, nil
}

func (s *NavService) GetItem(projectID, itemID string) (interface{}, error) {
	return nil, nil
}

func (s *NavService) CreateItem(projectID string, data interface{}) (interface{}, error) {
	return nil, nil
}

func (s *NavService) UpdateItem(projectID, itemID string, data interface{}) (interface{}, error) {
	return nil, nil
}

func (s *NavService) DeleteItem(projectID, itemID string) error {
	return nil
}
