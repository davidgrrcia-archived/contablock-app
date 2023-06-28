-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: db:3306
-- Generation Time: Jun 28, 2023 at 10:29 PM
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
('zara-auditor', '05841c70-7ef3-46fa-8b57-dfff5b08171a', '10057', '-7011.47999', '2023'),
('zara-accountant', '05d044b2-cc22-4922-8c8d-abf98c59497d', '10057', '-20246.96', '2023'),
('zara-auditor', '0616c1ab-a3cb-42f4-b7fc-67cdbebc9fcb', '10034', '420', '2023'),
('zara-viewer', '0fe46140-2468-45e0-95aa-9938366d4a80', '10034', '420', '2023'),
('zara-accountant', '1561e0a7-746f-45f9-ac69-62509f459353', '10057', '1000', '2023'),
('zara-accountant', '2731be97-9ce7-4cd2-ae7e-992a153bd967', '10092', '8291.84', '2023'),
('zara-accountant', '286a008a-18f1-4a8e-99f0-2053aee78154', '10057', '-7011.48', '2023'),
('zara-auditor', '32dfa568-704f-4fb2-9a03-a1d3e45bbfcf', '10090', '16435.75', '2023'),
('zara-viewer', '35011ad0-9a36-4745-990c-29133cecbe5c', '10057', '1000', '2023'),
('zara-viewer', '48d4a8f8-0c54-4b00-b1e7-1dbd6fcf17e8', '10090', '16435.75', '2023'),
('zara-auditor', '4aefa7fe-eefe-4058-928a-99dadd3c000c', '10092', '2530.84999', '2023'),
('zara-accountant', '50d3d507-b4b6-4d9f-b669-54acb75d150b', '10092', '2530.85', '2023'),
('zara-accountant', '5d8e08f2-5cec-4e72-9b33-8c4e6158f3df', '10034', '420', '2023'),
('zara-auditor', '75436dc9-f1be-4964-87e4-042a47d8e46d', '10057', '-20246.95999', '2023'),
('zara-accountant', '7aa974ff-0165-48ff-9e4a-b7bc3be40db5', '10090', '16435.75', '2023'),
('zara-viewer', '7b4c3794-341e-4ac2-baf2-cd51ba055b55', '10057', '-7011.47999', '2023'),
('zara-viewer', '7ce594d3-e1e7-488f-8367-d583d4ad7b0c', '10029', '1210', '2023'),
('zara-auditor', '85b75373-4cca-4a7b-93e4-1c6ce0ad3c84', '10087', '2420', '2023'),
('zara-viewer', '881a9ae4-c6ae-4100-9c45-9c33b3c03b58', '10092', '8291.84', '2023'),
('zara-auditor', '88376854-9fa7-4219-b97f-5e23c4da3ac6', '10057', '1000', '2023'),
('zara-accountant', '938d54e6-f4e6-4e50-a5d9-cd3105ea2e0a', '10029', '1210', '2023'),
('zara-auditor', '94802184-9cdd-4983-a569-48a5b453d072', '10092', '8291.84', '2023'),
('zara-accountant', '96ad57e0-aaad-4eab-8252-376657d4d498', '10092', '210', '2023'),
('zara-accountant', 'a74156db-af75-42a1-9c0d-d4270665bc2a', '10057', '-2000', '2023'),
('zara-accountant', 'ab4ca960-d1de-4f05-b835-fb9d47f87d66', '10087', '2420', '2023'),
('zara-viewer', 'b2cbc384-402e-4197-8d94-a97313d63048', '10092', '210', '2023'),
('zara-viewer', 'b7f264b8-8f6e-47d3-8fe5-ff4d855acd76', '10087', '2420', '2023'),
('zara-viewer', 'b953c464-615c-4882-9ecc-8567e168e34f', '10057', '-20246.95999', '2023'),
('zara-auditor', 'e150989d-84e0-4c3a-958b-a827a4181098', '10057', '-2000', '2023'),
('zara-auditor', 'e1d68abd-9e7b-4131-9fbe-55005aa75994', '10029', '1210', '2023'),
('zara-auditor', 'eba2b551-a580-44be-be6d-1e7105675259', '10092', '210', '2023'),
('zara-viewer', 'fccb91ab-da86-4c8a-80dd-c68959260da2', '10092', '2530.84999', '2023'),
('zara-viewer', 'fd052e6e-d4c8-4053-a80f-35058b7c3e6c', '10057', '-2000', '2023');

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
