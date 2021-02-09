-- phpMyAdmin SQL Dump
-- version 4.7.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Feb 09, 2021 at 04:41 AM
-- Server version: 10.1.30-MariaDB
-- PHP Version: 5.6.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bibit`
--

-- --------------------------------------------------------

--
-- Table structure for table `search_log`
--

CREATE TABLE `search_log` (
  `uuid` varchar(36) NOT NULL,
  `search_key` varchar(200) NOT NULL,
  `page` int(11) NOT NULL,
  `created_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `search_log`
--

INSERT INTO `search_log` (`uuid`, `search_key`, `page`, `created_at`) VALUES
('1065e0f1-de58-43f7-ac61-85696fe0a6d6', 'spiderman', 1, '2021-02-07 23:51:54'),
('4b53bc05-a028-494a-9033-33ea3ca2894f', 'spiderman', 1, '2021-02-07 23:46:27'),
('56f5d600-da43-43c4-899d-3e8a22027430', 'spiderman', 1, '2021-02-07 23:52:06'),
('8fc8ea79-5596-4abc-9e7f-4b6457919c01', 'spiderman', 1, '2021-02-07 23:30:30');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `search_log`
--
ALTER TABLE `search_log`
  ADD PRIMARY KEY (`uuid`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
