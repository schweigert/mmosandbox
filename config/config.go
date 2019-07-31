package config

// Web envs
func Web() *WebEnvs {
	return &WebEnvs{}
}

// Db envs
func Db() *DbEnvs {
	return &DbEnvs{}
}

// Client envs
func Client() *ClientEnvs {
	return &ClientEnvs{}
}

// Metric envs
func Metric() *MetricEnvs {
	return &MetricEnvs{}
}

// Service envs
func Service() *ServiceEnvs {
	return &ServiceEnvs{}
}
