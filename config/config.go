package config
import(
	"encoding/json"
	"os"
)
type DBConfig struct {
    User     string `json:"user"`
    Password string `json:"password"`
    DBName   string `json:"dbname"`
    Host     string `json:"host"`
    Port     int    `json:"port"`
}
type Config struct {
    Database DBConfig `json:"database"`
}
func LoadConfig(filename string) (*Config, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    decoder := json.NewDecoder(file)
    config := &Config{}
    err = decoder.Decode(config)
    if err != nil {
        return nil, err
    }
    return config, nil
}