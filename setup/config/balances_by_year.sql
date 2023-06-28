-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: db:3306
-- Generation Time: Jun 28, 2023 at 05:10 PM
-- Server version: 8.0.32
-- PHP Version: 8.1.16

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_accounting`
--

-- --------------------------------------------------------

--
-- Table structure for table `balances_by_year`
--

CREATE TABLE `balances_by_year` (
  `project` varchar(64) NOT NULL,
  `id` varchar(64) NOT NULL,
  `balance_id` varchar(64) NOT NULL,
  `amount` varchar(64) NOT NULL,
  `year` varchar(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `balances_by_year`
--

INSERT INTO `balances_by_year` (`project`, `id`, `balance_id`, `amount`, `year`) VALUES
('zara-viewer', '004783f1-aac5-42a4-9f4c-889fcea967f3', '10092', '8291', '2023'),
('zara-accountant', '07ce8512-4640-442a-abb2-66c5ac4b3586', '10057', '-20246.96', '2023'),
('zara-accountant', '11ccba9e-9aa0-4297-8675-04451cbb4423', '10092', '2530.85', '2023'),
('zara-auditor', '11d38987-9ced-4dc8-bd16-391b1a33a12f', '10034', '420', '2023'),
('zara-accountant', '13040bb9-cf3a-41f4-b234-b7fb712d893d', '10057', '-2000', '2023'),
('zara-auditor', '16cde015-e980-49d6-b65b-069ca6c85075', '10092', '210', '2023'),
('zara-auditor', '20d76b28-9fc9-4911-a750-700a83c42692', '10090', '16435', '2023'),
('zara-accountant', '290b8fc5-61c4-482a-9db0-c42d11c68519', '10057', '1000', '2023'),
('zara-auditor', '422e33fb-df66-48c9-aae7-aea675c0c8e0', '10087', '2420', '2023'),
('zara-auditor', '47219339-caf4-4b4e-bd70-8c923a2e6bde', '10092', '2530', '2023'),
('zara-viewer', '59f85f53-d9c0-4fbf-ae8e-8b2f9117fc01', '10029', '1210', '2023'),
('zara-accountant', '73111663-c9e6-4c4e-ba06-22e378790248', '10092', '210', '2023'),
('zara-accountant', '813c48a5-3cd3-49e2-9ae0-91e424aa92bf', '10029', '1210', '2023'),
('zara-accountant', '81c68e37-b87c-48d1-9d6b-5ce28ab189ee', '10034', '420', '2023'),
('zara-accountant', '82b50b07-c163-4118-8762-cb6f925f1a8f', '10090', '16435.75', '2023'),
('zara-viewer', '85c42cd6-72e4-4931-8adb-a5b48a68865d', '10090', '16435', '2023'),
('zara-accountant', '8eedae94-db28-4dae-9de2-0758c29e77ca', '10087', '2420', '2023'),
('zara-viewer', '908c9522-4d7b-4384-99ac-bad7791d1015', '10092', '2530', '2023'),
('zara-viewer', 'a3f7832f-080f-424b-8f65-767b80d994e0', '10092', '210', '2023'),
('zara-viewer', 'a7cc48da-00fb-48b1-b612-7c56a8815fa7', '10034', '420', '2023'),
('zara-accountant', 'c546e0fb-3bc6-457b-8520-58a04a0c9caa', '10092', '8291.84', '2023'),
('zara-auditor', 'd6a272ea-36bb-4f8e-9814-2e588cd8e816', '10029', '1210', '2023'),
('zara-accountant', 'e2bf72ea-0b06-46c9-b87d-54f9119883c3', '10057', '-7011.48', '2023'),
('zara-viewer', 'e6851a77-c984-4566-931b-cacf0d2e1b03', '10087', '2420', '2023'),
('zara-auditor', 'e8150354-b367-486b-b804-b757f158d87c', '10092', '8291', '2023');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `balances_by_year`
--
ALTER TABLE `balances_by_year`
  ADD PRIMARY KEY (`id`,`project`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
