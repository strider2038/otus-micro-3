-- +goose Up
-- +goose StatementBegin
ALTER TABLE "user" ADD COLUMN age smallint;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "user" DROP COLUMN age;
-- +goose StatementEnd
