
SET FOREIGN_KEY_CHECKS=1;

DROP TABLE IF EXISTS `temp_answer`;
CREATE TABLE `temp_answer`(
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `c` VARCHAR(1000) NOT NULL COMMENT '内容',
  PRIMARY KEY (`id`)
)AUTO_INCREMENT=1;

SELECT * FROM temp_answer;

DELETE FROM answer;

DELIMITER $$
DROP PROCEDURE IF EXISTS proc_batch_insert;
CREATE PROCEDURE proc_batch_insert()
  BEGIN
    DECLARE pre_user_id INT(11);
    DECLARE pre_question_id INT(11);
    DECLARE pre_content VARCHAR(1000);
    DECLARE i INT;
    SET pre_user_id=105;
    SET pre_question_id=108;
    SET pre_content='很好的问题，准备好再答';
    SET i=1;
    WHILE i < 10000 DO
      SELECT c INTO pre_content FROM temp_answer WHERE id=(i%30);
      INSERT INTO answer(`user_id`,`question_id`,`content`) VALUES (pre_user_id,pre_question_id,pre_content);
      SET pre_user_id=pre_user_id+1;
      SET pre_question_id=pre_question_id+1;
      SET i=i+1;
    END WHILE;
  END $$
DELIMITER ;
CALL proc_batch_insert();