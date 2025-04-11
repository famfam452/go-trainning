package stacks

import (
	"fmt"
	redisLib "github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"tanakrit.assignment4.task/cmd/api/config"
	services "tanakrit.assignment4.task/internal/core/service"
	"tanakrit.assignment4.task/internal/infrastructure/api/handlers"
	cacheRedis "tanakrit.assignment4.task/internal/infrastructure/cache/redis"
	postgresRepo "tanakrit.assignment4.task/internal/infrastructure/database/postgres"
)

func InitStack(properties config.Properties) *handlers.TaskHandler{
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		properties.Postgresql.Host,
		properties.Postgresql.Port,
		properties.Postgresql.Username,
		properties.Postgresql.Password,
		properties.Postgresql.Database,
		properties.Postgresql.Sslmode,
	)

	gormSession, _ := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	redis := redisLib.NewClient(&redisLib.Options{
		Addr: properties.Redis.Address,
		Password: properties.Redis.Password,
		DB: properties.Redis.Database,
		})
	_ = gormSession.AutoMigrate(&postgresRepo.Task{})
	repository := postgresRepo.NewTaskPostgresRepository(gormSession)
	cache := cacheRedis.NewTaskRedisCache(redis)
	service := services.NewTaskService(repository, cache)
	handler := handlers.NewTaskHandler(service)
	return handler
}