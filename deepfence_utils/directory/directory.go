package directory

import (
	"context"
	"os"
	"strconv"
	"sync"

	"github.com/deepfence/ThreatMapper/deepfence_utils/log"
)

const (
	GlobalDirKey   = NamespaceID("global")
	NonSaaSDirKey  = NamespaceID("default")
	DatabaseDirKey = NamespaceID("database")
	NamespaceKey   = "namespace"
)

type NamespaceID string

type RedisConfig struct {
	Endpoint string
	Password string
	Database int
}

type Neo4jConfig struct {
	Endpoint string
	Username string
	Password string
}

type PostgresqlConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	SslMode  string
}

type MinioConfig struct {
	Endpoint   string
	Username   string
	Password   string
	BucketName string
	Secure     bool
}

type DBConfigs struct {
	Redis    *RedisConfig
	Neo4j    *Neo4jConfig
	Postgres *PostgresqlConfig
	Minio    *MinioConfig
}

type namespaceDirectory struct {
	Directory map[NamespaceID]DBConfigs
	sync.RWMutex
}

var directory namespaceDirectory

func init() {
	directory = namespaceDirectory{
		Directory: map[NamespaceID]DBConfigs{},
	}
	minioCfg := initMinio()

	saasMode := false
	saasModeOn, has := os.LookupEnv("DEEPFENCE_SAAS_MODE")
	if !has {
		log.Warn().Msg("DEEPFENCE_SAAS_MODE defaults to: off")
	} else if saasModeOn == "on" {
		saasMode = true
	}

	directory.Lock()
	if !saasMode {
		redisCfg := initRedis()
		neo4jCfg := initNeo4j()
		postgresqlCfg := initPosgresql()
		directory.Directory[NonSaaSDirKey] = DBConfigs{
			Redis:    &redisCfg,
			Neo4j:    &neo4jCfg,
			Postgres: &postgresqlCfg,
			Minio:    nil,
		}
	}

	directory.Directory[GlobalDirKey] = DBConfigs{
		Redis:    nil,
		Neo4j:    nil,
		Postgres: nil,
		Minio:    &minioCfg,
	}
	directory.Unlock()
}

func GetAllNamespaces() []NamespaceID {
	directory.RLock()
	defer directory.RUnlock()
	var namespaces []NamespaceID
	for k, _ := range directory.Directory {
		if k != GlobalDirKey {
			namespaces = append(namespaces, k)
		}
	}
	return namespaces
}

func ForEachNamespace(applyFn func(ctx context.Context) (string, error)) {
	namespaces := GetAllNamespaces()
	var err error
	var msg string
	for _, ns := range namespaces {
		msg, err = applyFn(NewContextWithNameSpace(ns))
		if err != nil {
			log.Error().Msg(msg + ": " + err.Error())
		}
	}
}

func FetchNamespace(email string) NamespaceID {
	namespaces := GetAllNamespaces()
	if len(namespaces) == 1 && namespaces[0] == NonSaaSDirKey {
		return NonSaaSDirKey
	} else {
		// TODO: Fetch namespace for SaaS tenant
	}
	return ""
}

func FetchNamespaceFromID(namespaceID string) NamespaceID {
	namespaces := GetAllNamespaces()
	if len(namespaces) == 1 && namespaces[0] == NonSaaSDirKey {
		return NonSaaSDirKey
	} else {
		// TODO: Fetch namespace for SaaS tenant
	}
	return ""
}

func IsNonSaaSDeployment() bool {
	namespaces := GetAllNamespaces()
	if len(namespaces) == 1 && namespaces[0] == NonSaaSDirKey {
		return true
	}
	return false
}

func initRedis() RedisConfig {
	redisHost, has := os.LookupEnv("DEEPFENCE_REDIS_HOST")
	if !has {
		redisHost = "localhost"
		log.Warn().Msgf("DEEPFENCE_REDIS_HOST defaults to: %v", redisHost)
	}
	redisPort, has := os.LookupEnv("DEEPFENCE_REDIS_PORT")
	if !has {
		redisPort = "6379"
		log.Warn().Msgf("DEEPFENCE_REDIS_PORT defaults to: %v", redisPort)
	}
	redisEndpoint := redisHost + ":" + redisPort
	redisPassword := os.Getenv("DEEPFENCE_REDIS_PASSWORD")
	redisDbNumber := 0
	var err error
	redisDbNumberStr := os.Getenv("DEEPFENCE_REDIS_DB_NUMBER")
	if redisDbNumberStr != "" {
		redisDbNumber, err = strconv.Atoi(redisDbNumberStr)
		if err != nil {
			redisDbNumber = 0
		}
	}
	return RedisConfig{
		Endpoint: redisEndpoint,
		Password: redisPassword,
		Database: redisDbNumber,
	}
}

