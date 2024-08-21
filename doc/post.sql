create table post
(
    id         bigint auto_increment comment '主键id'
        primary key,
    user_id    bigint                  not null comment '作者id',
    title      varchar(255) default '' not null comment '标题',
    content    text                    not null comment '推文内容',
    status     tinyint      default 0  not null comment '状态 0 正常 1 已删除',
    views      int          default 0  not null comment '浏览数',
    likes      int          default 0  not null comment '点赞数',
    comments   int          default 0  not null comment '评论数',
    shares     int          default 0  not null comment '分享数',
    collects   int          default 0  not null comment '收藏数',
    created_at int          default 0  not null comment '发表时间',
    updated_at int          default 0  not null comment '更新时间',
    deleted_at int          default 0  not null comment '删除时间'
)
    comment '推文表' collate = utf8mb4_bin;

