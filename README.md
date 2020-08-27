
# Table of Contents

1.  [dht22-API](#orgfadf536)
2.  [Requirements](#org2ff2401)
3.  [Setup](#orge330080)
4.  [GET and POST requests description](#orge594679)
    1.  [GET request](#org5c43d03)
    2.  [POST request](#orgc2d9691)


<a id="orgfadf536"></a>

# dht22-API

A simple API written in Go for the temperature and humidity dht22 sensor.
It supports GET and POST requests


<a id="org2ff2401"></a>

# Requirements

-   Go 1.14 or higher
-   MongoDB
-   mongo-go-driver

OR

-   docker
-   docker-compose


<a id="orge330080"></a>

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


<a id="orge594679"></a>

# GET and POST requests description


<a id="org5c43d03"></a>

## GET request

<table border="2" cellspacing="0" cellpadding="6" rules="groups" frame="hsides">


<colgroup>
<col  class="org-left" />

<col  class="org-left" />
</colgroup>
<thead>
<tr>
<th scope="col" class="org-left">Path</th>
<th scope="col" class="org-left">Result</th>
</tr>
</thead>

<tbody>
<tr>
<td class="org-left">/measurements</td>
<td class="org-left">Returns a list with all measurements or an empty list</td>
</tr>


<tr>
<td class="org-left">/measurements/{id}</td>
<td class="org-left">Returns the one measurement with the specified {id} or 404</td>
</tr>


<tr>
<td class="org-left">/measurements/{name}</td>
<td class="org-left">Returns all measurements from the sensor with the {name} or an empty list</td>
</tr>
</tbody>
</table>


<a id="orgc2d9691"></a>

## POST request

POST request consists of three optional parameters (id is assigned automatically by the server).

Example of POST request with [cURL](https://en.wikipedia.org/wiki/CURL)

    curl -d '{
    
        "name":"basement",
        "temperature":"17.25",
        "humidity":"21.33"
    
    }' -H "Content-Type: application/json" -X POST http://localhost:8080/measurements

