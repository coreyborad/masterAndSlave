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
    1. Go to project directory
    2. Input ```docker-compose up -d```
    3. Waiting until finish
    4. Check container's status with ```docker-compose ps```
    5. <img src="./imgs/1.png"  width="75%">
    6. Client will auto connect to `master` 

- Start `master`
    1. Input ```docker-compose exec master go run main.go```
    2. Wait all client connect to `master`
    3. You will see the output as image here
    4. <img src="./imgs/2.png"  width="75%">
    5. When you get `Please input numbers`, you can start input the numbers

## Commands
- Input numbers, every number is split by one space
    1. like `20 30 40`
    <img src="./imgs/3.png"  width="50%">
    2. Also support float number `60.554 40.8 40.8 999`
    <img src="./imgs/4.png"  width="50%">
- Quit
    1. Input `quit`, you can exit this execution

## Architecture

- Launch server and client
<img src="./imgs/architecture-Page-1.png"  width="75%">

- Input number and broadcast
<img src="./imgs/architecture-Page-2.png"  width="75%">

- Return result to user
<img src="./imgs/architecture-Page-3.png"  width="75%">
