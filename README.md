# BalkanID Engineering Task

## Problem Statement
`Build a robust online book store to handle user authentication, authorization and access management.`

## Salient Features

### Combination of Hexagonal and Onion Architecture
### Background Workers
### Soft Delete
### Email Verification
### Search and Filtering with NLP in Postgres 
### Razorpay Integration


## Process Workflow

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
