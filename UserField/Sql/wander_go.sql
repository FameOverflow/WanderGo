-- MariaDB dump 10.19-11.1.2-MariaDB, for Win64 (AMD64)
--
-- Host: localhost    Database: spark_forge
-- ------------------------------------------------------
-- Server version	11.1.2-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `avatar`
--

DROP TABLE IF EXISTS `avatar`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `avatar` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_account` varchar(171) NOT NULL,
  `avatar_data` longblob NOT NULL,
  `avatar_uid` varchar(171) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_avatar_deleted_at` (`deleted_at`),
  KEY `idx_avatar_user_account` (`user_account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `avatar`
--

LOCK TABLES `avatar` WRITE;
/*!40000 ALTER TABLE `avatar` DISABLE KEYS */;
/*!40000 ALTER TABLE `avatar` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment`
--

DROP TABLE IF EXISTS `comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comment` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_account` varchar(171) NOT NULL,
  `date` varchar(171) NOT NULL,
  `photo_uid` varchar(171) DEFAULT NULL,
  `text` varchar(171) DEFAULT NULL,
  `comment_uid` bigint(20) unsigned NOT NULL,
  `star_cnt` bigint(20) DEFAULT 0,
  `photo_data` longblob DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_comment_deleted_at` (`deleted_at`),
  KEY `idx_comment_user_account` (`user_account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
/*!40000 ALTER TABLE `comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `photo`
--

DROP TABLE IF EXISTS `photo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `photo` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_account` varchar(171) NOT NULL,
  `photo_data` longblob NOT NULL,
  `photo_uid` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_photo_deleted_at` (`deleted_at`),
  KEY `idx_photo_user_account` (`user_account`),
  KEY `idx_photo_photo_uid` (`photo_uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `photo`
--

LOCK TABLES `photo` WRITE;
/*!40000 ALTER TABLE `photo` DISABLE KEYS */;
/*!40000 ALTER TABLE `photo` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `place`
--

DROP TABLE IF EXISTS `place`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `place` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `place_name` varchar(171) NOT NULL,
  `place_uid` bigint(20) NOT NULL,
  `top_left_point` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`top_left_point`)),
  `bottom_right_point` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`bottom_right_point`)),
  `center_point` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`center_point`)),
  PRIMARY KEY (`id`),
  KEY `idx_place_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `place`
--

LOCK TABLES `place` WRITE;
/*!40000 ALTER TABLE `place` DISABLE KEYS */;
INSERT INTO `place` VALUES
(1,'2023-12-06 19:27:41.166','2023-12-06 19:27:41.166',NULL,'	机电楼',0,'{\"x\":115.798058,\"y\":28.661722}','{\"x\":115.801668,\"y\":28.662683}','{\"x\":115.799863,\"y\":28.6622025}'),
(2,'2023-12-06 19:27:41.168','2023-12-06 19:27:41.168',NULL,'	建工楼',0,'{\"x\":115.798964,\"y\":28.662739}','{\"x\":115.802559,\"y\":28.66329}','{\"x\":115.8007615,\"y\":28.6630145}'),
(3,'2023-12-06 19:27:41.169','2023-12-06 19:27:41.169',NULL,'	一食堂',0,'{\"x\":115.803739,\"y\":28.66482}','{\"x\":115.804527,\"y\":28.664344}','{\"x\":115.804133,\"y\":28.664582}'),
(4,'2023-12-06 19:27:41.170','2023-12-06 19:27:41.170',NULL,'	润溪湖（1）',0,'{\"x\":115.805922,\"y\":28.66449}','{\"x\":115.804366,\"y\":28.665149}','{\"x\":115.805144,\"y\":28.6648195}'),
(5,'2023-12-06 19:27:41.171','2023-12-06 19:27:41.171',NULL,'	润溪湖（3）',0,'{\"x\":115.80722,\"y\":28.662584}','{\"x\":115.810235,\"y\":28.661492}','{\"x\":115.8087275,\"y\":28.662038}'),
(6,'2023-12-06 19:27:41.172','2023-12-06 19:27:41.172',NULL,'	艺术楼',0,'{\"x\":115.807285,\"y\":28.661303}','{\"x\":115.808915,\"y\":28.659976}','{\"x\":115.8081,\"y\":28.660639500000002}'),
(7,'2023-12-06 19:27:41.173','2023-12-06 19:27:41.173',NULL,'	外经楼',0,'{\"x\":115.805396,\"y\":28.661501}','{\"x\":115.80677,\"y\":28.66265}','{\"x\":115.806083,\"y\":28.6620755}'),
(8,'2023-12-06 19:27:41.173','2023-12-06 19:27:41.173',NULL,'	文法楼',0,'{\"x\":115.803744,\"y\":28.661209}','{\"x\":115.805332,\"y\":28.660023}','{\"x\":115.80453800000001,\"y\":28.660615999999997}'),
(9,'2023-12-06 19:27:41.174','2023-12-06 19:27:41.174',NULL,'	正气广场',0,'{\"x\":115.805461,\"y\":28.659948}','{\"x\":115.807285,\"y\":28.658969}','{\"x\":115.806373,\"y\":28.6594585}'),
(10,'2023-12-06 19:27:41.175','2023-12-06 19:27:41.175',NULL,'	白帆',0,'{\"x\":115.810444,\"y\":28.663709}','{\"x\":115.813459,\"y\":28.65967}','{\"x\":115.81195149999999,\"y\":28.6616895}'),
(11,'2023-12-06 19:27:41.176','2023-12-06 19:27:41.176',NULL,'	体育馆',0,'{\"x\":115.810755,\"y\":28.659557}','{\"x\":115.812515,\"y\":28.658682}','{\"x\":115.811635,\"y\":28.6591195}'),
(12,'2023-12-06 19:27:41.176','2023-12-06 19:27:41.176',NULL,'	游泳馆',0,'{\"x\":115.810187,\"y\":28.663709}','{\"x\":115.812043,\"y\":28.664829}','{\"x\":115.811115,\"y\":28.664269}'),
(13,'2023-12-06 19:27:41.177','2023-12-06 19:27:41.177',NULL,'	休闲运动场',0,'{\"x\":115.809919,\"y\":28.664886}','{\"x\":115.811689,\"y\":28.667747}','{\"x\":115.81080399999999,\"y\":28.6663165}'),
(14,'2023-12-06 19:27:41.178','2023-12-06 19:27:41.178',NULL,'	休闲广场',0,'{\"x\":115.808256,\"y\":28.665747}','{\"x\":115.806823,\"y\":28.665968}','{\"x\":115.80753949999999,\"y\":28.6658575}'),
(15,'2023-12-06 19:27:41.179','2023-12-06 19:27:41.179',NULL,'	休闲13栋',0,'{\"x\":115.806453,\"y\":28.665992}','{\"x\":115.807451,\"y\":28.666476}','{\"x\":115.806952,\"y\":28.666234}'),
(16,'2023-12-06 19:27:41.179','2023-12-06 19:27:41.179',NULL,'	树人广场',0,'{\"x\":115.802779,\"y\":28.657034}','{\"x\":115.803723,\"y\":28.656224}','{\"x\":115.803251,\"y\":28.656629000000002}'),
(17,'2023-12-06 19:27:41.180','2023-12-06 19:27:41.180',NULL,'	龙腾湖（1）',0,'{\"x\":115.800552,\"y\":28.656078}','{\"x\":115.802151,\"y\":28.654516}','{\"x\":115.8013515,\"y\":28.655297}'),
(18,'2023-12-06 19:27:41.181','2023-12-06 19:27:41.181',NULL,'	龙腾湖（2）',0,'{\"x\":115.802655,\"y\":28.656088}','{\"x\":115.803996,\"y\":28.655382}','{\"x\":115.8033255,\"y\":28.655735}'),
(19,'2023-12-06 19:27:41.182','2023-12-06 19:27:41.182',NULL,'	龙腾湖（3）',0,'{\"x\":115.805434,\"y\":28.657453}','{\"x\":115.806378,\"y\":28.657058}','{\"x\":115.805906,\"y\":28.657255499999998}'),
(20,'2023-12-06 19:27:41.182','2023-12-06 19:27:41.182',NULL,'	龙腾湖（4）',0,'{\"x\":115.806217,\"y\":28.65863}','{\"x\":115.809768,\"y\":28.657952}','{\"x\":115.80799250000001,\"y\":28.658291}'),
(21,'2023-12-06 19:27:41.183','2023-12-06 19:27:41.183',NULL,'	天健运动场',0,'{\"x\":115.793815,\"y\":28.653009}','{\"x\":115.796679,\"y\":28.654817}','{\"x\":115.79524699999999,\"y\":28.653913000000003}'),
(22,'2023-12-06 19:27:41.184','2023-12-06 19:27:41.184',NULL,'	医学实验大楼',0,'{\"x\":115.796829,\"y\":28.656088}','{\"x\":115.798213,\"y\":28.653216}','{\"x\":115.797521,\"y\":28.654652}'),
(23,'2023-12-06 19:27:41.185','2023-12-06 19:27:41.185',NULL,'	研究生',0,'{\"x\":115.793461,\"y\":28.65299}','{\"x\":115.79625,\"y\":28.651258}','{\"x\":115.7948555,\"y\":28.652124}');
/*!40000 ALTER TABLE `place` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `star`
--

DROP TABLE IF EXISTS `star`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `star` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_account` varchar(171) DEFAULT NULL,
  `comment_uid` varchar(171) DEFAULT NULL,
  `star_uid` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_star_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `star`
--

LOCK TABLES `star` WRITE;
/*!40000 ALTER TABLE `star` DISABLE KEYS */;
/*!40000 ALTER TABLE `star` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_name` varchar(171) NOT NULL,
  `user_account` varchar(171) NOT NULL,
  `user_password` varchar(171) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`),
  KEY `idx_user_user_account` (`user_account`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES
(1,'2023-12-06 14:28:39.154','2023-12-06 14:28:39.154',NULL,'haruhi','2391793239@qq.com','88c35f5e03f0af553ada4fc883f1a8cd');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-12-06 19:37:28
