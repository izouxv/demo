show databases ;
use reptile;
show tables ;
create database reptile default character set  utf8mb4 collate utf8mb4_general_ci;
use reptile;
CREATE TABLE `job` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `city` varchar(255) DEFAULT NULL,
  `district` varchar(255) DEFAULT '',
  `company_short_name` varchar(255) DEFAULT '',
  `company_full_name` varchar(255) DEFAULT '',
  `company_label_list` varchar(255) DEFAULT '',
  `company_size` varchar(255) DEFAULT '',
  `industry_field` varchar(255) DEFAULT '',
  `industry_lables` varchar(255) DEFAULT '',
  `position_name` varchar(255) DEFAULT '',
  `position_lables` varchar(255) DEFAULT '',
  `position_advantage` varchar(255) DEFAULT '',
  `finance_stage` varchar(255) DEFAULT '',
  `work_year` varchar(255) DEFAULT '',
  `education` varchar(255) DEFAULT '',
  `salary` varchar(255) DEFAULT '',
  `longitude` varchar(255) DEFAULT '',
  `latitude` varchar(255) DEFAULT '',
  `linestaion` varchar(255) DEFAULT '',
  `create_time` int(10) unsigned DEFAULT '0',
  `add_time` int(10) unsigned DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
