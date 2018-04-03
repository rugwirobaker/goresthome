
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE Articles(
    id        serial,
    title     VARCHAR(250),
    body      TEXT,
    createdOn DATETIME
    PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE Articles;
