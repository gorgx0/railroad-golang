-- +goose Up
create table stations(
                         id int primary key,
                         lng float,
                         lat float,
                         name varchar(50)
);

-- +goose Down
drop table stations;
