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
