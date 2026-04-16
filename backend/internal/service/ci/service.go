package ci

type CIService struct{}

func NewCIService() *CIService {
	return &CIService{}
}

func (s *CIService) GetPipelines(projectID string) ([]interface{}, error) {
	return []interface{}{}, nil
}

func (s *CIService) GetPipeline(projectID, pipelineID string) (interface{}, error) {
	return nil, nil
}

func (s *CIService) CreatePipeline(projectID string, data interface{}) (interface{}, error) {
	return nil, nil
}

func (s *CIService) UpdatePipeline(projectID, pipelineID string, data interface{}) (interface{}, error) {
	return nil, nil
}

func (s *CIService) DeletePipeline(projectID, pipelineID string) error {
	return nil
}

func (s *CIService) RunPipeline(projectID, pipelineID string) (interface{}, error) {
	return nil, nil
}
