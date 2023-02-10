-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Feb 10, 2023 at 01:51 PM
-- Server version: 10.4.27-MariaDB
-- PHP Version: 8.2.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `golang_rest_example`
--

-- --------------------------------------------------------

--
-- Table structure for table `armament`
--

CREATE TABLE `armament` (
  `id` int(11) NOT NULL,
  `spaceship_id` int(11) NOT NULL,
  `title` varchar(100) NOT NULL,
  `qty` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `armament`
--

INSERT INTO `armament` (`id`, `spaceship_id`, `title`, `qty`) VALUES
(1, 1, 'Turbo Laser', 60),
(2, 1, 'Ion Cannons', 60),
(3, 1, 'Tractor Beam', 10),
(4, 2, 'Turbo Laser 2', 60),
(5, 2, 'Ion Cannons 2', 60),
(6, 2, 'Tractor Beam 2', 10);

-- --------------------------------------------------------

--
-- Table structure for table `spaceships`
--

CREATE TABLE `spaceships` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `class` varchar(100) DEFAULT NULL,
  `crew` int(11) DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL,
  `value` decimal(10,2) DEFAULT NULL,
  `status` enum('operational','damaged','under maintenance') NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `spaceships`
--

INSERT INTO `spaceships` (`id`, `name`, `class`, `crew`, `image`, `value`, `status`) VALUES
(1, 'Devastator', 'Star Destroyer', 35000, 'https:\\url.to.image//1', '1999.99', 'operational'),
(2, 'Devastator 2', 'Star Destroyer 2', 35000, 'https:\\url.to.image/2', '2258.99', 'operational');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `password`) VALUES
(1, 'example_user', 'example_password');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `armament`
--
ALTER TABLE `armament`
  ADD PRIMARY KEY (`id`),
  ADD KEY `spaceship_id` (`spaceship_id`);

--
-- Indexes for table `spaceships`
--
ALTER TABLE `spaceships`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `armament`
--
ALTER TABLE `armament`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `spaceships`
--
ALTER TABLE `spaceships`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `armament`
--
ALTER TABLE `armament`
  ADD CONSTRAINT `armament_ibfk_1` FOREIGN KEY (`spaceship_id`) REFERENCES `spaceships` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
