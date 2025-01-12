package browse

import (
	"database/sql"
	"fmt"
	"sort"
)

func GetMatchingUsers(tx *sql.Tx, userID int, filter BrowseRequest) ([]MatchingUser, error) {

	rows, err := tx.Query(query,
		userID,
		filter.DistanceRange.Min*1000,
		filter.DistanceRange.Max*1000,
		filter.AgeRange.Min,
		filter.AgeRange.Max,
		filter.MinCommonTags,
		filter.MinFameRating,
	)
	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()

	var matchingUsers []MatchingUser
	for rows.Next() {
		var user MatchingUser
		err := rows.Scan(
			&user.Username,
			&user.Birthdate,
			&user.Area,
			&user.ProfileImagePath1,
			&user.CommonTagCount,
			&user.DistanceKm,
			&user.FameRating,
		)
		if err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}

		matchingUsers = append(matchingUsers, user)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}
	usersSort(matchingUsers, filter.SortOption, filter.SortOrder)
	return matchingUsers, nil
}

type SortFunc func(i, j int) bool

func usersSort(users []MatchingUser, sortOption SortOptionType, sortOrder SortOrder) {
	var less SortFunc
	switch sortOption {
	case Distance:
		less = func(i, j int) bool {
			return users[i].DistanceKm < users[j].DistanceKm
		}
	case Age:
		less = func(i, j int) bool {
			//return users[i].Birthdate.Before(users[j].Birthdate)
			return users[i].Birthdate.After(users[j].Birthdate)
		}
	case FameRating:
		less = func(i, j int) bool {
			return users[i].FameRating < users[j].FameRating
		}
	case Tag:
		less = func(i, j int) bool {
			return users[i].CommonTagCount < users[j].CommonTagCount
		}
	}
	if sortOrder == Descending {
		originalLess := less
		less = func(i, j int) bool { return !originalLess(j, i) }
	}
	sort.Slice(users, less)
}
