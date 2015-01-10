-- phpMyAdmin SQL Dump
-- version 4.4.0-dev
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: Jan 10, 2015 at 08:20 AM
-- Server version: 5.6.19
-- PHP Version: 5.4.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Database: `dalalstreet`
--

-- --------------------------------------------------------

--
-- Table structure for table `bank`
--

CREATE TABLE IF NOT EXISTS `bank` (
  `mortgageId` int(11) NOT NULL,
  `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `stockId` int(11) NOT NULL,
  `number` int(11) NOT NULL,
  `loanValue` double NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `buying_table`
--

CREATE TABLE IF NOT EXISTS `buying_table` (
  `buyId` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `stockId` int(11) NOT NULL,
  `num` int(11) NOT NULL,
  `value` double NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `selling_table`
--

CREATE TABLE IF NOT EXISTS `selling_table` (
  `sellId` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `stockId` int(11) NOT NULL,
  `num` int(11) NOT NULL,
  `value` double NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `stocks`
--

CREATE TABLE IF NOT EXISTS `stocks` (
  `stockId` int(11) NOT NULL,
  `name` varchar(200) NOT NULL,
  `marketValue` double NOT NULL,
  `exchangePrice` double NOT NULL,
  `lastTrade` double NOT NULL,
  `dayLow` double NOT NULL,
  `dayHigh` double NOT NULL,
  `numIssued` int(11) NOT NULL,
  `sharesInExchange` int(11) NOT NULL,
  `factor` float NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE IF NOT EXISTS `users` (
  `username` varchar(50) COLLATE utf8_bin NOT NULL,
  `email` varchar(50) COLLATE utf8_bin NOT NULL,
  `password` varchar(100) COLLATE utf8_bin NOT NULL,
  `name` varchar(50) COLLATE utf8_bin DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`username`, `email`, `password`, `name`) VALUES
('boopathi', 'me@boopathi.in', '$2a$10$q58PuPIqkOv.7ULslzfVhuIUsxPepu55vdRMmZc.g2aePyb83zXMq', 'boopathi'),
('nobelium', 'me@vignesh.info', '$2a$10$q58PuPIqkOv.7ULslzfVhuIUsxPepu55vdRMmZc.g2aePyb83zXMq', 'vignesh');

-- --------------------------------------------------------

--
-- Table structure for table `users_data`
--

CREATE TABLE IF NOT EXISTS `users_data` (
  `username` varchar(50) NOT NULL,
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `key` varchar(100) NOT NULL,
  `value` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `bank`
--
ALTER TABLE `bank`
  ADD PRIMARY KEY (`mortgageId`);

--
-- Indexes for table `buying_table`
--
ALTER TABLE `buying_table`
  ADD PRIMARY KEY (`buyId`);

--
-- Indexes for table `selling_table`
--
ALTER TABLE `selling_table`
  ADD PRIMARY KEY (`sellId`);

--
-- Indexes for table `stocks`
--
ALTER TABLE `stocks`
  ADD PRIMARY KEY (`stockId`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`username`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `bank`
--
ALTER TABLE `bank`
  MODIFY `mortgageId` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=3;
--
-- AUTO_INCREMENT for table `buying_table`
--
ALTER TABLE `buying_table`
  MODIFY `buyId` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `selling_table`
--
ALTER TABLE `selling_table`
  MODIFY `sellId` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=6;
--
-- AUTO_INCREMENT for table `stocks`
--
ALTER TABLE `stocks`
  MODIFY `stockId` int(11) NOT NULL AUTO_INCREMENT;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
