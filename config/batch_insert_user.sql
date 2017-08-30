DELIMITER $$
DROP PROCEDURE IF EXISTS proc_batch_insert;
CREATE PROCEDURE proc_batch_insert()
  BEGIN
    DECLARE pre_mobile BIGINT;
    DECLARE pre_pwd VARCHAR(20);
    DECLARE i INT;
    SET pre_mobile=13632000000;
    SET pre_pwd='123456';
    SET i=1;
    WHILE i < 100000 DO
      INSERT INTO user(`mobile`,`pwd`) VALUES (pre_mobile,pre_pwd);
      SET pre_mobile=pre_mobile+100;
      SET i=i+1;
    END WHILE;
  END $$
DELIMITER ;
CALL proc_batch_insert();