goblog
======

.. image:: https://api.travis-ci.org/ouyanggh/goblog.svg?branch=master
    :target: https://travis-ci.org/ouyanggh/goblog


Learn to write a blog framework from scratch in golang

Database
--------
Curently Sqlite and Mysql are supported.

To change database:

.. code-block:: bash

    $ sed -i 's#mysql#sqlite' main.go
    $ sed -i 's#mysql#sqlite' blog/blog.go

It can change back and forth as you wish.

**Sqlite**

no need setup.


**Mysql pre-configure**

1. Create databse 'test' for use:

.. code-block:: bash

    MariaDB [test]> CREATE DATABASE test;

2. Grant admin user with grant option:

.. code-block:: bash

    MariaDB [test]> GRANT ALL PRIVILEGES ON test.* TO admin@'localhost' IDENTIFIED
                  > BY 'password' WITH GRANT OPTION;

3. To set the default to UTF-8, you want to add the following to my.cnf

.. code-block:: bash

    [client]
    default-character-set=utf8

    [mysql]
    default-character-set=utf8

    [mysqld]
    collation-server = utf8_unicode_ci
    init-connect='SET NAMES utf8'
    character-set-server = utf8

and do below in db:

.. code-block:: bash

    MariaDB [test]> ALTER TABLE Table CONVERT TO CHARACTER SET utf8 COLLATE utf8_unicode_ci; 


Installation
------------

.. code-block:: bash

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
- http://shadynasty.biz/blog/2012/09/05/auth-and-sessions/
