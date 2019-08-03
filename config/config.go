package config

// Web envs
func Web() *WebEnvs {
	return &WebEnvs{}
}

// RPC envs
func RPC() *RPCEnvs {
	return &RPCEnvs{}
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

// Me envs
func Me() *MeEnvs {
	return &MeEnvs{}
}
