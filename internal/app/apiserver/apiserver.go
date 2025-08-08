package apiserver

type APIServer struct{
	config *Config
}

func New(Config *Config) *APIServer {
	return &APIServer{
		config: Config,
	}
}


func (s *APIServer) Start() error {
	return nil
}
