CREATE TABLE IF NOT EXISTS questions (
    question_id bigserial PRIMARY KEY,
    body text NOT NULL,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);