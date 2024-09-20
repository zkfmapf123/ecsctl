package base

type APIResources struct {
	Cluster    []string
	Services   []string
	Containers []string
	Tasks      []string
	Alb        []string
}

func GetAPIResources() APIResources {
	return APIResources{
		Cluster:    []string{"c", "cluster", "clu"},
		Services:   []string{"s", "service", "svc"},
		Containers: []string{"co", "containers", "con"},
		Tasks:      []string{"t", "tasks", "tsk"},
		Alb:        []string{"al", "alb"},
	}
}
