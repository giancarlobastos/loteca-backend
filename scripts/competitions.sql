-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: mysql
-- Generation Time: Dec 10, 2021 at 09:23 AM
-- Server version: 8.0.21
-- PHP Version: 7.4.20

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `loteca`
--

-- --------------------------------------------------------

--
-- Table structure for table `competition`
--

-- CREATE TABLE `competition` (
--   `id` int NOT NULL,
--   `name` varchar(255) DEFAULT NULL,
--   `division` varchar(255) DEFAULT NULL,
--   `logo` varchar(255) DEFAULT NULL,
--   `type` varchar(255) DEFAULT NULL,
--   `country` varchar(255) DEFAULT NULL
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `competition`
--

INSERT INTO `competition` (`id`, `name`, `division`, `logo`, `type`, `country`) VALUES
(0, '', NULL, '', '', ''),
(39, 'Premier League', NULL, 'https://media.api-sports.io/football/leagues/39.png', 'League', 'England'),
(40, 'Championship', NULL, 'https://media.api-sports.io/football/leagues/40.png', 'League', 'England'),
(41, 'League One', NULL, 'https://media.api-sports.io/football/leagues/41.png', 'League', 'England'),
(42, 'League Two', NULL, 'https://media.api-sports.io/football/leagues/42.png', 'League', 'England'),
(43, 'National League', NULL, 'https://media.api-sports.io/football/leagues/43.png', 'League', 'England'),
(44, 'FA WSL', NULL, 'https://media.api-sports.io/football/leagues/44.png', 'League', 'England'),
(45, 'FA Cup', NULL, 'https://media.api-sports.io/football/leagues/45.png', 'Cup', 'England'),
(46, 'EFL Trophy', NULL, 'https://media.api-sports.io/football/leagues/46.png', 'Cup', 'England'),
(47, 'FA Trophy', NULL, 'https://media.api-sports.io/football/leagues/47.png', 'Cup', 'England'),
(48, 'League Cup', NULL, 'https://media.api-sports.io/football/leagues/48.png', 'Cup', 'England'),
(49, 'National League - Play-offs', NULL, 'https://media.api-sports.io/football/leagues/49.png', 'League', 'England'),
(50, 'National League - North', NULL, 'https://media.api-sports.io/football/leagues/50.png', 'League', 'England'),
(51, 'National League - South', NULL, 'https://media.api-sports.io/football/leagues/51.png', 'League', 'England'),
(52, 'Non League Div One - Isthmian North', NULL, 'https://media.api-sports.io/football/leagues/52.png', 'League', 'England'),
(53, 'Non League Div One - Isthmian South', NULL, 'https://media.api-sports.io/football/leagues/53.png', 'League', 'England'),
(54, 'Non League Div One - Northern North', NULL, 'https://media.api-sports.io/football/leagues/54.png', 'League', 'England'),
(55, 'Non League Div One - Northern South', NULL, 'https://media.api-sports.io/football/leagues/55.png', 'League', 'England'),
(56, 'Non League Div One - Southern Central', NULL, 'https://media.api-sports.io/football/leagues/56.png', 'League', 'England'),
(57, 'Non League Div One - Southern SW', NULL, 'https://media.api-sports.io/football/leagues/57.png', 'League', 'England'),
(58, 'Non League Premier - Isthmian', NULL, 'https://media.api-sports.io/football/leagues/58.png', 'League', 'England'),
(59, 'Non League Premier - Northern', NULL, 'https://media.api-sports.io/football/leagues/59.png', 'League', 'England'),
(60, 'Non League Premier - Southern', NULL, 'https://media.api-sports.io/football/leagues/60.png', 'League', 'England'),
(61, 'Ligue 1', NULL, 'https://media.api-sports.io/football/leagues/61.png', 'League', 'France'),
(62, 'Ligue 2', NULL, 'https://media.api-sports.io/football/leagues/62.png', 'League', 'France'),
(63, 'National', NULL, 'https://media.api-sports.io/football/leagues/63.png', 'League', 'France'),
(64, 'Feminine Division 1', NULL, 'https://media.api-sports.io/football/leagues/64.png', 'League', 'France'),
(65, 'Coupe de la Ligue', NULL, 'https://media.api-sports.io/football/leagues/65.png', 'Cup', 'France'),
(66, 'Coupe de France', NULL, 'https://media.api-sports.io/football/leagues/66.png', 'Cup', 'France'),
(67, 'National 2 - Group A', NULL, 'https://media.api-sports.io/football/leagues/67.png', 'League', 'France'),
(68, 'National 2 - Group B', NULL, 'https://media.api-sports.io/football/leagues/68.png', 'League', 'France'),
(69, 'National 2 - Group C', NULL, 'https://media.api-sports.io/football/leagues/69.png', 'League', 'France'),
(70, 'National 2 - Group D', NULL, 'https://media.api-sports.io/football/leagues/70.png', 'League', 'France'),
(71, 'Serie A', NULL, 'https://media.api-sports.io/football/leagues/71.png', 'League', 'Brazil'),
(72, 'Serie B', NULL, 'https://media.api-sports.io/football/leagues/72.png', 'League', 'Brazil'),
(73, 'Copa Do Brasil', NULL, 'https://media.api-sports.io/football/leagues/73.png', 'Cup', 'Brazil'),
(74, 'Brasileiro Women', NULL, 'https://media.api-sports.io/football/leagues/74.png', 'League', 'Brazil'),
(75, 'Serie C', NULL, 'https://media.api-sports.io/football/leagues/75.png', 'League', 'Brazil'),
(76, 'Serie D', NULL, 'https://media.api-sports.io/football/leagues/76.png', 'League', 'Brazil'),
(77, 'Alagoano', NULL, 'https://media.api-sports.io/football/leagues/77.png', 'Cup', 'Brazil'),
(78, 'Bundesliga 1', NULL, 'https://media.api-sports.io/football/leagues/78.png', 'League', 'Germany'),
(79, 'Bundesliga 2', NULL, 'https://media.api-sports.io/football/leagues/79.png', 'League', 'Germany'),
(80, 'Liga 3', NULL, 'https://media.api-sports.io/football/leagues/80.png', 'League', 'Germany'),
(81, 'DFB Pokal', NULL, 'https://media.api-sports.io/football/leagues/81.png', 'Cup', 'Germany'),
(82, 'Women Bundesliga', NULL, 'https://media.api-sports.io/football/leagues/82.png', 'League', 'Germany'),
(83, 'Regionalliga - Bayern', NULL, 'https://media.api-sports.io/football/leagues/83.png', 'League', 'Germany'),
(84, 'Regionalliga - Nord', NULL, 'https://media.api-sports.io/football/leagues/84.png', 'League', 'Germany'),
(85, 'Regionalliga - Nordost', NULL, 'https://media.api-sports.io/football/leagues/85.png', 'League', 'Germany'),
(86, 'Regionalliga - SudWest', NULL, 'https://media.api-sports.io/football/leagues/86.png', 'League', 'Germany'),
(87, 'Regionalliga - West', NULL, 'https://media.api-sports.io/football/leagues/87.png', 'League', 'Germany'),
(128, 'Primera Division', NULL, 'https://media.api-sports.io/football/leagues/128.png', 'League', 'Argentina'),
(129, 'Primera B Nacional', NULL, 'https://media.api-sports.io/football/leagues/129.png', 'League', 'Argentina'),
(130, 'Copa Argentina', NULL, 'https://media.api-sports.io/football/leagues/130.png', 'Cup', 'Argentina'),
(131, 'Primera B Metropolitana', NULL, 'https://media.api-sports.io/football/leagues/131.png', 'League', 'Argentina'),
(132, 'Primera C', NULL, 'https://media.api-sports.io/football/leagues/132.png', 'League', 'Argentina'),
(133, 'Primera D', NULL, 'https://media.api-sports.io/football/leagues/133.png', 'League', 'Argentina'),
(134, 'Torneo Federal A', NULL, 'https://media.api-sports.io/football/leagues/134.png', 'League', 'Argentina'),
(135, 'Serie A', NULL, 'https://media.api-sports.io/football/leagues/135.png', 'League', 'Italy'),
(136, 'Serie B', NULL, 'https://media.api-sports.io/football/leagues/136.png', 'League', 'Italy'),
(137, 'Coppa Italia', NULL, 'https://media.api-sports.io/football/leagues/137.png', 'Cup', 'Italy'),
(138, 'Serie C', NULL, 'https://media.api-sports.io/football/leagues/138.png', 'League', 'Italy'),
(139, 'Serie A Women', NULL, 'https://media.api-sports.io/football/leagues/139.png', 'League', 'Italy'),
(140, 'La Liga', NULL, 'https://media.api-sports.io/football/leagues/140.png', 'League', 'Spain'),
(141, 'Segunda Division', NULL, 'https://media.api-sports.io/football/leagues/141.png', 'League', 'Spain'),
(142, 'Primera Division Women', NULL, 'https://media.api-sports.io/football/leagues/142.png', 'League', 'Spain'),
(143, 'Copa del Rey', NULL, 'https://media.api-sports.io/football/leagues/143.png', 'Cup', 'Spain'),
(426, 'Serie D - Girone A', NULL, 'https://media.api-sports.io/football/leagues/426.png', 'League', 'Italy'),
(427, 'Serie D - Girone B', NULL, 'https://media.api-sports.io/football/leagues/427.png', 'League', 'Italy'),
(428, 'Serie D - Girone C', NULL, 'https://media.api-sports.io/football/leagues/428.png', 'League', 'Italy'),
(429, 'Serie D - Girone D', NULL, 'https://media.api-sports.io/football/leagues/429.png', 'League', 'Italy'),
(430, 'Serie D - Girone E', NULL, 'https://media.api-sports.io/football/leagues/430.png', 'League', 'Italy'),
(431, 'Serie D - Girone F', NULL, 'https://media.api-sports.io/football/leagues/431.png', 'League', 'Italy'),
(432, 'Serie D - Girone G', NULL, 'https://media.api-sports.io/football/leagues/432.png', 'League', 'Italy'),
(433, 'Serie D - Girone H', NULL, 'https://media.api-sports.io/football/leagues/433.png', 'League', 'Italy'),
(434, 'Serie D - Girone I', NULL, 'https://media.api-sports.io/football/leagues/434.png', 'League', 'Italy'),
(435, 'Primera División RFEF - Group 1', NULL, 'https://media.api-sports.io/football/leagues/435.png', 'League', 'Spain'),
(436, 'Primera División RFEF - Group 2', NULL, 'https://media.api-sports.io/football/leagues/436.png', 'League', 'Spain'),
(437, 'Primera División RFEF - Group 3', NULL, 'https://media.api-sports.io/football/leagues/437.png', 'League', 'Spain'),
(438, 'Primera División RFEF - Group 4', NULL, 'https://media.api-sports.io/football/leagues/438.png', 'League', 'Spain'),
(439, 'Tercera División RFEF - Group 1', NULL, 'https://media.api-sports.io/football/leagues/439.png', 'League', 'Spain'),
(440, 'Tercera División RFEF - Group 2', NULL, 'https://media.api-sports.io/football/leagues/440.png', 'League', 'Spain'),
(441, 'Tercera División RFEF - Group 3', NULL, 'https://media.api-sports.io/football/leagues/441.png', 'League', 'Spain'),
(442, 'Tercera División RFEF - Group 4', NULL, 'https://media.api-sports.io/football/leagues/442.png', 'League', 'Spain'),
(443, 'Tercera División RFEF - Group 5', NULL, 'https://media.api-sports.io/football/leagues/443.png', 'League', 'Spain'),
(444, 'Tercera División RFEF - Group 6', NULL, 'https://media.api-sports.io/football/leagues/444.png', 'League', 'Spain'),
(445, 'Tercera División RFEF - Group 7', NULL, 'https://media.api-sports.io/football/leagues/445.png', 'League', 'Spain'),
(446, 'Tercera División RFEF - Group 8', NULL, 'https://media.api-sports.io/football/leagues/446.png', 'League', 'Spain'),
(447, 'Tercera División RFEF - Group 9', NULL, 'https://media.api-sports.io/football/leagues/447.png', 'League', 'Spain'),
(448, 'Tercera División RFEF - Group 10', NULL, 'https://media.api-sports.io/football/leagues/448.png', 'League', 'Spain'),
(449, 'Tercera División RFEF - Group 11', NULL, 'https://media.api-sports.io/football/leagues/449.png', 'League', 'Spain'),
(450, 'Tercera División RFEF - Group 12', NULL, 'https://media.api-sports.io/football/leagues/450.png', 'League', 'Spain'),
(451, 'Tercera División RFEF - Group 13', NULL, 'https://media.api-sports.io/football/leagues/451.png', 'League', 'Spain'),
(452, 'Tercera División RFEF - Group 14', NULL, 'https://media.api-sports.io/football/leagues/452.png', 'League', 'Spain'),
(453, 'Tercera División RFEF - Group 15', NULL, 'https://media.api-sports.io/football/leagues/453.png', 'League', 'Spain'),
(454, 'Tercera División RFEF - Group 16', NULL, 'https://media.api-sports.io/football/leagues/454.png', 'League', 'Spain'),
(455, 'Tercera División RFEF - Group 17', NULL, 'https://media.api-sports.io/football/leagues/455.png', 'League', 'Spain'),
(456, 'Tercera División RFEF - Group 18', NULL, 'https://media.api-sports.io/football/leagues/456.png', 'League', 'Spain'),
(461, 'National 3 - Group A', NULL, 'https://media.api-sports.io/football/leagues/461.png', 'League', 'France'),
(462, 'National 3 - Group B', NULL, 'https://media.api-sports.io/football/leagues/462.png', 'League', 'France'),
(463, 'National 3 - Group C', NULL, 'https://media.api-sports.io/football/leagues/463.png', 'League', 'France'),
(464, 'National 3 - Group D', NULL, 'https://media.api-sports.io/football/leagues/464.png', 'League', 'France'),
(465, 'National 3 - Group E', NULL, 'https://media.api-sports.io/football/leagues/465.png', 'League', 'France'),
(466, 'National 3 - Group F', NULL, 'https://media.api-sports.io/football/leagues/466.png', 'League', 'France'),
(467, 'National 3 - Group H', NULL, 'https://media.api-sports.io/football/leagues/467.png', 'League', 'France'),
(468, 'National 3 - Group I', NULL, 'https://media.api-sports.io/football/leagues/468.png', 'League', 'France'),
(469, 'National 3 - Group J', NULL, 'https://media.api-sports.io/football/leagues/469.png', 'League', 'France'),
(470, 'National 3 - Group K', NULL, 'https://media.api-sports.io/football/leagues/470.png', 'League', 'France'),
(471, 'National 3 - Group L', NULL, 'https://media.api-sports.io/football/leagues/471.png', 'League', 'France'),
(472, 'National 3 - Group M', NULL, 'https://media.api-sports.io/football/leagues/472.png', 'League', 'France'),
(475, 'Paulista - A1', NULL, 'https://media.api-sports.io/football/leagues/475.png', 'League', 'Brazil'),
(476, 'Paulista - A2', NULL, 'https://media.api-sports.io/football/leagues/476.png', 'League', 'Brazil'),
(477, 'Gaúcho - 1', NULL, 'https://media.api-sports.io/football/leagues/477.png', 'League', 'Brazil'),
(478, 'Gaúcho - 2', NULL, 'https://media.api-sports.io/football/leagues/478.png', 'League', 'Brazil'),
(483, 'Copa de la Superliga', NULL, 'https://media.api-sports.io/football/leagues/483.png', 'Cup', 'Argentina'),
(488, 'U19 Bundesliga', NULL, 'https://media.api-sports.io/football/leagues/488.png', 'League', 'Germany'),
(517, 'Trofeo de Campeones de la Superliga', NULL, 'https://media.api-sports.io/football/leagues/517.png', 'Cup', 'Argentina'),
(520, 'Acreano', NULL, 'https://media.api-sports.io/football/leagues/520.png', 'League', 'Brazil'),
(521, 'Amapaense', NULL, 'https://media.api-sports.io/football/leagues/521.png', 'League', 'Brazil'),
(522, 'Amazonense', NULL, 'https://media.api-sports.io/football/leagues/522.png', 'League', 'Brazil'),
(526, 'Trophée des Champions', NULL, 'https://media.api-sports.io/football/leagues/526.png', 'Cup', 'France'),
(528, 'Community Shield', NULL, 'https://media.api-sports.io/football/leagues/528.png', 'Cup', 'England'),
(529, 'Super Cup', NULL, 'https://media.api-sports.io/football/leagues/529.png', 'Cup', 'Germany'),
(547, 'Super Cup', NULL, 'https://media.api-sports.io/football/leagues/547.png', 'Cup', 'Italy'),
(556, 'Super Cup', NULL, 'https://media.api-sports.io/football/leagues/556.png', 'Cup', 'Spain'),
(602, 'Baiano - 1', NULL, 'https://media.api-sports.io/football/leagues/602.png', 'League', 'Brazil'),
(603, 'Paraibano', NULL, 'https://media.api-sports.io/football/leagues/603.png', 'League', 'Brazil'),
(604, 'Catarinense - 1', NULL, 'https://media.api-sports.io/football/leagues/604.png', 'League', 'Brazil'),
(605, 'Paulista - A3', NULL, 'https://media.api-sports.io/football/leagues/605.png', 'League', 'Brazil'),
(606, 'Paranaense - 1', NULL, 'https://media.api-sports.io/football/leagues/606.png', 'League', 'Brazil'),
(607, 'Roraimense', NULL, 'https://media.api-sports.io/football/leagues/607.png', 'League', 'Brazil'),
(608, 'Maranhense', NULL, 'https://media.api-sports.io/football/leagues/608.png', 'League', 'Brazil'),
(609, 'Cearense - 1', NULL, 'https://media.api-sports.io/football/leagues/609.png', 'League', 'Brazil'),
(610, 'Brasiliense', NULL, 'https://media.api-sports.io/football/leagues/610.png', 'League', 'Brazil'),
(611, 'Capixaba', NULL, 'https://media.api-sports.io/football/leagues/611.png', 'League', 'Brazil'),
(612, 'Copa do Nordeste', NULL, 'https://media.api-sports.io/football/leagues/612.png', 'Cup', 'Brazil'),
(613, 'Baiano - 2', NULL, 'https://media.api-sports.io/football/leagues/613.png', 'League', 'Brazil'),
(614, 'Paranaense - 2', NULL, 'https://media.api-sports.io/football/leagues/614.png', 'League', 'Brazil'),
(615, 'Rondoniense', NULL, 'https://media.api-sports.io/football/leagues/615.png', 'League', 'Brazil'),
(616, 'Potiguar', NULL, 'https://media.api-sports.io/football/leagues/616.png', 'League', 'Brazil'),
(617, 'Copa do Brasil U20', NULL, 'https://media.api-sports.io/football/leagues/617.png', 'Cup', 'Brazil'),
(618, 'São Paulo Youth Cup', NULL, 'https://media.api-sports.io/football/leagues/618.png', 'Cup', 'Brazil'),
(619, 'Mineiro - 2', NULL, 'https://media.api-sports.io/football/leagues/619.png', 'League', 'Brazil'),
(620, 'Cearense - 2', NULL, 'https://media.api-sports.io/football/leagues/620.png', 'League', 'Brazil'),
(621, 'Piauiense', NULL, 'https://media.api-sports.io/football/leagues/621.png', 'League', 'Brazil'),
(622, 'Pernambucano - 1', NULL, 'https://media.api-sports.io/football/leagues/622.png', 'League', 'Brazil'),
(623, 'Sul-Matogrossense', NULL, 'https://media.api-sports.io/football/leagues/623.png', 'League', 'Brazil'),
(624, 'Carioca - 1', NULL, 'https://media.api-sports.io/football/leagues/624.png', 'League', 'Brazil'),
(625, 'Carioca - 2', NULL, 'https://media.api-sports.io/football/leagues/625.png', 'League', 'Brazil'),
(626, 'Sergipano', NULL, 'https://media.api-sports.io/football/leagues/626.png', 'League', 'Brazil'),
(627, 'Paraense', NULL, 'https://media.api-sports.io/football/leagues/627.png', 'League', 'Brazil'),
(628, 'Goiano - 1', NULL, 'https://media.api-sports.io/football/leagues/628.png', 'League', 'Brazil'),
(629, 'Mineiro - 1', NULL, 'https://media.api-sports.io/football/leagues/629.png', 'League', 'Brazil'),
(630, 'Matogrossense', NULL, 'https://media.api-sports.io/football/leagues/630.png', 'League', 'Brazil'),
(631, 'Tocantinense', NULL, 'https://media.api-sports.io/football/leagues/631.png', 'League', 'Brazil'),
(632, 'Supercopa do Brasil', NULL, 'https://media.api-sports.io/football/leagues/632.png', 'Cup', 'Brazil'),
(670, 'Community Shield Women', NULL, 'https://media.api-sports.io/football/leagues/670.png', 'Cup', 'England'),
(692, 'Primera División RFEF - Group 5', NULL, 'https://media.api-sports.io/football/leagues/692.png', 'League', 'Spain'),
(695, 'U18 Premier League - North', NULL, 'https://media.api-sports.io/football/leagues/695.png', 'League', 'England'),
(696, 'U18 Premier League - South', NULL, 'https://media.api-sports.io/football/leagues/696.png', 'League', 'England'),
(697, 'WSL Cup', NULL, 'https://media.api-sports.io/football/leagues/697.png', 'Cup', 'England'),
(698, 'FA Women\'s Cup', NULL, 'https://media.api-sports.io/football/leagues/698.png', 'Cup', 'England'),
(699, 'Women\'s Championship', NULL, 'https://media.api-sports.io/football/leagues/699.png', 'League', 'England'),
(702, 'Premier League 2 Division One', NULL, 'https://media.api-sports.io/football/leagues/702.png', 'League', 'England'),
(703, 'Professional Development League', NULL, 'https://media.api-sports.io/football/leagues/703.png', 'League', 'England'),
(704, 'Coppa Italia Primavera', NULL, 'https://media.api-sports.io/football/leagues/704.png', 'Cup', 'Italy'),
(705, 'Campionato Primavera - 1', NULL, 'https://media.api-sports.io/football/leagues/705.png', 'League', 'Italy'),
(706, 'Campionato Primavera - 2', NULL, 'https://media.api-sports.io/football/leagues/706.png', 'League', 'Italy'),
(715, 'DFB Junioren Pokal', NULL, 'https://media.api-sports.io/football/leagues/715.png', 'Cup', 'Germany'),
(735, 'Copa Federacion', NULL, 'https://media.api-sports.io/football/leagues/735.png', 'Cup', 'Spain'),
(740, 'CBF Brasileiro U20', NULL, 'https://media.api-sports.io/football/leagues/740.png', 'Cup', 'Brazil'),
(741, 'Brasileiro de Aspirantes', NULL, 'https://media.api-sports.io/football/leagues/741.png', 'Cup', 'Brazil'),
(742, 'Copa Paulista', NULL, 'https://media.api-sports.io/football/leagues/742.png', 'Cup', 'Brazil'),
(744, 'Oberliga - Schleswig-Holstein', NULL, 'https://media.api-sports.io/football/leagues/744.png', 'League', 'Germany'),
(745, 'Oberliga - Hamburg', NULL, 'https://media.api-sports.io/football/leagues/745.png', 'League', 'Germany'),
(746, 'Oberliga - Mittelrhein', NULL, 'https://media.api-sports.io/football/leagues/746.png', 'League', 'Germany'),
(747, 'Oberliga - Westfalen', NULL, 'https://media.api-sports.io/football/leagues/747.png', 'League', 'Germany'),
(748, 'Oberliga - Niedersachsen', NULL, 'https://media.api-sports.io/football/leagues/748.png', 'League', 'Germany'),
(749, 'Oberliga - Bremen', NULL, 'https://media.api-sports.io/football/leagues/749.png', 'League', 'Germany'),
(750, 'Oberliga - Hessen', NULL, 'https://media.api-sports.io/football/leagues/750.png', 'League', 'Germany'),
(751, 'Oberliga - Niederrhein', NULL, 'https://media.api-sports.io/football/leagues/751.png', 'League', 'Germany'),
(752, 'Oberliga - Rheinland-Pfalz / Saar', NULL, 'https://media.api-sports.io/football/leagues/752.png', 'League', 'Germany'),
(753, 'Oberliga - Baden-Württemberg', NULL, 'https://media.api-sports.io/football/leagues/753.png', 'League', 'Germany'),
(754, 'Oberliga - Nordost-Nord', NULL, 'https://media.api-sports.io/football/leagues/754.png', 'League', 'Germany'),
(755, 'Oberliga - Nordost-Süd', NULL, 'https://media.api-sports.io/football/leagues/755.png', 'League', 'Germany'),
(810, 'Super Copa', NULL, 'https://media.api-sports.io/football/leagues/810.png', 'Cup', 'Argentina'),
(817, 'Super Cup Primavera', NULL, 'https://media.api-sports.io/football/leagues/817.png', 'Cup', 'Italy'),
(843, 'Copa Verde', NULL, 'https://media.api-sports.io/football/leagues/843.png', 'Cup', 'Brazil'),
(851, 'Carioca A2', NULL, 'https://media.api-sports.io/football/leagues/851.png', 'League', 'Brazil'),
(871, 'Premier League Cup', NULL, 'https://media.api-sports.io/football/leagues/871.png', 'Cup', 'England'),
(875, 'Segunda División RFEF - Group 1', NULL, 'https://media.api-sports.io/football/leagues/875.png', 'League', 'Spain'),
(876, 'Segunda División RFEF - Group 2', NULL, 'https://media.api-sports.io/football/leagues/876.png', 'League', 'Spain'),
(877, 'Segunda División RFEF - Group 3', NULL, 'https://media.api-sports.io/football/leagues/877.png', 'League', 'Spain'),
(878, 'Segunda División RFEF - Group 4', NULL, 'https://media.api-sports.io/football/leagues/878.png', 'League', 'Spain'),
(879, 'Segunda División RFEF - Group 5', NULL, 'https://media.api-sports.io/football/leagues/879.png', 'League', 'Spain'),
(891, 'Coppa Italia Serie C', NULL, 'https://media.api-sports.io/football/leagues/891.png', 'Cup', 'Italy'),
(892, 'Coppa Italia Serie D', NULL, 'https://media.api-sports.io/football/leagues/892.png', 'Cup', 'Italy');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `competition`
--
-- ALTER TABLE `competition`
--   ADD PRIMARY KEY (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
