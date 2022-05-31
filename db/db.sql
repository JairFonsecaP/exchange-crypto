create database if not exists `Exchange`;
use Exchange;

drop table if exists `Users`;
create table if not exists `Users`(
	`Id` INT  primary key not null auto_increment unique,
	`Name` varchar(50) not null, 
	`Username` varchar(16) not null,
	`Password` varchar(150) not null,
	`Email` varchar(32) not null
);
    
drop table if exists `Spot`;
create table if not exists `Spot`(
	`Id` INT primary key not null unique AUTO_INCREMENT unique,
    `Id_User` int not null,
    foreign key (`Id_User`)
		references `Users`(`Id`)
		ON UPDATE CASCADE ON DELETE RESTRICT
);

drop table if exists `SpotCoin`;
create table if not exists `SpotCoin`(
	`Id` INT primary key not null unique AUTO_INCREMENT unique,
    `Buy_Price` decimal,
    `Sell_Price` decimal,
    `Wallet_Spot` INT not null,
    foreign key (`Wallet_Spot`)
		references `Spot`(`Id`)
		ON UPDATE CASCADE ON DELETE RESTRICT
)
