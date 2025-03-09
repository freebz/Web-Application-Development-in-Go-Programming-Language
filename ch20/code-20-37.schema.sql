-- 코드 20.37 새로운 task 테이블 정의

CREATE TABLE `task`
(
    `id`      BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '태스크 식별자',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '태스크를 작성한 사용자의 ID',
    `title`    VARCHAR(128) NOT NULL COMMENT '태스크 타이틀',
    `status`   VARCHAR(20)  NOT NULL COMMENT '태스크 상태',
    `create`   DATETIME(6)  NOT NULL COMMENT '레코드 작성 시간',
    `modified` DATETIME(6)  NOT NULL COMMENT '레코드 수정 시간',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_user_id`
        FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
            ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='태스크';