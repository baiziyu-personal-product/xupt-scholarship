CREATE DATABASE  IF NOT EXISTS `xupt-scholarship` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_bin */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `xupt-scholarship`;
-- MySQL dump 10.13  Distrib 8.0.27, for Win64 (x86_64)
--
-- Host: 123.56.239.54    Database: xupt-scholarship
-- ------------------------------------------------------
-- Server version	8.0.26

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
-- Table structure for table `application_form`
--

DROP TABLE IF EXISTS `application_form`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `application_form` (
  `id` int NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL,
  `student_id` int NOT NULL,
  `info` json DEFAULT NULL,
  `editable` int NOT NULL DEFAULT '0' COMMENT '是否支持编辑\n1 支持\n0 不支持',
  `update_time` datetime NOT NULL COMMENT '最后更新时间',
  `year` year NOT NULL,
  `archive` int NOT NULL DEFAULT '0' COMMENT '是否归档\n1 归档\n0 未归档，流程处理中',
  `process_id` int NOT NULL,
  `list_id` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  KEY `apf_sid_idx` (`student_id`),
  KEY `ap_pid_idx` (`process_id`),
  KEY `apf_lid_idx` (`list_id`),
  CONSTRAINT `apf_lid` FOREIGN KEY (`list_id`) REFERENCES `application_list` (`id`),
  CONSTRAINT `apf_pid` FOREIGN KEY (`process_id`) REFERENCES `application_process` (`id`),
  CONSTRAINT `apf_sid` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='奖学金申请表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `application_form`
--

LOCK TABLES `application_form` WRITE;
/*!40000 ALTER TABLE `application_form` DISABLE KEYS */;
/*!40000 ALTER TABLE `application_form` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `application_list`
--

DROP TABLE IF EXISTS `application_list`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `application_list` (
  `id` int NOT NULL AUTO_INCREMENT,
  `process_id` int NOT NULL,
  `student_id` int NOT NULL,
  `user_id` int NOT NULL,
  `status` int NOT NULL COMMENT '状态\n0 正常\n-1 异常\n1 结束',
  `process` json NOT NULL,
  `create_time` datetime NOT NULL,
  `end_time` datetime NOT NULL,
  `status_info` longtext COLLATE utf8_bin COMMENT '状态信息，记录最后一次处理评语',
  UNIQUE KEY `id_UNIQUE` (`id`),
  KEY `pid_idx` (`process_id`),
  KEY `apl_sid_idx` (`student_id`),
  KEY `apl_uid_idx` (`user_id`),
  CONSTRAINT `apl_pid` FOREIGN KEY (`process_id`) REFERENCES `application_process` (`id`),
  CONSTRAINT `apl_sid` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`),
  CONSTRAINT `apl_uid` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='申请列表，学生个人的申请过程';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `application_list`
--

LOCK TABLES `application_list` WRITE;
/*!40000 ALTER TABLE `application_list` DISABLE KEYS */;
/*!40000 ALTER TABLE `application_list` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `application_process`
--

DROP TABLE IF EXISTS `application_process`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `application_process` (
  `id` int NOT NULL AUTO_INCREMENT,
  `status` int NOT NULL DEFAULT '0' COMMENT '当前状态',
  `create_time` datetime NOT NULL,
  `end_time` datetime NOT NULL,
  `manage_id` int NOT NULL COMMENT '发起者',
  `year` year NOT NULL,
  `info` json NOT NULL,
  `reviewers` longtext COLLATE utf8_bin,
  `students` longtext COLLATE utf8_bin,
  `list_ids` longtext COLLATE utf8_bin,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  KEY `mid_idx` (`manage_id`),
  CONSTRAINT `mid` FOREIGN KEY (`manage_id`) REFERENCES `reviewers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='奖学金流程';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `application_process`
--

LOCK TABLES `application_process` WRITE;
/*!40000 ALTER TABLE `application_process` DISABLE KEYS */;
/*!40000 ALTER TABLE `application_process` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reviewers`
--

DROP TABLE IF EXISTS `reviewers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `reviewers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `name` varchar(45) COLLATE utf8_bin NOT NULL,
  `professional` varchar(45) COLLATE utf8_bin NOT NULL,
  `institute` varchar(45) COLLATE utf8_bin NOT NULL COMMENT '学院',
  PRIMARY KEY (`id`,`user_id`,`institute`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `user_id_UNIQUE` (`user_id`),
  CONSTRAINT `rew_uid` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='参与评审的人员';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reviewers`
--

LOCK TABLES `reviewers` WRITE;
/*!40000 ALTER TABLE `reviewers` DISABLE KEYS */;
/*!40000 ALTER TABLE `reviewers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `students`
--

DROP TABLE IF EXISTS `students`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `students` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `name` varchar(45) COLLATE utf8_bin NOT NULL,
  `student_id` varchar(20) COLLATE utf8_bin NOT NULL,
  `gender` varchar(5) COLLATE utf8_bin NOT NULL DEFAULT 'woman' COMMENT '性别：man/woman',
  `professional` varchar(45) COLLATE utf8_bin NOT NULL COMMENT '专业',
  `class` int NOT NULL COMMENT '班级',
  `session` int NOT NULL COMMENT '年级\n',
  `receiving_status` int NOT NULL COMMENT '领取奖学金状态：\n1 已领取\n0 待领取\n-1 无需领取',
  PRIMARY KEY (`id`,`user_id`,`student_id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `user_id_UNIQUE` (`user_id`),
  UNIQUE KEY `student_id_UNIQUE` (`student_id`),
  CONSTRAINT `stu_uid` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='学生表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `students`
--

LOCK TABLES `students` WRITE;
/*!40000 ALTER TABLE `students` DISABLE KEYS */;
/*!40000 ALTER TABLE `students` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `email` varchar(45) COLLATE utf8_bin NOT NULL,
  `phone` varchar(45) COLLATE utf8_bin NOT NULL,
  `password` varchar(45) COLLATE utf8_bin NOT NULL,
  `avatar` longtext COLLATE utf8_bin,
  `identity` int NOT NULL DEFAULT '0' COMMENT '身份',
  `create_time` datetime NOT NULL,
  `institute` varchar(45) COLLATE utf8_bin NOT NULL COMMENT '学院',
  PRIMARY KEY (`id`,`email`,`phone`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `email_UNIQUE` (`email`),
  UNIQUE KEY `phone_UNIQUE` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
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

-- Dump completed on 2022-01-27 18:54:40
