package workflow

type WorkflowService struct{}

func NewWorkflowService() *WorkflowService {
	return &WorkflowService{}
}

func (s *WorkflowService) GetWorkflows(projectID string) ([]interface{}, error) {
	return []interface{}{}, nil
}

func (s *WorkflowService) GetWorkflow(projectID, workflowID string) (interface{}, error) {
	return nil, nil
}

func (s *WorkflowService) CreateWorkflow(projectID string, data interface{}) (interface{}, error) {
	return nil, nil
}

func (s *WorkflowService) UpdateWorkflow(projectID, workflowID string, data interface{}) (interface{}, error) {
	return nil, nil
}

func (s *WorkflowService) DeleteWorkflow(projectID, workflowID string) error {
	return nil
}

func (s *WorkflowService) ExecuteWorkflow(projectID, workflowID string) (interface{}, error) {
	return nil, nil
}
