# BalkanID Engineering Task

## Problem Statement
Build a robust online book store to handle user authentication, authorization and access management.

## Setup

1. Clone the repository repository
```
https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal.git
```

2. Install go-migrate tool
```
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
        sudo mv migrate.linux-amd64 /usr/bin/migrate
        which migrate
```

3. Now using docker and make file (make sure docker and make are installed)
```
make postgres
```
```
make createbd
```
```
make migrateup
```
```
make run
```

## Salient Features
Here are a list of features which I built appart from the demanded features

### Combination of Hexagonal and Onion Architecture
I have made the backend in a microservice architecture which is highly decoupled. The is because of the use of ports and adapters throughtout my code (basically different interfaces for different jobs). And took some inspiration from the onion architecture to build my logging service completly detached from my bussiness logic. 
![image](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/assets/95216160/ad1a6ec6-d076-497b-9503-ea48ac3580d0)
![image](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/assets/95216160/2f0ba9ee-fde5-4b54-9c30-e06dd26e59d9)


### Background Workers
I wrote the background workers from scratch for the task which should not block the user to further request my api. Like account deletion and email service. I have used semaphore to make a simple task queue for my worker. You can see the power of background workers here

[Screencast from 02-09-23 07:29:20 PM IST.webm](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/assets/95216160/2f64e370-73e7-4ca1-af2b-faa479bb9ba9)


### Soft Delete
Keeping data retention policy in mind I have implemented soft delete operation which basically marks the user as deleted but is not actualy deleted. The worker queue will delete the users in a batch operation

### Email Verification
Email verification for security reasons and can be used to send the receipt when order is done .

### Search and Filtering with NLP in Postgres 
I have made a complete fuzzy search implementation using NLP in postgres. Uses lexemes and similarity match in postgres to rank the results and give it to users. 
![image](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/assets/95216160/e139a29f-db1a-47b1-8909-94ebb250ba33)

You can see the demo here where I intensionaly put wrong spelling of `jane eyre` as `jane iyre` but still got the correct result
![image](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/assets/95216160/619a95b6-b54d-4e4a-8643-0ab7bfe5440e)

### Razorpay Integration
Complete razorpay integration for orders.

![image](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/assets/95216160/1f3f6136-2b28-4a87-806c-7b0070f881f7)

First adding to cart

[Screencast from 02-09-23 07:46:57 PM IST.webm](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/assets/95216160/5d568f94-29e4-43e2-baa2-a373e7ebcd66)

Second placing the order

[Screencast from 02-09-23 07:49:03 PM IST.webm](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/assets/95216160/47c3a84a-e0f4-4ae0-9fc5-939ef10e993d)


### Recommendations
users gets recommendations of the basis of their last purchase. Also I have added a hot-selling-books endpoint which you can try out

### Robust database design
![BalkanID-Book-Management-System](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/assets/95216160/766f3b3f-f1ee-4eee-af3c-396f849225a8)
