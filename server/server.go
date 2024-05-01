package server

import "github.com/chamodshehanka/gh-action-gen-api/config"

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("server.port"))
}
