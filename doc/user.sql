create table user
(
    id         BIGINT auto_increment comment '主键id',
    avatar     varchar(255) default '' not null comment '头像链接地址',
    nickname   char(32)     default '' not null comment '昵称',
    account    char(32)     default '' not null comment '账号',
    password   char(32)     default '' not null comment '密码',
    bio        varchar(255) default '' not null comment '个人简介 Biography',
    gender     TINYINT      default 2  not null comment '性别 0 女 1 男 2 未知',
    region     varchar(20)  default '' not null comment '地区',
    created_at int          default 0  not null comment '创建时间',
    updated_at int          default 0  not null comment '更新时间',
    deleted_at int          default 0  not null comment '删除时间',
    constraint user_pk
        primary key (id)
)
    comment '用户表' collate = utf8mb4_bin;