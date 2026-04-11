create table `like`
(
    id         bigint auto_increment comment '主键id'
        primary key,
    user_id    bigint                  not null comment '用户id',
    post_id    bigint                  not null comment '推文id',
    created_at int          default 0  not null comment '点赞时间',
    constraint like_pk
        unique (user_id, post_id)
)
    comment '点赞表' collate = utf8mb4_bin;

-- 索引
create index idx_like_post_id on `like`(post_id);
create index idx_like_user_id on `like`(user_id);
