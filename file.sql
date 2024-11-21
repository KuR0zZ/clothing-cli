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
	Email VARCHAR(50) UNIQUE,
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

-- insert data user
INSERT INTO UserAdmin (Email, Password) VALUES
('nafa@gmail.com', 'admin1'),
('yugo@gmail.com', 'admin2'),
('ferdinand@gmail.com', 'admin3');

-- insert data Customers
INSERT INTO Customers (Name, Address, Email, PhoneNumber) VALUES
('Alice Johnson', 'Bandung', 'alice.johnson@gmail.com', '+6212345678'),
('Bob Smith', 'Yogyakarta', 'bob.smith@example.com', '+6223458901'),
('Clara Davis', 'Surabaya', 'clara.davis@example.com', '+6256789012'),
('David Brown', 'Jakarta', 'david.brown@example.com', '+6245670123'),
('Eva White', 'Medan', 'eva.white@example.com', '+6256781234');

-- insert data Products
INSERT INTO Products (ProductName, Price, Stock) VALUES
('Kemeja Polo', 50000.00, 10),
('Daster Ibu Hamil', 40000.00, 20),
('Topi Anak', 25000.00, 100),
('Kaos Kaki Budak Korporat', 60000.00, 15),
('Dasi Bos', 2000000.00, 25);

-- insert data Transactions
INSERT INTO Transactions (CustomerId, Date) VALUES
  (1, '2024-11-01'),
  (2, '2024-11-12'),
  (1, '2024-11-01'),
  (3, '2024-11-02'),
  (3, '2024-11-15'),
  (4, '2024-11-04'),
  (3, '2024-11-03'),
  (5, '2024-11-05');

-- insert data transaction details
INSERT INTO TransactionsDetails (TransactionId, ProductId, Quantity, TotalPrice) VALUES
(1, 1, 1, 50000.00),
(1, 3, 2, 50000.00),
(2, 2, 1, 40000.00),
(3, 4, 1, 60000.00),
(3, 3, 1, 25000.00),
(4, 5, 2, 4000000.00),
(5, 1, 1, 50000.00),
(5, 2, 1, 40000.00),
(5, 3, 3, 75000.00),
(6, 1, 2, 100000.00),
(7, 2, 4, 160000.00),
(8, 5, 1, 2000000.00);
