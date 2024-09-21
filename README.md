# Datetime Server 

This repository implements http and Gin datetime servers.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)


## Installation

1. Clone the repository

   ```bash
   git clone github.com/codescalersinternships/Datetime-server-RawanMostafa/pkg
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
 To run a specific make target
   
   ```bash
    make <target_name>
   ```

### 2. Using docker-compose

   
  ```bash
    docker-compose up
  ```
### 3. Using main.go
   
  ```bash
    go run cmd/httpserver/main.go
  ```
  ```bash
    go run cmd/ginserver/main.go
  ```
