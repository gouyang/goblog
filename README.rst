goblog
======

.. image:: https://api.travis-ci.org/ouyanggh/goblog.svg?branch=master
    :target: https://travis-ci.org/ouyanggh/goblog


Learn to write a blog framework from scratch in golang

Database
--------
Curently Sqlite and Mysql are added for support.

To change database and vice versa:

.. code-block:: bash

    $ sed -i 's#mysql#sqlite' main.go
    $ sed -i 's#mysql#sqlite' blog/blog.go

**Sqlite**

.. code-block:: bash

    go get github.com/mattn/go-sqlite3

**MariaDB**

.. code-block:: bash

    go get github.com/go-sql-driver/mysql

1. Create databse 'test' for use:

.. code-block:: bash

    $ MariaDB [test]> CREATE DATABASE test;

2. Grant admin user with grant option:

.. code-block:: bash

    $ MariaDB [test]> GRANT ALL PRIVILEGES ON test.* TO admin@'localhost' IDENTIFIED
                  > BY 'password' WITH GRANT OPTION;

3. To set the default to UTF-8, you want to add the following to my.cnf

.. code-block:: bash

    $ vi /etc/my.cnf

    [client]
    default-character-set=utf8

    [mysql]
    default-character-set=utf8

    [mysqld]
    collation-server = utf8_unicode_ci
    init-connect='SET NAMES utf8'
    character-set-server = utf8

    $ MariaDB [test]> ALTER TABLE Table CONVERT TO CHARACTER SET utf8 COLLATE utf8_unicode_ci; 

Usage
-----

- enter: http://localhost:8080/
- admin page: http://localhost:8080/admin  admin/hello

TODO
----

- refactor the blog.go to use interface.
- enhance the front pages.
- auth and session
- http://shadynasty.biz/blog/2012/09/05/auth-and-sessions/
