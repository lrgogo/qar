DROP TABLE IF EXISTS `temp`;
CREATE TABLE `temp`(
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `t` VARCHAR(100) NOT NULL COMMENT '标题',
  PRIMARY KEY (`id`)
)AUTO_INCREMENT=1;

SELECT * FROM temp;



DELIMITER $$
DROP PROCEDURE IF EXISTS proc_batch_insert;
CREATE PROCEDURE proc_batch_insert()
  BEGIN
    DECLARE pre_user_id INT(11);
    DECLARE pre_title VARCHAR(100);
    DECLARE i INT;
    SET pre_user_id=101;
    SET pre_title='如何评价新发布的 Android Oreo?';
    SET i=1;
    WHILE i < 30000 DO
      SELECT t INTO pre_title FROM temp WHERE id=(i%40);
      INSERT INTO question(`user_id`,`title`) VALUES (pre_user_id,pre_title);
      SET pre_user_id=pre_user_id+1;
      SET i=i+1;
    END WHILE;
  END $$
DELIMITER ;
CALL proc_batch_insert();