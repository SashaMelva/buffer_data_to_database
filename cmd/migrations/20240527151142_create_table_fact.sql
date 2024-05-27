-- +goose Up
-- +goose StatementBegin
CREATE TABLE fact(
     id SERIAL PRIMARY KEY,
     period_start timestamp without time zone,
     period_end timestamp without time zone,
     comment character varying
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE fact;
-- +goose StatementEnd
