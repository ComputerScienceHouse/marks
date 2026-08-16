package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ComputerScienceHouse/marks/internal/logging"
	"github.com/ComputerScienceHouse/marks/internal/models"
	"github.com/redis/go-redis/v9"
)

func getLeaderboardEntryKey(eventID string, userUUID string) string {
	return fmt.Sprintf("leaderboard:entry:%s:%s", eventID, userUUID)
}

func getLeaderboardRankingKey(eventID string) string {
	return fmt.Sprintf("leaderboard:rankinglist:%s", eventID)
}

func marshalLeaderboardEntry(entry *models.LeaderboardEntry) ([]byte, error) {
	data, err := json.Marshal(*entry)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func unmarshalLeaderboardEntry(data []byte) (*models.LeaderboardEntry, error) {
	var entry models.LeaderboardEntry
	err := json.Unmarshal(data, &entry)
	if err != nil {
		return &models.LeaderboardEntry{}, err
	}
	return &entry, nil
}

func AddUserToLeaderboard(ctx context.Context, entry *models.LeaderboardEntry) error {
	entryKey := getLeaderboardEntryKey(entry.EventID, entry.UserUUID)

	entryMarshaled, err := marshalLeaderboardEntry(entry)
	if err != nil {
		return err
	}

	if err := client.Set(ctx, entryKey, entryMarshaled, 0).Err(); err != nil {
		return err
	}

	rankingKey := getLeaderboardRankingKey(entry.EventID)

	if err := client.ZAdd(ctx, rankingKey, redis.Z{
		Score:  float64(entry.Duration),
		Member: entry.UserUUID,
	}).Err(); err != nil {
		_ = client.Del(ctx, entryKey)
		return err
	}

	return nil
}

func RemoveUserFromLeaderboard(ctx context.Context, userUUID string, eventID string) error {
	entryKey := getLeaderboardEntryKey(eventID, userUUID)

	if err := client.Del(ctx, entryKey).Err(); err != nil {
		return err
	}

	rankingKey := getLeaderboardRankingKey(eventID)

	if err := client.ZRem(ctx, rankingKey, redis.Z{
		Member: userUUID,
	}).Err(); err != nil {
		return err
	}

	return nil
}

func GetTopOfLeaderboard(ctx context.Context, eventID string) ([]*models.LeaderboardEntry, error) {
	rankingKey := getLeaderboardRankingKey(eventID)
	result, err := client.ZRange(ctx, rankingKey, 0, 9).Result()
	if err != nil {
		return nil, err
	}

	entries := make([]*models.LeaderboardEntry, len(result))

	userKeys := make([]string, len(result))
	for i, userUUID := range result {
		userKeys[i] = getLeaderboardEntryKey(eventID, userUUID)
	}

	userEntries, err := client.MGet(ctx, userKeys...).Result()
	if err != nil {
		return nil, err
	}

	for i, dataAny := range userEntries {
		if dataAny == nil {
			logging.Logger.Warnf("Unable to collect data for userKey %s", userKeys[i])
			continue
		}

		data, ok := dataAny.(string)
		if !ok {
			return nil, errors.New("unable to convert data entry")
		}

		entry, err := unmarshalLeaderboardEntry([]byte(data))
		if err != nil {
			return nil, err
		}

		entries[i] = entry
	}

	return entries, nil
}
