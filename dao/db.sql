CREATE DATABASE `lottery` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

CREATE TABLE `dlt` (
  `id` int NOT NULL AUTO_INCREMENT,
  `num` varchar(45) DEFAULT NULL,
  `num_1` int DEFAULT NULL,
  `num_2` int DEFAULT NULL,
  `num_3` int DEFAULT NULL,
  `num_4` int DEFAULT NULL,
  `num_5` int DEFAULT NULL,
  `num_6` int DEFAULT NULL,
  `num_7` int DEFAULT NULL,
  `open_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20033 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
