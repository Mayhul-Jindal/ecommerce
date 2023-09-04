[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/LECuYE4o)
# BalkanID Engineering Task

## Problem Statement
Build a robust online book store to handle user authentication, authorization and access management.

## Setup

1. Clone the repository repository
```
https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal.git
```

2. Install docker and docker-compose (for ubunutu)
```
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

3. Run docker-compose up
```
docker-compose up
```

## Salient Features

### Combination of Hexagonal and Onion Architecture
My code is highly decoupled. This is possible because I have used `ports and adapters architecture` throughtout my code base(basically different interfaces for different jobs). Took some inspiration from the `onion architecture` to build my logging service completly detached from my bussiness logic and to make middlewares for my json api 

- Power of `ports and adapter architecture` (high level overview)
 ![ports and adapters example(1)](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/assets/95216160/ec096e8b-485a-465a-9461-8e43f008129b)

- Power of `onion architecture` (high level overview)
![onion architecture example(1)](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/assets/95216160/759e7fe9-409d-44bf-8ae4-9f20edf1d016)


### Background Workers
I wrote the `background workers from scratch` for the task which should not block the user to further request my api (eg:- while email is sent to user and when user deleted their account). I have used `semaphore` to make a simple task queue for my worker. You can see the power of background workers here

- Working of email background worker for `email verification`

[email_worker.webm](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/assets/95216160/feb4b57e-7036-418d-82b5-e091413f6539)


- Working of deletion operation in background to support `soft delete`



### Search and Filtering with NLP in Postgres 
I have made a complete `fuzzy search` implementation using NLP in postgres. Uses `stop word removal`, `lexeme` formation and `similarity match` in postgres to rank the results and give it to users. 

-- todo yaha image ayegi

### Razorpay Integration
Complete razorpay integration for orders. Operational flow is like

![image](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/assets/95216160/1f3f6136-2b28-4a87-806c-7b0070f881f7)

1. Add books to cart

[Screencast from 02-09-23 07:46:57 PM IST.webm](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/assets/95216160/5d568f94-29e4-43e2-baa2-a373e7ebcd66)

2. Place order

[Screencast from 02-09-23 07:49:03 PM IST.webm](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/assets/95216160/47c3a84a-e0f4-4ae0-9fc5-939ef10e993d)







## yeh kisi aur heading ke neeche ayenge

### Recommendations
users gets recommendations of the basis of their last purchase. Also I have added a hot-selling-books endpoint which you can try out

### Robust database design
![BalkanID-Book-Management-System](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/assets/95216160/766f3b3f-f1ee-4eee-af3c-396f849225a8)
