DROP TABLE IF EXISTS `a`;
CREATE TABLE `a`(
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `age` TINYINT NOT NULL COMMENT '年龄',
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `b`;
CREATE TABLE `b`(
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` VARCHAR(10) NOT NULL COMMENT '姓名',
  PRIMARY KEY (`id`)
);


DROP TABLE IF EXISTS `websites`;
CREATE TABLE `websites`(
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` VARCHAR(20) NOT NULL COMMENT '网站名称',
  `url` VARCHAR(50) NOT NULL COMMENT '网站url',
  `alexa` SMALLINT DEFAULT 0 COMMENT '排名',
  `country` VARCHAR(10) DEFAULT '' COMMENT '国家',
  PRIMARY KEY (`id`)
);


DROP TABLE IF EXISTS `apps`;
CREATE TABLE `apps`(
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `app_name` VARCHAR(20) NOT NULL COMMENT 'app名称',
  `url` VARCHAR(50) NOT NULL COMMENT 'app网址',
  `country` VARCHAR(10) DEFAULT '' COMMENT '国家',
  PRIMARY KEY (`id`)
);


SELECT * FROM a INNER JOIN b on a.id=b.id;

SELECT * FROM a,b WHERE a.id=b.id;

SELECT a.*,b.* FROM a LEFT JOIN b ON a.id=b.id WHERE b.id is NULL ORDER BY a.id;

SELECT a.id AS aid, a.age, b.id AS bid, b.name FROM a RIGHT JOIN b on a.id=b.id WHERE a.id IS NULL;


SELECT a.id AS aid, a.age, b.id AS bid, b.name FROM a LEFT JOIN b ON a.id=b.id WHERE b.id is null
UNION
SELECT a.id AS aid, a.age, b.id AS bid, b.name FROM a RIGHT JOIN b ON a.id=b.id WHERE a.id is NULL;


SELECT a.id AS aid, a.age, b.id AS bid, b.name FROM a CROSS JOIN b;


SELECT websites.country FROM websites WHERE websites.country='CN'
UNION ALL
SELECT apps.country FROM apps WHERE apps.country='CN'
ORDER BY country;