func initMinio() MinioConfig {
	minioHost, has := os.LookupEnv("DEEPFENCE_MINIO_HOST")
	if !has {
		minioHost = "deepfence-file-server"
		log.Warn().Msgf("DEEPFENCE_MINIO_HOST defaults to: %v", minioHost)
	}
	minioPort, has := os.LookupEnv("DEEPFENCE_MINIO_PORT")
	if !has {
		minioPort = "9000"
		log.Warn().Msgf("DEEPFENCE_MINIO_PORT defaults to: %v", minioPort)
	}
	minioEndpoint := minioHost + ":" + minioPort
	minioUser := os.Getenv("DEEPFENCE_MINIO_USER")
	minioPassword := os.Getenv("DEEPFENCE_MINIO_PASSWORD")
	minioBucket := os.Getenv("DEEPFENCE_MINIO_BUCKET")
	minioSecure := os.Getenv("DEEPFENCE_MINIO_SECURE")
	if minioSecure == "" {
		minioSecure = "false"
	}
	isSecure, err := strconv.ParseBool(minioSecure)
	if err != nil {
		isSecure = false
		log.Warn().Msgf("DEEPFENCE_MINIO_SECURE defaults to: %v (%v)", isSecure, err.Error())
	}
	return MinioConfig{
		Endpoint:   minioEndpoint,
		Username:   minioUser,
		Password:   minioPassword,
		BucketName: minioBucket,
		Secure:     isSecure,
	}
}

func initPosgresql() PostgresqlConfig {
	var err error
	postgresHost, has := os.LookupEnv("DEEPFENCE_POSTGRES_USER_DB_HOST")
	if !has {
		postgresHost = "localhost"
		log.Warn().Msgf("DEEPFENCE_POSTGRES_USER_DB_HOST defaults to: %v", postgresHost)
	}
	postgresPort := 5432
	postgresPortStr := os.Getenv("DEEPFENCE_POSTGRES_USER_DB_PORT")
	if postgresPortStr == "" {
		log.Warn().Msgf("DEEPFENCE_POSTGRES_USER_DB_PORT defaults to: %d", postgresPort)
	} else {
		postgresPort, err = strconv.Atoi(postgresPortStr)
		if err != nil {
			postgresPort = 5432
		}
	}
	postgresUsername := os.Getenv("DEEPFENCE_POSTGRES_USER_DB_USER")
	postgresPassword := os.Getenv("DEEPFENCE_POSTGRES_USER_DB_PASSWORD")
	postgresDatabase := os.Getenv("DEEPFENCE_POSTGRES_USER_DB_NAME")
	postgresSslMode := os.Getenv("DEEPFENCE_POSTGRES_USER_DB_SSLMODE")

	return PostgresqlConfig{
		Host:     postgresHost,
		Port:     postgresPort,
		Username: postgresUsername,
		Password: postgresPassword,
		Database: postgresDatabase,
		SslMode:  postgresSslMode,
	}
}

func initNeo4j() Neo4jConfig {
	neo4jHost, has := os.LookupEnv("DEEPFENCE_NEO4J_HOST")
	if !has {
		neo4jHost = "localhost"
		log.Warn().Msgf("DEEPFENCE_NEO4J_HOST defaults to: %v", neo4jHost)
	}
	neo4jBoltPort, has := os.LookupEnv("DEEPFENCE_NEO4J_BOLT_PORT")
	if !has {
		neo4jBoltPort = "7687"
		log.Warn().Msgf("DEEPFENCE_NEO4J_BOLT_PORT defaults to: %v", neo4jBoltPort)
	}
	neo4jEndpoint := "bolt://" + neo4jHost + ":" + neo4jBoltPort
	neo4jUsername := os.Getenv("DEEPFENCE_NEO4J_USER")
	neo4jPassword := os.Getenv("DEEPFENCE_NEO4J_PASSWORD")
	return Neo4jConfig{
		Endpoint: neo4jEndpoint,
		Username: neo4jUsername,
		Password: neo4jPassword,
	}
}
