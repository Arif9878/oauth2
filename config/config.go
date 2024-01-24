package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	echoserver "github.com/Arif9878/common/go/http/echo/server"
	"github.com/Arif9878/common/go/logger"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "products write microservice config path")
}

type Config struct {
	ServiceName string                 `mapstructure:"service-name"`
	Logger      *LoggerConfig          `mapstructure:"logger"`
	Echo        *echoserver.EchoConfig `mapstructure:"echo"`
	// Rabbitmq     *rabbitmq.RabbitMQConfig      `mapstructure:"rabbitmq"`
	// Grpc         *grpc.GrpcConfig              `mapstructure:"grpc"`
	// GormPostgres *gormpgsql.GormPostgresConfig `mapstructure:"gormPostgres"`
	// Jaeger       *otel.JaegerConfig            `mapstructure:"jaeger"`
	// 	Database      DatabaseConfig
	// 	Oauth         OauthConfig
	// 	Session       SessionConfig
	// 	IsDevelopment bool
}

type Context struct {
	Timeout int `mapstructure:"timeout"`
}

type LoggerConfig struct {
	LogLevel string `mapstructure:"level"`
}

func InitConfig() (*Config, *LoggerConfig, *echoserver.EchoConfig, error) {
	// , *otel.JaegerConfig, *gormpgsql.GormPostgresConfig,
	// *grpc.GrpcConfig, , *rabbitmq.RabbitMQConfig, error) {

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	if configPath == "" {
		configPathFromEnv := os.Getenv("CONFIG_PATH")
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			//https://stackoverflow.com/questions/31873396/is-it-possible-to-get-the-current-root-of-package-structure-as-a-string-in-golan
			//https://stackoverflow.com/questions/18537257/how-to-get-the-directory-of-the-currently-running-file
			d, err := dirname()
			if err != nil {
				return nil, nil, nil, err
			}

			configPath = d
		}
	}

	cfg := &Config{}

	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, nil, nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, nil, nil, errors.Wrap(err, "viper.Unmarshal")
	}

	return cfg, cfg.Logger, cfg.Echo, nil
	// cfg.Jaeger, cfg.GormPostgres, cfg.Grpc,  cfg.Rabbitmq, nil
}

// InitLogger Init logger
func InitLogger(cfg *LoggerConfig) logger.ILogger {

	l := &logger.AppLogger{Level: cfg.LogLevel}

	l.Logger = log.StandardLogger()

	logLevel := l.GetLevel()

	env := os.Getenv("APP_ENV")

	if env == "production" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		// The TextFormatter is default, you don't actually have to do this.
		log.SetFormatter(&log.TextFormatter{
			DisableColors: false,
			ForceColors:   true,
			FullTimestamp: true,
		})
	}

	log.SetLevel(logLevel)

	return logger.Logger
}

func GetMicroserviceName(serviceName string) string {
	return fmt.Sprintf("%s", strings.ToUpper(serviceName))
}

func filename() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	return filename, nil
}

func dirname() (string, error) {
	filename, err := filename()
	if err != nil {
		return "", err
	}
	return filepath.Dir(filename), nil
}
