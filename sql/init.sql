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
  `ship_type` varchar(45) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`,`user_id`,`procedure_id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  KEY `stu_id_PRIMARY` (`user_id`),
  KEY `fk_app_procedure_id` (`procedure_id`),
  CONSTRAINT `fk_app_procedure_id` FOREIGN KEY (`procedure_id`) REFERENCES `procedures` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_app_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='??????';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `applications`
--

LOCK TABLES `applications` WRITE;
/*!40000 ALTER TABLE `applications` DISABLE KEYS */;
INSERT INTO `applications` VALUES (30,1650604034,1650611212,'submit','{\"moral\": {\"list\": [{\"info\": \"?????????\", \"name\": \"????????????\", \"time\": \"2022-04-06T07:03:20.864Z\", \"files\": null, \"level\": \"provincial\", \"score\": 0}], \"score\": 0}, \"academic\": {\"award\": [{\"name\": \"??????\", \"time\": \"2022-04-22T04:54:34.925Z\", \"files\": [], \"level\": [\"provincial\", \"first\"], \"score\": 0}], \"score\": 0, \"publish\": [], \"scientific\": [], \"score_info\": {\"sum\": 0, \"moral\": 0, \"academic\": 0, \"practice\": 0}, \"dissertation\": []}, \"practice\": {\"score\": 0, \"result\": [], \"social\": {\"cadre\": [], \"score\": 0, \"activity\": []}, \"competition\": []}}','[{\"comment\": \"{\\\"moral\\\":2,\\\"practice\\\":3,\\\"academic\\\":1,\\\"sum\\\":6}\", \"edit_at\": \"2022-04-22 15:06:52\", \"user_id\": \"9542442\", \"identity\": \"manager\"}]',6,'04183180',6,'{\"sum\": 6, \"moral\": 2, \"academic\": 1, \"practice\": 3}','{\"comment\": \"{\\\"moral\\\":2,\\\"practice\\\":3,\\\"academic\\\":1,\\\"sum\\\":6}\", \"edit_at\": \"2022-04-22 15:06:52\", \"user_id\": \"9542442\", \"identity\": \"manager\"}',NULL),(31,1650728248,1650789741,'submit','{\"moral\": {\"list\": [], \"score\": 0}, \"academic\": {\"award\": [{\"name\": \"????????????\", \"time\": \"2022-04-23T15:34:37.057Z\", \"files\": [], \"level\": [\"vice_province\", \"first\"], \"score\": 0}], \"score\": 0, \"publish\": [], \"scientific\": [], \"score_info\": {\"sum\": 0, \"base\": 0, \"moral\": 0, \"academic\": 0, \"practice\": 0}, \"dissertation\": []}, \"practice\": {\"score\": 0, \"result\": [{\"name\": \"????????????\", \"time\": \"2022-04-23T15:34:37.057Z\", \"files\": [], \"level\": \"international\", \"order\": 1, \"score\": 0, \"partners\": 1}], \"social\": {\"cadre\": [{\"level\": [\"deputy_leader\", \"??????\"], \"score\": 0, \"department\": \"?????????\"}], \"score\": 0, \"activity\": [{\"name\": \"???????????????\", \"time\": [\"2022-04-23T15:34:37.057Z\", \"2022-04-23T15:34:37.057Z\"], \"files\": [], \"level\": \"college_activities\", \"score\": 0}]}, \"competition\": []}}','[]',47.5,'2003200002',6,'{\"sum\": 47.5, \"base\": 79.23, \"moral\": 13, \"academic\": 82, \"practice\": 13}',NULL,NULL),(32,1650730606,1650790394,'submit','{\"moral\": {\"list\": [{\"info\": \"\", \"name\": \"???????????????\", \"time\": \"2022-04-04T16:15:27.331Z\", \"files\": [], \"level\": \"school\", \"score\": 0}], \"score\": 0}, \"academic\": {\"award\": [{\"name\": \"????????????\", \"time\": \"2022-04-23T15:59:22.633Z\", \"files\": [], \"level\": [\"vice_province\", \"second\"], \"score\": 0}], \"score\": 0, \"publish\": [], \"scientific\": [], \"score_info\": {\"sum\": 0, \"base\": 0, \"moral\": 0, \"academic\": 0, \"practice\": 0}, \"dissertation\": []}, \"practice\": {\"score\": 0, \"result\": [{\"name\": \"sci??????\", \"time\": \"2022-04-23T15:59:22.633Z\", \"files\": [], \"level\": \"international\", \"order\": 3, \"score\": 0, \"partners\": 7}], \"social\": {\"cadre\": [{\"level\": [\"deputy_leader\", \"?????????????????????\"], \"score\": 0, \"department\": \"???????????????\"}], \"score\": 0, \"activity\": [{\"name\": \"?????????\", \"time\": [\"2022-04-23T15:59:22.633Z\", \"2022-04-23T15:59:22.633Z\"], \"files\": [], \"level\": \"college_activities\", \"score\": 0}, {\"name\": \"????????????\", \"time\": [\"2022-04-05T16:16:06.791Z\", \"2022-04-06T16:16:06.791Z\"], \"files\": null, \"level\": \"school_activities\", \"score\": 0}]}, \"competition\": []}}','[]',84.32,'2003200005',6,'{\"sum\": 84.32, \"base\": 84.32, \"moral\": 1, \"academic\": 3, \"practice\": 2}',NULL,NULL),(33,1650789900,1650790948,'submit','{\"moral\": {\"list\": [], \"score\": 0}, \"academic\": {\"award\": [], \"score\": 0, \"publish\": [], \"scientific\": [], \"score_info\": {\"sum\": 0, \"base\": 0, \"moral\": 0, \"academic\": 0, \"practice\": 0}, \"dissertation\": []}, \"practice\": {\"score\": 0, \"result\": [], \"social\": {\"cadre\": [], \"score\": 0, \"activity\": []}, \"competition\": []}}','[]',82.96,'2003200021',6,'{\"sum\": 82.96, \"base\": 82.96, \"moral\": 0, \"academic\": 0, \"practice\": 0}',NULL,'');
/*!40000 ALTER TABLE `applications` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comments`
--

DROP TABLE IF EXISTS `comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` varchar(20) COLLATE utf8_bin NOT NULL,
  `procedure_id` int NOT NULL,
  `create_at` bigint NOT NULL,
  `update_at` bigint NOT NULL,
  `content` longtext COLLATE utf8_bin,
  `reply_id` int NOT NULL,
  PRIMARY KEY (`id`,`user_id`,`procedure_id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  KEY `comment_user_fk` (`user_id`),
  KEY `comment_procedure_fk` (`procedure_id`),
  CONSTRAINT `comments_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`),
  CONSTRAINT `comments_ibfk_2` FOREIGN KEY (`procedure_id`) REFERENCES `procedures` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='?????????';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comments`
--

LOCK TABLES `comments` WRITE;
/*!40000 ALTER TABLE `comments` DISABLE KEYS */;
INSERT INTO `comments` VALUES (7,'2003200002',6,1650706584,1650706584,'@[??????_2003200002]   33300544',4),(8,'2003200002',6,1650706627,1650706627,'@[??????_2003200002]   sss',6),(9,'2003200002',6,1650706784,1650706784,'2022 ???????????????',0),(10,'2003200002',6,1650706794,1650706794,'@[??????_2003200002]   ???????????????????????????',9),(11,'2003200002',6,1650707479,1650707479,'@[??????_2003200002]   hello ??????',9),(12,'2003200002',6,1650708547,1650708547,'????????????',0),(13,'2003200002',6,1650708553,1650708553,'@[??????_2003200002]   ??????',12),(14,'04183180',6,1650725084,1650725084,'@[??????_2003200002]   hello????????????',9),(15,'9542442',6,1650725172,1650725172,'@[??????_2003200002]   ???????????????',9),(16,'9542442',6,1650725190,1650725190,'???????????????????????????',0),(17,'2003200005',6,1650758449,1650758449,'@[baiziyu_9542442]   ??????',16),(18,'2003200005',6,1650758483,1650758483,'@[??????_2003200002]   ????????????',12),(19,'2003200005',6,1650782522,1650782522,'???????????????',0);
/*!40000 ALTER TABLE `comments` ENABLE KEYS */;
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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `procedures`
--

