---
title: "A simple library management system using go language"
date: 2023-03-31
---

## Greetings and Post-Spring Break Excitement
Hello everyone! I hope you all had an amazing spring break. It's been a week since my university resumed classes, and I've been making great strides in learning the Go programming language. Our professor, Kate Holdener, has been challenging us by assigning tasks designed to improve our pragmatic software engineering skills. As a part of this learning process, we recently participated in a mini-project where we were tasked with building an application using our chosen programming language. In my case, it's the Go programming language. We formed groups based on students' preferences for learning a new language, and I teamed up with four other enthusiastic learners. Together, we built a library book management system focusing on two primary entities: books and users. My main contribution to this project was concentrated on the backend, setting up a database and implementing controller functions for the required actions.

## Embarking on the Development Journey
During our first group discussion, we brainstormed ideas and shared our prior experiences with Go. After much deliberation, we settled on creating a library book management system. The system we envisioned would allow the admin to access get, add, delete, and update functionalities, while the user would only have access to get_books and checkout (which would update the book's is_borrowed variable to true). I successfully implemented all backend functionalities as per the initial model proposed. However, the GUI part, which was divided among other teammates, didn't turn out exactly as planned. Although our team members were committed to creating a user-friendly GUI, we couldn't achieve what was initially proposed. Nonetheless, we made sure to satisfy all the project requirements, which is a noteworthy accomplishment.

Throughout this mini-project, I learned how to build a GUI in Go language using packages like Fyne and how to make the GUI interact with the backend using HTTP requests and responses.

## A Closer Look at the Code
The primary focus of our mini-project was the backend and its architecture, which required two models: book and user. After finalizing the models, we proceeded to create controllers to perform CRUD operations for getting, adding, deleting, and updating the models. I then integrated my previous code from a book management system project to use a MySQL database for updating the controllers. Meanwhile, other team members skilled in creating GUIs with Fyne built a basic login interface and utilized dummy instances to display data using CSV. We configured the GUI to connect to the backend server and ensured a seamless integration between the GUI and the backend. Finally, we refactored the code to remove redundancies and set separate routes for book and user controllers.

You can find the GitHub link to our project below. The README file for the application provides guidance on how to run the program.
<a href="https://github.com/SLUSE-Spring2022/miniproject-go">Library book management system</a>

## Overcoming Challenges
Working in a team of five presented certain communication challenges due to the increased number of communication paths. We tackled this by conducting more group discussions and maintaining close contact throughout the development process. Additionally, we faced language barrier issues, as one of our teammates had trouble understanding the initial idea and communicating his thoughts. To overcome this, we made sure to encourage open communication and provided additional clarification when necessary.

## Conclusion and Gratitude
In conclusion, I would like to express my sincere appreciation for my team's collaboration on this mini-project. Their dedication and commitment played a significant role in making this project a success. I am grateful for the opportunity to learn and grow alongside such a talented group of individuals. Thank you all for reading my blog. I will be back with more exciting content next week. Until then, have a fantastic weekend!

As we move forward, I'm eager to continue refining my skills in the Go programming language and further explore the potential of this powerful language in various applications. The experience gained from this mini-project has not only enriched my understanding of Go but also provided valuable lessons in teamwork, communication, and problem-solving.

If you're interested in learning more about the Go language or library management systems, feel free to leave a comment or share your thoughts. I'm always open to engaging in discussions and exchanging ideas with fellow developers and enthusiasts. Stay tuned for more updates on my learning journey and adventures in software engineering!