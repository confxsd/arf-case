# Currency Conversion Service Case Study for Arf

This document provides useful informations about the project and its deployment process.

**Table of Contents**

1. [About the project](#about-the-project)
2. [How to start](#how-to-start)

## [About the project](#about-the-project)

This project provides services which users are able to login with their credentials and convert the available currencies they want to. Also users have multiple wallets. Currency rates are static at the moment and defined in a config file.


## [How to start](#how-to-start)

First provide an `app.env` file. Use the `sample.env` format.

Run `make build` to build the project with `docker-compose`, then `make up` to run it.