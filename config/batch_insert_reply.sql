
SET FOREIGN_KEY_CHECKS=1;

DROP TABLE IF EXISTS `temp_reply`;
CREATE TABLE `temp_reply`(
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `c` VARCHAR(100) NOT NULL COMMENT '内容',
  PRIMARY KEY (`id`)
)AUTO_INCREMENT=1;

SELECT * FROM temp_reply;

DELETE FROM reply;

DELIMITER $$
DROP PROCEDURE IF EXISTS proc_batch_insert;
CREATE PROCEDURE proc_batch_insert()
  BEGIN
    DECLARE pre_user_id INT(11);
    DECLARE pre_answer_id INT(11);
    DECLARE pre_content VARCHAR(100);
    DECLARE i INT;
    SET pre_user_id=200;
    SET pre_answer_id=10103;
    SET pre_content='很好';
    SET i=0;
    WHILE i < 30000 DO
      SELECT c INTO pre_content FROM temp_reply WHERE id=(i%50+17);
      INSERT INTO reply(`user_id`,`answer_id`,`content`) VALUES (pre_user_id,pre_answer_id,pre_content);
      SET pre_user_id=pre_user_id+1;
      SET pre_answer_id=pre_answer_id+1;
      SET i=i+1;
    END WHILE;
  END $$
DELIMITER ;
CALL proc_batch_insert();