# BalkanID Engineering Task

## Problem Statement
Build a robust online book store to handle user authentication, authorization and access management.

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

You can see the demo here



### Razorpay Integration
Complete razorpay integration for orders

### Recommendations
users gets recommendations of the basis of their last purchase. Also I have added a hot-selling-books endpoint

## Process Workflow
The complete workflow is as follows:



## Technical Architecture

## Customizability and Maintainability




## Demand Features Status

### Key Features
- [X] Secure user registration and authentication
- [ ] Account Deactivation and Deletion: Allow users to deactivate or delete their accounts, if applicable. Implement a mechanism to handle account deletion securely while considering data retention policies.
- [X] Protection against vulnerabilities like SQL injection attacks
- [X] Have Proper system logging with retention policies upon system failure
- [ ] Users can easily search and filter books and add them to shopping cart
- [ ] Users can easily download their bought books and leave a review on the books they bought
- [ ] Admins have the ability to manage inventory and others 

### Requirements
- [ ] Make the necessary APIs to expose (go micro-service)
- [ ] SQL based database (PostgreSQL)
- [ ] Use a reverse proxy of your choice (nginx)

### Bonus Points
- [ ] You can use Docker and containerize your application code to run, including the database
- [ ] You can test your code by adding unit test cases and workflow test cases
- [ ] You can add recommendation system to recommend books to user

### Detailed Solution
- ports and adapter/ hex/ onion you can call it what you want. I just made the code maintainable
- why paesto instead of jwt. more secure and shit
- Dont need for uneccessary folders. In my opinion increases complexity, I like to keep it simple

- <iframe width="560" height="315" src='https://dbdiagram.io/embed/64e4cc4c02bd1c4a5e353140'> </iframe>

- why migrate tool is needed 
- database vs gorm vs other 

- database indexes filtering searching
- database ke models use honge saari jagah
- change name of database to postgres


- Soft delete allows you to mark records as deleted without actually removing them from the database. This can be useful for scenarios where you want to retain the data for auditing or historical purposes.

- tags se better seo, recommendations, search optimization 

- JWT give developers too many algorithms to choos, some algorithms are known to be vulnerable(ECDSA, RSA)
- trivial forgery "alg" to "none"
- it is crutial to check eaders jisse symetric keys aur asymteric keys wala issue na hoajye
- you can see there are many issue with jwt that is why use paesto (it follows the best practices)


- grpcs I dont think karne ki zaroorat hain
