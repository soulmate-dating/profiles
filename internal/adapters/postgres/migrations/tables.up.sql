CREATE SCHEMA profiles;

CREATE TYPE HABIT AS ENUM ('no', 'sometimes', 'yes', 'prefer not to say');
CREATE TYPE INTENTION AS ENUM (
    'life partner',
    'long-term relationship',
    'short-term relationship',
    'friendship',
    'figuring it out',
    'prefer not to say'
    );
CREATE TYPE GENDER AS ENUM ('man', 'woman');
CREATE TYPE PREFERRED_PARTNER AS ENUM ('man', 'woman', 'anyone');
CREATE TYPE FAMILY_PLANS AS ENUM ('don''t want children', 'want children', 'open to children', 'not sure yet', 'prefer not to say');

CREATE TABLE profiles.profiles
(
    user_id           uuid,
    first_name        TEXT,
    last_name         TEXT,
    birth_date        DATE,
    sex               GENDER,
    preferred_partner PREFERRED_PARTNER,
    intention         INTENTION,
    height            INTEGER,
    has_children      BOOLEAN,
    family_plans      TEXT,
    location          TEXT,
    drinks_alcohol    HABIT,
    smokes HABIT,
    PRIMARY KEY (user_id)
);

CREATE TABLE profiles.prompts
(
    id uuid,
    user_id uuid,
    position INTEGER,
    question TEXT,
    answer TEXT,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES profiles (user_id)
);