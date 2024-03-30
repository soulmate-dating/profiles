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
							    drinks_alcohol = $12, smokes = $13 WHERE user_id = $1 RETURNING *`
	createPromptQuery = `INSERT INTO profiles.prompts (uid, user_id, question, answer, position) VALUES ($1, $2, $3, $4, $5)`
	getPromptsQuery   = `SELECT * FROM profiles.prompts WHERE user_id = $1 ORDER BY position ASC`
	updatePromptQuery = `UPDATE profiles.prompts SET question = $2, answer = $3, position = $4 WHERE uid = $1 RETURNING *`
)
