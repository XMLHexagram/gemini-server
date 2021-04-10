package config

func ProvideLog() Log {
	return configService.Log
}

func ProvideGemini() Gemini {
	return configService.Gemini
}

func ProvideServer(serverName string) Server {
	return configService.ServerMap[serverName]
}

func ProvideTLS() TLS {
	return configService.TLS
}
