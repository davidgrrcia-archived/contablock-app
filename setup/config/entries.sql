-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: db:3306
-- Generation Time: Jun 28, 2023 at 10:30 PM
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
('zara-accountant', '00a80104-549b-4f0d-9aa7-6c3067811dc7', 'contablock-contract-id', 'zara-user-accountant-id', '642', 'eaf6c055-23e6-4af2-80e6-62f4ad1842bf', '286a008a-18f1-4a8e-99f0-2053aee78154', 'EUR', 'nomina mes', '2023-06-02T14:22:16Z', '2', '7011.48', '0', 'CREATED', NULL, '2023-06-28T22:23:53Z', '2023-06-28T22:23:53Z'),
('zara-auditor', '00a80104-549b-4f0d-9aa7-6c3067811dc7', 'contablock-contract-id', 'zara-user-accountant-id', '642', 'eaf6c055-23e6-4af2-80e6-62f4ad1842bf', '05841c70-7ef3-46fa-8b57-dfff5b08171a', 'EUR', 'nomina mes', '2023-06-02T16:22:16+02:00', '2', '7011.47999', '0', 'CREATED', NULL, '2023-06-29T00:23:53+02:00', '2023-06-29T00:23:53+02:00'),
('zara-viewer', '00a80104-549b-4f0d-9aa7-6c3067811dc7', 'contablock-contract-id', 'zara-user-accountant-id', '642', 'eaf6c055-23e6-4af2-80e6-62f4ad1842bf', '7b4c3794-341e-4ac2-baf2-cd51ba055b55', 'EUR', 'nomina mes', '2023-06-02T16:22:16+02:00', '2', '7011.47999', '0', 'CREATED', NULL, '2023-06-29T00:23:53+02:00', '2023-06-29T00:23:53+02:00'),
('zara-accountant', '3d4c1761-6272-40a9-b418-a0203dd1632c', 'contablock-contract-id', 'zara-user-accountant-id', '465', 'eaf6c055-23e6-4af2-80e6-62f4ad1842bf', '7aa974ff-0165-48ff-9e4a-b7bc3be40db5', 'EUR', 'nomina mes', '2023-06-02T14:22:16Z', '2', '0', '16435.75', 'CREATED', NULL, '2023-06-28T22:23:53Z', '2023-06-28T22:23:53Z'),
('zara-auditor', '3d4c1761-6272-40a9-b418-a0203dd1632c', 'contablock-contract-id', 'zara-user-accountant-id', '465', 'eaf6c055-23e6-4af2-80e6-62f4ad1842bf', '32dfa568-704f-4fb2-9a03-a1d3e45bbfcf', 'EUR', 'nomina mes', '2023-06-02T16:22:16+02:00', '2', '0', '16435.75', 'CREATED', NULL, '2023-06-29T00:23:53+02:00', '2023-06-29T00:23:53+02:00'),
('zara-viewer', '3d4c1761-6272-40a9-b418-a0203dd1632c', 'contablock-contract-id', 'zara-user-accountant-id', '465', 'eaf6c055-23e6-4af2-80e6-62f4ad1842bf', '48d4a8f8-0c54-4b00-b1e7-1dbd6fcf17e8', 'EUR', 'nomina mes', '2023-06-02T16:22:16+02:00', '2', '0', '16435.75', 'CREATED', NULL, '2023-06-29T00:23:53+02:00', '2023-06-29T00:23:53+02:00'),
('zara-accountant', '50d5cb7e-063d-4cbe-b475-2413fca30c82', 'contablock-contract-id', 'zara-user-accountant-id', '640', 'eaf6c055-23e6-4af2-80e6-62f4ad1842bf', '05d044b2-cc22-4922-8c8d-abf98c59497d', 'EUR', 'nomina mes', '2023-06-02T14:22:16Z', '2', '20246.96', '0', 'CREATED', NULL, '2023-06-28T22:23:53Z', '2023-06-28T22:23:53Z'),
('zara-auditor', '50d5cb7e-063d-4cbe-b475-2413fca30c82', 'contablock-contract-id', 'zara-user-accountant-id', '640', 'eaf6c055-23e6-4af2-80e6-62f4ad1842bf', '75436dc9-f1be-4964-87e4-042a47d8e46d', 'EUR', 'nomina mes', '2023-06-02T16:22:16+02:00', '2', '20246.95999', '0', 'CREATED', NULL, '2023-06-29T00:23:53+02:00', '2023-06-29T00:23:53+02:00'),
('zara-viewer', '50d5cb7e-063d-4cbe-b475-2413fca30c82', 'contablock-contract-id', 'zara-user-accountant-id', '640', 'eaf6c055-23e6-4af2-80e6-62f4ad1842bf', 'b953c464-615c-4882-9ecc-8567e168e34f', 'EUR', 'nomina mes', '2023-06-02T16:22:16+02:00', '2', '20246.95999', '0', 'CREATED', NULL, '2023-06-29T00:23:53+02:00', '2023-06-29T00:23:53+02:00'),
('zara-accountant', '678b5334-5315-4abd-8172-9228d0f6af6f', 'contablock-contract-id', 'zara-user-accountant-id', '472', '02e80740-0932-44a1-bc80-e5bb36261412', '5d8e08f2-5cec-4e72-9b33-8c4e6158f3df', 'EUR', 'compras', '2023-04-24T14:26:41Z', '2', '420', '0', 'CREATED', NULL, '2023-06-28T22:27:40Z', '2023-06-28T22:27:40Z'),
('zara-auditor', '678b5334-5315-4abd-8172-9228d0f6af6f', 'contablock-contract-id', 'zara-user-accountant-id', '472', '02e80740-0932-44a1-bc80-e5bb36261412', '0616c1ab-a3cb-42f4-b7fc-67cdbebc9fcb', 'EUR', 'compras', '2023-04-24T16:26:41+02:00', '2', '420', '0', 'CREATED', NULL, '2023-06-29T00:27:40+02:00', '2023-06-29T00:27:40+02:00'),
('zara-viewer', '678b5334-5315-4abd-8172-9228d0f6af6f', 'contablock-contract-id', 'zara-user-accountant-id', '472', '02e80740-0932-44a1-bc80-e5bb36261412', '0fe46140-2468-45e0-95aa-9938366d4a80', 'EUR', 'compras', '2023-04-24T16:26:41+02:00', '2', '420', '0', 'CREATED', NULL, '2023-06-29T00:27:40+02:00', '2023-06-29T00:27:40+02:00'),
('zara-accountant', '6bd50d09-463c-45e4-b3d5-8b00de9f526b', 'contablock-contract-id', 'zara-user-accountant-id', '476', 'eaf6c055-23e6-4af2-80e6-62f4ad1842bf', '2731be97-9ce7-4cd2-ae7e-992a153bd967', 'EUR', 'nomina mes', '2023-06-02T14:22:16Z', '2', '0', '8291.84', 'CREATED', NULL, '2023-06-28T22:23:53Z', '2023-06-28T22:23:53Z'),
('zara-auditor', '6bd50d09-463c-45e4-b3d5-8b00de9f526b', 'contablock-contract-id', 'zara-user-accountant-id', '476', 'eaf6c055-23e6-4af2-80e6-62f4ad1842bf', '94802184-9cdd-4983-a569-48a5b453d072', 'EUR', 'nomina mes', '2023-06-02T16:22:16+02:00', '2', '0', '8291.84', 'CREATED', NULL, '2023-06-29T00:23:53+02:00', '2023-06-29T00:23:53+02:00'),
('zara-viewer', '6bd50d09-463c-45e4-b3d5-8b00de9f526b', 'contablock-contract-id', 'zara-user-accountant-id', '476', 'eaf6c055-23e6-4af2-80e6-62f4ad1842bf', '881a9ae4-c6ae-4100-9c45-9c33b3c03b58', 'EUR', 'nomina mes', '2023-06-02T16:22:16+02:00', '2', '0', '8291.84', 'CREATED', NULL, '2023-06-29T00:23:53+02:00', '2023-06-29T00:23:53+02:00'),
('zara-accountant', '893ef9c1-3a97-4b51-868a-ebdfe0f046d3', 'contablock-contract-id', 'zara-user-accountant-id', '700', '00732959-fe23-4afd-92b4-de0b15cea192', '1561e0a7-746f-45f9-ac69-62509f459353', 'EUR', 'venta de mercaderías', '2023-03-23T09:24:58Z', '1', '0', '1000', 'CREATED', NULL, '2023-06-28T22:25:45Z', '2023-06-28T22:25:45Z'),
('zara-auditor', '893ef9c1-3a97-4b51-868a-ebdfe0f046d3', 'contablock-contract-id', 'zara-user-accountant-id', '700', '00732959-fe23-4afd-92b4-de0b15cea192', '88376854-9fa7-4219-b97f-5e23c4da3ac6', 'EUR', 'venta de mercaderías', '2023-03-23T10:24:58+01:00', '1', '0', '1000', 'CREATED', NULL, '2023-06-29T00:25:45+02:00', '2023-06-29T00:25:45+02:00'),
('zara-viewer', '893ef9c1-3a97-4b51-868a-ebdfe0f046d3', 'contablock-contract-id', 'zara-user-accountant-id', '700', '00732959-fe23-4afd-92b4-de0b15cea192', '35011ad0-9a36-4745-990c-29133cecbe5c', 'EUR', 'venta de mercaderías', '2023-03-23T10:24:58+01:00', '1', '0', '1000', 'CREATED', NULL, '2023-06-29T00:25:45+02:00', '2023-06-29T00:25:45+02:00'),
('zara-accountant', 'b29d93f2-77db-4f48-a570-2be8d0eb9c85', 'contablock-contract-id', 'zara-user-accountant-id', '400', '02e80740-0932-44a1-bc80-e5bb36261412', 'ab4ca960-d1de-4f05-b835-fb9d47f87d66', 'EUR', 'compras', '2023-04-24T14:26:41Z', '2', '0', '2420', 'CREATED', NULL, '2023-06-28T22:27:40Z', '2023-06-28T22:27:40Z'),
('zara-auditor', 'b29d93f2-77db-4f48-a570-2be8d0eb9c85', 'contablock-contract-id', 'zara-user-accountant-id', '400', '02e80740-0932-44a1-bc80-e5bb36261412', '85b75373-4cca-4a7b-93e4-1c6ce0ad3c84', 'EUR', 'compras', '2023-04-24T16:26:41+02:00', '2', '0', '2420', 'CREATED', NULL, '2023-06-29T00:27:41+02:00', '2023-06-29T00:27:41+02:00'),
('zara-viewer', 'b29d93f2-77db-4f48-a570-2be8d0eb9c85', 'contablock-contract-id', 'zara-user-accountant-id', '400', '02e80740-0932-44a1-bc80-e5bb36261412', 'b7f264b8-8f6e-47d3-8fe5-ff4d855acd76', 'EUR', 'compras', '2023-04-24T16:26:41+02:00', '2', '0', '2420', 'CREATED', NULL, '2023-06-29T00:27:41+02:00', '2023-06-29T00:27:41+02:00'),
('zara-accountant', 'b6f1594f-f208-4fe7-a90f-e804fabb4d62', 'contablock-contract-id', 'zara-user-accountant-id', '430', '00732959-fe23-4afd-92b4-de0b15cea192', '938d54e6-f4e6-4e50-a5d9-cd3105ea2e0a', 'EUR', 'venta de mercaderías', '2023-03-23T09:24:58Z', '1', '1210', '0', 'CREATED', NULL, '2023-06-28T22:25:45Z', '2023-06-28T22:25:45Z'),
('zara-auditor', 'b6f1594f-f208-4fe7-a90f-e804fabb4d62', 'contablock-contract-id', 'zara-user-accountant-id', '430', '00732959-fe23-4afd-92b4-de0b15cea192', 'e1d68abd-9e7b-4131-9fbe-55005aa75994', 'EUR', 'venta de mercaderías', '2023-03-23T10:24:58+01:00', '1', '1210', '0', 'CREATED', NULL, '2023-06-29T00:25:45+02:00', '2023-06-29T00:25:45+02:00'),
('zara-viewer', 'b6f1594f-f208-4fe7-a90f-e804fabb4d62', 'contablock-contract-id', 'zara-user-accountant-id', '430', '00732959-fe23-4afd-92b4-de0b15cea192', '7ce594d3-e1e7-488f-8367-d583d4ad7b0c', 'EUR', 'venta de mercaderías', '2023-03-23T10:24:58+01:00', '1', '1210', '0', 'CREATED', NULL, '2023-06-29T00:25:45+02:00', '2023-06-29T00:25:45+02:00'),
('zara-accountant', 'd920c863-36d4-41f0-9121-4845b8ca485d', 'contablock-contract-id', 'zara-user-accountant-id', '600', '02e80740-0932-44a1-bc80-e5bb36261412', 'a74156db-af75-42a1-9c0d-d4270665bc2a', 'EUR', 'compras', '2023-04-24T14:26:41Z', '2', '2000', '0', 'CREATED', NULL, '2023-06-28T22:27:40Z', '2023-06-28T22:27:40Z'),
('zara-auditor', 'd920c863-36d4-41f0-9121-4845b8ca485d', 'contablock-contract-id', 'zara-user-accountant-id', '600', '02e80740-0932-44a1-bc80-e5bb36261412', 'e150989d-84e0-4c3a-958b-a827a4181098', 'EUR', 'compras', '2023-04-24T16:26:41+02:00', '2', '2000', '0', 'CREATED', NULL, '2023-06-29T00:27:41+02:00', '2023-06-29T00:27:41+02:00'),
('zara-viewer', 'd920c863-36d4-41f0-9121-4845b8ca485d', 'contablock-contract-id', 'zara-user-accountant-id', '600', '02e80740-0932-44a1-bc80-e5bb36261412', 'fd052e6e-d4c8-4053-a80f-35058b7c3e6c', 'EUR', 'compras', '2023-04-24T16:26:41+02:00', '2', '2000', '0', 'CREATED', NULL, '2023-06-29T00:27:41+02:00', '2023-06-29T00:27:41+02:00'),
('zara-accountant', 'f3e57fcc-8c22-4945-8dcd-1988a0362fed', 'contablock-contract-id', 'zara-user-accountant-id', '4751', 'eaf6c055-23e6-4af2-80e6-62f4ad1842bf', '50d3d507-b4b6-4d9f-b669-54acb75d150b', 'EUR', 'nomina mes', '2023-06-02T14:22:16Z', '2', '0', '2530.85', 'CREATED', NULL, '2023-06-28T22:23:53Z', '2023-06-28T22:23:53Z'),
('zara-auditor', 'f3e57fcc-8c22-4945-8dcd-1988a0362fed', 'contablock-contract-id', 'zara-user-accountant-id', '4751', 'eaf6c055-23e6-4af2-80e6-62f4ad1842bf', '4aefa7fe-eefe-4058-928a-99dadd3c000c', 'EUR', 'nomina mes', '2023-06-02T16:22:16+02:00', '2', '0', '2530.84999', 'CREATED', NULL, '2023-06-29T00:23:53+02:00', '2023-06-29T00:23:53+02:00'),
('zara-viewer', 'f3e57fcc-8c22-4945-8dcd-1988a0362fed', 'contablock-contract-id', 'zara-user-accountant-id', '4751', 'eaf6c055-23e6-4af2-80e6-62f4ad1842bf', 'fccb91ab-da86-4c8a-80dd-c68959260da2', 'EUR', 'nomina mes', '2023-06-02T16:22:16+02:00', '2', '0', '2530.84999', 'CREATED', NULL, '2023-06-29T00:23:53+02:00', '2023-06-29T00:23:53+02:00'),
('zara-accountant', 'f514eba0-a0c3-4ac9-ad49-6c001679f8d9', 'contablock-contract-id', 'zara-user-accountant-id', '477', '00732959-fe23-4afd-92b4-de0b15cea192', '96ad57e0-aaad-4eab-8252-376657d4d498', 'EUR', 'venta de mercaderías', '2023-03-23T09:24:58Z', '1', '0', '210', 'CREATED', NULL, '2023-06-28T22:25:45Z', '2023-06-28T22:25:45Z'),
('zara-auditor', 'f514eba0-a0c3-4ac9-ad49-6c001679f8d9', 'contablock-contract-id', 'zara-user-accountant-id', '477', '00732959-fe23-4afd-92b4-de0b15cea192', 'eba2b551-a580-44be-be6d-1e7105675259', 'EUR', 'venta de mercaderías', '2023-03-23T10:24:58+01:00', '1', '0', '210', 'CREATED', NULL, '2023-06-29T00:25:45+02:00', '2023-06-29T00:25:45+02:00'),
('zara-viewer', 'f514eba0-a0c3-4ac9-ad49-6c001679f8d9', 'contablock-contract-id', 'zara-user-accountant-id', '477', '00732959-fe23-4afd-92b4-de0b15cea192', 'b2cbc384-402e-4197-8d94-a97313d63048', 'EUR', 'venta de mercaderías', '2023-03-23T10:24:58+01:00', '1', '0', '210', 'CREATED', NULL, '2023-06-29T00:25:45+02:00', '2023-06-29T00:25:45+02:00');

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
