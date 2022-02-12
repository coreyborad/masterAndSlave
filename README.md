## Description

This is a simple master and slave server with socket, input the numbers, and will output `Mean` `Mode` `Median`

## Table of Contents

- [Description](#description)
- [Table of Contents](#table-of-contents)
- [Preparation](#preparation)
- [Quick Start](#quick-start)
- [Commands](#commands)
- [Architecture](#architecture)

## Preparation

- [docker](https://www.docker.com) - `17.04.0+`
- [docker-compose](https://docs.docker.com/compose) - `1.27.0+`

## Quick Start

- Run docker container
    - Go to project directory
    - Input ```docker-compose up -d```
    - Waiting until finish
    - Check container's status with ```docker-compose ps```
    <img src="./imgs/1.png"  width="50%">
    
    - Client will auto connect to `master` 

- Start `master`
    - Input ```docker-compose exec master go run main.go```
    - Wait all client connect to `master`
    - You will see the output as image here
    <img src="./imgs/2.png"  width="40%">
    
    - When you get `Please input numbers`, you can start input the numbers

## Commands
- Input numbers, every number is split by one space
    - like `20 30 40`
    <img src="./imgs/3.png"  width="25%">
    
    - Also support float number `60.554 40.8 40.8 999`
    <img src="./imgs/4.png"  width="25%">
    
- Quit
    - Input `quit`, you can exit this execution

## Architecture

- Launch server and client
<img src="./imgs/architecture-Page-1.png"  width="75%">

- Input number and broadcast
<img src="./imgs/architecture-Page-2.png"  width="75%">

- Return result to user
<img src="./imgs/architecture-Page-3.png"  width="75%">
