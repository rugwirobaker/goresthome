
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE  Users ADD datejoined DATE;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE Users REMOVE datejoined;