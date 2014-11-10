goblog
======

Learn to write a blog framework from scratch in golang

**Mysql pre-configure**
The site will use mysql as its database, pre-setup:

1. Create databse test for use:
   MariaDB [test]> create database test;

2. Grant admin user with grant option:
   MariaDB [test]> grant all privileges on test.* to admin@'localhost' identified
   by'password' with grant option;

**Installation**

- $ cd $GOPATH
- $ go get github.com/ouyanggh/goblog 
- $ cd src/github.com/ouyanggh/goblog
- $ go run main.go

**Usage**

- enter: http://localhost:8080/
- admin portal: http://localhost:8080/admin  admin/hello

**TODO**

1. go though `Writing Web Applications`_

.. _Writing Web Applications: https://golang.org/doc/articles/wiki/

2. basic architecture of whole project

