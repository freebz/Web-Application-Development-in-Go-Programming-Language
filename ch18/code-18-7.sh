# 코드 18.7 make dry-migrate 명령 실행

$ make dry-migrate
mysqldef -u todo -p todo -h 127.0.0.1 -P 33306 todo --dry-run < ./_tools/mysql/schema.sql
-- dry run --
ALTER TABLE `user` CHANGE COLUMN `name` `name` varchar(20) NOT NULL COMMENT '사용자의 이름';