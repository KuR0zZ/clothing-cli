## ERD Title: Server clothing App

## Entities and Their Attributes:
Entity: UserAdmin, Products, Transactions, and Customers

## Attributes:
A. Entity: UserAdmin
Attributes:
- Id INT (PK, AI)
- Email Varchar (50) UNIQUE NOT NULL
- Password Varchar (50) NOT NULL

B. Entity: Customers
Attributes:
- Id INT (PK, AI)
- Name Varchar(50) NOT NULL
- Address Varchar (50) NOT NULL
- Email Varchar (50) UNIQUE
- PhoneNumber Varchar (15) NOT NULL - Customer's contact number (Format: +xxxxxx)

C. Entity: Products
Attributes:
- Id INT (PK, AI)
- ProductName Varchar(50) NOT NULL
- Price DECIMAL (10,2) NOT NULL
- Stock INT NOT NULL

D. Entity: Transactions
Attributes:
- Id INT (PK, AI)
- CustomerId INT (FK) NOT NULL - Foreign key referencing Customers(Id)
- Date DATE NOT NULL

E. Entity: TransactionDetails
Attributes:
- Id INT (PK, AI)
- TransactionId INT (FK) NOT NULL - Foreign key referencing Transactions(Id)
- ProductId INT (FK) NOT NULL - Foreign key referencing Products(Id)
- Quantity INT NOT NULL
- TotalPrice DECIMAL (10,2) NOT NULL

## Relationships:
- Transactions to TransactionDetails:

A. Type: One to Many
Description: One transaction can have multiple transaction details, but each transaction detail is for a single transaction.

- Customers to Transactions:

B. Type: One to Many
Description: One customer can make multiple transactions, but each transaction belongs to one customer.

- TransactionDetails to Products:

C. Type: One to Many
Description: One transaction detail can have multiple products, but each product belongs to one transaction detail.

## Integrity Constraints:
- Stock should be always greater than or equal to 0
- Quantity and Price should be always greater than to 0
- Unique: email in UserAdmin and Customers to prevent duplicate entries

## Additional Notes:
