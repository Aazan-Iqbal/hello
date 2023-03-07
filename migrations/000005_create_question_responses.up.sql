CREATE TABLE IF NOT EXISTS question_responses (
  question_response_id bigserial PRIMARY KEY,
  question_id BIGINT REFERENCES questions(question_id),
  response_code_id BIGINT REFERENCES response_codes(response_code_id),
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);
