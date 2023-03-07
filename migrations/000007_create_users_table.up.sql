CREATE TABLE users (
  user_id serial PRIMARY KEY,
  user_name text NOT NULL,
  email citext UNIQUE NOT NULL,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);