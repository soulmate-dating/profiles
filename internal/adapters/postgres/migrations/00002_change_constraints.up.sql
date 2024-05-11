ALTER TYPE FAMILY_PLANS RENAME VALUE 'don''t want children' TO 'do not want children';

ALTER TABLE profiles.prompts
    ADD CONSTRAINT unique_prompt_pair UNIQUE (user_id, question, type);

ALTER TABLE profiles.prompts
    DROP CONSTRAINT prompts_user_id_fkey;