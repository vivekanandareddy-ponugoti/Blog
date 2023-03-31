---
title: "A simple library management system using go language"
date: 2023-03-31
---

## Welcome Back After Spring Break
Hello everyone, I hope you had a fantastic spring break. It's been a week since my university resumed classes after the break, and I have made some progress in learning Go language. Our professor, Kate Holdener, has been challenging all the students by assigning tasks to help improve our pragmatic software engineering skills. We recently participated in a mini-project to build an application using the language we're currently learning - in my case, it's the Go programming language. We formed groups based on the students' preferences for learning a new language. My team consisted of five members, and together we built a library book management system with two primary objects - books and users. My main contribution to this project was focused on the backend, setting up a database and implementing controller functions for the required actions.

## Initial Stages of Development
During our first group discussion, we brainstormed ideas and shared our previous experiences working with Go. Eventually, we decided to create a library book management system. In this system, the admin would have access to get, add, delete, and update functionalities, while the user would have access to only get_books and checkout (which updates the book's is_borrowed variable to true). I successfully implemented all the backend functionalities as per the initial model proposed. However, the GUI part, which was divided among other teammates, did not turn out exactly as planned. Although our team members committed to creating a user-friendly GUI, we couldn't achieve what was initially proposed. Nonetheless, we made sure to satisfy all the project requirements, which is a commendable achievement.

Throughout this mini-project, I learned how to build a GUI in Go language using packages like Fyne, and how to make the GUI interact with the backend using HTTP requests and responses.

## Code Description
The main focus of our mini-project was the backend and its architecture, where we needed two models (book and user) with the following structure parameters:

type Book struct {
	gorm.Model
	Title      string `json:"title"`
	Author     string `json:"author"`
	Publisher  string `json:"publisher"`
	ISBN       string `json:"isbn"`
	Desc       string `json:"desc"`
	IsBorrowed bool   `json:"isBorrowed"`
	UserID     uint   `gorm:"index"`
}

type User struct {
	gorm.Model
	IsAdmin       bool   `json:"isAdmin"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Address       string `json:"address"`
	PhoneNumber   string `json:"phoneNumber"`
	BorrowedBooks []Book `gorm:"foreignkey:UserID" json:"borrowedBooks"`
}

After finalizing the models, we proceeded to create controllers to perform CRUD operations for getting, adding, deleting, and updating the models. I then integrated my previous code from a book management system project to use a MySQL database for updating the controllers. Meanwhile, other team members experienced in creating GUIs with Fyne built a basic login interface and utilized dummy instances to display data using CSV. We configured the GUI to connect to the backend server and ensured a cohesive program between the GUI and backend. Finally, we refactored the code to remove redundancies and set separate routes for book and user controllers.

You can find the GitHub link to our project below. The README file for the application provides guidance on how to run the program.
<a href="https://github.com/SLUSE-Spring2022/miniproject-go">Library book management system</a>

## Challenges Faced
Working in a team of five presented some communication challenges due to the increased number of communication paths. We addressed this by conducting more group discussions and maintaining close contact during development. Additionally, we faced language barrier issues, as one of our teammates had trouble understanding the initial idea and communicating his thoughts.

## Conclusion
In conclusion, I would like to express my appreciation for my team's collaboration on this mini-project. I am grateful for their dedication and commitment to making this project a success. Thank you all for reading my blog. I will be back with more content next week. Until then, have a great weekend!