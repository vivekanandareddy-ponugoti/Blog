---
title: "Exploring the Versatile World of Go Language and Gobot for Drone Controller Application Development"
date: 2023-03-10
---

Spring break is around the corner, and I am thrilled to share my exciting journey with Go language and IoT application development. As I continue to explore this powerful language, I have discovered the extensive range of external libraries and packages it offers, making it an ideal choice for building robust and efficient IoT applications.

One of the most impressive environments available for IoT development is Gobot, which provides a plethora of dependencies to support diverse IoT devices. It has an extensive library of platform-specific drivers for different IoT devices and protocols, making it easier to create customized applications. I found Gobot to be highly intuitive, user-friendly, and easy to use, especially with its comprehensive documentation and pre-built code samples.

To start building applications with Gobot, I first initialized a Go module in a new folder using the command <code>gomod init file_name</code>. This allowed me to create a Gobot environment by installing all the required packages through the command <code>go get -d -u gobot.io/x/gobot/</code>. This creates a go.sum file, which enables me to develop applications seamlessly. The documentation for Gobot is highly detailed and offers a wealth of information, including pre-built code samples that helped me understand the basics and get started with building applications.

Gobot offers support for a wide range of devices, including drones, robots, and other IoT devices, making it highly versatile. I found building a basic drone controller application with Gobot to be a breeze thanks to its comprehensive platform support for various drivers, including Mavlink, DJI Tello, Parrot, and Sphero. Each of these drivers provides support for different protocols and sensors, making it easy to create customized applications for different types of drones.

As part of our software engineering class, we are building a Mavlink protocol model of an application to communicate with the drone. The Mavlink protocol is widely used in most drone controller boards and provides a simple and robust method for communicating with drones.

Another aspect of Gobot that I found particularly impressive is its built-in support for concurrency. Go's lightweight Goroutines make it possible to run multiple tasks concurrently without the overhead of traditional threads. This is particularly useful in IoT applications, where efficient resource utilization is critical for maintaining responsiveness and real-time performance.

Additionally, Go has excellent support for networking and communication protocols, such as MQTT and CoAP, which are widely used in IoT applications. This makes it easier to implement and manage the communication between IoT devices and the cloud or other devices in the network.

Gobot's community support is another vital aspect that helped me throughout my learning journey. The active community around the project ensures that developers can quickly get answers to their questions and share their experiences, which significantly improves the overall learning process.

Furthermore, Go's strong typing and static compilation make it easier to catch errors early in the development process, ensuring high-quality code and reducing the chances of runtime errors. This is particularly important when developing applications for IoT devices, where reliability is crucial.

In conclusion, my experience with learning the Go language and exploring its applications in IoT devices has been incredibly enjoyable. Although there are relatively fewer online resources for Go compared to other popular programming languages, the documentation is comprehensive and user-friendly. The Gobot website also provides boilerplate code for all platforms, making it easier to get started with building applications. The powerful combination of Go's efficiency, Gobot's extensive support for IoT devices, and the active community make it a top choice for developers looking to build innovative and reliable IoT applications.
