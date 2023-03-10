---
title: "Exploring the Versatile World of Go Language and Gobot for Drone Controller Application Development"
date: 2023-03-10
---

Spring break is just around the corner, and my Go language learning journey has been going strong. As I delved deeper into this powerful language, I discovered the vast array of external libraries it offers for building robust and efficient IoT applications. One of the most impressive environments available for IoT development is Gobot, which comes loaded with a plethora of dependencies to support diverse IoT devices.

To start building applications with Gobot, we first need to initialize a Go module in a new folder using the command: <code>gomod init file_name</code>. This sets the stage for creating our Gobot environment by installing all the required packages through the command: <code>go get -d -u gobot.io/x/gobot/</code>. This creates a go.sum file, which allows us to start developing applications seamlessly.

Developing a basic drone controller application with Gobot is a breeze thanks to its comprehensive platform support for various drivers:
- <a href="https://gobot.io/documentation/platforms/mavlink/">mavlink</a>.
- <a href="https://gobot.io/documentation/platforms/tello/">DJI tello</a>
- <a href="https://gobot.io/documentation/platforms/ardrone/">parrot</a>
- <a href="https://gobot.io/documentation/platforms/sphero/">sphero</a>

Each of these drivers provides support for different protocols and sensors, making it easy to create customized applications for different types of drones.

In this blog, I will focus on using the mavlink protocol for drone controller application development, which is widely used in most drone controller boards. As part of our software engineering class, we are building a mavlink protocol model of an application to communicate with the drone.

Overall, my experience with learning Go language and exploring its applications in IoT devices has been incredibly enjoyable. Despite the relatively fewer online resources compared to other popular programming languages, the documentation is comprehensive and user-friendly. Additionally, the Gobot website provides boilerplate code for all platforms, as seen here: <a href="https://gobot.io/documentation/platforms/mavlink/#:~:text=Guides-,Platforms,-Drivers">link</a>. As we plan to continue exploring Go after spring break, I encourage you to join me in discovering the endless possibilities of this versatile language. Wishing you all a pleasant break and thank you for reading!"
