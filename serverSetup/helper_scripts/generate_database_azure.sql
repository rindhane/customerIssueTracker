CREATE SCHEMA DataSchema;
GO


CREATE TABLE DataSchema.UserDetails (
	ID           INT IDENTITY(1,1) NOT NULL PRIMARY KEY,
	Name         NVARCHAR(200),
    Email        NVARCHAR(200),
	PasswordHash NVARCHAR(80),
    Phone        NVARCHAR(20),
	AccessLevel  INT
);
GO

CREATE TABLE DataSchema.IssueDetails (
	ID           INT IDENTITY(1,1) NOT NULL PRIMARY KEY,
	IssueType    INT,
    Description  TEXT,
    Status       INT,
    UserID       INT,
    LastUpdate  INT,
);
GO

CREATE TABLE DataSchema.TaskMembers (
    ID    INT IDENTITY(1,1) NOT NULL PRIMARY KEY,
    Name  NVARCHAR(200), 
    Email NVARCHAR(200)
)

SELECT * FROM DataSchema.IssueDetails;
GO

