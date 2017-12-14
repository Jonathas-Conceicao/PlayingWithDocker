# PlayingWithDocker
A simple application to play with Docker

## Application Running

A simple **Go**(lang) application that prints a message to a _html_ page and the number of access it has received.

## Docker compose

The docker compose will build the image for the go web application and launch a container for it and for the redis image.

## Running
**docker-compose build**             # will build the image for the Go web app
**docker-compose up**                # will run both the images. The App will be avaliabe at port 8080
**docker-compose down --rmi local**  # will remove the containers created and remove the go web app image
