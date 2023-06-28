-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: db:3306
-- Generation Time: Jun 28, 2023 at 05:11 PM
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
-- Table structure for table `entries`
--

CREATE TABLE `entries` (
  `project` varchar(64) NOT NULL,
  `id` varchar(64) NOT NULL,
  `contract_id` varchar(64) NOT NULL,
  `user_id` varchar(64) NOT NULL,
  `account_id` varchar(64) NOT NULL,
  `multiple_entry_id` varchar(64) NOT NULL,
  `balance_by_year_id` varchar(64) DEFAULT NULL,
  `currency` varchar(3) NOT NULL,
  `concept` varchar(256) DEFAULT NULL,
  `date` varchar(64) NOT NULL,
  `quarter` varchar(1) NOT NULL,
  `debit_balance` varchar(64) NOT NULL,
  `credit_balance` varchar(64) NOT NULL,
  `state` varchar(64) NOT NULL,
  `state_reason` varchar(256) DEFAULT NULL,
  `creation_date` varchar(64) NOT NULL,
  `last_update` varchar(64) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `entries`
--

INSERT INTO `entries` (`project`, `id`, `contract_id`, `user_id`, `account_id`, `multiple_entry_id`, `balance_by_year_id`, `currency`, `concept`, `date`, `quarter`, `debit_balance`, `credit_balance`, `state`, `state_reason`, `creation_date`, `last_update`) VALUES
('zara-accountant', '00b0e779-581d-4636-9c99-69b724ebae6f', 'contablock-contract-id', 'zara-user-accountant-id', '400', '1dd47dcf-3887-4753-91a9-45372e31b33e', '8eedae94-db28-4dae-9de2-0758c29e77ca', 'EUR', 'compras', '2023-04-17T09:08:08Z', '2', '0', '2420', 'CREATED', NULL, '2023-06-28T17:09:47Z', '2023-06-28T17:09:47Z'),
('zara-auditor', '00b0e779-581d-4636-9c99-69b724ebae6f', 'contablock-contract-id', 'zara-user-accountant-id', '400', '1dd47dcf-3887-4753-91a9-45372e31b33e', '422e33fb-df66-48c9-aae7-aea675c0c8e0', 'EUR', 'compras', '2023-04-17T11:08:08+02:00', '2', '0', '2420', 'CREATED', NULL, '2023-06-28T19:09:47+02:00', '2023-06-28T19:09:47+02:00'),
('zara-viewer', '00b0e779-581d-4636-9c99-69b724ebae6f', 'contablock-contract-id', 'zara-user-accountant-id', '400', '1dd47dcf-3887-4753-91a9-45372e31b33e', 'e6851a77-c984-4566-931b-cacf0d2e1b03', 'EUR', 'compras', '2023-04-17T11:08:08+02:00', '2', '0', '2420', 'CREATED', NULL, '2023-06-28T19:09:47+02:00', '2023-06-28T19:09:47+02:00'),
('zara-accountant', '5346ce59-96e1-4037-bc5b-79d50e2393ad', 'contablock-contract-id', 'zara-user-accountant-id', '4751', '53592761-8b61-454e-aeef-898fd84a5ffd', '11ccba9e-9aa0-4297-8675-04451cbb4423', 'EUR', 'nomina mes', '2023-06-28T17:05:07Z', '2', '0', '2530.85', 'CREATED', NULL, '2023-06-28T17:06:35Z', '2023-06-28T17:06:35Z'),
('zara-auditor', '5346ce59-96e1-4037-bc5b-79d50e2393ad', 'contablock-contract-id', 'zara-user-accountant-id', '4751', '53592761-8b61-454e-aeef-898fd84a5ffd', '47219339-caf4-4b4e-bd70-8c923a2e6bde', 'EUR', 'nomina mes', '2023-06-28T19:05:07+02:00', '2', '0', '2530', 'CREATED', NULL, '2023-06-28T19:06:35+02:00', '2023-06-28T19:06:35+02:00'),
('zara-viewer', '5346ce59-96e1-4037-bc5b-79d50e2393ad', 'contablock-contract-id', 'zara-user-accountant-id', '4751', '53592761-8b61-454e-aeef-898fd84a5ffd', '908c9522-4d7b-4384-99ac-bad7791d1015', 'EUR', 'nomina mes', '2023-06-28T19:05:07+02:00', '2', '0', '2530', 'CREATED', NULL, '2023-06-28T19:06:35+02:00', '2023-06-28T19:06:35+02:00'),
('zara-accountant', '635c0543-d4dc-4a39-8e08-45598fdf43ce', 'contablock-contract-id', 'zara-user-accountant-id', '476', '53592761-8b61-454e-aeef-898fd84a5ffd', 'c546e0fb-3bc6-457b-8520-58a04a0c9caa', 'EUR', 'nomina mes', '2023-06-28T17:05:07Z', '2', '0', '8291.84', 'CREATED', NULL, '2023-06-28T17:06:35Z', '2023-06-28T17:06:35Z'),
('zara-auditor', '635c0543-d4dc-4a39-8e08-45598fdf43ce', 'contablock-contract-id', 'zara-user-accountant-id', '476', '53592761-8b61-454e-aeef-898fd84a5ffd', 'e8150354-b367-486b-b804-b757f158d87c', 'EUR', 'nomina mes', '2023-06-28T19:05:07+02:00', '2', '0', '8291', 'CREATED', NULL, '2023-06-28T19:06:35+02:00', '2023-06-28T19:06:35+02:00'),
('zara-viewer', '635c0543-d4dc-4a39-8e08-45598fdf43ce', 'contablock-contract-id', 'zara-user-accountant-id', '476', '53592761-8b61-454e-aeef-898fd84a5ffd', '004783f1-aac5-42a4-9f4c-889fcea967f3', 'EUR', 'nomina mes', '2023-06-28T19:05:07+02:00', '2', '0', '8291', 'CREATED', NULL, '2023-06-28T19:06:36+02:00', '2023-06-28T19:06:36+02:00'),
('zara-accountant', '79c23590-bf92-481a-b5be-23f4eec636e1', 'contablock-contract-id', 'zara-user-accountant-id', '642', '53592761-8b61-454e-aeef-898fd84a5ffd', 'e2bf72ea-0b06-46c9-b87d-54f9119883c3', 'EUR', 'nomina mes', '2023-06-28T17:05:07Z', '2', '7011.48', '0', 'CREATED', NULL, '2023-06-28T17:06:35Z', '2023-06-28T17:06:35Z'),
('zara-auditor', '79c23590-bf92-481a-b5be-23f4eec636e1', 'contablock-contract-id', 'zara-user-accountant-id', '642', '53592761-8b61-454e-aeef-898fd84a5ffd', NULL, 'EUR', 'nomina mes', '2023-06-28T19:05:07+02:00', '2', '7011', '0', 'CREATED', NULL, '2023-06-28T19:06:36+02:00', '2023-06-28T19:06:36+02:00'),
('zara-viewer', '79c23590-bf92-481a-b5be-23f4eec636e1', 'contablock-contract-id', 'zara-user-accountant-id', '642', '53592761-8b61-454e-aeef-898fd84a5ffd', NULL, 'EUR', 'nomina mes', '2023-06-28T19:05:07+02:00', '2', '7011', '0', 'CREATED', NULL, '2023-06-28T19:06:36+02:00', '2023-06-28T19:06:36+02:00'),
('zara-accountant', 'c4c74b9e-ef65-4295-869e-3b1d124e364a', 'contablock-contract-id', 'zara-user-accountant-id', '477', '2805e718-e3ea-4f4e-8354-c151c8f08099', '73111663-c9e6-4c4e-ba06-22e378790248', 'EUR', 'venta de mercaderías', '2023-03-23T09:01:47Z', '1', '0', '210', 'CREATED', NULL, '2023-06-28T17:07:55Z', '2023-06-28T17:07:55Z'),
('zara-auditor', 'c4c74b9e-ef65-4295-869e-3b1d124e364a', 'contablock-contract-id', 'zara-user-accountant-id', '477', '2805e718-e3ea-4f4e-8354-c151c8f08099', '16cde015-e980-49d6-b65b-069ca6c85075', 'EUR', 'venta de mercaderías', '2023-03-23T10:01:47+01:00', '1', '0', '210', 'CREATED', NULL, '2023-06-28T19:07:55+02:00', '2023-06-28T19:07:55+02:00'),
('zara-viewer', 'c4c74b9e-ef65-4295-869e-3b1d124e364a', 'contablock-contract-id', 'zara-user-accountant-id', '477', '2805e718-e3ea-4f4e-8354-c151c8f08099', 'a3f7832f-080f-424b-8f65-767b80d994e0', 'EUR', 'venta de mercaderías', '2023-03-23T10:01:47+01:00', '1', '0', '210', 'CREATED', NULL, '2023-06-28T19:07:55+02:00', '2023-06-28T19:07:55+02:00'),
('zara-accountant', 'd515d02b-12ce-43d3-88c2-cf76523b1419', 'contablock-contract-id', 'zara-user-accountant-id', '472', '1dd47dcf-3887-4753-91a9-45372e31b33e', '81c68e37-b87c-48d1-9d6b-5ce28ab189ee', 'EUR', 'compras', '2023-04-17T09:08:08Z', '2', '420', '0', 'CREATED', NULL, '2023-06-28T17:09:47Z', '2023-06-28T17:09:47Z'),
('zara-auditor', 'd515d02b-12ce-43d3-88c2-cf76523b1419', 'contablock-contract-id', 'zara-user-accountant-id', '472', '1dd47dcf-3887-4753-91a9-45372e31b33e', '11d38987-9ced-4dc8-bd16-391b1a33a12f', 'EUR', 'compras', '2023-04-17T11:08:08+02:00', '2', '420', '0', 'CREATED', NULL, '2023-06-28T19:09:47+02:00', '2023-06-28T19:09:47+02:00'),
('zara-viewer', 'd515d02b-12ce-43d3-88c2-cf76523b1419', 'contablock-contract-id', 'zara-user-accountant-id', '472', '1dd47dcf-3887-4753-91a9-45372e31b33e', 'a7cc48da-00fb-48b1-b612-7c56a8815fa7', 'EUR', 'compras', '2023-04-17T11:08:08+02:00', '2', '420', '0', 'CREATED', NULL, '2023-06-28T19:09:47+02:00', '2023-06-28T19:09:47+02:00'),
('zara-accountant', 'e3ce48ae-cc73-455c-98ff-33aa7746dddc', 'contablock-contract-id', 'zara-user-accountant-id', '430', '2805e718-e3ea-4f4e-8354-c151c8f08099', '813c48a5-3cd3-49e2-9ae0-91e424aa92bf', 'EUR', 'venta de mercaderías', '2023-03-23T09:01:47Z', '1', '1210', '0', 'CREATED', NULL, '2023-06-28T17:07:55Z', '2023-06-28T17:07:55Z'),
('zara-auditor', 'e3ce48ae-cc73-455c-98ff-33aa7746dddc', 'contablock-contract-id', 'zara-user-accountant-id', '430', '2805e718-e3ea-4f4e-8354-c151c8f08099', 'd6a272ea-36bb-4f8e-9814-2e588cd8e816', 'EUR', 'venta de mercaderías', '2023-03-23T10:01:47+01:00', '1', '1210', '0', 'CREATED', NULL, '2023-06-28T19:07:55+02:00', '2023-06-28T19:07:55+02:00'),
('zara-viewer', 'e3ce48ae-cc73-455c-98ff-33aa7746dddc', 'contablock-contract-id', 'zara-user-accountant-id', '430', '2805e718-e3ea-4f4e-8354-c151c8f08099', '59f85f53-d9c0-4fbf-ae8e-8b2f9117fc01', 'EUR', 'venta de mercaderías', '2023-03-23T10:01:47+01:00', '1', '1210', '0', 'CREATED', NULL, '2023-06-28T19:07:55+02:00', '2023-06-28T19:07:55+02:00'),
('zara-accountant', 'e83f84ce-dead-4325-8b26-5645dee78cf7', 'contablock-contract-id', 'zara-user-accountant-id', '465', '53592761-8b61-454e-aeef-898fd84a5ffd', '82b50b07-c163-4118-8762-cb6f925f1a8f', 'EUR', 'nomina mes', '2023-06-28T17:05:07Z', '2', '0', '16435.75', 'CREATED', NULL, '2023-06-28T17:06:35Z', '2023-06-28T17:06:35Z'),
('zara-auditor', 'e83f84ce-dead-4325-8b26-5645dee78cf7', 'contablock-contract-id', 'zara-user-accountant-id', '465', '53592761-8b61-454e-aeef-898fd84a5ffd', '20d76b28-9fc9-4911-a750-700a83c42692', 'EUR', 'nomina mes', '2023-06-28T19:05:07+02:00', '2', '0', '16435', 'CREATED', NULL, '2023-06-28T19:06:36+02:00', '2023-06-28T19:06:36+02:00'),
('zara-viewer', 'e83f84ce-dead-4325-8b26-5645dee78cf7', 'contablock-contract-id', 'zara-user-accountant-id', '465', '53592761-8b61-454e-aeef-898fd84a5ffd', '85c42cd6-72e4-4931-8adb-a5b48a68865d', 'EUR', 'nomina mes', '2023-06-28T19:05:07+02:00', '2', '0', '16435', 'CREATED', NULL, '2023-06-28T19:06:36+02:00', '2023-06-28T19:06:36+02:00'),
('zara-accountant', 'eeb69a06-5fbf-4fc6-b419-ce2a25341706', 'contablock-contract-id', 'zara-user-accountant-id', '700', '2805e718-e3ea-4f4e-8354-c151c8f08099', '290b8fc5-61c4-482a-9db0-c42d11c68519', 'EUR', 'venta de mercaderías', '2023-03-23T09:01:47Z', '1', '0', '1000', 'CREATED', NULL, '2023-06-28T17:07:55Z', '2023-06-28T17:07:55Z'),
('zara-auditor', 'eeb69a06-5fbf-4fc6-b419-ce2a25341706', 'contablock-contract-id', 'zara-user-accountant-id', '700', '2805e718-e3ea-4f4e-8354-c151c8f08099', NULL, 'EUR', 'venta de mercaderías', '2023-03-23T10:01:47+01:00', '1', '0', '1000', 'CREATED', NULL, '2023-06-28T19:07:55+02:00', '2023-06-28T19:07:55+02:00'),
('zara-viewer', 'eeb69a06-5fbf-4fc6-b419-ce2a25341706', 'contablock-contract-id', 'zara-user-accountant-id', '700', '2805e718-e3ea-4f4e-8354-c151c8f08099', NULL, 'EUR', 'venta de mercaderías', '2023-03-23T10:01:47+01:00', '1', '0', '1000', 'CREATED', NULL, '2023-06-28T19:07:55+02:00', '2023-06-28T19:07:55+02:00'),
('zara-accountant', 'f40ba157-5972-4316-9d27-b452d760f650', 'contablock-contract-id', 'zara-user-accountant-id', '600', '1dd47dcf-3887-4753-91a9-45372e31b33e', '13040bb9-cf3a-41f4-b234-b7fb712d893d', 'EUR', 'compras', '2023-04-17T09:08:08Z', '2', '2000', '0', 'CREATED', NULL, '2023-06-28T17:09:47Z', '2023-06-28T17:09:47Z'),
('zara-auditor', 'f40ba157-5972-4316-9d27-b452d760f650', 'contablock-contract-id', 'zara-user-accountant-id', '600', '1dd47dcf-3887-4753-91a9-45372e31b33e', NULL, 'EUR', 'compras', '2023-04-17T11:08:08+02:00', '2', '2000', '0', 'CREATED', NULL, '2023-06-28T19:09:47+02:00', '2023-06-28T19:09:47+02:00'),
('zara-viewer', 'f40ba157-5972-4316-9d27-b452d760f650', 'contablock-contract-id', 'zara-user-accountant-id', '600', '1dd47dcf-3887-4753-91a9-45372e31b33e', NULL, 'EUR', 'compras', '2023-04-17T11:08:08+02:00', '2', '2000', '0', 'CREATED', NULL, '2023-06-28T19:09:47+02:00', '2023-06-28T19:09:47+02:00'),
('zara-accountant', 'f68106e0-bcfa-4044-8f3a-ab473cf58ed0', 'contablock-contract-id', 'zara-user-accountant-id', '640', '53592761-8b61-454e-aeef-898fd84a5ffd', '07ce8512-4640-442a-abb2-66c5ac4b3586', 'EUR', 'nomina mes', '2023-06-28T17:05:07Z', '2', '20246.96', '0', 'CREATED', NULL, '2023-06-28T17:06:35Z', '2023-06-28T17:06:35Z'),
('zara-auditor', 'f68106e0-bcfa-4044-8f3a-ab473cf58ed0', 'contablock-contract-id', 'zara-user-accountant-id', '640', '53592761-8b61-454e-aeef-898fd84a5ffd', NULL, 'EUR', 'nomina mes', '2023-06-28T19:05:07+02:00', '2', '20246', '0', 'CREATED', NULL, '2023-06-28T19:06:35+02:00', '2023-06-28T19:06:35+02:00'),
('zara-viewer', 'f68106e0-bcfa-4044-8f3a-ab473cf58ed0', 'contablock-contract-id', 'zara-user-accountant-id', '640', '53592761-8b61-454e-aeef-898fd84a5ffd', NULL, 'EUR', 'nomina mes', '2023-06-28T19:05:07+02:00', '2', '20246', '0', 'CREATED', NULL, '2023-06-28T19:06:35+02:00', '2023-06-28T19:06:35+02:00');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `entries`
--
ALTER TABLE `entries`
  ADD PRIMARY KEY (`id`,`project`),
  ADD UNIQUE KEY `project` (`project`,`id`,`multiple_entry_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
