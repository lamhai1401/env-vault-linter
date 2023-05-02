package env

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/joho/godotenv"
)

type Mapper struct {
	PostgrestDatabase string `config:"SECRET_POSTGRES_DATABASE,required"`
	PostgrestHostname string `config:"SECRET_POSTGRES_HOSTNAME,required"`
	PostgrestPort     string `config:"SECRET_POSTGRES_PORT,required"`
	PostgrestPassword string `config:"SECRET_POSTGRES_PASSWORD,required"`
	PostgrestUser     string `config:"SECRET_POSTGRES_USER,required"`
	// SendGridAPIKey    string `config:"SENDGRID_API_KEY, required"`
	// AWS config
	AWSConfigFile      string `config:"aws_access_key_id"`
	AWSCredentialsFile string `config:"aws_secret_access_key"`
}

var envData = Mapper{}

// Secret is an exported getter function.

func Load() Mapper {
	loadEnv(os.Getenv("ENV"))
	return envData
}

func loadEnv(target string) {
	// get filepath of this file
	_, base, _, _ := runtime.Caller(1)
	base = filepath.Dir(base)
	f := filepath.Join(base, "..", "..", fmt.Sprintf(".env.%s", target))

	if f == filepath.Join("..", "..", ".env.") {
		fmt.Println("NOTE: did you set 'ENV' environment variable? For testing, set 'ENV=test'.")
	}

	if err := godotenv.Load(f); err != nil {
		// When dotenv files are not found, print warning and continue the context.
		fmt.Println("NOTE:", err)
		fmt.Println("NOTE: falling back to environment variable checks..")
	}

	loader := confita.NewLoader(
		env.NewBackend(),
	)

	if err := loader.Load(context.Background(), &envData); err != nil {
		panic(err)
	}
}

func Clear() {
	envData = Mapper{}
}
