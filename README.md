# autoshenas-mhap/socket

## Overview
`autoshenas-mhap/socket` is a real-time chat application that uses WebSockets for efficient bi-directional communication. This project demonstrates how WebSockets can be used to create a simple, yet functional chat interface where users can send and receive messages instantly. 

## Features
- Real-time messaging
- Simple and intuitive user interface
- WebSocket-based communication
- Lightweight and scalable design

## Requirements
- A modern web browser supporting JavaScript and WebSockets
- Go (for running the WebSocket server)
- Visual Studio Code with the 'Go Live' extension (for running the frontend)

## Installation
To set up the project locally, follow these steps:

1. **Clone the Repository**
    ```bash
    git clone https://github.com/autoshenas-mhap/socket.git
    cd socket
    ```

2. **Setting Up the Server**
    - Make sure you have Go installed on your system.
    - Run the WebSocket server using the command:
      ```bash
      go run main.go
      ```
    - This will start the server, which listens for incoming WebSocket connections.

3. **Running the Frontend**
    - Open the project in Visual Studio Code.
    - Use the 'Go Live' extension to serve the frontend files.
    - Open `index.html` in a web browser through the live server to launch the chat application.
