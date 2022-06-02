create database if not exists `Exchange`;
use Exchange;

drop table if exists `User`;
create table if not exists `User`(
	`Id` INT  primary key not null auto_increment unique,
	`Name` varchar(50) not null, 
	`Username` varchar(16) not null,
	`Password` varchar(150) not null,
	`Email` varchar(32) not null
);
    
drop table if exists `Spot`;
create table if not exists `Spot`(
	`Id` INT primary key not null unique AUTO_INCREMENT unique
);

drop table if exists `SpotCoin`;
create table if not exists `SpotCoin`(
	`Id` INT primary key not null unique AUTO_INCREMENT unique,
    `Spot_Pocket` int not null,
    `Id_Coin` varchar(20) not null,
    `Quantity` decimal(20,10) not null,
    `Buy_Price` decimal(10,2) not null,
    `Sell_Price` decimal(10,2),
    `Buy_Date` datetime not null,
    `Sell_Date` datetime,
    foreign key (`Spot_Pocket`)
		references `Spot`(`Id`)
);

drop table if exists `Earn`;
create table if not exists `Earn`(
	`Id` INT primary key not null unique AUTO_INCREMENT unique
);

drop table if exists `EarnCoins`;
create table if not exists `EarnCoins`(
	`Id` INT primary key not null unique AUTO_INCREMENT unique,
    `Earn_Pocket` int not null,
    `Id_Coin` varchar(20) not null,
    `Quantity` decimal(20,10) not null,
    `Buy_Price` decimal(10,2) not null,
    `Sell_Price` decimal(10,2),
    `Buy_Date` datetime not null,
    `Sell_Date` datetime,
    foreign key (`Earn_Pocket`)
		references `Earn`(`Id`)
);

drop table if exists `Fiat`;
create table if not exists `Fiat`(
	`Id` INT primary key not null unique AUTO_INCREMENT unique,
    `Total` Decimal not null
);

drop table if exists `Wallet`;
create table if not exists `Wallet`(
	`Id` INT primary key not null unique AUTO_INCREMENT unique,
    `Spot_Pocket` int not null,
    `Earn_Pocket` int not null,
    `Fiat_Pocket` INT not null,
    `Id_User` int not null,
    foreign key (`Spot_Pocket`)
		references `Spot`(`Id`),
	foreign key (`Earn_Pocket`)
		references `Earn`(`Id`),
	foreign key (`Fiat_Pocket`)
		references `Fiat`(`Id`),
	foreign key (`Id_User`)
		references `User`(`Id`)
)
