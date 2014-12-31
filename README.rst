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

3. To set the default to UTF-8, you want to add the following to my.cnf

    [client]

    default-character-set=utf8

    [mysql]

    default-character-set=utf8

    [mysqld]

    collation-server = utf8_unicode_ci

    init-connect='SET NAMES utf8'

    character-set-server = utf8

and do below in db:

    ALTER TABLE Table CONVERT TO CHARACTER SET utf8 COLLATE utf8_unicode_ci; 


Installation
------------

    $ cd $GOPATH

    $ go get github.com/ouyanggh/goblog 

    $ cd src/github.com/ouyanggh/goblog

    $ go run main.go

**Usage**

- enter: http://localhost:8080/
- admin portal: http://localhost:8080/admin  admin/hello

TODO
----

- refactor the blog.go to use interface.
- enhance the front pages.
