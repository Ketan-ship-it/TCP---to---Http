# 🔵 HTTP Server from Scratch in Go

![TCP](https://img.shields.io/badge/Protocol-TCP-green)
![HTTP](https://img.shields.io/badge/Protocol-HTTP-blue)
![Go](https://img.shields.io/badge/Language-Go-lightblue)

---

## 📌 Overview
This project is an implementation of a simple **HTTP server** built **from scratch** on top of **TCP sockets** in Go.  
Instead of relying on Go's standard `net/http` package, the server directly uses low-level TCP connections to parse HTTP requests and send responses manually.  

This helps in understanding:
- How **TCP** provides the reliable transport layer.
- How **HTTP** is built on top of TCP.
- How web servers process client requests and send structured responses.

---

## 🟢 What is TCP?
- **Transmission Control Protocol (TCP)** is a core protocol of the Internet protocol suite.
- It provides a **reliable, ordered, and error-checked** stream of communication between applications.
- Key features:
  - 🔗 Connection-oriented (requires handshake before communication).  
  - 📦 Guarantees delivery of packets in order.  
  - 🌍 Used by many application protocols (e.g., HTTP, FTP, SMTP).  

---

## 🟣 What is HTTP?
- **Hypertext Transfer Protocol (HTTP)** is an **application layer protocol** that runs on top of TCP.
- It defines rules for communication between **clients (browsers)** and **servers** on the web.
- HTTP messages consist of:
  - 📝 A **request line** (e.g., `GET /index.html HTTP/1.1`).  
  - 📑 **Headers** (metadata like `Host`, `Content-Type`, etc.).  
  - 📦 An optional **body** (used in `POST` requests).  
- Servers reply with:
  - ✅ A **status line** (e.g., `HTTP/1.1 200 OK`).  
  - 📑 **Headers**.  
  - 📂 A **response body** (HTML, JSON, files, etc.).  

---

## 🎯 Project Goal
- Implement a **basic HTTP server** in Go using raw TCP sockets.  
- Handle incoming connections, parse raw HTTP requests, and return HTTP responses.  
- Gain deeper insight into how higher-level protocols (HTTP) are built over lower-level protocols (TCP).  

---

## 🛠️ Tech Stack
- **Language**: Go (Golang)  
- **Protocols**: TCP (transport layer), HTTP (application layer)  

---

## 🚀 Future Enhancements
- Support for multiple routes.  
- Handling different HTTP methods (`GET`, `POST`, etc.).  
- Parsing and responding with JSON.  
- Basic concurrency for handling multiple clients.  

---
