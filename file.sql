-- Create new database
CREATE DATABASE clothing_cli;

-- Use newly created database
USE clothing_cli;

-- Create table UserAdmin
CREATE TABLE IF NOT EXISTS UserAdmin (
	id INT AUTO_INCREMENT PRIMARY KEY,
	Email VARCHAR(50) UNIQUE NOT NULL,
	Password VARCHAR(50) NOT NULL
);

-- Create table Customers
CREATE TABLE IF NOT EXISTS Customers (
	id INT AUTO_INCREMENT PRIMARY KEY,
	Name VARCHAR(50) NOT NULL,
	Address VARCHAR(50) NOT NULL,
	Email VARCHAR(50) UNIQUE NOT NULL,
	PhoneNumber VARCHAR(15) NOT NULL
);

-- Create table Transactions
CREATE TABLE IF NOT EXISTS Transactions (
	id INT AUTO_INCREMENT PRIMARY KEY,
	CustomerId INT NOT NULL,
	Date Date NOT NULL,
	FOREIGN KEY (CustomerId) REFERENCES Customers(id)
);

-- Create table Products
CREATE TABLE IF NOT EXISTS Products (
	id INT AUTO_INCREMENT PRIMARY KEY,
	ProductName VARCHAR(50) NOT NULL,
	Price DECIMAL(10,2) NOT NULL,
	Stock INT NOT NULL
);

-- Create table TransactionsDetails
CREATE TABLE IF NOT EXISTS TransactionsDetails (
	id INT AUTO_INCREMENT PRIMARY KEY,
	TransactionId INT NOT NULL,
	ProductId INT NOT NULL,
	Quantity INT NOT NULL,
	TotalPrice DECIMAL(10,2) NOT NULL,
	FOREIGN KEY (TransactionId) REFERENCES Transactions(id),
	FOREIGN KEY (ProductId) REFERENCES Products(id)
);