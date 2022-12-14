# Currency Conversion Service Case Study for Arf

This document provides useful informations about the project and its deployment process.

**Table of Contents**

1. [About the project](#about-the-project)
2. [How to start](#how-to-start)
    1. [Run repo via git-clone](#run-repo)
    2. [Test sample endpoints](#sample-endpoints)
    3. [Swagger docs](#docs)
4. [Business Logic](#business-logic)
5. [System Design](#system-design)
    1. [Data models](#data-models)
    2. [API Layers](api-layers)


## [About the project](#about-the-project)

This project provides services which users are able to login with their credentials and convert the available currencies they want to. Also users have multiple wallets. Currency rates are static at the moment and defined in a config file.


## [How to start](#how-to-start)

### [Run repo via git-clone](#run-repo)

First provide an `app.env` file. Use the `sample.env` format. You can copy or change the name of  `sample.dev.env`.

System user created initialy as provided in the `app.env` file. It has all the wallets with sufficient balances.

To test the system, create a user & wallets with different currencies.

Use `docker exec -it postgres psql -U ${USER} -d ${DB} -W` to access the db. So you can add balance to the users & modify records.

Run `make build` to build the project with `docker-compose`, then `make up` to run it.


### [Sample endpoints](#sample-endpoints)
I tested endpoints in Postman, so here you can see the sample requests.

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/23017395-02e9cd2a-6e96-4c7a-b5f8-2c6318662f99?action=collection%2Ffork&collection-url=entityId%3D23017395-02e9cd2a-6e96-4c7a-b5f8-2c6318662f99%26entityType%3Dcollection%26workspaceId%3Db5818ecc-197b-4659-a846-3c1c9d4ad497#?env%5Blocal%5D=W3sia2V5IjoiYmFzZV91cmwiLCJ2YWx1ZSI6ImxvY2FsaG9zdDozMzMzIiwiZW5hYmxlZCI6dHJ1ZSwidHlwZSI6ImRlZmF1bHQifV0=)

### [Swagger docs](#docs)

Hit `{BASE_URL}/swagger/index.html` to see the generated swagger endpoint list.


## [System Design](#system-design)

### [Data Models](#data-models)
Basic data blocks used in the system:

![Arf case - Data models](https://user-images.githubusercontent.com/105401425/203045174-62f2a1a4-40f6-49c0-86af-75801a07e65f.png)

### [API Layers: Offer service](api-layers)
`Offer` service implemented as a 2-step process which requires an approval after it. 

Endpoints are isolated to separate DB domains.

![arf case - API Layers - Offer](https://user-images.githubusercontent.com/105401425/203050320-420d1c00-8451-4a87-a76c-736bbdf1210c.png)


## [Business Logic](#business-logic)

The main point of the conversion is creating `Offer`, which has the required conversion info. `rate` is tricky since I assumed that frontend fetches the rate from db & use it to request an offer. Then we check if it's same with the one in our system.

### How to determine rates according to the markup rate
I used a simple logic as below:

```
e.g

from: USD
to: TRY
amount: 100
rate: 10
markupRate: 0.1
rate * (1 - markupRate) = 9

system should have at least (100) TRY
user should have at least (100/9) = 11.1 USD
(it's supposed to be 10 USD if no markup applied)
```

Basically, instead of changing the currency rate, I charged user more according to the markup rate. So user pays more to get the desired amount. There could be alternatives but I thought it would still make us profitable & also practical to implement.


### What about freezing the currency rate?

I used `Offer` object with a `timestamp`. System requests the approval of this offer. Since `Offer` has a rate and timestamp in it, I check the timestamp difference when approve request hit. If 3+ minutes have passed, I reject that offer by responding an offer timeout. 



### What about tests

I really tried, but couldn't cover all of them due to the lack of time & some platform issues. I still think I could cover most of them over a wider period of time since I implemented the test suits, mocks & stubs. 


Cheers ^^