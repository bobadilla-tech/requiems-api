package config

import "os"

type Config struct {
	Port               string
	DatabaseURL        string
	RedisURL           string
	VPNDatabasePath    string
	VPNASNDatabasePath string
	IPCityDatabasePath string
	NominatimURL       string
	PostalCodeDBPath   string
	CitiesDBPath       string
	EnabledServices    string // comma-separated list of services to mount; empty = all
	Environment        string // "development" | "staging" | "production"
	SentryDSN          string
	LanguageToolURL    string // base URL for LanguageTool spell-check service
}

func Load() Config {
	return Config{
		Port:               envOrDefault("PORT", "8080"),
		DatabaseURL:        envOrDefault("DATABASE_URL", "postgres://requiem:requiem@localhost:5432/requiem?sslmode=disable"),
		RedisURL:           envOrDefault("REDIS_URL", "redis://localhost:6379/0"),
		VPNDatabasePath:    envOrDefault("VPN_DATABASE_PATH", "dbs/IP2PROXY-LITE-PX2.BIN"),
		VPNASNDatabasePath: envOrDefault("VPN_ASN_DATABASE_PATH", "dbs/GeoLite2-ASN.mmdb"),
		IPCityDatabasePath: envOrDefault("IP_CITY_DATABASE_PATH", "dbs/GeoLite2-City.mmdb"),
		NominatimURL:       envOrDefault("NOMINATIM_URL", "https://nominatim.openstreetmap.org"),
		PostalCodeDBPath:   envOrDefault("POSTAL_CODE_DB_PATH", "dbs/postal_codes.txt"),
		CitiesDBPath:       envOrDefault("CITIES_DB_PATH", "dbs/cities15000.txt"),
		EnabledServices:    envOrDefault("ENABLED_SERVICES", ""),
		Environment:        envOrDefault("ENVIRONMENT", "development"),
		SentryDSN:          envOrDefault("SENTRY_DSN", "https://df9023edf98442d4887d1040de81b768@issues.bobadilla.tech/1"),
		LanguageToolURL:    envOrDefault("LANGUAGETOOL_URL", "http://localhost:8010"),
	}
}

func envOrDefault(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return def
}
