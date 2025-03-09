# 코드 20.38 make dry-migrate 명령 실행

$ make dry-migrate
mysqldef -u todo -p todo -h 127.0.0.1 -P 33306 todo --dry-run < ./_tools/mysql/schema.sql
-- dry run --
ALTER TABLE `task` ADD COLUMN `user_id` bigint UNSIGNED NOT NULL COMMENT '태스크를 작성한 사용자의 ID' AFTER `id`;
ALTER TABLE `task` ADD COLUMN `create` datetime(6) NOT NULL COMMENT '레코드 작성 시간' AFTER `status`;
ALTER TABLE `task` ADD CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `task` DROP COLUMN `created`;