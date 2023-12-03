-- MySQL dump 10.13  Distrib 8.0.29, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: mydb-g5
-- ------------------------------------------------------
-- Server version	8.0.35

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
-- Table structure for table `dentist`
--
CREATE DATABASE IF NOT EXISTS `mydb-g5`;

USE `mydb-g5`;

DROP TABLE IF EXISTS `dentist`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `dentist` (
  `dentist_id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `last_name` varchar(50) DEFAULT NULL,
  `license` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`dentist_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dentist`
--

LOCK TABLES `dentist` WRITE;
/*!40000 ALTER TABLE `dentist` DISABLE KEYS */;
INSERT INTO `dentist` VALUES (1,'Keith','Jones','ABC123'),(3,'Aquiles','Bailo','BBB222'),(4,'Aquiles','Vailo','BBB222');
/*!40000 ALTER TABLE `dentist` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `patient`
--

DROP TABLE IF EXISTS `patient`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `patient` (
  `patient_id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `last_name` varchar(50) DEFAULT NULL,
  `address` varchar(100) DEFAULT NULL,
  `dni` varchar(50) DEFAULT NULL,
  `discharge_date` date DEFAULT NULL,
  PRIMARY KEY (`patient_id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `patient`
--

LOCK TABLES `patient` WRITE;
/*!40000 ALTER TABLE `patient` DISABLE KEYS */;
INSERT INTO `patient` VALUES (1,'Pepe','Lota','calle falsa 123','48655324','2024-11-30'),(2,'Larry','Cachón','Avenida Inventada 923','5364212','2023-03-25'),(3,'John','Doe','123 Main St','48651324','2023-10-27'),(4,'Sancho','Panza','Calle Siempreviva 742','50345678','2023-05-03');
/*!40000 ALTER TABLE `patient` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `turn`
--

DROP TABLE IF EXISTS `turn`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `turn` (
  `turn_id` bigint NOT NULL AUTO_INCREMENT,
  `dentist_id` bigint DEFAULT NULL,
  `patient_id` bigint DEFAULT NULL,
  `date_time` datetime DEFAULT NULL,
  `details` text,
  PRIMARY KEY (`turn_id`),
  KEY `id_patient_dentist_idx` (`dentist_id`),
  KEY `id_turno_patient_idx` (`patient_id`),
  CONSTRAINT `id_turno_dentist` FOREIGN KEY (`dentist_id`) REFERENCES `dentist` (`dentist_id`),
  CONSTRAINT `id_turno_patient` FOREIGN KEY (`patient_id`) REFERENCES `patient` (`patient_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `turn`
--

LOCK TABLES `turn` WRITE;
/*!40000 ALTER TABLE `turn` DISABLE KEYS */;
INSERT INTO `turn` VALUES (1,1,1,'2024-08-02 16:45:00','from postman'),(2,3,1,'2024-08-02 16:45:00','from postman[PATCHED]'),(4,4,1,'2024-08-02 16:45:00','from postman'),(5,3,4,'2024-03-16 10:30:00','from postman[UPDATED]'),(7,4,3,'2024-08-02 16:45:00','from postman'),(8,4,3,'2024-08-02 16:45:00','from postman');
/*!40000 ALTER TABLE `turn` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-12-03  8:27:26
