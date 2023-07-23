# Host: 127.0.01  (Version: 5.7.26)
# Date: 2023-07-23 15:09:25
# Generator: MySQL-Front 5.3  (Build 4.234)

/*!40101 SET NAMES utf8 */;

#
# Structure for table "head_table"
#

DROP TABLE IF EXISTS `head_table`;
CREATE TABLE `head_table` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `userid` int(11) NOT NULL DEFAULT '0',
  `username` varchar(255) DEFAULT NULL,
  `imgpath` varchar(255) NOT NULL DEFAULT '/static/img/default.png',
  `in_use` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
