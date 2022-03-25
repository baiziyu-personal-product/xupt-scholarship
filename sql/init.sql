CREATE DATABASE  IF NOT EXISTS `xupt-scholarship` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_bin */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `xupt-scholarship`;
-- MySQL dump 10.13  Distrib 8.0.27, for Win64 (x86_64)
--
-- Host: localhost    Database: xupt-scholarship
-- ------------------------------------------------------
-- Server version	8.0.27

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `applications`
--

DROP TABLE IF EXISTS `applications`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `applications` (
  `id` int NOT NULL AUTO_INCREMENT,
  `create_at` bigint NOT NULL,
  `update_at` bigint NOT NULL,
  `creator` varchar(20) COLLATE utf8_bin NOT NULL,
  `status` varchar(30) COLLATE utf8_bin NOT NULL DEFAULT '"save"' COMMENT 'save\\nsubmit\\nfinish\\ndroped\\nerror',
  `form` json NOT NULL,
  `step` varchar(320) COLLATE utf8_bin NOT NULL COMMENT '当前进行步骤',
  `history` json NOT NULL,
  PRIMARY KEY (`id`,`creator`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='申请';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `applications`
--

LOCK TABLES `applications` WRITE;
/*!40000 ALTER TABLE `applications` DISABLE KEYS */;
/*!40000 ALTER TABLE `applications` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `procedures`
--

DROP TABLE IF EXISTS `procedures`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `procedures` (
  `id` int NOT NULL AUTO_INCREMENT,
  `steps` json DEFAULT NULL,
  `current_step` json DEFAULT NULL,
  `creator` varchar(20) COLLATE utf8_bin NOT NULL,
  `create_at` bigint NOT NULL,
  `update_at` bigint NOT NULL,
  PRIMARY KEY (`id`,`creator`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `procedures`
--

LOCK TABLES `procedures` WRITE;
/*!40000 ALTER TABLE `procedures` DISABLE KEYS */;
/*!40000 ALTER TABLE `procedures` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) COLLATE utf8_bin NOT NULL,
  `email` varchar(320) COLLATE utf8_bin NOT NULL,
  `phone` varchar(45) COLLATE utf8_bin NOT NULL,
  `password` varchar(45) COLLATE utf8_bin NOT NULL,
  `identity` varchar(60) COLLATE utf8_bin NOT NULL DEFAULT '"学生"' COMMENT '身份信息',
  `manage_id` varchar(20) COLLATE utf8_bin NOT NULL DEFAULT '',
  `student_id` varchar(20) COLLATE utf8_bin NOT NULL DEFAULT '',
  `update_at` bigint NOT NULL,
  `avatar` longtext COLLATE utf8_bin,
  `course_credit` json DEFAULT NULL,
  `create_at` bigint NOT NULL,
  PRIMARY KEY (`id`,`email`,`phone`,`manage_id`,`student_id`),
  UNIQUE KEY `email_UNIQUE` (`email`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (2,'白子煜','baiziyu-fe@outlook.com','18740312553','586014BZYbzy','','','04183180',0,'https://developer.harmonyos.com/resource/image/release2/home/HarmonyOS_Developer_logo.png',NULL,1645861276),(3,'白子煜','1737586014@qq.com','13299130443','Admin12345','','','04183199',0,'https://consumer-img.huawei.com/content/dam/huawei-cbg-site/greate-china/cn/mkt/emui-11/emui-11-new/img/pc/huawei-emui-11-logo.svg',NULL,1645861899),(4,'王著名','1923790788@qq.com','1923790788','Admin12345','','','04183187',0,'https://consumer-img.huawei.com/content/dam/huawei-cbg-site/greate-china/cn/mkt/emui-11/emui-11-new/img/pc/huawei-emui-11-logo.svg',NULL,1645861899),(5,'王勇','1258001867@qq.com','869270569','Admin12345','','','04183175',0,'https://consumer-img.huawei.com/content/dam/huawei-cbg-site/greate-china/cn/mkt/emui-11/emui-11-new/img/pc/huawei-emui-11-logo.svg',NULL,1645861899),(6,'王洋','614931274@qq.com','614931274','Admin12345','','','04183177',0,'https://consumer-img.huawei.com/content/dam/huawei-cbg-site/greate-china/cn/mkt/emui-11/emui-11-new/img/pc/huawei-emui-11-logo.svg',NULL,1645861899),(7,'王勇','869270569@qq.com','314561235','Admin12345','','','04183198',0,'https://consumer-img.huawei.com/content/dam/huawei-cbg-site/greate-china/cn/mkt/emui-11/emui-11-new/img/pc/huawei-emui-11-logo.svg',NULL,1645861899),(8,'杨航','1258001276@qq.com','1923790788','Admin12345','','','04183181',0,'https://consumer-img.huawei.com/content/dam/huawei-cbg-site/greate-china/cn/mkt/emui-11/emui-11-new/img/pc/huawei-emui-11-logo.svg',NULL,1645861899),(9,'赵世宇','1399015539@qq.com','192379078834','Admin12345','','','04183178',0,'https://consumer-img.huawei.com/content/dam/huawei-cbg-site/greate-china/cn/mkt/emui-11/emui-11-new/img/pc/huawei-emui-11-logo.svg',NULL,1645861899);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-03-19 18:54:54
