goblog
======

Learn to write a blog framework from scratch in golang

Database
--------
Curently Sqlite and Mysql are supported.

To change database:

$sed -i 's#mysql#sqlite' main.go

$sed -i 's#mysql#sqlite' blog/blog.go

It can change back and forth as you wish.

**Sqlite**

no need setup.


**Mysql pre-configure**

1. Create databse 'test' for use:

   MariaDB [test]> create database test;

2. Grant admin user with grant option:

   MariaDB [test]> grant all privileges on test.* to admin@'localhost' identified
   by'password' with grant option;


Installation
------------

- $ cd $GOPATH
- $ go get github.com/ouyanggh/goblog 
- $ cd src/github.com/ouyanggh/goblog
- $ go run main.go

**Usage**

- enter: http://localhost:8080/
- admin portal: http://localhost:8080/admin  admin/hello

TODO
----

- refactor the blog.go to use interface.
- enhance the front pages.
