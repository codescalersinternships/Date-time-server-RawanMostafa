# Datetime Server 

This repository implements http and Gin datetime servers.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)


## Installation

1. Clone the repository

   ```bash
   git clone https://github.com/codescalersinternships/Datetime-server-RawanMostafa.git
   ```

2. Install the dependencies
    ```bash
    go mod download
    ```

## Usage

### 1. Using Makefile

 To run all make targets
   
  ```bash
    make all
  ```

### 2. Using docker-compose

   
  ```bash
    docker-compose up
  ```

### 2. Using kubernetes deployed server

   
  ```bash
    curl http://185.206.122.17:30100/
  ```

### 3. Using main.go
   
  ```bash
    go run cmd/httpserver/main.go
  ```
  ```bash
    go run cmd/ginserver/main.go
  ```
