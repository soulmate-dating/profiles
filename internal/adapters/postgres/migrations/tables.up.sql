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
CREATE TYPE PREFERRED_PARTNER AS ENUM ('men', 'women', 'anyone');
CREATE TYPE FAMILY_PLANS AS ENUM ('don''t want children', 'want children', 'open to children', 'not sure yet', 'prefer not to say');

CREATE TABLE profiles
(
    user_id           TEXT,
    first_name        TEXT,
    last_name         TEXT,
    birth_date        DATE,
    sex               GENDER,
    preferred_partner PREFERRED_PARTNER,
    intention         INTENTION,
    height            INTEGER CHECK (height > 0 && height < 300),
    has_children      BOOLEAN,
    family_plans      TEXT,
    location          TEXT,
    drinks_alcohol    HABIT,
    smokes HABIT,
    PRIMARY KEY (user_id)
);