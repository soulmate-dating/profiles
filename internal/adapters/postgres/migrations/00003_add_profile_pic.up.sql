ALTER TABLE profiles.profiles
    ADD COLUMN fk_main_pic_prompt uuid REFERENCES profiles.prompts(id);

ALTER TABLE profiles.prompts
    ADD CONSTRAINT prompts_user_id_fkey FOREIGN KEY (user_id) REFERENCES profiles.profiles(user_id);
