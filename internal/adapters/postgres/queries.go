package postgres

const (
	getProfileByIDQuery = `SELECT * FROM profiles.profiles WHERE user_id = $1`
	createProfileQuery  = `INSERT INTO profiles.profiles (
                      		user_id, first_name, last_name, birth_date, sex, preferred_partner, intention, 
    						height, has_children, family_plans, location,
    						drinks_alcohol, smokes
    						) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	updateProfileQuery = `UPDATE profiles.profiles 
							SET first_name = $2, last_name = $3, birth_date = $4,
							    sex = $5, preferred_partner = $6, intention = $7, height = $8,
							    has_children = $9, family_plans = $10, location = $11,
							    drinks_alcohol = $12, smokes = $13, fk_main_pic_prompt = $14 WHERE user_id = $1 RETURNING *`
	getRandomProfileBySexAndPreferenceQuery = `SELECT * FROM profiles.profiles WHERE (user_id != $1 AND (sex = $2 OR sex = $3) AND (preferred_partner = $4 OR preferred_partner = 'anyone')) ORDER BY RANDOM() LIMIT 1`
	getMultipleProfilesByIDsQuery           = `SELECT * FROM profiles.profiles WHERE user_id = ANY($1)`
	createPromptQuery                       = `INSERT INTO profiles.prompts (id, user_id, question, content, type, position) VALUES ($1, $2, $3, $4, $5, $6)`
	getPromptsByUserQuery                   = `SELECT * FROM profiles.prompts WHERE user_id = $1 ORDER BY position ASC`
	getPromptByIDQuery                      = `SELECT * FROM profiles.prompts WHERE id = $1`
	getPromptByUserQuestionAndTypeQuery     = `SELECT * FROM profiles.prompts WHERE user_id = $1 AND question = $2 AND type = $3`
	updatePromptQuery                       = `UPDATE profiles.prompts SET question = $2, content = $3, position = $4 WHERE id = $1 RETURNING *`
	getPromptsByIDsQuery                    = `SELECT * FROM profiles.prompts WHERE id = ANY($1)`
	updatePromptsPositionQuery              = `
		UPDATE profiles.prompts
		SET position = updated.new_position
		FROM (SELECT unnest($1::uuid[]) as new_id, unnest($2::int[]) as new_position) as updated
		WHERE id = updated.new_id`
)