LOCK TABLES `procedures` WRITE;
/*!40000 ALTER TABLE `procedures` DISABLE KEYS */;
INSERT INTO `procedures` VALUES (6,'{\"form\": [{\"date\": [\"2022-04-19\", \"2022-04-20\"], \"desc\": \"\", \"step\": \"deployment_mobilization_phase\", \"mentions\": [\"baiziyu-fe@outlook.com\", \"1737586014@qq.com\"]}, {\"date\": [\"2022-04-20\", \"2022-04-21\"], \"desc\": \"\", \"step\": \"individual_application_stage\", \"mentions\": [\"1737586014@qq.com\", \"admin@123.com\"]}, {\"date\": [\"2022-04-21\", \"2022-04-22\"], \"desc\": \"\", \"step\": \"policies_for_all_grades\", \"mentions\": [\"admin@123.com\", \"yanghang@123.com\"]}, {\"date\": [\"2022-04-22\", \"2022-04-26\"], \"desc\": \"\", \"step\": \"first_self_assessment\", \"mentions\": [\"admin@123.com\", \"yanghang@123.com\"]}, {\"date\": [\"2022-04-26\", \"2022-04-30\"], \"desc\": \"\", \"step\": \"first_class_announcement\", \"mentions\": [\"yanghang@123.com\", \"admin@123.com\"]}, {\"date\": [\"2022-04-30\", \"2022-05-03\"], \"desc\": \"\", \"step\": \"second_personal_self_assessment\", \"mentions\": [\"admin@123.com\"]}, {\"date\": [\"2022-05-03\", \"2022-05-06\"], \"desc\": \"\", \"step\": \"second_class_announcement\", \"mentions\": [\"admin@123.com\"]}, {\"date\": [\"2022-05-06\", \"2022-05-10\"], \"desc\": \"\", \"step\": \"grade_announcement\", \"mentions\": [\"admin@123.com\"]}, {\"date\": [\"2022-05-06\", \"2022-05-09\"], \"desc\": \"\", \"step\": \"examination_and_review_of_the_discipline_office\", \"mentions\": [\"admin@123.com\", \"1737586014@qq.com\"]}, {\"date\": [\"2022-05-09\", \"2022-05-13\"], \"desc\": \"\", \"step\": \"verification_and_deliberation_by_the_scholarship_evaluation_group\", \"mentions\": [\"admin@123.com\", \"yanghang@123.com\"]}, {\"date\": [\"2022-05-13\", \"2022-05-16\"], \"desc\": \"\", \"step\": \"school_review_stage\", \"mentions\": [\"1737586014@qq.com\", \"admin@123.com\"]}, {\"date\": [\"2022-05-16\", \"2022-05-24\"], \"desc\": \"\", \"step\": \"finish\", \"mentions\": [\"admin@123.com\"]}], \"upload\": {}}','{\"step\": \"first_self_assessment\", \"start_at\": \"2022-04-22 22:58:58\"}','9542442',1650380338,1650380338,'[{\"step\": \"individual_application_stage\", \"start_at\": \"2022-04-19 22:58:58\"}]');
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
  `identity` set('student','manager') CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT 'student' COMMENT '????????????',
  `update_at` bigint NOT NULL,
  `avatar` longtext CHARACTER SET utf8 COLLATE utf8_bin,
  `info` json DEFAULT NULL,
  `create_at` bigint NOT NULL,
  `user_id` varchar(20) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `course` float NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`,`email`,`phone`,`user_id`),
  UNIQUE KEY `email_UNIQUE` (`email`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `user_id_UNIQUE` (`user_id`),
  KEY `user_id_INDEX` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=509 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='?????????';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (20,'?????????','baiziyu-fe@outlook.com','18740312553','586014BZYbzy','student',1650118672,'http://127.0.0.1:8096/avatars/967672.jpg','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022-04-16T14:17:12.223Z\", \"college\": \"????????????\", \"professional\": \"????????????\"}',1648550339,'04183180',0),(22,'baiziyu','1737586014@qq.com','12445544123','586014BZYbzy','manager',1650725151,'http://127.0.0.1:8096/avatars/883652.jpg','{\"office\": \"ruanjian\", \"position\": \"laoshi\", \"department\": \"???????????????\"}',1648555501,'9542442',0),(23,'?????????','admin@123.com','13353453345','adminA123','student',1650113526,'http://127.0.0.1:8096/avatars/1086722.png',NULL,1650111300,'04183179',0),(24,'??????','yanghang@123.com','18433432534','yangH123','student,manager',1650115379,'http://127.0.0.1:8096/avatars/repository-open-graph-template.png','{\"type\": \"profession_degree\", \"class\": 1, \"grade\": \"2022-04-16T13:22:43.404Z\", \"college\": \"???????????????\", \"professional\": \"????????????\"}',1650113905,'04183181',0),(25,'??????','2003200002@stu.xiyou.edu.cn','-','Jxj2003200002','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2020\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650639696,'2003200002',79.23),(26,'??????','2003200004@stu.xiyou.edu.cn','-','Jxj2003200004','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650639696,'2003200004',75.19),(27,'??????','2003200005@stu.xiyou.edu.cn','17743778941','Jxj2003200005','student',1650782511,'http://127.0.0.1:8096/avatars/1033287.jpg','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2021-12-31T16:00:00.000Z\", \"college\": \"??????????????????\", \"professional\": \"????????????????????????\"}',1650639696,'2003200005',84.32),(28,'??????','2003200006@stu.xiyou.edu.cn','-','Jxj2003200006','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650639696,'2003200006',81.81),(29,'?????????','2003200007@stu.xiyou.edu.cn','-','Jxj2003200007','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 3, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650639696,'2003200007',84.28),(30,'?????????','2003200008@stu.xiyou.edu.cn','-','Jxj2003200008','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 2, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650639696,'2003200008',75.32),(31,'??????','2003200009@stu.xiyou.edu.cn','-','Jxj2003200009','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 4, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650639696,'2003200009',84.85),(32,'??????','2003200010@stu.xiyou.edu.cn','-','Jxj2003200010','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650639696,'2003200010',89.09),(33,'?????????','2003200011@stu.xiyou.edu.cn','-','Jxj2003200011','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 4, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650639696,'2003200011',83.78),(34,'??????','2003200012@stu.xiyou.edu.cn','-','Jxj2003200012','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650639696,'2003200012',85.29),(35,'?????????','2003200013@stu.xiyou.edu.cn','-','Jxj2003200013','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650639696,'2003200013',80.04),(36,'?????????','2003200014@stu.xiyou.edu.cn','-','Jxj2003200014','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650639696,'2003200014',89.58),(37,'?????????','2003200015@stu.xiyou.edu.cn','-','Jxj2003200015','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650639696,'2003200015',82.65),(38,'?????????','2003200016@stu.xiyou.edu.cn','-','Jxj2003200016','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650639696,'2003200016',80.81),(39,'??????','2003200017@stu.xiyou.edu.cn','-','Jxj2003200017','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 3, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650639696,'2003200017',82.19),(40,'?????????','2003200018@stu.xiyou.edu.cn','-','Jxj2003200018','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650639696,'2003200018',86.27),(41,'?????????','2003200019@stu.xiyou.edu.cn','-','Jxj2003200019','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650639696,'2003200019',81.68),(42,'??????','2003200020@stu.xiyou.edu.cn','-','Jxj2003200020','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 2, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650639696,'2003200020',83.46),(43,'??????','2003200021@stu.xiyou.edu.cn','-','Jxj2003200021','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650639696,'2003200021',82.96),(44,'?????????','2003200022@stu.xiyou.edu.cn','-','Jxj2003200022','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 2, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650639696,'2003200022',82.32),(498,'?????????','2003200023@stu.xiyou.edu.cn','-','Jxj2003200023','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650727941,'2003200023',83.69),(499,'?????????','2003200024@stu.xiyou.edu.cn','-','Jxj2003200024','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 2, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650727941,'2003200024',75.05),(500,'??????','2003200025@stu.xiyou.edu.cn','-','Jxj2003200025','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650727941,'2003200025',79.65),(501,'?????????','2003200026@stu.xiyou.edu.cn','-','Jxj2003200026','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650727941,'2003200026',86.6),(502,'?????????','2003200027@stu.xiyou.edu.cn','-','Jxj2003200027','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650727941,'2003200027',84.28),(503,'?????????','2003200028@stu.xiyou.edu.cn','-','Jxj2003200028','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650727941,'2003200028',89.19),(504,'?????????','2003200029@stu.xiyou.edu.cn','-','Jxj2003200029','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650727941,'2003200029',85.21),(505,'??????','2003200030@stu.xiyou.edu.cn','-','Jxj2003200030','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650727941,'2003200030',85.28),(506,'??????','2003200031@stu.xiyou.edu.cn','-','Jxj2003200031','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650727941,'2003200031',74.77),(507,'?????????','2003200032@stu.xiyou.edu.cn','-','Jxj2003200032','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650727941,'2003200032',80.72),(508,'?????????','2003200034@stu.xiyou.edu.cn','-','Jxj2003200034','student',1650729742,'','{\"type\": \"bachelor_degree\", \"class\": 1, \"grade\": \"2022\", \"college\": \"\", \"professional\": \"????????????????????????\"}',1650727941,'2003200034',79.58);
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

-- Dump completed on 2022-04-24 17:12:06
