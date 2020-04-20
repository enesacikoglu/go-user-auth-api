-- Sequences
CREATE SEQUENCE SEQ_USERS
  START WITH 1
  INCREMENT BY 1
  MINVALUE -9223372036854775808
  MAXVALUE 9223372036854775807
  CACHE;

CREATE SEQUENCE SEQ_ROLES
  START WITH 1
  INCREMENT BY 1
  MINVALUE -9223372036854775808
  MAXVALUE 9223372036854775807
  CACHE;

CREATE SEQUENCE SEQ_PERMISSIONS
  START WITH 1
  INCREMENT BY 1
  MINVALUE -9223372036854775808
  MAXVALUE 9223372036854775807
  CACHE;

CREATE SEQUENCE SEQ_APPLICATIONS
  START WITH 1
  INCREMENT BY 1
  MINVALUE -9223372036854775808
  MAXVALUE 9223372036854775807
  CACHE;

CREATE SEQUENCE SEQ_USER_ROLES
  START WITH 1
  INCREMENT BY 1
  MINVALUE -9223372036854775808
  MAXVALUE 9223372036854775807
  CACHE;

CREATE SEQUENCE SEQ_ROLE_PERMISSIONS
  START WITH 1
  INCREMENT BY 1
  MINVALUE -9223372036854775808
  MAXVALUE 9223372036854775807
  CACHE;


create table Applications
(
  Id           int          not null
    constraint PK_APPLICATIONS
      primary key,
  Name         varchar(100) not null
    constraint UC_Applications_Name
      unique,
  CreatedBy    varchar(50)  not null,
  ModifiedBy   varchar(50)  not null,
  CreatedDate  datetime2(3) not null,
  ModifiedDate datetime2(3) not null
)
go

create table Permissions
(
  Id             int          not null
    constraint PK_Permissions
      primary key,
  Name           varchar(100) not null
    constraint UC_Permissions_Name
      unique,
  CreatedBy      varchar(50)  not null,
  ModifiedBy     varchar(50)  not null,
  CreatedDate    datetime2(3) not null,
  ModifiedDate   datetime2(3) not null
)
go

create table RolePermissions
(
  Id            int           not null
    constraint PK_RolePermissions
      primary key,
  RoleId        int           not null,
  PermissionId  int           not null,
  CreatedBy     varchar(50)   not null,
  CreatedDate   datetime2(3)  not null,
  ApplicationId int default 1 not null,
  constraint UC_RolePermissions_RoleId_PermissionId
    unique (RoleId, PermissionId)
)
go

create table Roles
(
  Id           int          not null
    constraint PK_Roles
      primary key,
  Name         varchar(100) not null
    constraint UC_Roles_Name
      unique,
  CreatedBy    varchar(50)  not null,
  ModifiedBy   varchar(50)  not null,
  CreatedDate  datetime2(3) not null,
  ModifiedDate datetime2(3) not null
)
go

create table UserRoles
(
  Id          int          not null
    constraint PK_UserRoles
      primary key,
  UserId      int          not null,
  RoleId      int          not null,
  CreatedBy   varchar(50)  not null,
  CreatedDate datetime2(3) not null,
  constraint UC_UserRoles_UserId_RoleId
    unique (UserId, RoleId)
)
go

create table Users
(
  Id           int          not null
    constraint PK_Users
      primary key,
  Email        varchar(100) not null
    constraint UC_Users_Email
      unique,
  Name         varchar(100) not null,
  Surname      varchar(100) not null,
  CreatedBy    varchar(50)  not null,
  ModifiedBy   varchar(50)  not null,
  CreatedDate  datetime2(3) not null,
  ModifiedDate datetime2(3) not null
)
go