-- +migrate Up
create table `user` (
    `id` char(26) not null unique comment 'ユーザーの一意な識別子ULID',
    primary key (`id`)
) engine=InnoDB default charset=utf8mb4 comment 'リソース: 登録したユーザー全てを記録';

-- +migrate Down
DROP TABLE IF EXISTS `user`;
