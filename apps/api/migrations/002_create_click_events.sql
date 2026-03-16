CREATE TABLE IF NOT EXISTS click_events (
    id         UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    link_id    UUID        NOT NULL REFERENCES links(id) ON DELETE CASCADE,
    user_agent TEXT,
    referrer   TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS click_events_link_id_idx ON click_events (link_id);
