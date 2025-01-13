package main

import (
	"time"

	"github.com/qwerqy/api-starter/internal/env"
	"github.com/qwerqy/api-starter/internal/ratelimiter"
	"go.uber.org/zap"
)

const version = "0.0.1"

// TODO: Change to your issuer name
const issuer = "api-starter"

//	@title			Starter API
//	@description	Starter API
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/v1

// @securityDefinitions.apiKey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {
	cfg := config{
		addr:   env.GetString("ADDR", ":8080"),
		apiURL: env.GetString("EXTERNAL_URL", "localhost:8080"),
		// frontendURL: env.GetString("FRONTEND_URL", "http://localhost:3000"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		redisCfg: redisConfig{
			addr:    env.GetString("REDIS_ADDR", "localhost:6379"),
			pw:      env.GetString("REDIS_PW", ""),
			db:      env.GetInt("REDIS_DB", 0),
			enabled: env.GetBool("REDIS_ENABLED", false),
		},
		env: env.GetString("ENV", "development"),
		// mail: mailConfig{
		// 	exp:       time.Hour * 24 * 3, // 3 days
		// 	fromEmail: env.GetString("FROM_EMAIL", ""),
		// 	sendGrid: sendGridConfig{
		// 		apiKey: env.GetString("SENDGRID_API_KEY", ""),
		// 	},
		// },
		auth: authConfig{
			basic: basicConfig{
				user: env.GetString("AUTH_BASIC_USER", "admin"),
				pass: env.GetString("AUTH_BASIC_PASS", "admin"),
			},
			token: tokenConfig{
				secret: env.GetString("AUTH_TOKEN_SECRET", "example"),
				exp:    time.Hour * 24 * 3,
				iss:    issuer,
			},
		},
		rateLimiter: ratelimiter.Config{
			RequestsPerTimeFrame: env.GetInt("RATELIMITER_REQUESTS_COUNT", 20),
			TimeFrame:            time.Second * 5,
			Enabled:              env.GetBool("RATE_LIMITER_ENABLED", false),
		},
	}

	logger := zap.Must(zap.NewProduction()).Sugar()

	defer logger.Sync()

	// DATABASE
	// db, err := db.New(
	// 	cfg.db.addr,
	// 	cfg.db.maxOpenConns,
	// 	cfg.db.maxIdleConns,
	// 	cfg.db.maxIdleTime,
	// )

	// if err != nil {
	// 	logger.Fatal(err)
	// }

	// defer db.Close()
	// logger.Info("Database connection pool established")

	// CACHE
	// var rdb *redis.Client
	// if cfg.redisCfg.enabled {
	// 	rdb = cache.NewRedisClient(cfg.redisCfg.addr, cfg.redisCfg.pw, cfg.redisCfg.db)
	// 	logger.Info("Redis connection pool established")
	// }
	// cacheStorage := cache.NewRedisStorage(rdb)

	// mailer := mailer.NewSendGrid(cfg.mail.sendGrid.apiKey, cfg.mail.fromEmail)

	// jwtAuthenticator := auth.NewJWTAuthenticator(cfg.auth.token.secret, cfg.auth.token.iss, cfg.auth.token.iss)

	// RATE LIMITING
	rateLimiter := ratelimiter.NewFixedWindowRateLimiter(
		cfg.rateLimiter.RequestsPerTimeFrame,
		cfg.rateLimiter.TimeFrame,
	)

	app := &application{
		config:      cfg,
		logger:      logger,
		rateLimiter: rateLimiter,
	}

	mux := app.mount()

	logger.Fatal(app.run(mux))
}
