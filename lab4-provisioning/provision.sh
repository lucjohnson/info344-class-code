#!/usr/bin/env bash

# Provisioning script for the authentication challenge
echo "Hello, world!"

# use noninteractive mode since this is an automated script
export DEBIAN_FRONTEND=noninteractive

# update the package database, -E tells it to pass the environment variables (like the DEBIAN_FRONTEND variable)
sudo -E apt-get update

# install git, -y represents yes for when apt-get asks if you're sure you want to download it
sudo -E apt-get install -y git

# install Node.js v4.x
curl -sL https://deb.nodesource.com/setup_4.x | sudo -E bash -
sudo -E apt-get install -y nodejs

# install build-essential for Node modules w/ native code
sudo -E apt-get install -y build-essential

# allow Node.js servers to bind to low ports
sudo -E apt-get install -y chase
sudo -E apt-get install -y libcap2-bin
sudo -E setcap cap_net_bind_service=+ep $(chase $(which node))

# install MySQL
sudo -E apt-get install -y mysql-server

# run the SQL commands for the mysql_secure_installation script
# except for setting the root password (so that we don't embed)
# the root password in our provisioning script
mysql -u root <<-EOF
DELETE FROM mysql.user WHERE User='root' AND Host NOT IN ('localhost', '127.0.0.1', '::1');
DELETE FROM mysql.user WHERE User='';
DELETE FROM mysql.db WHERE Db='test' OR Db='test\_%';
FLUSH PRIVILEGES;
EOF

# install mongo
sudo -E apt-key adv --keyserver hkp://keyserver.ubuntue.com:80 --recv EA312927
echo "deb http://repo.mongodb.org/apt/ubuntu trusty/mongodb-org/3.2 multiverse" | sudo -E tee /etc/apt/sources.list.d/mongodb-org-3.2.list
sudo -E apt-get update
sudo -E apt-get install -y mongodb-org

# install the kerberos library
# which is needed to build the MongoDB
# Node.js driver during npm install
sudo -E apt-get install -y libkrb5-dev

# install pm2 utility for managing node servers
sudo -E npm install -g pm2 --loglevel error

# install the tsd utility for installing
# Visual Studio Code typings files
# gives statement completion and parameter hinting
sudo -E npm install -g tsd --loglevel error