-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE blacklist (
    address CIDR PRIMARY KEY
);

CREATE TABLE whitelist (
    address CIDR PRIMARY KEY
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE blacklist;
DROP TABLE whitelist;
