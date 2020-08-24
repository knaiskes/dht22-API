
# Table of Contents

1.  [dht22-API](#org974c2b4)
2.  [Requirements](#orgdbe5deb)
3.  [Setup](#org867a2e3)


<a id="org974c2b4"></a>

# dht22-API

A simple API written in Go for the temperature and humidity dht22 sensor.
It supports GET and POST requests


<a id="orgdbe5deb"></a>

# Requirements

-   Go 1.14 or higher
-   MongoDB
-   mongo-go-driver

OR

-   docker
-   docker-compose


<a id="org867a2e3"></a>

# Setup

Due to a [license](https://lists.archlinux.org/pipermail/arch-dev-public/2019-January/029430.html)
issue many Linux distributions removed MongoDB from their official repositories,
as result it is easier to set up the project with Docker rather
than compiling MongoDB ourselves.

1.  Install docker and docker-compose on Arch based distributions

        $ pacman -S docker docker-compose

2.  Get the source code

        $ git clone git@github.com:KNaiskes/dht22-API.git
        $ cd dht22-API

3.  Run project

    1.  Start docker
    
            $ sudo systemctl start docker
    
    2.  Run project
    
        It will take some time the first time you run this command as it will have to pull
        the Go and MongoDB images
        
            $ docker-compose up # You may have to run this command with sudo
        
        Visit [localhost:8080/measurements](http://localhost:8080/measurements)
        to verify that everything work as they should, you are expected to see an empty list

