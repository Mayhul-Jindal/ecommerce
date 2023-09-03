[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/LECuYE4o)
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
make createdb
```
```
make migrateup
```
```
make run
```

4. Refer to postman if any issues faced while requesting
```
https://www.postman.com/mission-physicist-26981670/workspace/balkanid-book-store
```

## Salient Features
Here are a list of features which I built appart from the demanded features

### Combination of Hexagonal and Onion Architecture
I have made the backend in a microservice style which is highly decoupled. This is because ,I used ports and adapters architecture throughtout my code (basically different interfaces for different jobs). And took some inspiration from the onion architecture to build my logging service completly detached from my bussiness logic. 

[Screencast from 02-09-23 08:16:42 PM IST.webm](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/assets/95216160/81555b0c-c800-4506-bfa4-a52abc393144)


I have implemented logs with tracing which can help export the logs to fluentd for further analysis
![image](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/assets/95216160/47677e88-2ac1-4bc0-a6e2-bd2368df247f)

All the errors are handled at a single place so you dont have to look here are there in the code to handle the errors
![image](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/assets/95216160/38e7745e-37e5-45d6-b94e-e60d302d4754)

### Background Workers
I wrote the background workers from scratch for the task which should not block the user to further request my api. Like account deletion and email service. I have used semaphore to make a simple task queue for my worker. You can see the power of background workers here

[Screencast from 02-09-23 07:29:20 PM IST.webm](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/assets/95216160/2f64e370-73e7-4ca1-af2b-faa479bb9ba9)

The semaphore implementation is below. Lokks very small but it works like a charm (basically prevents task to enter into the criticial section if the channel if full)
![image](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/assets/95216160/ad8e67c6-6114-48c8-82dd-bb285fcd6bf5)



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
