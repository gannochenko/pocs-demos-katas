package interfaces

type TemporalWorkflowGroup interface {
	GetWorkflows() map[string]any
}

type TemporalActivityGroup interface {
	GetActivities() map[string]any
}
