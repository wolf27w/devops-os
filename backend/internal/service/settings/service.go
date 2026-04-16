package settings

type SettingsService struct{}

func NewSettingsService() *SettingsService {
	return &SettingsService{}
}

func (s *SettingsService) GetProjectSettings(projectID string) (interface{}, error) {
	return nil, nil
}

func (s *SettingsService) UpdateProjectSettings(projectID string, data interface{}) (interface{}, error) {
	return nil, nil
}

func (s *SettingsService) GetUserSettings(userID string) (interface{}, error) {
	return nil, nil
}

func (s *SettingsService) UpdateUserSettings(userID string, data interface{}) (interface{}, error) {
	return nil, nil
}
