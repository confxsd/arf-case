# Currency Conversion Service Case Study for Arf

This document provides useful informations about the project and its deployment process.

**Table of Contents**

1. [About the project](#about-the-project)
2. [How to start](#how-to-start)
3. [Business Logic](#business-logic)

## [About the project](#about-the-project)

This project provides services which users are able to login with their credentials and convert the available currencies they want to. Also users have multiple wallets. Currency rates are static at the moment and defined in a config file.


## [How to start](#how-to-start)

First provide an `app.env` file. Use the `sample.env` format.

System user created initialy as provided in the `app.env` file. It has all the wallets with sufficient balances.

To test the system, create a user & wallets with different currencies.

Use `docker exec -it postgres psql -U ${USER} -d ${DB} -W` to access the db. So you can add balance to the users & modify records.

Run `make build` to build the project with `docker-compose`, then `make up` to run it.

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