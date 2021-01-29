CREATE TABLE `jokes` (
  `id` SERIAL PRIMARY KEY,
  `autor` varchar(255) NOT NULL,
  `date` date NOT NULL,
  `header` varchar(255) NOT NULL,
  `content` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `jokes` (`id`, `autor`, `date`, `header`, `content`) VALUES
(NULL, 'Автор', '2020-04-20', 'Школа', 'Сейчас был на родительском собрании. На удалёнке. Первый раз присутствую на собрании в рубашке и трусах. Иногда отключал камеру, чтобы глотнуть пива. Причём, судя по мельканию камер, не я один такой. Мне понравилось :)'),
(NULL, 'Автор', '2020-04-22', 'Кино', 'Внимание! В фильме содержатся сцены рукопожатий, прикосновений к лицу и пребывания на улице без уважительной причины.');
