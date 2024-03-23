package postgres

const (
	getProfileByIDQuery = `SELECT * FROM profiles WHERE user_id = $1`
	createProfileQuery  = `INSERT INTO profiles (
                      		user_id, first_name, last_name, birth_date, sex, preferred_partner, intention, 
    						height, has_children, family_plans, location, education_level, 
    						drinks_alcohol, smokes_cigarettes
    						) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`
)
