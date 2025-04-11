package redis

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"strconv"
	"tanakrit.assignment4.task/internal/core/domain"
	"time"
)

type TaskRedisCache struct {
	client *redis.Client
}

func NewTaskRedisCache(client *redis.Client) *TaskRedisCache{
	return &TaskRedisCache{client}
}
func (cache *TaskRedisCache) Save(task domain.Task) (bool, error) {
	jsonData, err := json.Marshal(task)
	if err != nil {
		return false, err
	}
	_, err = cache.client.Set(strconv.Itoa(int(task.ID)), string(jsonData), time.Minute*10).Result()
	if err != nil {
		return false, err
	}
	return true, nil
}
func (cache *TaskRedisCache) Update(task domain.Task) (bool, error) {
	cachedValue, err := cache.client.Get(strconv.Itoa(int(task.ID))).Result()
	if err != nil {
		return false, err
	}
	var oldTask domain.Task

	if err = json.Unmarshal([]byte(cachedValue), &oldTask); err == nil {
		title := task.GetTitle()
		if title == "" || title == oldTask.Title {
			title = oldTask.Title
		}
		description := task.GetDescription()
		if description == "" || description == oldTask.Description {
			description = oldTask.Description
		}
		completed := task.GetCompleted()
		updatedTask := domain.Task{ID: task.GetId(), Title: title, Description: description, Completed: completed}
		newCached, uErr := json.Marshal(updatedTask)
		if uErr != nil {
			return false, uErr
		}
		_, uErr = cache.client.Set(strconv.Itoa(int(task.ID)), string(newCached), time.Minute*10).Result()
		if uErr != nil {
			return false, uErr
		}
		return true, nil
	}
	return false, err
}
func (cache *TaskRedisCache) Delete(id uint) (bool, error) {
	_, err := cache.client.Del(strconv.Itoa(int(id))).Result()
	if err != nil {
		return false, err
	}
	return true, nil
}