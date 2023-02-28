---
title: "Book management system using go language"
date: 2023-02-22
---

Welcome to the sixth week of my blog post. In the process of learning go for software engineering i have a clear understanding of go syntax. So, in the past week, I have made progress by building a book management system using a mySQL database.

## Using GoLang in building web severs to perform CRUD operations:
At first, I had some understanding of databases and their basic operations, but I had never used Go language to interact with databases. I decided to learn about Go's built-in database/sql package, which provides a set of interfaces to interact with relational databases.

## A book management system using go language and MySQL:
After acquiring all the necessary information, I have decided to make a book management system that will perform CRUD operations to create, remove, update, and delete data (books) from a MySQL database. The packages used in this project are simple and straightforward, with the Gorilla Mux package being used to create custom routes for HTTP endpoints and handle various HTTP methods via an API. I also used the GORM package, which consists of multiple database functions like mySQL, SQLLite, and PostgreSQL functions provided by Dr. Jinzhu at <code>"github.com/jinzhu/gorm"</code>. I chose MySQL for this project as it is a relational database, which is kind of suitable for this project, and I also used <code>encoding/json</code> package, which provides functionality for encoding and decoding data in JSON format. In this package, JSON marshall and unmarshall functions can be used to serialize and de-serialize data from and into JSON format, respectively. This will make it easier to exchange data between different systems that use JSON. To design the structure of the project, I have used the basic model-view-controller architecture to make the project better coupled and cohesive. I made custom routes for their HTTP endpoints to handle different HTTP methods (e.g., GET, POST, PUT, DELETE), and extract parameters from URLs. Then, I created a connection to the mySQL database, which has a table of books with details like name, author, publication, date created, date deleted, and date updated. Here is the link to my code, and you can find some images of output in the repository. 
<<<<<<< HEAD
<a href="https://github.com/vivekanandareddy-ponugoti/Blog/tree/main/code/CRUD-Book-Management-System">CRUD-Book_Management-System</a>
=======
<code>https://github.com/vivekanandareddy-ponugoti/Blog/tree/main/code/CRUD-Book-Management-System</code>

>>>>>>> refs/remotes/origin/weeklyupdate
I also tested the whole backend using Postman and verified if the database was updated.

## My learnings while building a web server:
As I reflect upon my experience in creating a book-management system in Go language, I find myself thoroughly satisfied with the straightforwardness of the process. The availability of the gorm/sql package, which offers a simple and consistent interface for interacting with databases, proved to be a valuable asset in the development process. Its integration with Go language facilitated a seamless transition in performing basic CRUD operations, which allowed me to concentrate on the core aspects of the project.

In addition to the ease of use of the gorm/sql package, the comprehensiveness of its documentation made it possible to quickly comprehend its workings and implement its features effectively. The examples provided were clear and concise, making it possible to build on my knowledge of Go and databases with relative ease. The package's flexibility and modularity ensured that I could create a customized database management system that met my specific project requirements.

As I progress in my knowledge of Go and databases, I look forward to exploring the more advanced features of this powerful language and the wide variety of databases that it can interact with. The possibilities for optimization and increased efficiency in data management are endless, and I am excited to continue learning and expanding my skill set in this area. Ultimately, the creation of this book-management system has not only enhanced my understanding of Go and databases, but also opened up new doors for innovative and efficient data management solutions in the future.

## Challanges faced in my learnings
In building a book management system, I encountered several challenges, primarily in the design of the architecture and establishing a connection to the database. Although the issues largely involved syntactical errors, I was able to overcome them by identifying and applying the correct syntax.

## Conclusion:
In conclusion, my journey to learn the Go programming language has been an enlightening experience. From the basics of syntax and data types to more advanced concepts such as concurrency and error handling, I have gained a comprehensive understanding of the language and efficiently manage memory for building microservices and other distributed systems.
