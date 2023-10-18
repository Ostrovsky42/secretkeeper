-- +goose Up
-- +goose StatementBegin
CREATE TYPE data_type AS enum (
    'login_password',
    'text_data',
    'binary_data',
    'credit_card'
);

CREATE TABLE users(
  user_id uuid PRIMARY KEY,
  username varchar,
  password_hash varchar,
);

CREATE TABLE private_data(
  data_id uuid PRIMARY KEY,
  user_id uuid,
  data_type data_type,
  data bytea,
  meta_info json,
  updated_at timestamptz,
  FOREIGN KEY (user_id) REFERENCES users(user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE private_data;
DROP TABLE users;
DROP TYPE data_type;
-- +goose StatementEnd
