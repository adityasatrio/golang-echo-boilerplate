-- +goose Up
-- +goose StatementBegin
INSERT INTO example (username,email)
VALUES ('joko','joko@email.com');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM example WHERE username like '%joko%';
-- +goose StatementEnd
