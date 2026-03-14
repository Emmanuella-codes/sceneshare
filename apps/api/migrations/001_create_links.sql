CREATE TABLE IF NOT EXISTS links (
    id          UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    short_code  VARCHAR(20)  UNIQUE NOT NULL,
    platform    VARCHAR(50)  NOT NULL,
    content_id  VARCHAR(200) NOT NULL,
    timestamp_s INTEGER      NOT NULL DEFAULT 0,
    title       TEXT,
    thumbnail   TEXT,
    owner_token VARCHAR(50)  NOT NULL,
    created_by  TEXT         NOT NULL DEFAULT '',
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    expires_at  TIMESTAMPTZ,
    click_count INTEGER      NOT NULL DEFAULT 0
);

CREATE INDEX IF NOT EXISTS links_short_code_idx ON links (short_code);
