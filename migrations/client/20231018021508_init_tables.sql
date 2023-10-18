-- +goose Up
-- +goose StatementBegin
CREATE TYPE data_type AS enum (
    'login_password',
    'text_data',
    'binary_data',
    'credit_card'
);

CREATE TABLE private_data(
  data_id uuid PRIMARY KEY,
  name varchar,
  data_type data_type,
  data bytea,
  meta_info varchar,
  updated_at timestamptz,
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE private_data;
-- +goose StatementEnd
