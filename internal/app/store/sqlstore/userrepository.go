package sqlstore

import (
	"fmt"
	"testapi/internal/app/response"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) GetUsersList(username string) ([]response.Profile, error) {
	query := `SELECT 
				user.id,
				user.username,
				user_profile.first_name,
				user_profile.last_name,
				user_profile.city,
				user_data.school
			FROM user
				INNER JOIN user_profile ON user.id = user_profile.user_id
				INNER JOIN user_data ON user.id = user_data.user_id
	`

	if len(username) > 0 {
		query = fmt.Sprintf("%s WHERE user.username = \"%s\"", query, username)
	}

	rows, err := r.store.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	profiles := []response.Profile{}

	for rows.Next() {
		profile := response.Profile{}

		if err := rows.Scan(
			&profile.ID,
			&profile.Username,
			&profile.First_name,
			&profile.Last_name,
			&profile.City,
			&profile.School,
		); err != nil {
			return profiles, err
		}

		profiles = append(profiles, profile)
	}

	return profiles, nil
}
