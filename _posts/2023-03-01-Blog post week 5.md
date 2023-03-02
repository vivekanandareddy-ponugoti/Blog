---
title: "Go Language learning report 6"
date: 2023-03-01
---

Welcome to my blog post for week six. I feel highly motivated as you are still reading my blogposts. In the process of learning go language and software engineering I have realized that, for the sake of the semester end project, I need to pick a framework that will be suitable for the creation of robotics and IOT projects and make my overall coding experience better.

## Learning a new framework
After exploring the basics of the Go programming language and its applications, I now wanted to move towards a different goal and have started learning about the creation of robotics and IoT projects using the Go programming language. For this, GoLanguage provides an open-source framework named GoBot. In my time exploring Gobot, I’ve learned a lot about the framework and the capabilities it offers for creating innovative projects.

## Writing my first robot(Button and LED)
In order to initialize the Gobot framework, first we need to create a module and initialize Go Mod into it using <code>go mod init file_name</code> After initializing the module, you can now use the command <code>go get -d -u gobot.io/x/gobot</code> to initialize gobot libraries in your module. To create a button and LED robot using GoBot, I have declared the main package and imported some important dependencies like gpio and firmata, which are basically drivers and platforms to run GoBot applications.

Here is the github link to the code I have written. <a href="https://github.com/vivekanandareddy-ponugoti/Blog/tree/weeklyupdate/code/gobot">Button-LED-Robot</a>

## Advantages in using Gobot:
- <a href="https://gobot.io/documentation/examples/">Gobot Examples</a> In this link, there are so many examples of GoBot on many different implementations.
- Gobot is cross-platform compatible with a wide range of hardware platforms and operating systems.
- Gobot is highly extensible, allowing developers to create their own drivers and adaptors for new hardware platforms and devices.
- Gobot supports multiple communication protocols such as Bluetooth Low Energy (BLE), HTTP, MQTT, and more, making it easy to connect to a wide range of devices.
- Gobot has excellent documentation, including examples and tutorials, which makes it easy for developers to get started and learn how to use the framework.

## Challanges in learning Gobot:
- While Gobot has an active community of developers, the framework's documentation may not be as comprehensive as that for other, more popular platforms.
- Debugging code in a robotics or IoT context seems a little challenging as a lot of external libraries are used in making connections possible.

## Conclusion:
I am thoroughly enjoying learning about Go language and exploring its vast array of features. My software engineering class is particularly fascinating as it has sparked my interest in the exciting fields of robotics and IoT, which I have limited experience in. Moving forward, I am committed to continuing my progress in software engineering and will keep updating my progress as I learn more.