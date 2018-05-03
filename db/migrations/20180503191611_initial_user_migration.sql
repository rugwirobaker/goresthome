
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE Users(
    id serial,
    fname    TEXT,
    lname    TEXT,
    email    TEXT,
    passhash VARCHAR(60)
    --salt     VARCHAR(32)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE Users();
