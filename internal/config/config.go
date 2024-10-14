// package config

// import "time"
// 		"os"

// type Config struct {
// 	Env         string `yaml:"env" env:"ENV" env-default:"local"`
// 	StoragePath string `yaml:"storage_path" env-required:"true"`
// 	HTTPServer `yaml`: "http_server"
// }

// type HTTPServer struct {
// 	Address     string        `yaml:"address" env-default:"localhost:8081"`
// 	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
// 	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
// }

// func MustLoad() {
// 	configPath := os.Getenv("CONFIG_PATH")
// 	if configPath == "" {
// 		log.Fatal("CONFIG_PATH is not set")
// 	}

// 	if _, err := os.Stat(configPath); os.IsNotExist(err) {
// 		log.Fatalf("config file %s does not exist", configPath)
// 	}
// }