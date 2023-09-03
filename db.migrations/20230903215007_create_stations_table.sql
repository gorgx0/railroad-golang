-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table stations(
                         id int primary key,
                         lng float,
                         lat float,
                         name varchar(50)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table stations;
-- +goose StatementEnd
