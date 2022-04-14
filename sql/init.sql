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
-- Table structure for table `actions`
--

DROP TABLE IF EXISTS `actions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `actions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `create_at` bigint NOT NULL,
  `update_at` bigint NOT NULL,
  `user_id` varchar(20) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `apply_id` int NOT NULL,
  `info` json NOT NULL,
  PRIMARY KEY (`id`,`user_id`,`apply_id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  KEY `fk_action_apply_id` (`apply_id`),
  KEY `fk_action_user_id` (`user_id`),
  CONSTRAINT `fk_action_apply_id` FOREIGN KEY (`apply_id`) REFERENCES `applications` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_action_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `actions`
--

LOCK TABLES `actions` WRITE;
/*!40000 ALTER TABLE `actions` DISABLE KEYS */;
/*!40000 ALTER TABLE `actions` ENABLE KEYS */;
UNLOCK TABLES;

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
  `status` varchar(30) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '"save"' COMMENT 'save\\nsubmit\\nfinish\\ndroped\\nerror',
  `info` json NOT NULL,
  `history` json NOT NULL,
  `score` float NOT NULL DEFAULT '0',
  `user_id` varchar(20) COLLATE utf8_bin NOT NULL,
  `procedure_id` int NOT NULL,
  `score_info` json DEFAULT NULL,
  `step` json DEFAULT NULL,
  PRIMARY KEY (`id`,`user_id`,`procedure_id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  KEY `stu_id_PRIMARY` (`user_id`),
  KEY `fk_app_procedure_id` (`procedure_id`),
  CONSTRAINT `fk_app_procedure_id` FOREIGN KEY (`procedure_id`) REFERENCES `procedures` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_app_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='申请';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `applications`
--

LOCK TABLES `applications` WRITE;
/*!40000 ALTER TABLE `applications` DISABLE KEYS */;
INSERT INTO `applications` VALUES (29,1649402894,1649403931,'submit','{\"moral\": {\"list\": [{\"info\": \"\", \"name\": \"国家奖学金\", \"time\": \"2022-04-06T07:27:47.870Z\", \"files\": null, \"level\": \"national\", \"score\": 0}], \"score\": 0}, \"score\": 0, \"academic\": {\"award\": [{\"name\": \"静思成功\", \"time\": \"2022-04-01T07:18:06.802Z\", \"files\": [], \"level\": [\"provincial\", \"second\"], \"score\": 0}], \"score\": 0, \"publish\": [{\"name\": \"js导论\", \"time\": \"2022-04-07T07:39:18.842Z\", \"files\": null, \"level\": \"a_publishing_house\", \"score\": 0, \"fonts_count\": 2, \"publish_house_name\": \"工业出版社\"}], \"scientific\": [{\"name\": \"挑战书\", \"time\": [\"2022-04-01T07:38:33.915Z\", \"2022-04-02T07:38:33.915Z\"], \"files\": null, \"level\": \"national\", \"order\": 1, \"score\": 0, \"partners\": 1, \"funds_due\": 1, \"distribute\": 1, \"funds_actually_received\": 1}], \"dissertation\": [{\"name\": \"sci\", \"time\": \"2022-04-05T07:38:59.337Z\", \"files\": null, \"level\": \"SCI_2\", \"score\": 0, \"id_number\": \"25\"}]}, \"practice\": {\"score\": 0, \"result\": [{\"name\": \"互联网竞赛\", \"time\": \"2022-04-04T07:27:58.096Z\", \"files\": null, \"level\": \"national\", \"order\": 1, \"score\": 0, \"partners\": 1}], \"social\": {\"cadre\": [{\"level\": [\"leader\", \"研究生会主席\"], \"score\": 0, \"department\": \"部门负责人\"}], \"score\": 0, \"activity\": [{\"name\": \"校园行\", \"time\": [\"2021-04-15T07:18:06.802Z\", \"2021-04-22T07:18:06.802Z\"], \"files\": [], \"level\": \"college_activities\", \"score\": 0}]}, \"competition\": [{\"name\": \"竞赛杯\", \"time\": \"2022-04-08T07:18:06.802Z\", \"files\": [], \"level\": [\"national\", \"first\"], \"order\": 1, \"score\": 0, \"partners\": 1}]}}','{}',0,'04183180',5,NULL,NULL);
/*!40000 ALTER TABLE `applications` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `logs`
--

DROP TABLE IF EXISTS `logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `logs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `create_at` bigint NOT NULL,
  `update_at` bigint NOT NULL,
  `data` longtext CHARACTER SET utf8 COLLATE utf8_bin,
  `info` json DEFAULT NULL,
  `user_id` varchar(20) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`id`,`user_id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  KEY `fk_log_user_id` (`user_id`),
  CONSTRAINT `fk_log_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `logs`
--

LOCK TABLES `logs` WRITE;
/*!40000 ALTER TABLE `logs` DISABLE KEYS */;
/*!40000 ALTER TABLE `logs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `procedures`
--

DROP TABLE IF EXISTS `procedures`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `procedures` (
  `id` int NOT NULL AUTO_INCREMENT,
  `info` json DEFAULT NULL,
  `current_step` json DEFAULT NULL,
  `user_id` varchar(20) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `create_at` bigint NOT NULL,
  `update_at` bigint NOT NULL,
  `history` json DEFAULT NULL,
  PRIMARY KEY (`id`,`user_id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  KEY `fk_procedure_user_id` (`user_id`),
  CONSTRAINT `fk_procedure_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `procedures`
--

LOCK TABLES `procedures` WRITE;
/*!40000 ALTER TABLE `procedures` DISABLE KEYS */;
INSERT INTO `procedures` VALUES (5,'{\"form\": {\"finish\": {\"date\": [\"2022-04-15\", \"2022-04-16\"], \"desc\": \"\", \"mentions\": [\"baiziyu-fe@outlook.com\", \"1737586014@qq.com\"]}, \"grade_announcement\": {\"date\": [\"2022-04-09\", \"2022-04-11\"], \"desc\": \"\", \"mentions\": []}, \"school_review_stage\": {\"date\": [\"2022-04-13\", \"2022-04-15\"], \"desc\": \"\", \"mentions\": []}, \"first_self_assessment\": {\"date\": [\"2022-04-02\", \"2022-04-04\"], \"desc\": \"\", \"mentions\": []}, \"policies_for_all_grades\": {\"date\": [\"2022-04-01\", \"2022-04-02\"], \"desc\": \"\", \"mentions\": [\"1737586014@qq.com\"]}, \"first_class_announcement\": {\"date\": [\"2022-04-04\", \"2022-04-05\"], \"desc\": \"\", \"mentions\": []}, \"second_class_announcement\": {\"date\": [\"2022-04-07\", \"2022-04-09\"], \"desc\": \"\", \"mentions\": []}, \"individual_application_stage\": {\"date\": [\"2022-03-31\", \"2022-04-01\"], \"desc\": \"\", \"mentions\": [\"baiziyu-fe@outlook.com\"]}, \"deployment_mobilization_phase\": {\"date\": [\"2022-03-30\", \"2022-03-31\"], \"desc\": \"\", \"mentions\": [\"baiziyu-fe@outlook.com\", \"1737586014@qq.com\"]}, \"second_personal_self_assessment\": {\"date\": [\"2022-04-05\", \"2022-04-07\"], \"desc\": \"\", \"mentions\": []}, \"examination_and_review_of_the_discipline_office\": {\"date\": [\"2022-04-09\", \"2022-04-11\"], \"desc\": \"\", \"mentions\": []}, \"verification_and_deliberation_by_the_scholarship_evaluation_group\": {\"date\": [\"2022-04-11\", \"2022-04-13\"], \"desc\": \"\", \"mentions\": []}}, \"upload\": {}}','[]','9542442',1648650765,1648650765,'[]');
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
  `name` varchar(45) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `email` varchar(320) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `phone` varchar(45) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `password` varchar(45) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `identity` set('student','manager') CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT 'student' COMMENT '身份信息',
  `update_at` bigint NOT NULL,
  `avatar` longtext CHARACTER SET utf8 COLLATE utf8_bin,
  `info` json DEFAULT NULL,
  `create_at` bigint NOT NULL,
  `user_id` varchar(20) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`id`,`email`,`phone`,`user_id`),
  UNIQUE KEY `email_UNIQUE` (`email`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `user_id_UNIQUE` (`user_id`),
  KEY `user_id_INDEX` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (20,'白子煜','baiziyu-fe@outlook.com','18740312553','586014BZYbzy','student',1648550454,'http://127.0.0.1:8096/avatars/967672.jpg',NULL,1648550339,'04183180'),(22,'baiziyu','1737586014@qq.com','12445544123','586014BZYbzy','manager',1648555521,'http://127.0.0.1:8096/avatars/967672.jpg',NULL,1648555501,'9542442');
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

-- Dump completed on 2022-04-14 21:41:30
