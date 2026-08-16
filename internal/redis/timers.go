package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/ComputerScienceHouse/marks/internal/models"
	"github.com/redis/go-redis/v9"
)

const MAX_TIMER_DURATION = time.Hour * 6

func getUserTimerKey(eventID string, userUUID string) string {
	return fmt.Sprintf("timer:%s:%s", eventID, userUUID)
}

func marshalTimerEntry(entry *models.TimerEntry) ([]byte, error) {
	data, err := json.Marshal(entry)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func unmarshalTimerEntry(data []byte) (*models.TimerEntry, error) {
	var entry models.TimerEntry
	err := json.Unmarshal(data, &entry)
	if err != nil {
		return &models.TimerEntry{}, err
	}
	return &entry, nil
}

func GetUserData(ctx context.Context, eventID string, userUUID string) (*models.TimerEntry, error) {
	key := getUserTimerKey(eventID, userUUID)

	data, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	entry, err := unmarshalTimerEntry([]byte(data))
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func StartUserTimer(ctx context.Context, startTime models.TimerValue, eventID string, userUUID string) error {
	key := getUserTimerKey(eventID, userUUID)

	entry, err := GetUserData(ctx, eventID, userUUID)
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return err
		}

		entry = &models.TimerEntry{
			TimerStatus: models.TimerActive,

			StartTime: &startTime,
			EndTime:   nil,
			Duration:  nil,
		}
	}

	if entry.TimerStatus == models.TimerActive {
		return nil //maybe could be an error, convo needed
	}
	entry.TimerStatus = models.TimerActive
	entry.StartTime = &startTime
	entry.EndTime = nil
	entry.Duration = nil

	entryMarshaled, err := marshalTimerEntry(entry)
	if err != nil {
		return err
	}

	return client.Set(ctx, key, entryMarshaled, MAX_TIMER_DURATION).Err()
}

func StopUserTimer(ctx context.Context, endTime models.TimerValue, eventID string, userUUID string) error {
	key := getUserTimerKey(eventID, userUUID)

	entry, err := GetUserData(ctx, eventID, userUUID)
	if err != nil {
		return err
	}

	if entry.TimerStatus == models.TimerInactive {
		return nil //maybe could be an error, convo needed
	}

	if entry.StartTime == nil {
		return errors.New("start time was not set")
	}
	entry.TimerStatus = models.TimerActive
	entry.EndTime = &endTime

	duration := *entry.EndTime - *entry.StartTime
	entry.Duration = &duration

	entryMarshaled, err := marshalTimerEntry(entry)
	if err != nil {
		return err
	}

	return client.Set(ctx, key, entryMarshaled, MAX_TIMER_DURATION).Err()
}

func RemoveUserTimer(ctx context.Context, eventID string, userUUID string) error {
	key := getUserTimerKey(eventID, userUUID)
	return client.Del(ctx, key).Err()
}
