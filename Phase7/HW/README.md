# Docker Homework Guide

Welcome! This guide will walk you step-by-step through essential Docker tasks. Each section explains the concepts, commands, and best practices, so you can confidently use Docker in real-world DevOps environments.

---

## 1. Build and Run Your First Docker Container

**Docker** is a platform for building, running, and managing containers.  
A **container** is a lightweight, portable package that includes everything needed to run an application (code, runtime, libraries, and settings).

**Steps:**

1. **Write a Dockerfile:**  
   A `Dockerfile` is a text file with instructions to build a Docker image.

   ```dockerfile
   # Example Dockerfile
   FROM ubuntu:latest
   CMD ["echo", "Hello, Docker!"]
   ```

2. **Build the Image:**  
   Run this command in your terminal (in the folder with your Dockerfile):

   ```bash
   docker build -t my-first-image .
   ```

   - `docker build`: Command to build an image.
   - `-t my-first-image`: Tags the image with a name.
   - `.`: The build context (current directory).

3. **Run the Container:**  

   ```bash
   docker run my-first-image
   ```

   - `docker run`: Starts a container from the image.

**Best Practice:**  
Always use descriptive image names and tags.

---

## 2. Use Volumes to Persist Data

**Volumes** are Docker-managed directories on your host machine used to store data outside the container’s filesystem.  
This ensures your data isn’t lost when containers stop or are deleted.

**Steps:**

1. **Run a Container with a Volume:**

   ```bash
   docker run -v mydata:/data ubuntu echo "Hello" > /data/hello.txt
   ```

   - `-v mydata:/data`: Creates a volume named `mydata` and mounts it to `/data` inside the container.

2. **Check Volume Data:**

   ```bash
   docker run -v mydata:/data ubuntu cat /data/hello.txt
   ```

**Best Practice:**  
Use volumes for databases, logs, or any persistent data.

---

## 3. Publish Your Container's Port

**Port publishing** lets you access services running inside your container from your host or network.

**Steps:**

1. **Run a Web Server Container and Publish Port:**

   ```bash
   docker run -p 8080:80 nginx
   ```

   - `-p 8080:80`: Maps port 80 inside the container to port 8080 on your host.

2. **Access the Service:**  
   Open `http://localhost:8080` in your browser.

**Best Practice:**  
Choose non-conflicting host ports and document them.

---

## 4. Write and Run Your First Docker Compose

**Docker Compose** is a tool for defining and running multi-container Docker applications using a `docker-compose.yml` file.

**Steps:**

1. **Create a `docker-compose.yml`:**

   ```yaml
   version: '3'
   services:
     web:
       image: nginx
       ports:
         - "8080:80"
   ```

2. **Start the Application:**

   ```bash
   docker compose up
   ```

   - This starts all services defined in the file.

3. **Stop the Application:**

   ```bash
   docker compose down
   ```

**Best Practice:**  
Use Compose for apps with multiple services (e.g., web + database).

---

## 5. Containerize the Application Developed in the "Go" Phase

**Containerizing** means packaging your application and its dependencies into a Docker image.

**Steps:**

1. **Write a Dockerfile for Your Go App:**

   ```dockerfile
   FROM golang:1.21-alpine
   WORKDIR /app
   COPY . .
   RUN go build -o myapp
   CMD ["./myapp"]
   ```

2. **Build the Image:**

   ```bash
   docker build -t go-app .
   ```

3. **Run the Container:**

   ```bash
   docker run go-app
   ```

**Best Practice:**  
Use multi-stage builds to keep your image small and secure.

---

## Summary

- **Docker** helps you build, run, and manage containers.
- **Volumes** persist data.
- **Port publishing** exposes container services.
- **Docker Compose** manages multi-container apps.
- **Containerization** packages your app for portability and consistency.

If you have questions about any step, check the official Docker documentation or ask your senior
