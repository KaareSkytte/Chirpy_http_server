# Chirpy

Chirpy is a minimalist Twitter-like REST API built in Go, designed for sharing short, text-based messages called "chirps." This project was developed as part of a Boot.dev course to reinforce backend web development and API design using Go.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
- [Why Chirpy?](#why-chirpy)

## Overview

Chirpy is a backend service that provides endpoints to create and retrieve chirps. It does not include a frontend – the API can be tested with command-line tools like `curl` or API clients like Postman.

## Features

- Create, read, and list chirps via HTTP endpoints
- RESTful API built with idiomatic Go
- Simple project structure for easy understanding and future expansion

## Installation

1. **Clone the repository**
    ```sh
    git clone <your-repo-url>
    cd chirpy
    ```

2. **Build the application**
    ```sh
    go build -o chirpy
    ```

3. **Run the application**
    ```sh
    ./chirpy
    ```
    The server will start (by default) on port 8080.

This project demonstrates practical Go skills, including:

  ## Why Chirpy?

This project demonstrates practical Go skills, including:

- Designing and implementing RESTful APIs
- Structuring Go projects for clarity and maintainability
- Handling JSON payloads and HTTP routing
