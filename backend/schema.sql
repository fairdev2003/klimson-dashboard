-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Wrz 14, 2025 at 05:21 PM
-- Wersja serwera: 10.3.39-MariaDB-0ubuntu0.20.04.2
-- Wersja PHP: 8.1.32

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `www1277_aplikacja`
--

-- --------------------------------------------------------

--
-- Struktura tabeli dla tabeli `data`
--

CREATE TABLE `data` (
  `id` int(11) NOT NULL,
  `poziom` varchar(500) NOT NULL,
  `czas` varchar(500) NOT NULL,
  `ilo` varchar(500) NOT NULL,
  `data` varchar(500) NOT NULL,
  `aktualizacja` varchar(500) NOT NULL,
  `autor` varchar(500) NOT NULL,
  `title` varchar(500) NOT NULL,
  `opis` varchar(500) NOT NULL,
  `graj` int(11) NOT NULL,
  `src` varchar(500) NOT NULL,
  `quiz_photo` varchar(100) NOT NULL,
  `quiz_time` int(11) NOT NULL,
  `isReady` tinyint(1) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_polish_ci;

--
-- Dumping data for table `data`
--

INSERT INTO `data` (`id`, `poziom`, `czas`, `ilo`, `data`, `aktualizacja`, `autor`, `title`, `opis`, `graj`, `src`, `quiz_photo`, `quiz_time`, `isReady`) VALUES
(2, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(52, 'Średni', '4', '9', '23.01.2020', '25.01.2021', 'Paweł Cyngot ', 'Symbolika Wędrowniczej Watry', 'Pytania tego quizu dotyczą symboliki Watry', 553, 'symbolika_wedrowniczej_watry/naramienik_Wedrowniczy.jpg', '', 25, 1),
(53, 'Łatwy', '12', '12', '23.01.2018', '15.02.2022', 'Paweł Cyngot', 'AdminQuizTest- AWARIA!!! Strony', 'Jeśli widzisz ten quiz oznacza to ze: \n- Jest przerwa techniczna lub awaria na strony', 18, 'sssadminquiztest-_awaria!!!_strony/test.png', '', 12, 1),
(55, 'Łatwy', '9', '9', '23.01.2021', '18.02.2022', 'Paweł Cyngot', 'Symbolika Krzyża Harcerskiego', 'Pytania tego quizu dotyczą symboliki krzyża i lilijki. Życzę miłej zabawy ', 596, 'symbolika_krzyza_harcerskiego/krzyz.jpg', '', 30, 1),
(56, 'Średni', '4', '8', '23.01.2021', '25.01.2021', 'Paweł Cyngot', 'Symbolika WAGGGS', 'Pytania tego quizu dotyczą symboliki WAGGGS.\nŻyczę miłej zabawy\nPytania tego quizu dotyczą symboliki WAGGGS.\nŻyczę miłej zabawy\nPytania tego quizu dotyczą symboliki WAGGGS.\nŻyczę miłej zabawy\nPytania tego quizu dotyczą symboliki WAGGGS.\nŻyczę miłej zabawy\n', 253, 'symbolika_wagggs/waggs.jpg', '', 25, 1),
(57, 'Średni', '4', '7', '23.01.2020', '02.01.2022', 'Paweł Cyngot', 'Symbolika WOSM', 'Pytania tego quizu dotyczą symboliki WOSM.\nŻyczę miłej zabawy\n', 114, 'symbolika_wosm/wosm.jpg', '', 30, 1),
(58, 'Łatwy', '2', '5', '23.01.2020', '02.01.2022', 'Paweł Cyngot', 'Symbolika Znaczka Zucha', 'Pytania tego quizu dotyczą symboliki znaczka zucha.\nŻyczę miłej zabawy', 219, 'symbolika_znaczka_zucha/zuch.jpg', '', 35, 1),
(59, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'Symbolika Flagi oraz Herbu Polski', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(60, 'Średni', '8', '15', '16.02.2021', '16.11.2021', 'Paweł Cyngot', 'Znaki Patrolowe cz. 2', 'Pytania tego quizu dotyczą znaków patrolowych. Życzę udanej zabawy', 38, 'znaki_patrolowe_cz__2/Znaki_Patrolowe_02.jpg', '', 20, 1),
(61, 'Średni', '6', '15', '25.04.2021', '16.11.2021', 'Paweł Cyngot', 'Znaki Patrolowe cz. 3', 'Pytania tego quizu dotyczą znaków patrolowych. Życzę udanej zabawy', 26, 'znaki_patrolowe_cz__3/Znaki_Patrolowe_03.jpg', '', 20, 1),
(68, 'Średni', '', '', '07.03.2021', '24.10.2021', 'Paweł Cyngot', 'Znaki Patrolowe cz. 1 (Podstawowe)', 'Pytania tego quizu dotyczą podstawowych znaków patrolowych. Życzę miłej zabawy', 203, 'znaki_patrolowe_cz__1_(podstawowe)/znaki_Patrolowe_1.jpg', '', 20, 1),
(70, 'Średni', '', '', '04.04.2021', '02.01.2022', 'Paweł Cyngot', 'Znane Postacie Harcerskie cz. 1  (Podstawowe)', 'Pytania tego quizu dotyczą znanych ważnych postaci dla harcerstwa i Skautingi. Życzę udanej zabawy', 206, 'znane_postacie_harcerskie_cz__1__(podstawowe)/Znane_postacie_1.jpg', '', 25, 1),
(74, 'Średni', '', '', '27.07.2021', '27.07.2021', 'Paweł Cyngot', 'Znane Postacie Harcerskie cz. 2', '', 0, 'znane_postacie_harcerskie_cz__2/Znane_postacie_2.jpg', '', 25, 1),
(75, 'Średni', '', '', '27.07.2021', '27.07.2021', 'Paweł Cyngot', 'Znane Postacie Harcerskie cz. 3', '', 0, 'znane_postacie_harcerskie_cz__3/Znane_postacie_3.jpg', '', 25, 1),
(76, 'Łatwy', '', '11', '17.10.2021', '18.10.2021', 'Paweł Cyngot', 'Stopnie Harcerskie ZHP', 'Pytania tego quizu dotyczą stopni harcerskich. Życzę miłej zabawy', 332, 'stopnie_harcerskie_zhp/Stopnie_harcerskie_ikona.jpg', '', 25, 1),
(78, 'Łatwy', '', '', '02.11.2021', '02.11.2021', '', 'Sznury Funkcyjne cz. 1 (Podstawowe)', 'W Fazie tworzenia ', 17, 'sznury_funkcyjne_cz__1_(podstawowe)/', '', 25, 1),
(79, 'Poziom trudności', '', '', '02.11.2021', '02.11.2021', '', 'Sznury Funkcyjne cz. 2', '', 0, 'sznury_funkcyjne_cz__2/', '', 0, 1),
(80, 'Poziom trudności', '', '', '02.11.2021', '02.11.2021', '', 'Sznury Funkcyjne Z Suwakami cz. 3 ', '', 0, 'sznury_funkcyjne_z_suwakami_cz__3_/', '', 0, 1),
(81, 'Średni', '', '', '20.12.2021', '02.01.2022', 'Paweł Cyngot', 'Znaki Służby ', 'Quiz sprawdza twoją wiedze z zakresu znajomości znaków służby.  Powodzenia i miłej zabawy. ', 75, 'znaki_sluzby_/Znaki_Sluzby_new.jpg', '', 20, 1),
(82, 'Łatwy', '', '', '01.01.2022', '04.02.2022', 'Paweł Cyngot', 'Prawo Zucha', 'Quiz sprawdza twoją wiedze z zakresu znajomości prawa zuchowego.  Powodzenia i miłej zabawy. ', 141, 'prawo_zucha/prawo_zucha.jpg', '', 30, 1),
(83, 'Trudny', '', '', '01.01.2022', '02.01.2022', 'Paweł Cyngot', 'Plakietki Specjalności Drużyn ', 'Quiz sprawdza twoją wiedze z zakresu znajomości plakietek specjalności drużyn.  Powodzenia i miłej zabawy. \n ', 144, 'plakietki_specjalnosci_druzyn_/Plakietki_Specjalnosci_druzyn.jpg', '', 30, 1),
(86, 'Łatwy', '', '', '18.02.2022', '18.02.2022', 'Paweł Cyngot ', 'Symbolika Lilijki Harcerskiej', 'Pytania tego quizu dotyczą symboliki lilijki. Życzę miłej zabawy', 66, 'symbolika_lilijki_harcerskiej/lilijka.jpg', '', 25, 1),
(87, 'Średni', '', '', '20.12.2021', '02.01.2022', 'Paweł Cyngot ', 'Znaki Służby ', 'Quiz sprawdza twoją wiedze z zakresu znajomości znaków służby.  Powodzenia i miłej zabawy. Wędrowniczej. Życzę miłej zabawy', 73, 'znaki_sluzby_/Znaki_Sluzby_new.jpg', '', 20, 1),
(88, 'Średni', '', '', '20.12.2021', '02.01.2022', 'Paweł Cyngot', 'Znaki Służby ', 'Quiz sprawdza twoją wiedze z zakresu znajomości znaków służby.  Powodzenia i miłej zabawy. ', 73, 'znaki_sluzby_/Znaki_Sluzby_new.jpg', '', 20, 1),
(90, 'Średni', '', '', '20.12.2021', '02.01.2022', 'Paweł Cyngot', 'Znaki Służby ', 'Quiz sprawdza twoją wiedze z zakresu znajomości znaków służby.  Powodzenia i miłej zabawy. ', 73, 'znaki_sluzby_/Znaki_Sluzby_new.jpg', '', 20, 1),
(91, 'Średni', '', '', '20.12.2021', '02.01.2022', 'Paweł Cyngot', 'Znaki Służby ', 'Quiz sprawdza twoją wiedze z zakresu znajomości znaków służby.  Powodzenia i miłej zabawy. ', 73, 'znaki_sluzby_/Znaki_Sluzby_new.jpg', '', 20, 1),
(92, 'Średni', '4', '9', '23.01.2020', '25.01.2021', 'Paweł Cyngot', 'Symbolika Wędrowniczej Watry', 'Pytania tego quizu dotyczą symboliki krzyża i lilijki. Życzę miłej zabawy ', 553, 'symbolika_wedrowniczej_watry2/naramienik_Wedrowniczy.jpg', '', 25, 0),
(94, 'Łatwy', '0', '0', '02.11.2024', '02.11.2024', 'testowy', 'testowy', 'test', 0, '', '', 12, 1),
(96, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(97, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(98, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(99, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(100, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(101, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(102, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(103, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(104, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(105, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(106, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(107, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(108, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(109, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(110, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(111, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(112, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(113, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(114, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(115, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(116, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(117, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(118, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(119, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(120, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(121, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(122, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(123, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(124, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1),
(125, 'Łatwy', '3', '6', '23.01.2019', '21.10.2021', 'Paweł Cyngot', 'DO TESTOWANIA JEDNO PYTANIE', 'Pytania tego quizu dotyczą symboliki Flagi oraz Herbu. Polski Życzę miłej zabawy', 559, 'symbolika_flagi_oraz_herbu_polski/falga_herb_polski.jpg', '', 25, 1);

-- --------------------------------------------------------

--
-- Struktura tabeli dla tabeli `extraPytania`
--

CREATE TABLE `extraPytania` (
  `question_id` int(11) NOT NULL,
  `quiz_id` int(11) NOT NULL,
  `question` mediumtext NOT NULL,
  `questionOrder` int(10) NOT NULL DEFAULT 20,
  `link` varchar(100) NOT NULL,
  `link2` text NOT NULL,
  `link3` text NOT NULL,
  `link4` text NOT NULL,
  `right_answer` varchar(200) NOT NULL,
  `photo` varchar(100) NOT NULL,
  `photo2` text NOT NULL,
  `photo3` text NOT NULL,
  `photo4` text NOT NULL,
  `linkOrder` int(10) NOT NULL DEFAULT 0,
  `link2Order` int(10) NOT NULL DEFAULT 1,
  `link3Order` int(10) NOT NULL DEFAULT 2,
  `link4Order` int(10) NOT NULL DEFAULT 3,
  `photoOrder` int(10) NOT NULL DEFAULT 4,
  `photo2Order` int(10) NOT NULL DEFAULT 5,
  `photo3Order` int(10) NOT NULL DEFAULT 6,
  `photo4Order` int(10) NOT NULL DEFAULT 7,
  `text` text NOT NULL,
  `textOrder` int(10) NOT NULL DEFAULT 10,
  `text2` text NOT NULL,
  `text2Order` int(10) NOT NULL DEFAULT 11,
  `text3` text NOT NULL,
  `text3Order` int(10) NOT NULL DEFAULT 12,
  `text4` text NOT NULL,
  `text4Order` int(10) NOT NULL DEFAULT 13
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_polish_ci;

--
-- Dumping data for table `extraPytania`
--

INSERT INTO `extraPytania` (`question_id`, `quiz_id`, `question`, `questionOrder`, `link`, `link2`, `link3`, `link4`, `right_answer`, `photo`, `photo2`, `photo3`, `photo4`, `linkOrder`, `link2Order`, `link3Order`, `link4Order`, `photoOrder`, `photo2Order`, `photo3Order`, `photo4Order`, `text`, `textOrder`, `text2`, `text2Order`, `text3`, `text3Order`, `text4`, `text4Order`) VALUES
(142, 35, 'Jesteście gotowi ?', 3, '', '', '', '', 'tak / start /gotowi / #819', 'harcerska_grodziska_gra_terenowa/02WG.png', 'harcerska_grodziska_gra_terenowa/', 'harcerska_grodziska_gra_terenowa/', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 1, 0, 0, 0, 'Witajcie na Harcerskiej  Grodziskiej Grze terenowej !\n\n', 2, '', 0, '', 0, '', 0),
(144, 35, 'Wypisz znaki z \npierwszego wersu od góry \n8,14,15 oraz z\npiątego wersu od góry \n10,12,17', 4, '', '', '', '', '15OTBK / #775 ', 'harcerska_grodziska_gra_terenowa/1map.png', 'harcerska_grodziska_gra_terenowa/03WG.png', 'harcerska_grodziska_gra_terenowa/', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 1, 3, 0, 0, 'Przypomnienie :\nznaki polskie oraz mała wielka litera nie robi znaczenia  Znaki pisz jeden za drugim bez przecinków (jeśli nie jest wymagany) ', 2, '', 0, '', 0, '', 0),
(153, 35, 'Jaki wynik powstaje po dodaniu wysokości kapliczki oraz jej roku powstania ? ', 5, '', '', '', '', '1864,9 / #686 ', 'harcerska_grodziska_gra_terenowa/04RWG.png', 'harcerska_grodziska_gra_terenowa/05RWG.png', 'harcerska_grodziska_gra_terenowa/2map.png', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 2, 4, 1, 0, 'Odszukaj informacje na temat wysokości kapliczki   oraz przypuszczalny rok powstania \n', 3, '', 0, '', 0, '', 0),
(155, 35, 'Jaka jest suma wszystkich licz z tablicy jaki jest pseudonim ?\n', 6, '', '', '', '', ' 3923Niedzwiadek /#768 ', 'harcerska_grodziska_gra_terenowa/3map.png', 'harcerska_grodziska_gra_terenowa/07WG.png', 'harcerska_grodziska_gra_terenowa/08WG.png', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 1, 3, 5, 0, 'HISTORIA:\nKamień z tablicą upamiętniającą pobyt w Grodzisku Maz. ostatniego dowódcy Armii Krajowej gen. bryg. Leopolda Okulickiego. Tablica usytuowana w Parku Skarbków, w pobliżu grodziskiej kwatery generała przy ul. Przejazdowej 8 została odsłonięta 20 grudnia 1990 r. Tablicę odsłonili: bratanica generała Irena Okulicka-Kaczmarska i zastępca Kwatermistrza WP gen. bryg. Roman Pusiak, a jej poświęcenia dokonał bp Zbigniew J. Kraszewski, duszpasterz kombatantów RP. Wykonawcą tablicy był Roman Reichert. Z inicjatywy kombatantów AK gen. bryg. Leopold Okulicki stał się patronem przebiegającej w tym miejscu ulicy, wcześniej noszącej imię gen. Karola Świerczewskiego.\n\nLeopold Okulicki ps. Kobra, Niedźwiadek (1898-1946)\ngenerał bryg. WP; w okresie II wojny światowej działacz konspiracji, od 1939 r. w Służbie Zwycięstwu Polski, następnie w Związku Walki Zbrojnej, komendant ZWZ na terenach okupacji sowieckiej, w 1941 r. aresztowany przez sowietów, po wybuchu wojny niemiecko-radzieckiej zwolniony z więzienia, 1941-42 szef sztabu armii polskiej w ZSRR, 1942-43 dowódca 7 Dywizji Piechoty; cichociemny; od 1944, po przerzuceniu do kraju, w Armii Krajowej, zastępca szefa sztabu Komendy Głównej, komendant organizacji „Nie”, 1944-45 dowódca AK; podstępnie aresztowany przez NKWD, skazany w tzw. procesie szesnastu w Moskwie, prawdopodobnie zamordowany 24 grudnia 1946 r. w sowieckim więzieniu.', 2, 'ZAPIS:  0000PSEUDONIM ', 4, '', 0, '', 0),
(156, 35, 'Ile jest wszystkich grajków na Muralu  ? \n(Mowa o tym największym ) nie liczą się małe garaże ', 5, '', '', '', '', '  35 / #323 ', 'harcerska_grodziska_gra_terenowa/4map.png', 'harcerska_grodziska_gra_terenowa/09WG.png', 'harcerska_grodziska_gra_terenowa/', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 1, 3, 0, 0, 'Mural Orkiestra namalowany w 2014 roku wg. projektu Edmunda Dudka , jest to najdłuższy z grodziskich murali i robi nam paparara.\n\n', 2, 'ZADANIE DODATKOWE:\nzadzwoń wrazi pominięcia zadania ', 4, '', 0, '', 0),
(157, 35, 'W którym roku zaprojektowano wnętrze \ndworku  oraz kiedy odbyła się ostatni remont ', 6, '', '', '', '', '1761,2003-2004 /  #204  ', 'harcerska_grodziska_gra_terenowa/5map.png', 'harcerska_grodziska_gra_terenowa/10WG.png', 'harcerska_grodziska_gra_terenowa/', 'harcerska_grodziska_gra_terenowa/12WG.jpg', 0, 0, 0, 0, 1, 2, 0, 4, '', 0, 'W centrum Grodziska w pobliżu skrzyżowania głównych ulic znajduje się tzw. Dwór Skarbków. Został on wybudowany w drugiej połowie XVIII w. dla rodziny Mokronoskich.  Wskazują na to polichromie malowane przez Jana Bogumiła Plerscha w 1782 r. Podobno projekt dworu został opracowany przez architekta Charlesa Pierre Coustou w 1761 r. i po kilkunastu latach wcielony w życie. Źródła historyczne wspominają, że wcześniej stała w tym miejscu inna rezydencja, na której założeniach wybudowano nową.\nParterowy budynek na planie prostokąta, kryty dwuspadowym dachem jest utrzymany w w stylu barokowo-klasycystycznym.\nDwór był własnością kolejnych właścicieli miasta - Mokronoskich, a następnie Skarbków. Kilkukrotnie przebudowywany, a obecny wygląd zawdzięcza Izabeli z Poniatowskich Branickiej, która po śmierci męża - hetmana Jana Klemensa Branickiego zamieszkała w Grodzisku z Andrzejem Mokronoskim.\n\nBudynek przeszedł gruntowny remont w latach 1952-1956, a ostatnio w 2003-2004 r. We wnętrzu zachowało się fragmenty oryginalnego wyposażenia i dekoracji, m.in. wspomniane już polichromie, dwie szafy w westybulu, pokryte rysunkami o motywach arabeskowych z medalionami, w których przedstawiono postacie mitologiczne, trzy kominki: w westybulu rokokowo-klasycystyczny z piaskowca, w salach przylegających do westybulu barokowo-klasycystyczne.\n\nObecnie mieści się tu Państwowa Szkoła Muzyczna I stopnia im. T. Bairda', 3, 'ZAPIS:  \nXXXX,XXXX-XXXX', 5, '', 0),
(158, 35, 'Ile okien dachowych mam dworek ?\n(Do okoła) ', 5, '', '', '', '', '10 / #397 ', 'harcerska_grodziska_gra_terenowa/6map.png', 'harcerska_grodziska_gra_terenowa/11WG.jpg', 'harcerska_grodziska_gra_terenowa/', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 1, 3, 0, 0, 'Spodobało wam się liczenie myślę ze tak \npowodzenia :) \n', 2, 'ZADANIE DODATKOWE:\nzadzwoń wrazi pominięcia zadania ', 4, '', 0, '', 0),
(159, 35, 'Pierwszy wers słowo drugie i trzecie \ntrzeci wers znaki:\n9,11,15,22\n\n', 5, '', '', '', '', 'LEONIDTELIGACA,Z / #953', 'harcerska_grodziska_gra_terenowa/7map.jpg', 'harcerska_grodziska_gra_terenowa/13WG.jpg', 'harcerska_grodziska_gra_terenowa/14WG.jpg', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 1, 2, 4, 0, 'ZAPIS: \nPisz wszystko razem pamiętaj , to też znak ', 3, '', 0, '', 0, '', 0),
(160, 35, 'Jakie jest ostatnie słowo z muralu?', 4, '', '', '', '', 'bażant  /  #173 ', 'harcerska_grodziska_gra_terenowa/8map.jpg', 'harcerska_grodziska_gra_terenowa/15WG.jpg', 'harcerska_grodziska_gra_terenowa/', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 1, 2, 0, 0, 'HISTORIA:\nTAJNY MAGAZYN BRONI  AK\nDom rodziny Szmajdowiczów i Osuchowskich był jednym z wielu lokali na terenie miasta, w którym przechowywano broń stanowiącą arsenał Ośrodka „Gąbka” - „Osa” oraz rezerwy sprzętu militarnego dla całego Obwodu „Bażant”. W grobli na Stawach Goliana wykonany został magazyn przejściowy, z którego nocą transportowano znajdujące się tam wyposażenie do magazynu głównego ulokowanego w domu przy ul. Sienkiewicza. Zamaskowane wejście do magazynu prowadziło z posesji Wiktora Goliana ps. Nałęcz. Wskutek denuncjacji 31 lipca 1944 r.. magazyn został odkryty i zdobyty przez okupantów niemieckich.', 3, '', 0, '', 0, '', 0),
(161, 35, 'W którym roku zbudowano kolej  EKD \noraz przez kogo ?', 8, '', '', '', '', '1927,SiłaiŚwiatło /   #185 ', 'harcerska_grodziska_gra_terenowa/9map.jpg', 'harcerska_grodziska_gra_terenowa/17WG.jpg', 'harcerska_grodziska_gra_terenowa/16WG.jpg', 'harcerska_grodziska_gra_terenowa/18WG.jpg', 0, 0, 0, 0, 1, 2, 4, 6, '', 0, 'HISTORIA \nJedno z najsilniejszych zgrupowań ZWZ/AK w Ośrodku „Gąbka” - „Osa” istniało w Elektrycznych Kolejach Dojazdowych (ob. WKD). Do zadań kolejarzy zaprzysiężonych w AK należał przerzut broni, wywiad i kontrwywiad. Personel EKD starał się także chronić pasażerów ostrzegając przed rewizjami i łapankami oraz wydając fałszywe legitymacje dla tzw. niedzielnych konduktorów. Podczas akcji masowego wywożenia mieszkańców stolicy po wybuchu powstania warszawskiego kolejarzom udało się uratować znaczną grupę wysiedlonych przed osadzeniem w obozie przejściowym Dulag 121 w Pruszkowie. Kolejką EKD ewakuowano także do Grodziska i Milanówka warszawski Szpital Dzieciątka Jezus.\nEKD i jej załoga w czasie okupacji stanowili pierwszy ośrodek organizacyjny ruchu oporu. Tu zostały założone pierwsze komórki polskiego podziemia. Całość prac organizacyjnych w konspiracji koncentrowała się w ośrodku dyspozycyjnym w Komendzie Obwodu „Bekas” - późniejszy kryptonim „Bażant”, zaś bezpośrednie zwierzchnictwo nad działalnością pełniła Komenda Ośrodka „Gąbka” - późniejszy kryptonim „Osa”......', 3, 'Odszukaj  tablicy na stacji ', 5, 'ZAPISZ:\n0000,NAZWA', 7),
(162, 35, 'Jztn rwt aizjdujc snę iz kzylneb  ?', 3, '', '', '', '', '1924  / #541', 'harcerska_grodziska_gra_terenowa/10map.jpg', 'harcerska_grodziska_gra_terenowa/19WG.jpg', 'harcerska_grodziska_gra_terenowa/', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 1, 2, 0, 0, '', 0, '', 0, '', 0, '', 0),
(163, 35, 'Znajdź ukryte  znaki na słupach wejścia ', 4, '', '', '', '', 'XDKMB1914 / #131', 'harcerska_grodziska_gra_terenowa/20WG.jpg', 'harcerska_grodziska_gra_terenowa/11map.jpg', 'harcerska_grodziska_gra_terenowa/', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 2, 1, 0, 0, '', 0, '', 0, '', 0, '', 0),
(164, 35, 'Odszukaj na muralu  \ndrugi wers ?\n', 5, '', '', '', '', 'PROJEKT-539 / #821', 'harcerska_grodziska_gra_terenowa/12map.jpg', 'harcerska_grodziska_gra_terenowa/21WG.jpg', 'harcerska_grodziska_gra_terenowa/', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 1, 2, 0, 0, 'Odszukajcie mural ', 3, 'ZAPIS \nXXXXXXX-XXX', 4, '', 0, '', 0),
(165, 35, 'Kto zlecił budowę wili oraz w którym roku ?', 5, '', '', '', '', ' Zacharija Putiato 1889  / #125', 'harcerska_grodziska_gra_terenowa/22WG.jpg', 'harcerska_grodziska_gra_terenowa/13map.jpg', 'harcerska_grodziska_gra_terenowa/', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 2, 1, 0, 0, 'HISTORIA\nWilla Radogoszcz został wybudowana w \n------- roku na zlecenie ---------  -------, naczelnika powiatu sochaczewskiego, z przeznaczeniem na letnią rezydencję carskiego urzędnika.  Po 1914 roku dom został zakupiony przez Rozalię i Pawła Helcmanów, warszawskich aptekarzy. W latach międzywojennych budynek dość często zmieniał właścicieli, którzy nadawali mu rozmaite nazwy. Według zachowanych źródeł, jednym z mieszkańców willi miał być miejscowy rabin.  Z willi korzystał również Bolesław Saltz, burmistrz miasta, jednakże okoliczności  jego pobytu  – ze względu na brak dokumentów – pozostają niejasne. Po 1945 roku swoją siedzibę miało tu powołane przez artystów osiadłych w Grodzisku i okolicach Stowarzyszenie Miłośników Sztuki. W ramach działalności kulturalnej stowarzyszenia  prowadzono szereg sekcji związanych z rozmaitymi formami kształcenia artystycznego, m.in. sekcję fotograficzną i teatralną, a także ogniska – plastyczne i muzyczne, które dały początek placówkom publicznym, do dziś funkcjonującym w mieście.  W latach PRL-u willa służyła również jako miejsce spotkań różnych środowisk i organizacji działających w mieście, w tym rzemieślników i kupców grodziskich.', 3, 'ZAPISZ:\n Imię nazwisko rok \n(z spacjami ) ', 4, '', 0, '', 0),
(166, 35, 'Ami bwł  Fclmas  Dzmcrżkutysam ?', 4, '', '', '', '', 'Muzyk / kompozytor / dyrygent / kapelmistrz / pedagog / #366 ', 'harcerska_grodziska_gra_terenowa/14map.jpg', 'harcerska_grodziska_gra_terenowa/23WG.jpg', 'harcerska_grodziska_gra_terenowa/', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 1, 2, 0, 0, 'Szukajcie tablic ', 3, '', 0, '', 0, '', 0),
(167, 35, 'Jaki jest rok założenia cmentarza oraz\njakie nr teksów cmentarza były pierwsze ?\n', 4, '', '', '', '', '1817-VI-IX /  #429 ', 'harcerska_grodziska_gra_terenowa/15map.jpg', 'harcerska_grodziska_gra_terenowa/25WG.png', 'harcerska_grodziska_gra_terenowa/', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 1, 2, 0, 0, 'ZAPISZ:\nRok-nr-nr', 3, '', 0, '', 0, '', 0),
(168, 35, 'Odszukaj znaków na tablicy \nWers drugi : 2.5,14\nWers czwarty :3,7,10\nWers ósmy : 5,10,12', 3, '', '', '', '', 'AAZAWANMA / #366 ', 'harcerska_grodziska_gra_terenowa/16map.jpg', 'harcerska_grodziska_gra_terenowa/25WG.png', 'harcerska_grodziska_gra_terenowa/', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 1, 2, 0, 0, '', 0, '', 0, '', 0, '', 0),
(169, 35, 'Ek uk zm pkaino?', 5, '', '', '', '', 'Pomnik Wolności i Zwycięstwa  / #739 ', 'harcerska_grodziska_gra_terenowa/17map.jpg', 'harcerska_grodziska_gra_terenowa/26WG.jpg', 'harcerska_grodziska_gra_terenowa/27WG.jpg', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 1, 2, 4, 0, 'HISTORIA\nPo odzyskaniu przez Polskę niepodległości powstało wiele obiektów upamiętniających tę okoliczność. W Grodzisku Mazowieckim w 1919 r. wzniesiono pomnik na Placu Wolności. Obelisk w dwudziestoleciu międzywojennym był miejscem uroczystości patriotycznych. 1 września 1939 r., w pierwszym dniu II wojny światowej, podczas nalotu i bombardowania niemieckiego pomnik legł w gruzach. W 1984 r. obelisk zrekonstruowano.', 3, '', 0, '', 0, '', 0),
(170, 35, '•-••/•//•---/•/•••/-//--/••/•/-•-•/--••/-•--//-•/•-//•--•/---/--/-•/••/-•-/••-// ?\nUbikl ulsy pklroszl mkbsyw nb śckbnkl ?', 5, '', '', '', '', 'dwa-Warszawa / #592', 'harcerska_grodziska_gra_terenowa/18map.jpg', 'harcerska_grodziska_gra_terenowa/28WG.jpg', 'harcerska_grodziska_gra_terenowa/', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 1, 2, 0, 0, 'HISTORIA\nPomnik ku czci żołnierzy polskich poległych na frontach II wojny światowej wg proj. Czesława Piwowarczyka odsłonięty 25 września 1966 r. Symboliczne mauzoleum upamiętnia miejsca heroicznej walki narodu polskiego z okupantem w latach 1939-1945. Pomnik znajduje się w miejscu egzekucji dokonanej przez wojsko niemieckie na 20 więźniach Pawiaka 18 listopada 1943 r., o czym przypomina tablica z nazwiskami ofiar oraz ślady po kulach widoczne na murze pobliskiej kamienicy. W 1989 r. złożono tu urnę z prochami żołnierzy polskich zamordowanych w Katyniu. W 1998 r. pomnik uzupełniony został o tablice upamiętniające żołnierzy Armii Krajowej i ofiary sowieckiego bombardowania miasta z 16 stycznia 1945 r. Dnia 18 listopada 1943 r. na grodziskim rynku, w miejscu po spalonym domu żydowskiego kupca o nazwisku Szydło, niemieccy okupanci rozstrzelali 20 osób. Ofiarami egzekucji byli więźniowie Pawiaka straceni w odwecie za śmierć Edwarda Sommera, dyrektora pozostającej pod niemieckim nadzorem fabryki tarcz ściernych „Haeberle i S-ka”, na którym wyrok wykonała 28 października komórka likwidacyjna AK. Początkowo za ten czyn odpowiedzieć mieli mieszkańcy Grodziska, którzy zostali aresztowani i osadzeni w tutejszym więzieniu. Aby nie dopuścić do ich zamordowania, komenda Ośrodka „Gąbka”- „Osa” sporządziła listę miejscowych folksdojczów, którą dostarczono na posterunek niemieckiej policji z ostrzeżeniem, że w razie egzekucji grodziszczan umieszczeni na niej kolaboranci będą zlikwidowani. Na skutek tej groźby zakładnicy zostali zwolnieni, zaś ich miejsce zajęli więźniowie z Pawiaka przywiezieni do Grodziska. Egzekucja miała wywrzeć wrażenie na ludności – ofiary ze związanymi rękami i zakneblowanymi ustami prowadzono piątkami pod mur kamienicy i rozstrzeliwano salwą. Zwłoki pomordowanych, z których połowę stanowili żołnierze Obwodu AK „Skowronek” z Sochaczewa, zabrano z miejsca kaźni i wywieziono w nieznanym kierunku. „Rynek zapełnił się natychmiast głównie kobietami, które cisnęły się do miejsca stracenia, żeby maczać chustki we krwi rozstrzelanych i palić świeczki na miejscu ich męczeństwa – notował w Dzienniku okupacyjnym Stanisław Rembek, naoczny świadek zdarzenia – Nasiąkłą krwią ziemię zebrano starannie i zaniesiono na cmentarz. Świeczki paliły się jeszcze przez szereg dni.” Po zakończeniu okupacji w „miejscu straceń” stanął krzyż, zaś we wrześniu 1966 r. odsłonięto tu pomnik wg projektu Czesława Piwowarczyka ps. Olszyna. Ofiary egzekucji dokonanej 18 listopada 1943 r. przez okupantów niemieckich:\nTadeusz Anikiewicz, Franciszek Bolczak, Konstanty Borkowski, Paweł Buza, Stanisław Dolubowski, Franciszek Dąbrowski, Marian Duński, Józef Gdowiak, Józef Grabski, Henryk Jabłoński, Antoni Jaskólski, Konstanty Killer, Jan Kowalski, Franciszek Luc, Daniel Matusiak, Bolesław Pawłowski, Mieczysław Pliszkiewicz, Mieczysław Raczkowski, Antoni Stasikowski, Stanisław Sztark. We wtorek 16 stycznia 1945 r., na rynek w Grodzisku wypełniony do niemożliwości kłębiącym się tłumem spadły bomby zrzucone z samolotów radzieckich, poprzedzających zbliżającą się do Warszawy piechotę 1 Frontu Białoruskiego Armii Czerwonej. Bezprecedensowy i nieuzasadniony z przyczyn militarnych atak na ludność cywilną, do którego użyto bomb rozpryskowych i burzących oraz serii z broni pokładowej, przyniósł nadzwyczaj krwawe żniwo. Tego dnia, w godzinach popołudniowych, zginęło blisko 200 osób, a drugie tyle zmarło wkrótce z powodu odniesionych ran. Szok spowodowany nalotem, ilość rannych, mrok zapadający nad miastem utrudniający akcję ratunkową oraz obawa przed powtórnym bombardowaniem sprawiły, że wydarzenie miało charakter istnej hekatomby. Oddajmy głos Bogdanowi Lewandowskiemu, świadkowi traumatycznego wydarzenia: „Rosjanie strzelali i bombardowali wypełnione tysiącami warszawiaków centrum miasta. Ludzie padali pokotem. Na rynku rozszarpane ciała ludzkie pomieszane były ze szczątkami straganów. Przez wiele godzin umierali na ulicach ranni, pozbawieni pomocy. Następnego dnia, ponad pięćdziesiąt nierozpoznanych zwłok ułożono na chodniku obok cmentarza, a tłum przechodził, aby znaleźć swoich bliskich lub znajomych” (Banwar1944: Wypędzeni z Warszawy 1944. Losy dzieci, www.banwar1944.eu). Wspomnienia tego tragicznego w skutkach zdarzenia wryły się w pamięć wielu świadków, którzy w swoich relacjach powtarzają informacje na temat dziesiątków zwłok przypominających kłody drzewa, układanych w alejach cmentarnych i w pobliżu cmentarza, a także na podwórkach szpitali. Pogrzeby ofiar nalotu trwały kilkanaście dni, a niezidentyfikowane ofiary – osoby, które zginęły wskutek bombardowania bądź też zmarły tego samego dnia w szpitalu – złożono w zbiorowych mogiłach na miejscowym cmentarzu (do dziś zachowała się jedna z nich).', 3, 'ZAPIS:  (Słowny)nie cyfrowy\nFORMA: Odpowiedz-Odpowiedz', 4, '', 0, '', 0),
(171, 35, 'PYTANIE: \n', 6, '', '', '', '', '11  /  #300 ', 'harcerska_grodziska_gra_terenowa/29WG.jpg', 'harcerska_grodziska_gra_terenowa/19map.jpg', 'harcerska_grodziska_gra_terenowa/30WG.jpg', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 1, 2, 7, 0, '', 0, 'ZADANIE DODATKOWE:\nzadzwoń wrazi pominięcia zadania ', 4, 'HISTORIA:\nMiejsce pierwszej zbiorowej egzekucji dokonanej przez oddział SS na mieszkańcach Grodziska Maz. 20 września 1939 r. na terenie garbarni, w odwecie za rzekome zabójstwo niemieckiego żołnierza, rozstrzelano 11 osób. Zginęli: Jan Kamiński, Marcin Michalski, Jan Pobiedziński, Wincenty Samoraj, Albin Szatan, Józef Somczyński, Wiesław Somczyński, Jan Zagajewski, Henryk Zagajewski oraz dwaj niezidentyfikowani obywatele miasta. W 1970 r. w miejscu egzekucji odsłonięto tablicę upamiętniającą ofiary masakry. W okresie okupacji garbarnia stała się azylem dla uchodźców: ludności wysiedlonej z woj. poznańskiego, Dzieci Zamojszczyzny i uciekinierów z grodziskiego getta.\n\nWspomnienie o egzekucji z 20 września 1939 r. autorstwa Jerzego Kowalczyka opublikowane w książce M. Cabanowskiego „Domy i ludzie” (Grodzisk Maz. 1998)\n\n„Tego właśnie dnia około godziny piątej rano mieszkańcy, sąsiadujących z Garbarnią domów zostali obudzeni serią strzałów karabinowych. Były one skierowane do jednego ze współwłaścicieli Garbarni, który akurat wyszedł z domu i przechodził przez teren fabryczny. Strzały okazały się niecelne. Mężczyzna ów zdołał uciec i skryć się w szafce roboczej na terenie fabryki. Niemcy, natomiast, poszli w kierunku sąsiedniej posesji, należącej do Jana Zagajewskiego. Właściciel był już na podwórzu i przygotowywał do otwarcia swój warsztat kowalski.', 3, 'ZAPIS :  (Cyfrowy) ', 5),
(172, 35, 'Jaką produkcje zaczęła fabryka w  1940r ?\n', 5, '', '', '', '', 'bomby / bomba /  #882', 'harcerska_grodziska_gra_terenowa/20map.jpg', 'harcerska_grodziska_gra_terenowa/31WG.jpg', 'harcerska_grodziska_gra_terenowa/', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 1, 2, 0, 0, 'ZAPIS: \nskłada się z 1 słowa \n ', 4, 'HISTORIA :\nW 1940 r. na terenie fabryki rozpoczęła się tajna produkcja (-----------------------------------------------------------). Dnia 28 października 1943 r. żołnierze komórki likwidacyjnej AK wykonali wyrok śmierci na Edwardzie Sommerze, dyrektorze fabryki. W odwecie aresztowani została grupa pracowników fabryki, którzy wraz z innymi mieszkańcami miasta mieli być rozstrzelani. W wyniku prowokacji Ośrodka AK, zakładnicy zostali zwolnieni. Zamiast nich śmierć poniosło 20 więźniów Pawiaka, których egzekucja odbyła się na rynku w tzw. Miejscu Straceń.', 3, '', 0, '', 0),
(173, 35, 'Wypisz znaki z \npierwszego wersu od góry \n11,14,19 znak \npiątego wersu od góry \n10,12,17 znak ', 4, '', '', '', '', 'NROYNS / #491', 'harcerska_grodziska_gra_terenowa/21map.jpg', 'harcerska_grodziska_gra_terenowa/32WG.jpg', 'harcerska_grodziska_gra_terenowa/', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 1, 2, 0, 0, 'HISTORIA: \nOd 1941 r. w więzieniu działała tajna drukarnia prowadzona przez żołnierzy Armii Krajowej pod kierunkiem lekarza więziennego Jerzego Szpakowskiego. Drukowano w niej prasę konspiracyjną, redagowano podziemne periodyki, prowadzono nasłuch radiowy i kolportaż. W sierpniu 1943 r. w wyniku aresztowania dr. Szpakowskiego drukarnia została przeniesiona do Milanówka, gdzie pracowała do końca wojny. Do sierpnia 1943 r. na terenie więzienia znajdował się tajny magazyn leków i materiałów sanitarnych Ośrodka AK, przeniesiony do nieczynnej cegielni w Natolinie.\n\nDr med. Jerzy Szpakowski ps. Bronisław, kierownik referatu propagandy komendy Obwodu „Bażant”, aresztowany 21 sierpnia 1943 r. przez Gestapo, zamordowany na Pawiaku. Jego imię nosi ulica przebiegająca przy murze więzienia.', 3, '', 0, '', 0, '', 0),
(174, 35, 'Jaka jest pełna data przyjazdu pierwszego pociągu ?', 5, '', '', '', '', '14.06.1845r /  #193', 'harcerska_grodziska_gra_terenowa/22map.jpg', 'harcerska_grodziska_gra_terenowa/33WG.jpg', 'harcerska_grodziska_gra_terenowa/', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 1, 2, 0, 0, 'ZAPIS:\n00.00.0000r', 4, 'HISTORIA:\nW pierwszych dniach wojny na dworcu uruchomiony został punkt pomocy dla polskich żołnierzy obsługiwany przez harcerki i uczennice grodziskiego gimnazjum, zrzeszone następnie w Wojskowej Służbie Kobiet (WSK). Komórka konspiracyjna ZWZ/AK na dworcu PKP została utworzona na początku 1940 r. Złożona z pracowników kolejowych wszystkich szczebli sekcja należała do najsilniejszych zgrupowań tajnej organizacji na terenie miasta. Poza szkoleniem w zakresie posługiwania się bronią i materiałami wybuchowymi, kolportażem prasy podziemnej i ochroną podróżnych przed łapankami, uczestnicy konspiracji kolejowej wykonywali akcje wywiadowcze i sabotażowe, a po wybuchu powstania warszawskiego nieśli pomoc wysiedlanej ludności stolicy. Pod koniec wojny kolejarze nie dopuścili do zniszczenia obiektów kolejowych, zaminowanych przez wycofujące się wojsko niemieckie', 3, '', 0, '', 0),
(175, 35, 'Jaki jest rok powstania pomnika? ', 4, '', '', '', '', '2015r / #700', 'harcerska_grodziska_gra_terenowa/34WG.jpg', 'harcerska_grodziska_gra_terenowa/23map.jpg', 'harcerska_grodziska_gra_terenowa/', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 2, 1, 0, 0, 'ZAPISZ:\n0000r', 3, '', 0, '', 0, '', 0),
(176, 35, 'Ile obrazów Chełmońskiego znajduje się  na \"Deptaku\" w Grodzisku ', 4, '', '', '', '', '11 / #444', 'harcerska_grodziska_gra_terenowa/24map.jpg', 'harcerska_grodziska_gra_terenowa/35WG.jpg', 'harcerska_grodziska_gra_terenowa/', 'harcerska_grodziska_gra_terenowa/', 0, 0, 0, 0, 1, 2, 0, 0, 'Zapis: Cyfrowy ', 3, '', 0, '', 0, '', 0),
(185, 36, 'SIEMA NOWY', 0, '12', '', '', '', '123', 'harcerska_gra_grodziskowa_test/colorbluje.png', 'harcerska_gra_grodziskowa_test/GM_PERFEKT_3.png', 'harcerska_gra_grodziskowa_test/004-pictogram-free.jpg', 'harcerska_gra_grodziskowa_test/GJB.png', 0, 0, 0, 0, 0, 0, 0, 0, '1212', 0, '', 0, '', 0, '', 0),
(186, 36, 'nowe', 0, '', '', '', '', '123', 'GJB2.png', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, '12', 0, '', 0, '', 0, '', 0),
(187, 38, 'asd', 0, '', '', '', '', '12', 'harcerska_gra_grodziskowa_test_nowa/GM_PERFEKT_1.png', 'harcerska_gra_grodziskowa_test_nowa/GM_PERFEKT.png', 'harcerska_gra_grodziskowa_test_nowa/', 'harcerska_gra_grodziskowa_test_nowa/', 0, 0, 0, 0, 0, 0, 0, 0, '12', 0, '', 0, '', 0, '', 0),
(188, 38, 'asd', 0, '', '', '', '', '12', 'harcerska_gra_grodziskowa_test_nowa/GM_PERFEKT_1.png', 'harcerska_gra_grodziskowa_test_nowa/GM_PERFEKT.png', 'harcerska_gra_grodziskowa_test_nowa/', 'harcerska_gra_grodziskowa_test_nowa/', 0, 0, 0, 0, 0, 0, 0, 0, '12', 0, '', 0, '', 0, '', 0),
(189, 38, 'asd', 0, '', '', '', '', '12', 'harcerska_gra_grodziskowa_test_nowa/GM_PERFEKT_1.png', 'harcerska_gra_grodziskowa_test_nowa/GM_PERFEKT.png', 'harcerska_gra_grodziskowa_test_nowa/', 'harcerska_gra_grodziskowa_test_nowa/', 0, 0, 0, 0, 0, 0, 0, 0, '12', 0, '', 0, '', 0, '', 0),
(190, 38, 'asd', 0, '', '', '', '', '12', 'harcerska_gra_grodziskowa_test_nowa/GM_PERFEKT_1.png', 'harcerska_gra_grodziskowa_test_nowa/GM_PERFEKT.png', 'harcerska_gra_grodziskowa_test_nowa/', 'harcerska_gra_grodziskowa_test_nowa/', 0, 0, 0, 0, 0, 0, 0, 0, '12', 0, '', 0, '', 0, '', 0),
(191, 41, 'asd', 0, '', '', '', '', '12', '12/GM_PERFEKT_1.png', '12/GM_PERFEKT.png', '12/', '12/', 0, 0, 0, 0, 0, 0, 0, 0, '12', 0, '', 0, '', 0, '', 0),
(192, 41, 'asd', 0, '', '', '', '', '12', '12/GM_PERFEKT_1.png', '12/GM_PERFEKT.png', '12/', '12/', 0, 0, 0, 0, 0, 0, 0, 0, '12', 0, '', 0, '', 0, '', 0),
(193, 41, 'asd', 0, '', '', '', '', '12', '12/GM_PERFEKT_1.png', '12/GM_PERFEKT.png', '12/', '12/', 0, 0, 0, 0, 0, 0, 0, 0, '12', 0, '', 0, '', 0, '', 0),
(194, 41, 'asd', 0, '', '', '', '', '12', '12/GM_PERFEKT_1.png', '12/GM_PERFEKT.png', '12/', '12/', 0, 0, 0, 0, 0, 0, 0, 0, '12', 0, '', 0, '', 0, '', 0),
(195, 41, 'asd', 0, '', '', '', '', '12', '12/GM_PERFEKT_1.png', '12/GM_PERFEKT.png', '12/', '12/', 0, 0, 0, 0, 0, 0, 0, 0, '12', 0, '', 0, '', 0, '', 0),
(196, 41, 'asd', 0, '', '', '', '', '12', '12/GM_PERFEKT_1.png', '12/GM_PERFEKT.png', '12/', '12/', 0, 0, 0, 0, 0, 0, 0, 0, '12', 0, '', 0, '', 0, '', 0),
(197, 41, 'asd', 0, '', '', '', '', '12', '12/GM_PERFEKT_1.png', '12/GM_PERFEKT.png', '12/', '12/', 0, 0, 0, 0, 0, 0, 0, 0, '12', 0, '', 0, '', 0, '', 0),
(198, 41, 'asd', 0, '', '', '', '', '12', '12/GM_PERFEKT_1.png', '12/GM_PERFEKT.png', '12/', '12/', 0, 0, 0, 0, 0, 0, 0, 0, '12', 0, '', 0, '', 0, '', 0),
(199, 41, 'asd', 0, '', '', '', '', '12', '12/GM_PERFEKT_1.png', '12/GM_PERFEKT.png', '12/', '12/', 0, 0, 0, 0, 0, 0, 0, 0, '12', 0, '', 0, '', 0, '', 0),
(200, 41, 'asd 123', 0, '', '', '', '', '12', '12/GM_PERFEKT_1.png', '12/GM_PERFEKT.png', '12/', '12/', 0, 0, 0, 0, 0, 0, 0, 0, '12', 0, '', 0, '', 0, '', 0),
(201, 41, 'asd 123', 0, '', '', '', '', '12', '12/GM_PERFEKT_1.png', '12/GM_PERFEKT.png', '12/', '12/', 0, 0, 0, 0, 0, 0, 0, 0, '12', 0, '', 0, '', 0, '', 0),
(202, 41, 'asd 123', 0, '', '', '', '', '12', '12/GM_PERFEKT_1.png', '12/GM_PERFEKT.png', '12/', '12/', 0, 0, 0, 0, 0, 0, 0, 0, '12', 0, '', 0, '', 0, '', 0),
(203, 41, 'asd 123', 0, '', '', '', '', '12', '12/GM_PERFEKT_1.png', '12/GM_PERFEKT.png', '12/', '12/', 0, 0, 0, 0, 0, 0, 0, 0, '12', 0, '', 0, '', 0, '', 0),
(204, 41, 'asd 123', 0, '', '', '', '', '12', '12/GM_PERFEKT_1.png', '12/GM_PERFEKT.png', '12/', '12/', 0, 0, 0, 0, 0, 0, 0, 0, '12', 0, '', 0, '', 0, '', 0),
(205, 41, 'asd 123', 0, '', '', '', '', '12', '12/GM_PERFEKT_1.png', '12/GM_PERFEKT.png', '12/', '12/', 0, 0, 0, 0, 0, 0, 0, 0, '12', 0, '', 0, '', 0, '', 0),
(208, 42, 'test', 0, '', '', '', '', 'dobra odpowiedz', 'nowy_quiz_do_testow_123_456_cos_cos/GJB2.png', 'nowy_quiz_do_testow_123_456_cos_cos/GM_PERFEKT_3 (2).png', 'nowy_quiz_do_testow_123_456_cos_cos/colorbluje.png', 'nowy_quiz_do_testow_123_456_cos_cos/004-pictogram-free.jpg', 0, 0, 0, 0, 0, 0, 0, 0, 'aasasd', 0, '', 0, '', 0, '', 0),
(220, 0, '', 0, '', '', '', '', '', '/', '/', '/', '/', 0, 0, 0, 0, 0, 0, 0, 0, '', 0, '', 0, '', 0, '', 0),
(221, 0, '', 0, '', '', '', '', '', '/', '/', '/', '/', 0, 0, 0, 0, 0, 0, 0, 0, '', 0, '', 0, '', 0, '', 0),
(224, 0, 'pytnia 1234', 20, 'link 14', 'link 24', 'link 34', 'link 44', 'poprawne odpowiedzi 123 4124 edited', 'nowy_quiz_testowy_22222/image_2.png', 'nowy_quiz_testowy_22222/image_1.png', 'nowy_quiz_testowy_22222/image_3.jpg', 'nowy_quiz_testowy_22222/prawe_polano_Naramienik.jpg', 14, 24, 34, 44, 14, 24, 34, 44, 'quizTitle 15', 14, 'quizTitle 25', 24, 'quizTitle 35', 34, 'quizTitle 45', 44),
(225, 0, 'pytnia 123 4', 124, 'link 14', 'link 24', 'link 34', 'link 44', 'poprawne odpowiedzi 123 4124 44444444444', 'nowy_quiz_testowy_22222/image_2.png', 'nowy_quiz_testowy_22222/image_1.png', 'nowy_quiz_testowy_22222/image_3.jpg', 'nowy_quiz_testowy_22222/prawe_polano_Naramienik.jpg', 14, 24, 34, 44, 14, 24, 34, 44, 'quizTitle 14', 14, 'quizTitle 24', 24, 'quizTitle 34', 34, 'quizTitle 44', 44);

-- --------------------------------------------------------

--
-- Struktura tabeli dla tabeli `extraQuiz`
--

CREATE TABLE `extraQuiz` (
  `id` int(11) NOT NULL,
  `poziom` varchar(500) NOT NULL,
  `czas` varchar(500) NOT NULL,
  `ilo` varchar(500) NOT NULL,
  `data` varchar(500) NOT NULL,
  `aktualizacja` varchar(500) NOT NULL,
  `autor` varchar(500) NOT NULL,
  `title` varchar(500) NOT NULL,
  `opis` varchar(500) NOT NULL,
  `graj` int(11) NOT NULL,
  `src` varchar(500) NOT NULL,
  `quiz_photo` varchar(100) NOT NULL,
  `quiz_time` int(11) NOT NULL,
  `timeOn` tinyint(1) NOT NULL DEFAULT 0,
  `isReady` tinyint(1) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_polish_ci;

--
-- Dumping data for table `extraQuiz`
--

INSERT INTO `extraQuiz` (`id`, `poziom`, `czas`, `ilo`, `data`, `aktualizacja`, `autor`, `title`, `opis`, `graj`, `src`, `quiz_photo`, `quiz_time`, `timeOn`, `isReady`) VALUES
(32, 'Łatwy', '120', '5', '20.12.2019', '20.12.2019', 'Paweł \"nCor3\"  Cyngot ', 'RN-2019-BetaTesty', 'Rajda Niepodległości 2019  BetaTesty ', 58, '4kot.jpg', '', 60, 0, 0),
(35, 'Trudny', '150', '25', '20.12.2019', '07.03.2021', 'Paweł \"nCor3\"  Cyngot ', 'Harcerska Grodziska Gra Terenowa', 'Zapraszam Czas trwania 2-3,5h. Trasa w całości wykonana przez Harc-Quiz  \nautor zdjęć Paweł Cyngot ', 136, 'harcerska_grodziska_gra_terenowa/01WG.png', '', 999, 0, 0);

-- --------------------------------------------------------

--
-- Struktura tabeli dla tabeli `news`
--

CREATE TABLE `news` (
  `id` int(11) NOT NULL,
  `title` text CHARACTER SET utf8 COLLATE utf8_polish_ci NOT NULL,
  `structure` longtext CHARACTER SET utf8 COLLATE utf8_polish_ci NOT NULL,
  `isReady` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `news`
--

INSERT INTO `news` (`id`, `title`, `structure`, `isReady`) VALUES
(26, 'Nowa Aktualizacja ! - Aktualności', '<!DOCTYPE html>\n<html>\n<head>\n</head>\n<body>\n<p>...</p>\n</body>\n</html>', 1),
(27, 'Projekty Graficzne - Zlecenia ( Strefa klienta )', '<!DOCTYPE html>\n<html>\n<head>\n</head>\n<body>\n<h2><strong><span style=\"color: #ffffff;\">&nbsp;Projekty Graficzne - Zlecenia (</span><span style=\"color: #ffd700;\"> Strefa klienta </span><span style=\"color: #ffffff;\">)</span></strong></h2>\n<h3><strong><span style=\"color: #f7f2f2;\">&nbsp;Cennik&nbsp;</span></strong><span style=\"background-color: #ffffff; color: #ff0000;\">(Dona</span><span style=\"background-color: #ffffff; color: #ff0000;\">te)</span><strong><span style=\"color: #f7f2f2;\">:</span></strong></h3>\n<ul>\n<li><span style=\"color: #f7f2f2;\"><strong>25 zł + Dyplomy&nbsp;</strong><strong>[ Grafika rastowa]</strong></span></li>\n<li><span style=\"color: #f7f2f2;\"><strong><strong>35 zł +&nbsp;</strong><strong>Plakietka&nbsp;</strong><strong>do&nbsp;</strong><strong>[</strong><strong>&nbsp;Sublimacja ]</strong><strong>&nbsp;</strong><strong>[</strong><strong>&nbsp;Druk ] [ Grafika rastowa]</strong></strong></span></li>\n<li><span style=\"color: #f7f2f2;\"><strong><strong><strong>40 zł + Plakietka do&nbsp;</strong><strong>[</strong><strong>&nbsp;Haft ] [ Grafika rastowa]</strong></strong></strong></span></li>\n<li><span style=\"color: #f7f2f2;\"><strong>50 zł + Inne projekty&nbsp;</strong><strong>[ Grafika rastowa]</strong><strong>[ Graf</strong><strong>ika Wektorowa]</strong></span></li>\n<li><span style=\"color: #f7f2f2;\"><strong>0&nbsp; &nbsp;zł + Plakietki , Dyplomy itp&nbsp;</strong></span><strong style=\"color: #f7f2f2;\">[ Grafika rastowa]</strong><strong style=\"color: #f7f2f2;\">[ Graf</strong><strong style=\"color: #f7f2f2;\">ika Wektorowa]</strong></li>\n</ul>\n<ul>\n<li><span style=\"background-color: #ffffff;\"><em><span style=\"color: #ff0000;\">Kwoty są to dobrowolne donate wsperające rozw&oacute;j strony oraz jej utrzymanie&nbsp;</span></em></span></li>\n</ul>\n<p>&nbsp;</p>\n<pre><strong><span style=\"color: #ffd700;\">&nbsp;<em>Portfolio:</em></span></strong></pre>\n<p>&nbsp;</p>\n<p><img src=\"http://harc-quiz.pl/news_images/9986f890-01f4-4266-b27a-7e7846d0f30e/Projekt_28 (1).png\" alt=\"\" width=\"196\" height=\"196\" /><img src=\"http://harc-quiz.pl/news_images/9986f890-01f4-4266-b27a-7e7846d0f30e/Wersja_03_Gradient.png\" alt=\"\" width=\"202\" height=\"202\" /><img src=\"http://harc-quiz.pl/news_images/9986f890-01f4-4266-b27a-7e7846d0f30e/graficzkalesne.png\" alt=\"\" width=\"205\" height=\"205\" /></p>\n<p style=\"padding-left: 30px;\"><img src=\"http://harc-quiz.pl/news_images/9986f890-01f4-4266-b27a-7e7846d0f30e/imagetools1.png\" alt=\"\" width=\"164\" height=\"164\" /><img src=\"http://harc-quiz.pl/news_images/9986f890-01f4-4266-b27a-7e7846d0f30e/Oboz_01.png\" alt=\"\" width=\"200\" height=\"200\" /><img src=\"http://harc-quiz.pl/news_images/9986f890-01f4-4266-b27a-7e7846d0f30e/iniemamocni.png\" alt=\"\" width=\"181\" height=\"181\" /></p>\n<pre><strong><span style=\"color: #ffd700;\">&nbsp;<em>Programy graficzne:<br /><br /></em></span></strong></pre>\n<ul>\n<li>\n<pre><span style=\"color: #ffffff;\"><strong>Gimp 2&nbsp;-</strong><strong> Grafika rastowa</strong></span></pre>\n</li>\n<li>\n<pre><span style=\"color: #ffffff;\"><strong><strong>Photoshop - Grafika rastowa</strong></strong></span></pre>\n</li>\n<li>\n<pre><span style=\"color: #ffffff;\"><strong>Affinity Photo - Grafika rastowa</strong></span></pre>\n</li>\n<li>\n<pre><span style=\"color: #ffffff;\"><strong>Affinity Designer&nbsp;-</strong><strong> Grafika Wektorowa</strong></span></pre>\n</li>\n<li>\n<pre><span style=\"color: #ffffff;\"><strong>Inkscape&nbsp;- Grafika Wektorowa</strong></span></pre>\n</li>\n<li>\n<pre><span style=\"color: #ffffff;\"><strong>Ilustrator - Grafika Wektorowa</strong></span></pre>\n</li>\n</ul>\n</body>\n</html>', 1),
(28, 'HarcMonopoly', 'testowy xd', 1),
(29, 'HarcMonopoly', '<!DOCTYPE html>\n<html>\n<head>\n</head>\n<body>\n<p>&nbsp;</p>\n<h1><em><span style=\"color: #ffcc00;\"><strong>Aktualności z 2021-10-17!</strong></span></em><img style=\"display: block; margin-left: auto; margin-right: auto;\" src=\"../news_images/cbc8d51a-ea1c-44b5-9255-9f2df2e41186/blobid2.png\" alt=\"\" width=\"702\" height=\"501\" /></h1>\n<p style=\"text-align: center;\"><span style=\"color: #ffcc00;\"><strong>ETAP I Aktualny postęp: 100/100%</strong></span></p>\n<p style=\"text-align: center;\"><strong><span style=\"color: #ffffff;\"><span style=\"color: #f0e9e9;\"><span style=\"color: #f5f2f2;\">- Karty</span>&nbsp;</span> <span style=\"color: #339966;\">✔</span></span></strong></p>\n<p style=\"text-align: center;\"><strong><span style=\"color: #ffffff;\"><span style=\"color: #f5f2f2;\">- Plansza</span> <span style=\"color: #339966;\">✔</span></span></strong></p>\n<p style=\"text-align: center;\"><strong><span style=\"color: #ffffff;\"><span style=\"color: #f5f2f2;\">- Banknoty&nbsp;</span> <span style=\"color: #339966;\">✔</span></span></strong></p>\n<p style=\"text-align: center;\"><strong><span style=\"color: #f5f0f0;\">- Plansza rozwoju postaci</span>&nbsp; <span style=\"color: #ffffff;\"><span style=\"color: #339966;\">✔</span></span></strong></p>\n<p style=\"text-align: center;\"><strong><span style=\"color: #f5f0f0;\">- Karty rozwoju postaci</span>&nbsp;<span style=\"color: #ffffff;\"> <span style=\"color: #339966;\">✔</span></span></strong></p>\n<p style=\"text-align: center;\"><strong><span style=\"color: #ffffff;\"><span style=\"color: #f5f0f0;\">- 5 karty specjalnych p&oacute;l&nbsp; &nbsp;</span><span style=\"color: #339966;\">✔</span></span></strong></p>\n<p style=\"text-align: center;\"><strong><span style=\"color: #ffffff;\"><span style=\"color: #f5f0f0;\">- Instrukcja i zasady&nbsp;</span> <span style=\"color: #339966;\">✔</span></span></strong></p>\n<p style=\"text-align: center;\"><strong><span style=\"color: #ffffff;\"><span style=\"color: #f5f0f0;\">- Pudełko</span> <span style=\"color: #339966;\">✔</span></span></strong></p>\n<p style=\"text-align: center;\"><span style=\"color: #ffcc00;\"><strong>ETAP II Aktualny postęp: 38/100%</strong></span></p>\n<p style=\"text-align: center;\"><span style=\"color: #f5f0f0;\"><strong>(Zapisywanie w pdf do Druku )</strong></span></p>\n<p style=\"text-align: center;\"><strong><span style=\"color: #ffffff;\"><span style=\"color: #f5f0f0;\">- Karty&nbsp;</span> <span style=\"color: #339966;\">✔</span></span></strong></p>\n<p style=\"text-align: center;\"><strong><span style=\"color: #ffffff;\"><span style=\"color: #f5f0f0;\">- Plansza</span> <span style=\"color: #ff0000;\">✖</span></span></strong></p>\n<p style=\"text-align: center;\"><strong><span style=\"color: #ffffff;\"><span style=\"color: #f5f0f0;\">- Banknoty&nbsp;</span> <span style=\"color: #339966;\">✔</span></span></strong></p>\n<p style=\"text-align: center;\"><strong><span style=\"color: #ffffff;\"><span style=\"color: #f5f0f0;\">- Plansza rozwoju postaci&nbsp;</span> <span style=\"color: #ff0000;\">✖</span></span></strong></p>\n<p style=\"text-align: center;\"><strong><span style=\"color: #ffffff;\"><span style=\"color: #f5f0f0;\">- Karty rozwoju postaci&nbsp;</span> <span style=\"color: #ff0000;\">✖</span></span></strong></p>\n<p style=\"text-align: center;\"><strong><span style=\"color: #ffffff;\"><span style=\"color: #f5f0f0;\">- 5 karty specjalnych p&oacute;l</span> <span style=\"color: #ff0000;\">✖</span>&nbsp;&nbsp;</span></strong></p>\n<p style=\"text-align: center;\"><strong><span style=\"color: #ffffff;\"><span style=\"color: #f5f0f0;\">- Instrukcja i zasady&nbsp;</span> <span style=\"color: #339966;\">✔</span></span></strong></p>\n<p style=\"text-align: center;\"><strong><span style=\"color: #ffffff;\"><span style=\"color: #f5f0f0;\">- Pudełko</span> <span style=\"color: #339966;\">✔</span></span></strong></p>\n<p style=\"text-align: center;\"><span style=\"color: #ffcc00;\"><strong>ETAP III Aktualny postęp: 0/100%</strong></span></p>\n<p style=\"text-align: center;\"><span style=\"color: #f5f0f0;\"><strong>( Data 2021/2022 )</strong></span></p>\n<p style=\"text-align: center;\"><strong><span style=\"color: #ffffff;\"><span style=\"color: #f5f0f0;\">- Publikacja</span> <span style=\"color: #ff0000;\">✖</span></span></strong></p>\n<p style=\"text-align: center;\"><span style=\"color: #f5f0f0;\"><strong>Ps: ✎✎ Dziękuje osobom trzecim za drobną pomoc przy kartach&nbsp; ✐✐</strong></span></p>\n<p style=\"text-align: center;\"><span style=\"color: #f5f0f0;\"><strong>⚜ FAQ - Najczęściej zadawane pytania ! ⚜</strong></span></p>\n<p style=\"text-align: center;\"><span style=\"color: #f5f0f0;\"><strong>*</strong><strong>&nbsp;Kiedy zakończę projekty ?&nbsp;</strong></span></p>\n<p style=\"text-align: center;\"><span style=\"color: #f5f0f0;\"><strong>&nbsp; &nbsp; &nbsp;</strong><em><strong>~&nbsp;</strong><strong>&nbsp;Nie wiem wszystko zależy od mojego wolnego czasu&nbsp;</strong></em></span></p>\n<p style=\"text-align: center;\"><span style=\"color: #f5f0f0;\"><strong>*&nbsp;</strong><strong>&nbsp;Gdzie udostępnię pliki plansz&oacute;wki ?</strong></span></p>\n<p style=\"text-align: center;\"><em><span style=\"color: #f5f0f0;\"><strong>&nbsp; &nbsp;&nbsp;</strong><strong>~&nbsp;</strong><strong>&nbsp;Na swojej stronie i może na centralny bank pomysł&oacute;w zhp.</strong></span></em></p>\n<p style=\"text-align: center;\"><span style=\"color: #f5f0f0;\"><strong>*&nbsp;</strong><strong>&nbsp;Czy udostępnię projekty ?&nbsp;</strong></span></p>\n<p style=\"text-align: center;\"><span style=\"color: #f5f0f0;\"><strong>&nbsp; &nbsp; <em>&nbsp;</em></strong><em><strong>~</strong><strong>&nbsp;Tak ale będą 2 wersje darmowa i płatna.</strong></em></span></p>\n<p style=\"text-align: center;\"><span style=\"color: #f5f0f0;\"><strong>*&nbsp;</strong><strong>&nbsp;Jak długo robię projekt tej plansz&oacute;wki?&nbsp;</strong></span></p>\n<p style=\"text-align: center;\"><span style=\"color: #f5f0f0;\"><strong>&nbsp; &nbsp; <em>&nbsp;&nbsp;</em></strong><em><strong>~</strong><strong>&nbsp;</strong><strong>&nbsp;Projekt zacząłem w Styczniu 2019.</strong></em></span></p>\n', 1),
(30, 'Nowa Aktualizacja ! - Aktualności', '<!DOCTYPE html>\r\n<html>\r\n<head>\r\n</head>\r\n<body>\r\n<p>...</p>\r\n</body>\r\n</html>', 1);

-- --------------------------------------------------------

--
-- Struktura tabeli dla tabeli `pytania`
--

CREATE TABLE `pytania` (
  `question_id` int(11) NOT NULL,
  `quiz_id` int(11) NOT NULL,
  `question` mediumtext NOT NULL,
  `A` varchar(100) NOT NULL,
  `B` varchar(100) NOT NULL,
  `C` varchar(100) NOT NULL,
  `D` varchar(100) NOT NULL,
  `right_answer` varchar(1) NOT NULL,
  `photo` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_polish_ci;

--
-- Dumping data for table `pytania`
--

INSERT INTO `pytania` (`question_id`, `quiz_id`, `question`, `A`, `B`, `C`, `D`, `right_answer`, `photo`) VALUES
(197, 54, 'Co symbolizuje prawe polano?', 'Służbę ', 'praca nad sobą ', 'Wytrwałość ', 'Szukanie miejsca w społeczeństwie', 'B', 'symbolika_wedrowniczej_watry/prawe_polano_Naramienik.jpg'),
(198, 54, 'Co symbolizuje mały płomień ogniska?', 'Siła ciała', 'Siła rozumu', 'Siła ducha', 'Siła wytrwałości', 'A', 'symbolika_wedrowniczej_watry/maly_ogien_Naramienik.jpg'),
(199, 54, 'Co jest symbolem służby, szukania miejsca w społeczeństwie i pracy nad sobą?', 'Ogień ', 'Kamień', 'Zielone tło', 'Trzy polana ', 'D', 'symbolika_wedrowniczej_watry/naramienik_Wedrowniczy.jpg'),
(200, 54, 'Co symbolizuje środkowe polano?', 'praca nad sobą', 'Nadzieje ', 'Szukanie miejsca w społeczeństwie', 'Służbę', 'C', 'symbolika_wedrowniczej_watry/srodkowe_polano_Naramienik.jpg'),
(201, 54, 'Co symbolizuje duży płomień ogniska?', 'Siła ducha', 'Siła rozumu', 'Siła wytrwałości', 'Siła ciała', 'A', 'symbolika_wedrowniczej_watry/duzy_ogien_Naramienik.jpg'),
(202, 54, 'Gdzie nosimy naramiennik wędrowniczy?', 'Obojętnie', 'Na lewym ramieniu', 'Na prawym ramieniu', 'Na prawej kieszeni', 'B', 'symbolika_wedrowniczej_watry/naramienik_Wedrowniczy.jpg'),
(203, 54, 'Co symbolizuje lewe polano?', 'Męstwo', 'Szukanie miejsca w społeczeństwie ', 'Służbę ', 'praca nad sobą', 'C', 'symbolika_wedrowniczej_watry/lewe_polano_Naramienik.jpg'),
(204, 54, 'Co oznacza zielone tło?', 'Nie mam znaczenia ', 'Nadzieję i wolności', 'Las , puszczaństwo', 'sumienność', 'C', 'symbolika_wedrowniczej_watry/naramienik_Wedrowniczy.jpg'),
(205, 54, 'Co symbolizuje średni płomień ogniska?', 'Siła wytrwałości', 'Siła rozumu', 'Siła ciała', 'Siła ducha', 'B', 'symbolika_wedrowniczej_watry/sredni_ogien_Naramienik.jpg'),
(206, 56, 'Co oznacza skrót WAGGGS po polsku ?', 'Światowe Stowarzyszenie harcerek i Skautek', 'Światowe Stowarzyszenie Przewodniczek i Skautek', 'Światowe Stowarzyszenie Przewodniczek i Skautów', 'Światowe Stowarzyszenie Skautów i Skautek', 'B', 'symbolika_wagggs/waggs.jpg'),
(207, 56, 'Co symbolizuje przerwa w okręgu do ?\n', 'Otwartość świata', 'Nadzieje', 'Nic nie symbolizuję ', 'Wolne miejsce dla każdego ', 'D', 'symbolika_wagggs/waggs.jpg'),
(208, 56, 'Co symbolizuje kolor złoty i niebieski?', 'Najwyższe ideały i pokój na świecie ', 'Słońce świecące dzieciom na całym świecie ', 'Spokój i dobrobyt na świecie ', 'Uduchowienie i dobrobyt', 'B', 'symbolika_wagggs/waggs.jpg'),
(210, 56, 'Co symbolizuje łodyga koniczyny?', 'Płomień miłości międzyludzkiej', 'Zawiłą drogę harcerską ', 'Płomień harcerski ', 'Prostą drogę harcerską ', 'A', 'symbolika_wagggs/waggs.jpg'),
(211, 56, 'Co symbolizuje Żyłka koniczyny ?\n', 'Nic nie symbolizuje ', 'Symbol laski skautowej ', 'Igłę kompasu \"trudną droga\"', 'Igłę kompasu \"Dobra droga\"', 'D', 'symbolika_wagggs/waggs.jpg'),
(212, 56, '	Co symbolizuje okrąg dokoła koniczynki?\n\n', 'Symbol Męstwa', 'Symbol doskonałości', 'Braterstwo', 'Jedność na całym świecie ', 'B', 'symbolika_wagggs/waggs.jpg'),
(213, 56, 'Co symbolizują dwie gwiazdki?\n\n', 'Siłę i męstwo ', 'Dwie wojny światowe ', 'Prawo i Przyrzeczenie', 'Odwagę ', 'C', 'symbolika_wagggs/waggs.jpg'),
(214, 58, 'Co symbolizuje biały orzeł?', 'Waleczność i dzielność', 'Patriotyzm', 'Odwagę, męstwo i dzielność', 'Odwagę', 'C', 'symbolika_znaczka_zucha/zuch.jpg'),
(215, 58, 'Co symbolizuje słońce?', 'Radość', 'Wytrwałość', 'Dzielność ', 'Pogodę ducha', 'A', 'symbolika_znaczka_zucha/zuch.jpg'),
(216, 58, 'Co symbolizuje czyste niebo?', 'Spokój ducha', 'Radość', 'Spokój ', 'Pogodę ducha', 'D', 'symbolika_znaczka_zucha/zuch.jpg'),
(217, 58, 'Co symbolizuje Promienie słońca?', 'Gotowość niesienia pomocy, świecenie przykładem', 'Dążenie do doskonałości', 'Gotowość do działania', 'Gotowość niesienia pomocy', 'D', 'symbolika_znaczka_zucha/zuch.jpg'),
(218, 58, 'Co symbolizuje napis \"ZUCH\"', 'Przynależność do ZHP/ZHR', 'Przynależność Gromady', 'Bycie zuchem', 'Nic nie symbolizuje', 'A', 'symbolika_znaczka_zucha/zuch.jpg'),
(219, 57, 'Co oznacza skrót WOSM po polsku ?', 'Światowa Organizacja Ruchu Skautek', 'Światowa Organizacja Ruchu Skautowego', 'Światowa Organizacja Ruchu harcerskiego', 'Światowa Organizacja Ruchu skautów i skautek', 'B', 'symbolika_wosm/wosm.jpg'),
(220, 57, 'Co symbolizuje węzeł płaski?', 'Jedność ', 'Braterstwo ', 'Dążenie do doskonałości ', 'Jedność i braterstwo ', 'D', 'symbolika_wosm/wosm.jpg'),
(221, 57, 'Co symbolizuje biały kolor?', 'Przywództwo', 'Czystość', 'Spokój', 'Niewinność', 'B', 'symbolika_wosm/wosm.jpg'),
(222, 57, 'Co symbolizują trzy ramiona lilijki?', 'Dziś, jutro, pojutrze ', 'Trzech założycieli WAGGGS', 'Obowiązki wobec boga ,bliźnich i siebie ', 'Służbę Bogu, nauce i ojczyźnie', 'C', 'symbolika_wosm/wosm.jpg'),
(223, 57, 'Co symbolizuje purpurowy kolor?', 'Odpowiedzialność', 'Spokój', 'Duchowieństwo', 'Czystość ', 'A', 'symbolika_wosm/wosm.jpg'),
(224, 57, 'Co symbolizuje lina?', 'Symbol Męstwa', 'Jedność ludzi na świecie', 'Symbol doskonałości', 'Braterstwo ', 'D', 'symbolika_wosm/wosm.jpg'),
(225, 57, 'Co symbolizują dwie gwiazdki?', 'Prawo i przyrzeczenie ', 'Odwagę ', 'Siłę i męstwo', 'Dwie wojny światowe ', 'A', 'symbolika_wosm/wosm.jpg'),
(226, 55, 'Na czym wzrowany był krzyż harcerski?', 'Orderem Odrodzenia Polski ', 'Krzyżem Walecznych', 'Orderem Orła Białego', 'Orderem Wojenny Virtuti Militari ', 'D', 'symbolika_krzyza_harcerskiego/krzyz.jpg'),
(227, 55, 'Co symbolizują puste miejsca między ziarnkami piasku?', 'Miejsce dla nowych harcerzy', 'Miejsce dla nowych Instruktorów', 'Harcerzy co odeślij ', 'Ogromną ilość harcerzy', 'A', 'symbolika_krzyza_harcerskiego/krzyz.jpg'),
(228, 55, 'Co symbolizują promienie rozchodzące się od lilijki?', 'Nadzieję', 'Waleczność ', 'Promieniowanie harcerskimi wartościami na świat', 'Harcerskie ideały ', 'C', 'symbolika_krzyza_harcerskiego/krzyz.jpg'),
(229, 55, 'Co symbolizują trzy żołędzie na wieńcach dębowych?', 'Trzy wojny', 'Trzy rozbiory', 'Trzy pokolenia harcerzy  ', 'Nic nie symbolizuje ', 'B', 'symbolika_krzyza_harcerskiego/krzyz.jpg'),
(230, 55, 'Co symbolizują cztery ramiona krzyża ?', 'Ciężką drogę oraz gotowość do waliki i poświecenia \"Cztery strony świata\"', 'Cztery  żywioły', 'Cztery metodyki ', 'Róże wiatrów ', 'A', 'symbolika_krzyza_harcerskiego/krzyz.jpg'),
(231, 55, 'Co symbolizuje wieniec z liści dębowych?', 'Waleczność', 'Zwycięstwa', 'Siłę, męstwo i odwagę', 'Siłę i męstwo', 'C', 'symbolika_krzyza_harcerskiego/krzyz.jpg'),
(232, 55, 'Co symbolizuje hasło \"Czuwaj\"?', 'Bycie czujnym', 'Pozdrowienie harcerskie', 'Służbę', 'Stałą gotowości do służby', 'D', 'symbolika_krzyza_harcerskiego/krzyz.jpg'),
(233, 55, 'Co symbolizują ziarenka piasku na ramionach?', 'Ogromną ilość instruktorów ', 'Ogromną ilość harcerzy', 'Nic nie symbolizują ', 'Ilość drużyn', 'B', 'symbolika_krzyza_harcerskiego/krzyz.jpg'),
(234, 55, 'Co symbolizują dwie gwiazdki na ramionach lilijki?', 'Harcerzy i harcerki', 'Dwie wojny światowe', 'Prawo i przyrzeczenie', 'Odwagę i męstwo', 'C', 'symbolika_krzyza_harcerskiego/krzyz.jpg'),
(235, 55, 'Co symbolizuje wieniec z liśćmi laurowymi?', 'Zwycięstwo ', 'Nadzieje', 'Uczciwość', 'Siłę, męstwo i odwagę', 'A', 'symbolika_krzyza_harcerskiego/krzyz.jpg'),
(236, 55, 'Co symbolizują dwie oliwki na wieńcu laurowym?', 'Nic nie symbolizuje ', 'Dwie wojny światowe', 'Harcerki i harcerzy', 'Prawo i przyrzeczenie ', 'B', 'symbolika_krzyza_harcerskiego/krzyz.jpg'),
(241, 59, 'Co symbolizuje biały kolor na fladze ?', 'Niewinności i spokój', 'Niewinności i szczerość ', 'Czystości i niewinności', 'Czystość i szczerość ', 'D', 'symbolika_flagi_oraz_herbu_polski/falga_pl.jpg'),
(242, 59, 'Co symbolizuje czerwony kolor na fladze ?', 'Krew poległych w obronie Ojczyzny i małość do niej', 'Siłę i potęgę Ojczyzny ', 'Miłość do Ojczyzny ', 'Siłę i gniew', 'A', 'symbolika_flagi_oraz_herbu_polski/falga_pl.jpg'),
(243, 59, 'Co symbolizuje biały orzeł na godle ?', 'Szlachetność', 'Niezależność państwa od innych krajów ', 'Niepodległość ', 'Czystość i szczerość ', 'B', 'symbolika_flagi_oraz_herbu_polski/falga_herb_pl.jpg'),
(244, 59, 'Co symbolizuje korona na głowie orła?', 'Przywództwo ', 'Symbol królewski  i Przywództwo ', 'Symbol królewski oraz Niepodległość', 'Niepodległość', 'C', 'symbolika_flagi_oraz_herbu_polski/falga_herb_pl.jpg'),
(245, 59, 'Co symbolizuje czerwień na godle?', 'Krew poległych w obronie Ojczyzny i małość do niej', 'Waleczność i Ogień', 'Miłość do Ojczyzny ', 'Ogień, odwaga oraz waleczność', 'D', 'symbolika_flagi_oraz_herbu_polski/falga_herb_pl.jpg'),
(246, 59, 'Co symbolizuje biel na godle?', 'Niezależność państwa od innych krajów ', 'Srebro, wodę. duchownie-czystość i niepokalanie ', 'Miłość do Ojczyzny ', 'Niepodległość ', 'B', 'symbolika_flagi_oraz_herbu_polski/falga_herb_pl.jpg'),
(263, 68, 'Co oznacza ten znak patrolowy?', 'Iść', 'Biec', 'Iść szybko', 'iść wolno', 'A', 'znaki_patrolowe_cz__1_(podstawowe)/isc.jpg'),
(264, 68, 'Co oznacza ten znak patrolowy?', 'Czekać 30 min', 'Czekać 10 min', 'Czekać 20 min', 'Czekać 15 min', 'D', 'znaki_patrolowe_cz__1_(podstawowe)/czekac_15min.jpg'),
(265, 68, 'Co oznacza ten znak patrolowy?', 'Spotkamy się tu za 1 godzinę', 'Nie iść tędy', 'Iść ostrożnie ', 'Podzieliliśmy się', 'C', 'znaki_patrolowe_cz__1_(podstawowe)/isc_ostroznie.jpg'),
(266, 68, 'Co oznacza ten znak patrolowy?', 'List 4 kroki stąd ', '4 osoby poszły w tą stronę ', 'Obóz 4 kroki stąd ', '4 listy ukryte w tym kierunku ', 'A', 'znaki_patrolowe_cz__1_(podstawowe)/list_4_kroki_stad.jpg'),
(267, 68, 'Co oznacza ten znak patrolowy?', 'Iść', 'Iść szybko', 'Wracać ', 'Biec', 'B', 'znaki_patrolowe_cz__1_(podstawowe)/isc_szybko.jpg'),
(268, 68, 'Co oznacza ten znak patrolowy?', 'Wróg', 'Zła droga', 'Niebezpieczeństwo', 'Uciekać ', 'C', 'znaki_patrolowe_cz__1_(podstawowe)/niebezpieczenstwo.jpg'),
(269, 68, 'Co oznacza ten znak patrolowy?', 'Czekać 10 min', 'Czekać 15 min', 'Czekać 1 min', 'Czekać 5 min', 'D', 'znaki_patrolowe_cz__1_(podstawowe)/czekac_5min.jpg'),
(270, 68, 'Co oznacza ten znak patrolowy?', 'Niebezpieczeństwo ', 'Wróg ', 'Zła droga', 'Zawracać ', 'C', 'znaki_patrolowe_cz__1_(podstawowe)/zla_droga.jpg'),
(271, 68, 'Co oznacza ten znak patrolowy?', 'Nie iść tędy', 'Spotkamy się tu za 2 godziny', 'Szukać innej drogi', 'Podzieliliśmy się', 'A', 'znaki_patrolowe_cz__1_(podstawowe)/nie_isc_tendy.jpg'),
(272, 68, 'Co oznacza ten znak patrolowy?', 'Iść szybko', 'Biec', 'Iść ', 'Wracać Biegiem', 'B', 'znaki_patrolowe_cz__1_(podstawowe)/biec.jpg'),
(273, 68, 'Co oznacza ten znak patrolowy?', 'Niebezpieczna woda', 'Woda pitna', 'Woda niepitna', 'Zakaz kąpieli ', 'C', 'znaki_patrolowe_cz__1_(podstawowe)/woda_niepitna.jpg'),
(274, 68, 'Co oznacza ten znak patrolowy?', 'Poszukać innej drogi', 'Wracać Szybko', 'Rozdzielić się ', 'Wracać ', 'D', 'znaki_patrolowe_cz__1_(podstawowe)/wracac.jpg'),
(275, 68, 'Co oznacza ten znak patrolowy?', 'Wróg ', 'Nieprzyjaciel', 'Niebezpieczeństwo ', 'Przyjaciel', 'A', 'znaki_patrolowe_cz__1_(podstawowe)/wrog.jpg'),
(276, 68, 'Co oznacza ten znak patrolowy?', 'Iść szybko', 'Wracać szybko', 'Wracać', 'Wracać biegiem', 'B', 'znaki_patrolowe_cz__1_(podstawowe)/wracac_szybko.jpg'),
(277, 68, 'Co oznacza ten znak patrolowy?', 'Wracać', 'Szukać innej drogi', 'Rozwidlenie dróg ', 'Podzieliliśmy się ', 'D', 'znaki_patrolowe_cz__1_(podstawowe)/podzielilismy_sie.jpg'),
(278, 68, 'Co oznacza ten znak patrolowy?', 'Góry w kierunku ', 'List w kierunku ', 'Obóz w kierunku ', 'Namiot w kierunku ', 'C', 'znaki_patrolowe_cz__1_(podstawowe)/oboz_w_kierunku.jpg'),
(279, 68, 'Co oznacza ten znak patrolowy?', 'Szukać innej drogi', 'Rozwidlenie dróg ', 'Zawracać', 'Podzieliliśmy się', 'A', 'znaki_patrolowe_cz__1_(podstawowe)/szukac_innej_drogi.png'),
(280, 68, 'Co oznacza ten znak patrolowy?', 'Woda nie pitna', 'Niebezpieczna woda ', 'Woda pitna', 'Kąpielisko ', 'C', 'znaki_patrolowe_cz__1_(podstawowe)/woda_pitna.jpg'),
(281, 68, 'Co oznacza ten znak patrolowy?', 'Spotkamy się tu za 1 godzinę ', 'nie iść tędy', 'Spotkamy się tu za 10 min', 'Iść ostrożnie ', 'A', 'znaki_patrolowe_cz__1_(podstawowe)/spotkamy_sie_tu_za_1h.jpg'),
(282, 68, 'Co oznacza ten znak patrolowy?', 'Wracać w podskokach ', 'Wracać szybko', 'Iść biegiem', 'Wracać biegiem', 'D', 'znaki_patrolowe_cz__1_(podstawowe)/wracac_biegiem.jpg'),
(283, 60, 'Co oznacza ten znak patrolowy?', 'Wracać', 'Zła droga', 'Podzieliliśmy się', 'Wracać szybko ', 'A', 'znaki_patrolowe_cz__2/wracac.jpg'),
(284, 60, 'Co oznacza ten znak patrolowy?', 'Czekać 15 min', 'Czekać 1 godzinne ', 'Czekać 10 min', 'Czekać 5 min', 'C', 'znaki_patrolowe_cz__2/czekac_5min.jpg'),
(285, 60, 'Co oznacza ten znak patrolowy?', 'Wysokie Fale', 'Świeża woda ', 'Wybrzeże ', 'Rzeka', 'B', 'znaki_patrolowe_cz__2/swieza_woda.jpg'),
(286, 60, 'Co oznacza ten znak patrolowy?', 'Gromada', 'Drużyna ', 'Zastęp', 'Obóz ', 'C', 'znaki_patrolowe_cz__2/zastep.jpg'),
(287, 60, 'Co oznacza ten znak patrolowy?', 'Iść do dołu ', 'Obniżenie terenu ', 'Uwaga dziura ', 'Iść do góry', 'A', 'znaki_patrolowe_cz__2/iss_do_dolu.jpg'),
(289, 60, 'Co oznacza ten znak patrolowy?', 'Śnieg', 'Burza', 'Pogoda', 'Deszcz', 'D', 'znaki_patrolowe_cz__2/deszcz.jpg'),
(290, 60, 'Co oznacza ten znak patrolowy?', 'Wracać ', 'Iść', 'Biec', 'Iść szybko', 'B', 'znaki_patrolowe_cz__2/isc.jpg'),
(291, 60, 'Co oznacza ten znak patrolowy?', 'Przyjaciel', 'Wojna', 'Pokój', 'Można zrywać rośliny', 'C', 'znaki_patrolowe_cz__2/pokoj.jpg'),
(292, 60, 'Co oznacza ten znak patrolowy?', 'Woda nie do przejścia', 'Most', 'Rzeka', 'Woda do przejścia', 'A', 'znaki_patrolowe_cz__2/woda_nie_do_przejscia.jpg'),
(293, 60, 'Co oznacza ten znak patrolowy?', 'Czekać 15 min', 'Czekać 20 min', 'Czekać 10 min', 'Czekać 30 min', 'D', 'znaki_patrolowe_cz__2/czekac_30min.jpg'),
(294, 60, 'Co oznacza ten znak patrolowy?', 'Dzień', 'Wschód ', 'Zachód ', 'Noc', 'C', 'znaki_patrolowe_cz__2/zachod.jpg'),
(295, 60, 'Co oznacza ten znak patrolowy?', 'List w kierunku', 'Obóz w kierunku', 'Harcerze w kierunku', 'Góry w kierunku', 'B', 'znaki_patrolowe_cz__2/oboz_w_kierunku.jpg'),
(296, 60, 'Co oznacza ten znak patrolowy?', 'Spotkamy się za 2 godziny', 'Zła droga ', 'Nie iść tędy ', 'Iść ostrożnie ', 'C', 'znaki_patrolowe_cz__2/nie_isc_tendy.jpg'),
(297, 60, 'Co oznacza ten znak patrolowy?', 'Zatrzymać się tutaj', 'Nie', 'Tak', 'Ukryj się tutaj', 'D', 'znaki_patrolowe_cz__2/ukryj_sie_tutaj.jpg'),
(298, 60, 'Co oznacza ten znak patrolowy?', 'Woda pitna', 'Rzeka', 'Kąpielisko ', 'Woda nie pitna', 'A', 'znaki_patrolowe_cz__2/woda_pitna.jpg'),
(299, 60, 'Co oznacza ten znak patrolowy?', 'Doliny ', 'Góry', 'Wzgórza', 'Wzgórza , doliny ', 'D', 'znaki_patrolowe_cz__2/wzgoza.doliny.jpg'),
(300, 60, 'Co oznacza ten znak patrolowy?', 'Wróg', 'Znalazłem , cieszę się , zwycierzyłem', 'Przyjaciel', 'Tu byli harcerze', 'B', 'znaki_patrolowe_cz__2/znalazlem_ciesze_sie_zwyciezylem.jpg'),
(301, 61, 'Co oznacza ten znak patrolowy?', 'Iść', 'Biec', 'Iść szybko', 'Wracać', 'B', 'znaki_patrolowe_cz__3/biec.jpg'),
(302, 61, 'Co oznacza ten znak patrolowy?', 'Deszcz', 'Pogoda', 'Burza', 'Śnieg', 'C', 'znaki_patrolowe_cz__3/burza.jpg'),
(303, 61, 'Co oznacza ten znak patrolowy?', 'Czekać 15 min', 'Czekać 20 min', 'Czekać 30 min', 'Czekać 10 min', 'D', 'znaki_patrolowe_cz__3/czekac_10min.jpg'),
(304, 61, 'Co oznacza ten znak patrolowy?', 'Drużyna', 'Obóz', 'Zastęp', 'Ukryta wiadomość', 'A', 'znaki_patrolowe_cz__3/druzyna.jpg'),
(305, 61, 'Co oznacza ten znak patrolowy?', 'Spotkamy się za 2 godziny', 'Spotkamy się za 1 godzinę', 'Nie iść tędy', 'Iść ostrożnie ', 'D', 'znaki_patrolowe_cz__3/isc_ostroznie.jpg'),
(306, 61, 'Co oznacza ten znak patrolowy?', 'Zachód ', 'Iść do góry', 'Wschód ', 'Iść w dół', 'B', 'znaki_patrolowe_cz__3/iss_do_gory.jpg'),
(307, 61, 'Co oznacza ten znak patrolowy?', 'Podzieliliśmy się ', 'Las liściasty', 'Biec', 'Las', 'B', 'znaki_patrolowe_cz__3/las.jpg'),
(308, 61, 'Co oznacza ten znak patrolowy?', 'Iść 4 kroki', '4 listy ukryte w tym kierunku ', 'List 4 kroki stąd ', '4 osoby poszły w tą stronę ', 'C', 'znaki_patrolowe_cz__3/list_4_kroki_stad.jpg'),
(309, 61, 'Co oznacza ten znak patrolowy?', 'Tak', 'Czekać 1 godzine', 'Nie', 'Zatrzymać się', 'C', 'znaki_patrolowe_cz__3/nie.jpg'),
(310, 61, 'Co oznacza ten znak patrolowy?', 'Zawracać', 'Zła droga', 'Wróg', 'Niebezpieczeństwo', 'D', 'znaki_patrolowe_cz__3/niebezpieczenstwo.jpg'),
(311, 61, 'Co oznacza ten znak patrolowy?', 'Dzień', 'Połódnie', 'Pogoda', 'Słońce', 'B', 'znaki_patrolowe_cz__3/polodnie.jpg'),
(312, 61, 'Co oznacza ten znak patrolowy?', 'Iść ostrożnie', 'Na pomoc', 'Spotkamy się tu za 2 godziny', 'Nie iść tędy ', 'C', 'znaki_patrolowe_cz__3/spotkamy_sie_tu_za_2h.jpg'),
(313, 61, 'Co oznacza ten znak patrolowy?', 'Szukać', 'Wracać', 'Rozdzieliliśmy się', 'Szukać innej drogi', 'D', 'znaki_patrolowe_cz__3/szykac_innej_drogi.png'),
(314, 61, 'Co oznacza ten znak patrolowy?', 'Świeża woda', 'Woda nie do przejścia', 'Woda podziemna', 'Most na rzeką', 'C', 'znaki_patrolowe_cz__3/woda_podziemna.jpg'),
(315, 61, 'Co oznacza ten znak patrolowy?', 'Czekać 20 min', 'Czekać 10 min', 'Czekać 15 min', 'Czekać 30 min', 'A', 'znaki_patrolowe_cz__3/czekac_20min.jpg'),
(316, 61, 'Co oznacza ten znak patrolowy?', 'Zagrożenie', 'Pokój', 'Wojna', 'Niebezpieczeństwo', 'C', 'znaki_patrolowe_cz__3/wojna.jpg'),
(317, 61, 'Co oznacza ten znak patrolowy?', 'Wracać biegiem', 'Wracać', 'Biec', 'Iść  szybko', '', 'znaki_patrolowe_cz__3/wracac_biegiem.jpg'),
(318, 61, 'Co oznacza ten znak patrolowy?', 'Woja', 'Zła droga', 'Niebezpieczeństwo', 'Skrzyżowanie dróg', 'B', 'znaki_patrolowe_cz__3/zla_droga.jpg'),
(348, 76, 'Ile jest stopni harcerskich?', '7', '6', '4', '8', 'B', 'stopnie_harcerskie_zhp/stopnie_0.jpg'),
(349, 76, 'Jaka jest poprawna kolejność stopni harcerskich?', '4,1,5,2,3,6', '6,2,4,1,3,5', '6,1,4,2,3,5', '1,2,3,4,5,6', 'C', 'stopnie_harcerskie_zhp/stopnie_1.jpg'),
(350, 76, 'Jaki stopnień przedstawia ilustracja? ', 'Harcerka orla - Harcerz orli', 'Ochotniczka - Młodzik', 'Tropicielka - Wywiadowca', 'Samarytanka - Ćwik', 'D', 'stopnie_harcerskie_zhp/stopnie_2.jpg'),
(351, 76, 'Jak nazywa się 5 stopień harcerski?', 'tropicielka - Wywiadowca', 'Harcerka orla - Harcerz orli', 'Pionierka - Odkrywca', 'Ochotniczka - Młodzik', 'B', 'stopnie_harcerskie_zhp/stopnie_0.jpg'),
(352, 76, 'Jak nazywa się żeński wersja stopnia z jedną krokiewką ? ', 'Pionierka', 'Samarytanka', 'Tropicielka', 'Ochotniczka', 'A', 'stopnie_harcerskie_zhp/stopnie_3.jpg'),
(353, 76, 'Jaki stopnień przedstawia ilustracja?', 'Samarytanka - Ćwik', 'Ochotniczka - Młodzik', 'Samarytanka - Ćwik', 'Tropicielka - Wywiadowca', 'B', 'stopnie_harcerskie_zhp/stopnie_4.jpg'),
(354, 76, 'Jak oznacza się stopnień Tropicielka - Wywiadowca?', 'Jedna belka', 'Jedna krokiewka', 'Dwie Gwiazdki', 'Dwie belki', 'D', 'stopnie_harcerskie_zhp/stopnie_0.jpg'),
(355, 76, 'Jak nazywa się 3 stopień harcerski?', 'Samarytanka - Ćwik', 'Tropicielka - Wywiadowca', 'Pionierka - Odkrywca', 'Harcerka Rzeczypospolitej - Harcerz Rzeczypospolitej', 'C', 'stopnie_harcerskie_zhp/stopnie_0.jpg'),
(356, 76, 'Jak nazywa się męska wersja stopnia Pionierki?', 'Odkrywca', 'Wywiadowca', 'Ćwik', 'Harcerz Rzeczypospolitej', 'A', 'stopnie_harcerskie_zhp/stopnie_0.jpg'),
(357, 76, 'Jak oznacza się stopnień Harcerka Rzeczypospolitej - Harcerz Rzeczypospolitej?', 'Jedna Gwiazdka', 'Dwie Gwiazdki', 'Dwie Belki', 'Dwie Krokiewki', 'B', 'stopnie_harcerskie_zhp/stopnie_0.jpg'),
(358, 76, 'Jakim nr oznaczono stopień Tropicielka - Wywiadowca?', '2', '1', '5', '3', 'A', 'stopnie_harcerskie_zhp/stopnie_5.jpg'),
(359, 70, 'Kto znajduje się na ilustracji?', 'Robert Baden-Powell', 'Ernest Thompson Seton', 'Andrzej Małkowski', 'Aleksander Kamiński', 'A', 'znane_postacie_harcerskie_cz__1__(podstawowe)/Robert_Baden_Powell.jpg'),
(360, 70, 'Kto znajduje się na ilustracji?', 'Stanisław Broniewski', 'ks. Stefan Wincenty Frelichowski', 'Maciej Aleksy Dawidowski', 'ks.Kazimierz Lutosławski', 'B', 'znane_postacie_harcerskie_cz__1__(podstawowe)/Wincenty_Frelichowski.jpg'),
(361, 70, 'Kto znajduje się na ilustracji?', 'Stanisław Broniewski', 'Ernest Thompson Seton', 'Maciej Aleksy Dawidowski', 'Aleksander Kamiński', 'D', 'znane_postacie_harcerskie_cz__1__(podstawowe)/Aleksander Kamiński.jpg'),
(362, 70, 'Kto znajduje się na ilustracji?', 'Jan Bytnar', 'Tadeusz Zawadzki', 'Maciej Aleksy Dawidowski', 'Stanisław Broniewski', 'C', 'znane_postacie_harcerskie_cz__1__(podstawowe)/Aleksy_Dawidowski.jpg'),
(363, 70, 'Kto znajduje się na ilustracji?', 'Stanisław Kostka', 'Święty Jerzy', 'Zawisza Czarny', 'Święty  Łukasz', 'B', 'znane_postacie_harcerskie_cz__1__(podstawowe)/Jerzy.jpg'),
(364, 70, 'Kto znajduje się na ilustracji?', 'Olga Drahonowska-Małkowska', 'Jadwiga Falkowska', 'Jadwiga Zwolakowska', 'Olave Baden-Powell', 'D', 'znane_postacie_harcerskie_cz__1__(podstawowe)/Olave_Baden-Powell.jpg'),
(365, 70, 'Kto znajduje się na ilustracji?', 'Tadeusz Zawadzki', 'Aleksander Kamiński', 'Maciej Aleksy Dawidowski', 'Jan Bytnar', 'A', 'znane_postacie_harcerskie_cz__1__(podstawowe)/Tadeusz_Zawadzki.jpg'),
(366, 70, 'Kto znajduje się na ilustracji?', 'Stefan Mirowski', 'Andrzej Małkowski', 'gen. Józef Haller', 'ks. Kazimierz Lutosławski', 'B', 'znane_postacie_harcerskie_cz__1__(podstawowe)/Andrzej_Małkowski.jpg'),
(367, 70, 'Kto znajduje się na ilustracji?', 'Stefan Mirowski', 'Ernest Thompson Seto', 'gen. Józef Haller', 'Stanisław Broniewski', 'C', 'znane_postacie_harcerskie_cz__1__(podstawowe)/Józef_Haller.jpg'),
(368, 70, 'Kto znajduje się na ilustracji?', 'Stefan Mirowski', 'Stanisław Broniewski', 'Ernest Thompson Seton', 'Jan Bytnar', 'B', 'znane_postacie_harcerskie_cz__1__(podstawowe)/Stanisław_Broniewski.jpg'),
(369, 70, 'Kto znajduje się na ilustracji?', 'Stefan Mirowski', 'Tadeusz Zawadzki', 'Ernest Thompson Seton', 'Jan Bytnar', 'D', 'znane_postacie_harcerskie_cz__1__(podstawowe)/Jan_Bytnar.jpg'),
(370, 70, 'Kto znajduje się na ilustracji?', 'Olave Baden-Powell', 'Jadwiga Zwolakowska ', 'Olga Drahonowska-Małkowska', 'Jadwiga Falkowska', 'C', 'znane_postacie_harcerskie_cz__1__(podstawowe)/Olga_Drahonowska.jpg'),
(371, 70, 'Kto znajduje się na ilustracji?', 'Andrzej Małkowski', 'Ernest Thompson Seton', 'ks. Kazimierz Lutosławski', 'Stefan Mirowski', 'D', 'znane_postacie_harcerskie_cz__1__(podstawowe)/Stefan_Mirowski.jpg'),
(372, 70, 'Kto znajduje się na ilustracji?', 'Aleksander Kamiński', 'Ernest Thompson Seton', 'ks.Kazimierz Lutosławski', 'gen. Józef Haller', 'B', 'znane_postacie_harcerskie_cz__1__(podstawowe)/Thompson_Seton.jpg'),
(373, 70, 'Kto znajduje się na ilustracji?', 'ks.Kazimierz Lutosławski', 'Święty Jerzy', 'ks. Stefan Wincenty Frelichowski', 'gen. Józef Haller', 'A', 'znane_postacie_harcerskie_cz__1__(podstawowe)/Kazimierz_Lutosławski.jpg'),
(385, 56, 'Co symbolizują ramiona gwiazdek?', '10 punktów prawa harcerskiego', 'Strony świata', '10 założycieli organizacji WAGGGS', 'Nic nie symbolizują', 'A', 'symbolika_wagggs/waggs.jpg'),
(386, 78, 'Jaka to funkcja granatowy sznur z ramienia z dwoma węzłami lub dwoma granatowymi suwakami ?', '1', '2', '3', '4', 'A', 'sznury_funkcyjne_cz__1_(podstawowe)/falga_herb_polski.jpg'),
(387, 81, 'Jaki znak służby przedstawia ilustracja?', 'Znak służby wspólnocie lokalnej', 'Znak służby kulturze', 'Znak służby przyjaźni', 'Znak służby dziecku', 'D', 'znaki_sluzby_/znak_słuzby_dziecku.jpg'),
(388, 81, 'Jaki znak służby przedstawia ilustracja?', 'Znak służby nauce', 'Znak służby kulturze', 'Znak służby gospodarce', 'Znak służby wspólnocie lokalnej', 'C', 'znaki_sluzby_/znak_sluzby_gospodarce.jpg'),
(389, 81, '\nJaki znak służby przedstawia ilustracja?', 'Znak służby zdrowiu', 'Znak służby życia', 'Znak służby pomocy', 'Znak służby nauce', 'A', 'znaki_sluzby_/znak_sluzby_zdrowiu.jpg'),
(391, 81, '\nJaki znak służby przedstawia ilustracja?\n', 'Znak służby historii ', 'Znak służby kulturze', 'Znak służby wspólnocie lokalnej', 'Znak służby pamięci', 'B', 'znaki_sluzby_/znak_sluzby_kulturze.jpg'),
(392, 81, '\nJaki znak służby przedstawia ilustracja?\n', 'Znak służby ludziom ', 'Znak służby przyjaźni ', 'Znak służby wspólnocie lokalnej', 'Znak służby dzieciom ', 'C', 'znaki_sluzby_/znak_sluzby_ wspolnocie_lokalnej.jpg'),
(393, 81, '\nJaki znak służby przedstawia ilustracja?\n', 'Znak służby światu', 'Znak służby przyrodzie', 'Znak służby ziemi', 'Znak służby turystyce', 'D', 'znaki_sluzby_/znak_sluzby_turystyce.jpg'),
(394, 81, '\nJaki znak służby przedstawia ilustracja?\n', 'Znak służby pamięci', 'Znak służby nauce', 'Znak służby kulturze', 'Znak służby historii', 'B', 'znaki_sluzby_/znak_sluzby_nauce.jpg'),
(395, 81, '\nJaki znak służby przedstawia ilustracja?\n', 'Znak służby przyjaźni', 'Znak służby dzieciom', 'Znak służby światu', 'Znak służby harcerzom', 'A', 'znaki_sluzby_/znak_sluzby_przyjazni.jpg'),
(396, 81, '\nJaki znak służby przedstawia ilustracja?\n', 'Znak służby bogu ', 'Znak służby historii ', 'Znak służby pamięci', 'Znak służby poległym', 'C', 'znaki_sluzby_/znak_sluzby_pamieci.jpg'),
(397, 81, '\nJaki znak służby przedstawia ilustracja?\n', 'Znak służby ekologi', 'Znak służby rośliną', 'Znak służby ziemi', 'Znak służby przyrodzie', 'D', 'znaki_sluzby_/znak_sluzby_przyrodzie.jpg'),
(398, 83, 'Jaką specjalność przedstawia plakietka?', 'Naukową', 'Jeździecką i Kawaleryjską', 'Olimpijską', 'Sportową', 'D', 'plakietki_specjalnosci_druzyn_/p_sportowa.jpg'),
(399, 83, 'Jaką specjalność przedstawia plakietka?', 'Straży Granicznej', 'Pożarniczą ', 'Proobronna', 'Strażacka', 'B', 'plakietki_specjalnosci_druzyn_/p_strazacka.jpg'),
(400, 83, 'Jaką specjalność przedstawia plakietka?', 'Naukowa', 'Artystyczną ', 'Kulturalną ', 'Teatralna i artystyczną', 'B', 'plakietki_specjalnosci_druzyn_/p_teatralna_artystyczna.jpg'),
(401, 83, ' Jaką specjalność przedstawia plakietka?', 'Pocztową (Poczta Harcerska)', 'Łącznościową ', 'Artystyczną ', 'Muzyczną ', 'A', 'plakietki_specjalnosci_druzyn_/p_poczta_harcerska.jpg'),
(402, 83, ' Jaką specjalność przedstawia plakietka?', 'Medyczną', 'Ratowniczą', 'Krwiodawstwa ', 'Krwiodawstwa  I  Medyczną', 'B', 'plakietki_specjalnosci_druzyn_/p_medyczna.jpg'),
(403, 83, ' Jaką specjalność przedstawia plakietka?', 'Proobronną', 'Biblioteczną', 'Kulturalną', 'Naukową', 'D', 'plakietki_specjalnosci_druzyn_/p_naukowa.jpg'),
(404, 83, ' Jaką specjalność przedstawia plakietka?', 'Straży Granicznej', 'Baloniarską', 'Lotniczą', 'Spadochroniarką ', 'C', 'plakietki_specjalnosci_druzyn_/p_lotnicza.jpg'),
(405, 83, ' Jaką specjalność przedstawia plakietka?', 'Techniczną ', 'Mechaniczną', 'Krawiecką', 'Nie istniej taka', 'A', 'plakietki_specjalnosci_druzyn_/p_techniczna.jpg'),
(406, 83, ' Jaką specjalność przedstawia plakietka?', 'Lądową - wodną', 'Żeglarską i Wodną', 'Lądową ', 'Morską  i Wodną', 'B', 'plakietki_specjalnosci_druzyn_/p_zeglarska.jpg'),
(407, 83, ' Jaką specjalność przedstawia plakietka?', 'Krajoznawczą ', 'Ekologiczną', 'Turystyczną', 'Straży Granicznej', 'C', 'plakietki_specjalnosci_druzyn_/p_turystyczna.jpg'),
(408, 83, 'Jaką specjalność przedstawia plakietka?', 'Straży Granicznej', 'Turystyczną', 'Lotniczą', 'Proobronną (obronną)', 'D', 'plakietki_specjalnosci_druzyn_/p_proobronna.jpg'),
(409, 83, 'Jaką specjalność przedstawia plakietka?', 'Spadochroniarską ', 'łącznościowo-informatyczną', 'Muzyczną', 'Straży Granicznej', 'B', 'plakietki_specjalnosci_druzyn_/p_lacznosci.jpg'),
(410, 83, 'Jaką specjalność przedstawia plakietka?', 'Lotniczą', 'Baloniarską', 'Jeździecka i Kawaleryjska', 'Spadochroniarką', 'C', 'plakietki_specjalnosci_druzyn_/p_jezdziecka.jpg'),
(411, 83, 'Jaką specjalność przedstawia plakietka?', 'Historyczno-Rekonstrukcyjną ', 'Historyczną', 'Ratowniczą', 'Krwiodawczą', 'D', 'plakietki_specjalnosci_druzyn_/krwiodawstwo.jpg'),
(412, 83, 'Jaką specjalność przedstawia plakietka?', 'Ekologiczną (Przyrodniczą)', 'Naukową', 'Turystyczną', 'Krajoznawczą', 'A', 'plakietki_specjalnosci_druzyn_/p_ekologiczna.jpg'),
(413, 83, 'Jaką specjalność przedstawia plakietka?', 'Harcerskiej szkoły ratowniczej', 'Ruchu drogowego', 'Ruchu lądowego', 'Nie istnieje taka', 'B', 'plakietki_specjalnosci_druzyn_/p_hsr.jpg'),
(414, 83, 'Jaką specjalność przedstawia plakietka?', 'Proobronną', 'Krajoznawczą', 'Służby granicznej', 'Służby narodowej', 'C', 'plakietki_specjalnosci_druzyn_/p_granicza.jpg'),
(415, 82, 'Ile punktów mam prawo zuchowe?', '6', '5', '4', '3', 'A', 'prawo_zucha/prawo_zuchowe_ile.jpg'),
(416, 82, 'Jak brzmi 1 punkt prawa zuchowego?', 'Zuch jest dzielny.', 'Zuch kocha Boga i Polskę.', 'Zuch pamięta o swoich obowiązkach.', 'Zuch stara się być coraz lepszy.', 'B', 'prawo_zucha/prawo_zuchowe_1.jpg'),
(417, 82, 'Jaki to punkt prawa zuchowego? ', '3', '4', '5', '2', 'C', 'prawo_zucha/Prawo_harcerskie_5.jpg'),
(418, 82, 'Jak brzmi 2 punkt prawa zuchowego?', 'Zuch pamięta o swoich obowiązkach.', 'Wszystkim jest z zuchem dobrze.', 'Zuch stara się być coraz lepszy.', 'Zuch jest dzielny.', 'D', 'prawo_zucha/prawo_zuchowe_2.jpg'),
(419, 82, 'Jaki to punkt prawa zuchowego?', '6', '5', '2', '3', 'A', 'prawo_zucha/Prawo_harcerskie_6.jpg'),
(420, 82, 'Jak brzmi 3 punkt prawa zuchowego?', 'Zuch pamięta o swoich obowiązkach.', 'Zuch mówi prawdę.', 'Wszystkim jest z zuchem dobrze.', 'Zuch stara się być coraz lepszy.', 'B', 'prawo_zucha/prawo_zuchowe_3.jpg'),
(421, 82, 'Jaki to punkt prawa zuchowego?', '3', '2', '1', '4', 'C', 'prawo_zucha/Prawo_harcerskie_1_1.jpg'),
(422, 82, 'Jak brzmi 4 punkt prawa zuchowego?', 'Wszystkim jest z zuchem dobrze.', 'Zuch kocha Boga i Polskę.', 'Zuch mówi prawdę.', 'Zuch pamięta o swoich obowiązkach.', 'D', 'prawo_zucha/prawo_zuchowe_4.jpg'),
(439, 86, 'Na czym wzorowana była lilijka ?', 'igłą magnetyczną', 'Różą wiatrów ', 'Grotem strzały', 'Nie była na niczym wzorowana', 'A', 'symbolika_lilijki_harcerskiej/lilijka_1.jpg'),
(440, 86, 'Co symbolizuje litera O na ramieniu lilijki?', 'Ofiarność', 'Odpowiedzialność', 'Odwagę', 'Ojczyzna', 'D', 'symbolika_lilijki_harcerskiej/lilijka_1.jpg'),
(441, 86, 'Co symbolizują 3 ramiona lilijki?', 'Służbę bliźnim i Ojczyźnie', 'Służbę Bogu, Kulturze i Ojczyźnie', 'Służbę Bogu, Ojczyźnie i bliźnim', 'Służbę Bogu, Nauce  i bliźnim ', 'C', 'symbolika_lilijki_harcerskiej/lilijka_1.jpg'),
(442, 86, 'Co symbolizuje litera N na ramieniu lilijki?', 'Niewinność ', 'Naukę', 'Nadzieje', 'Narodowosć', 'B', 'symbolika_lilijki_harcerskiej/lilijka_1.jpg'),
(443, 86, 'Co symbolizuje igła magnetyczna na której wzorwana byla lilijka ?', 'Dążenie do ideałów harcerskich \"Dobra droga\"', 'Wędrowanie stała wędrówka ', 'Szukanie miejsa dla siebie wśród ludzi ', 'Nic nie symbolizuje', 'A', 'symbolika_lilijki_harcerskiej/lilijka_1.jpg'),
(444, 86, 'Co symbolizuje litera C na ramieniu lilijki?', 'Cierpliwość', 'Ciekawość', 'Chciwość', 'Cnotę', 'D', 'symbolika_lilijki_harcerskiej/lilijka_1.jpg'),
(445, 55, 'Co symbolizuje okrąg pośrodku krzyża?', 'Odpowiedzialność ', 'Jedność ', 'Braterstwo', 'Doskonałość', 'D', 'symbolika_krzyza_harcerskiego/krzyz1.jpg'),
(457, 53, 'Czy polska jest w poslka', 'tak', 'nie', 'nie wiem', 'możliwe', 'B', 'sssadminquiztest-_awaria!!!_strony/Screenshot_1638387926.png'),
(463, 52, 'Czy polska jest w poslka', 'tak', 'nie', 'nie wiem', 'możliwe', 'B', 'sssadminquiztest-_awaria!!!_strony/Screenshot_1638387926.png');

-- --------------------------------------------------------

--
-- Struktura tabeli dla tabeli `quotations`
--

CREATE TABLE `quotations` (
  `id` int(11) NOT NULL,
  `content` varchar(300) NOT NULL,
  `author` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `quotations`
--

INSERT INTO `quotations` (`id`, `content`, `author`) VALUES
(1, 'Nie błądzi tylko, ten kto nic nie robi.', 'Robert Baden-Powell'),
(2, 'Starajcie się zostawić ten świat troszkę lepszym niż go zastaliście.', 'Robert Baden-Powell'),
(3, 'Każdy osioł potrafi być dobrym człowiekiem podczas pogody', 'Robert Baden-Powell'),
(4, 'Jesteśmy powołani do tego, żeby być oazą. Oaza jest znakiem nadziei na horyzoncie.', 'Phill Bosman'),
(5, 'Nie czyńcie skautingu zbyt łatwym... Stawiajcie wyzwania.', 'Jan Paweł II'),
(6, 'Najtrudniej być porządnym człowiekiem, gdy nikt nie widzi.', 'Stefan Mirowski'),
(7, 'Siła w nas jest, na którą Polska liczy, o której potrzebuje.', 'Halina Paliwodzianka'),
(8, 'Nie błądzi tylko, ten kto nic nie robi.', 'Robert Baden-Powell'),
(9, 'Starajcie się zostawić ten świat troszkę lepszym niż go zastaliście.', 'Robert Baden-Powell'),
(10, 'Każdy osioł potrafi być dobrym człowiekiem podczas pogody', 'Robert Baden-Powell'),
(11, 'Jesteśmy powołani do tego, żeby być oazą. Oaza jest znakiem nadziei na horyzoncie.', 'Phill Bosman'),
(12, 'Nie czyńcie skautingu zbyt łatwym... Stawiajcie wyzwania.', 'Jan Paweł II'),
(13, 'Najtrudniej być porządnym człowiekiem, gdy nikt nie widzi.', 'Stefan Mirowski'),
(14, 'Siła w nas jest, na którą Polska liczy, o której potrzebuje.', 'Halina Paliwodzianka'),
(15, 'praszamy do przeczesywania bazy literackich cytatów. Najpopularniejsze cytaty o miłości, cytaty o życiu, cytaty motywacyjne, aforyzmy o przyjaźni, no i oczywiście cytaty z konkretnych książek, jak c', 'dwd');

-- --------------------------------------------------------

--
-- Struktura tabeli dla tabeli `stats`
--

CREATE TABLE `stats` (
  `id` int(11) NOT NULL,
  `date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `quiz_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `stats`
--

INSERT INTO `stats` (`id`, `date`, `quiz_id`) VALUES
(1, '2024-11-03 00:38:33', 3),
(2, '2024-11-03 13:05:09', 5),
(3, '2024-11-03 14:07:01', 59),
(4, '2024-11-03 14:12:12', 59),
(5, '2024-11-03 16:17:37', 59),
(6, '2024-11-03 16:55:24', 59),
(7, '2024-11-03 17:09:20', 59),
(8, '2024-11-03 17:17:07', 59),
(9, '2024-11-03 17:17:50', 59),
(10, '2024-11-24 20:30:42', 57),
(11, '2024-11-24 20:31:05', 57),
(12, '2024-11-25 16:13:55', 53),
(13, '2024-11-25 16:18:15', 53),
(14, '2024-11-25 17:43:53', 53),
(15, '2024-11-25 17:44:03', 53),
(16, '2024-11-25 17:44:51', 53),
(17, '2024-11-25 18:37:09', 76),
(29, '2024-11-03 00:38:33', 3),
(30, '2024-11-03 00:38:33', 3),
(31, '2024-11-03 00:38:33', 3),
(32, '2024-11-03 00:38:33', 3),
(33, '2024-11-03 00:38:33', 3),
(34, '2024-11-03 00:38:33', 3),
(35, '2024-11-03 00:38:33', 3),
(36, '2024-11-03 00:38:33', 3),
(37, '2024-11-03 00:38:33', 3),
(38, '2024-11-03 00:38:33', 3),
(39, '2024-11-03 00:38:33', 3),
(40, '2024-11-03 00:38:33', 3),
(41, '2024-11-03 00:38:33', 3),
(42, '2024-11-03 00:38:33', 3),
(43, '2024-11-03 00:38:33', 3),
(44, '2024-11-03 00:38:33', 3),
(45, '2024-11-03 00:38:33', 3),
(46, '2024-11-03 00:38:33', 3),
(47, '2024-11-03 00:38:33', 3),
(48, '2024-11-03 00:38:33', 3),
(49, '2024-11-03 00:38:33', 3),
(50, '2024-11-03 00:38:33', 3),
(51, '2024-11-03 00:38:33', 3),
(52, '2024-11-03 00:38:33', 3),
(53, '2024-11-03 00:38:33', 3),
(54, '2024-11-03 00:38:33', 3),
(55, '2024-11-03 00:38:33', 3),
(56, '2024-11-03 00:38:33', 3),
(57, '2024-11-03 00:38:33', 3),
(58, '2024-11-03 00:38:33', 3),
(59, '2024-11-03 00:38:33', 3),
(60, '2024-11-03 00:38:33', 3),
(61, '2024-11-03 00:38:33', 3),
(62, '2024-11-03 00:38:33', 3),
(63, '2024-11-03 00:38:33', 3),
(64, '2024-11-03 00:38:33', 3),
(65, '2024-11-03 00:38:33', 3),
(66, '2024-11-03 00:38:33', 3),
(67, '2024-11-03 00:38:33', 3),
(68, '2024-11-03 00:38:33', 3),
(69, '2024-11-03 00:38:33', 3),
(70, '2024-11-03 00:38:33', 3),
(71, '2024-11-03 00:38:33', 3),
(72, '2024-11-03 00:38:33', 3),
(73, '2024-11-03 00:38:33', 3),
(74, '2024-11-03 00:38:33', 3),
(75, '2024-11-03 00:38:33', 3),
(76, '2024-11-03 00:38:33', 3),
(77, '2024-11-03 00:38:33', 3),
(78, '2024-11-03 00:38:33', 3),
(79, '2024-11-03 00:38:33', 3),
(80, '2024-11-03 00:38:33', 3),
(81, '2024-11-03 00:38:33', 3),
(82, '2024-11-03 00:38:33', 3),
(83, '2024-11-03 00:38:33', 3),
(84, '2024-11-03 00:38:33', 3),
(85, '2024-11-03 00:38:33', 53),
(86, '2024-11-03 00:38:33', 53),
(87, '2024-11-03 00:38:33', 53),
(88, '2024-11-03 00:38:33', 53),
(89, '2024-11-03 00:38:33', 53),
(90, '2024-11-03 00:38:33', 53),
(91, '2024-11-03 00:38:33', 53),
(92, '2024-11-03 00:38:33', 53),
(93, '2024-11-03 00:38:33', 53),
(94, '2024-11-03 00:38:33', 53),
(95, '2024-11-03 00:38:33', 53),
(96, '2024-11-03 00:38:33', 53),
(97, '2024-11-03 00:38:33', 53),
(98, '2024-11-03 00:38:33', 53),
(99, '2024-11-03 00:38:33', 53),
(100, '2024-11-03 00:38:33', 53),
(101, '2024-11-03 00:38:33', 53),
(102, '2024-11-03 00:38:33', 53),
(103, '2024-11-03 00:38:33', 53),
(104, '2024-11-03 00:38:33', 53),
(105, '2024-11-03 00:38:33', 53),
(106, '2024-11-03 00:38:33', 53),
(107, '2024-11-03 00:38:33', 53),
(108, '2024-11-03 00:38:33', 53),
(109, '2024-11-03 00:38:33', 53),
(110, '2024-11-03 00:38:33', 53),
(111, '2024-11-03 00:38:33', 53),
(112, '2024-11-03 00:38:33', 53),
(113, '2024-11-03 00:38:33', 53),
(114, '2024-11-03 00:38:33', 53),
(115, '2024-11-03 00:38:33', 53),
(116, '2024-11-03 00:38:33', 53),
(117, '2024-11-03 00:38:33', 53),
(118, '2024-11-03 00:38:33', 53),
(119, '2024-11-03 00:38:33', 53),
(120, '2024-11-03 00:38:33', 53),
(121, '2024-11-03 00:38:33', 53),
(122, '2024-11-03 00:38:33', 53),
(123, '2024-11-03 00:38:33', 53),
(124, '2024-11-03 00:38:33', 53),
(125, '2024-11-03 00:38:33', 53),
(126, '2024-11-03 00:38:33', 53),
(127, '2024-11-03 00:38:33', 53),
(128, '2024-11-03 00:38:33', 53),
(129, '2024-11-03 00:38:33', 53),
(130, '2024-11-03 00:38:33', 53),
(131, '2024-11-03 00:38:33', 53),
(132, '2024-11-03 00:38:33', 53),
(133, '2024-11-03 00:38:33', 53),
(134, '2024-11-03 00:38:33', 53),
(135, '2024-11-03 00:38:33', 53),
(136, '2024-11-03 00:38:33', 53),
(137, '2024-11-03 00:38:33', 53),
(138, '2024-11-03 00:38:33', 53),
(139, '2024-11-03 00:38:33', 53),
(140, '2024-11-03 00:38:33', 53),
(141, '2024-11-03 00:38:33', 53),
(142, '2024-11-03 00:38:33', 53),
(143, '2024-11-03 00:38:33', 53),
(144, '2024-11-03 00:38:33', 53),
(145, '2024-11-03 00:38:33', 53),
(146, '2024-11-03 00:38:33', 53),
(147, '2024-11-03 00:38:33', 53),
(148, '2024-11-03 00:38:33', 53),
(149, '2024-11-03 00:38:33', 53),
(150, '2024-12-09 00:38:33', 53),
(151, '2024-12-09 00:38:33', 53),
(152, '2024-12-09 00:38:33', 53),
(153, '2024-12-09 00:38:33', 53),
(154, '2024-12-09 00:38:33', 53),
(155, '2024-12-09 00:38:33', 53),
(156, '2024-12-09 00:38:33', 53),
(157, '2024-12-09 00:38:33', 53),
(158, '2024-12-09 00:38:33', 53),
(159, '2024-12-09 00:38:33', 53),
(160, '2024-12-09 00:38:33', 53),
(161, '2024-12-09 00:38:33', 53),
(162, '2024-12-09 00:38:33', 53),
(163, '2024-12-09 00:38:33', 53),
(164, '2024-12-09 00:38:33', 53),
(165, '2024-12-09 00:38:33', 53),
(166, '2024-12-09 00:38:33', 53),
(167, '2024-12-09 00:38:33', 53),
(168, '2024-12-09 00:38:33', 53),
(169, '2024-12-09 00:38:33', 53),
(170, '2024-12-09 00:38:33', 53),
(171, '2024-12-09 00:38:33', 53),
(172, '2024-12-09 00:38:33', 53),
(173, '2024-12-09 00:38:33', 53),
(174, '2024-12-09 00:38:33', 53),
(175, '2024-12-09 00:38:33', 53),
(176, '2024-12-09 00:38:33', 53),
(177, '2024-12-09 00:38:33', 53),
(178, '2024-12-09 00:38:33', 53),
(179, '2024-12-09 00:38:33', 53),
(180, '2024-12-09 00:38:33', 53),
(181, '2024-12-09 00:38:33', 53),
(182, '2024-12-09 00:38:33', 53),
(183, '2024-12-09 00:38:33', 53),
(184, '2024-12-09 00:38:33', 53),
(185, '2024-12-09 00:38:33', 53),
(186, '2024-12-09 00:38:33', 53),
(187, '2024-12-09 00:38:33', 53),
(188, '2024-12-09 00:38:33', 53),
(189, '2024-12-09 00:38:33', 53),
(190, '2024-12-09 00:38:33', 53),
(191, '2024-12-09 00:38:33', 53),
(192, '2024-12-09 00:38:33', 53),
(193, '2024-12-09 00:38:33', 53),
(194, '2024-12-09 00:38:33', 53),
(195, '2024-12-09 00:38:33', 53),
(196, '2024-12-09 00:38:33', 53),
(197, '2024-12-09 00:38:33', 53),
(198, '2024-12-09 00:38:33', 53),
(199, '2024-12-09 00:38:33', 53),
(200, '2024-12-09 00:38:33', 53),
(201, '2024-12-09 00:38:33', 53),
(202, '2024-12-09 00:38:33', 53),
(203, '2024-12-09 00:38:33', 53),
(204, '2024-12-09 00:38:33', 53),
(205, '2024-12-09 00:38:33', 53),
(206, '2024-12-09 00:38:33', 53),
(207, '2024-12-09 00:38:33', 53),
(208, '2024-12-09 00:38:33', 53),
(209, '2024-12-09 00:38:33', 53),
(210, '2024-12-09 00:38:33', 53),
(211, '2024-12-09 00:38:33', 53),
(212, '2024-12-09 00:38:33', 53),
(213, '2024-12-09 00:38:33', 53),
(214, '2025-04-24 10:38:33', 2);

--
-- Indeksy dla zrzutów tabel
--

--
-- Indeksy dla tabeli `data`
--
ALTER TABLE `data`
  ADD PRIMARY KEY (`id`);

--
-- Indeksy dla tabeli `extraPytania`
--
ALTER TABLE `extraPytania`
  ADD PRIMARY KEY (`question_id`),
  ADD KEY `quiz_id` (`quiz_id`);

--
-- Indeksy dla tabeli `extraQuiz`
--
ALTER TABLE `extraQuiz`
  ADD PRIMARY KEY (`id`);

--
-- Indeksy dla tabeli `news`
--
ALTER TABLE `news`
  ADD PRIMARY KEY (`id`);

--
-- Indeksy dla tabeli `pytania`
--
ALTER TABLE `pytania`
  ADD PRIMARY KEY (`question_id`),
  ADD KEY `quiz_id` (`quiz_id`);

--
-- Indeksy dla tabeli `quotations`
--
ALTER TABLE `quotations`
  ADD PRIMARY KEY (`id`);

--
-- Indeksy dla tabeli `stats`
--
ALTER TABLE `stats`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `data`
--
ALTER TABLE `data`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=126;

--
-- AUTO_INCREMENT for table `extraPytania`
--
ALTER TABLE `extraPytania`
  MODIFY `question_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=226;

--
-- AUTO_INCREMENT for table `extraQuiz`
--
ALTER TABLE `extraQuiz`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=36;

--
-- AUTO_INCREMENT for table `news`
--
ALTER TABLE `news`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=31;

--
-- AUTO_INCREMENT for table `pytania`
--
ALTER TABLE `pytania`
  MODIFY `question_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=465;

--
-- AUTO_INCREMENT for table `quotations`
--
ALTER TABLE `quotations`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT for table `stats`
--
ALTER TABLE `stats`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=215;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
