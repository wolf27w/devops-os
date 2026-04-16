package db

type DBService struct{}

func NewDBService() *DBService {
	return &DBService{}
}

func (s *DBService) GetConnections(projectID string) ([]interface{}, error) {
	return []interface{}{}, nil
}

func (s *DBService) GetConnection(projectID, connectionID string) (interface{}, error) {
	return nil, nil
}

func (s *DBService) CreateConnection(projectID string, data interface{}) (interface{}, error) {
	return nil, nil
}

func (s *DBService) UpdateConnection(projectID, connectionID string, data interface{}) (interface{}, error) {
	return nil, nil
}

func (s *DBService) DeleteConnection(projectID, connectionID string) error {
	return nil
}
