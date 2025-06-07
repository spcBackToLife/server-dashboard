# Server Management Panel

This project is a web-based server management panel similar to bt.cn. It is built with a Flask backend and a JavaScript/HTML/CSS frontend.

## Table of Contents
- [Key Features](#key-features)
- [How to Start](#how-to-start)
  - [Backend Setup](#backend-setup)
  - [Frontend Setup](#frontend-setup)
- [Deployment](#deployment)
  - [Backend (Flask Application)](#backend-flask-application)
  - [Frontend Files](#frontend-files)
  - [Hosting Environment](#hosting-environment)
- [Screenshots](#screenshots)

## Key Features
Key features (currently simulated for most backend operations):
- User registration and login.
- Adding and listing servers.
- Getting server status (mocked CPU, memory, disk usage).
- File upload/download to/from a server (simulated).
- Listing files on a server (simulated).
- Listing databases and tables on a server (simulated).
- Executing SQL queries (simulated, distinguishes SELECT from other DML).
- Executing shell commands on a server (simulated for commands like `ls`, `whoami`, `pwd`).
- Managing Docker images and containers (list, pull image, get container logs - all simulated).
- Installing software packages on a server (simulated).

## How to Start

### Backend Setup
1. From the root of the project, navigate to the `backend` directory: `cd backend`
2. Create a virtual environment: `python -m venv venv`
3. Activate the virtual environment:
    - On macOS and Linux: `source venv/bin/activate`
    - On Windows: `venv\\Scripts\\activate`
4. Install dependencies: `pip install -r requirements.txt`
5. Run the Flask application: `python app.py`
6. The backend will be running on `http://localhost:5000` by default.

### Frontend Setup
1. Open the `frontend/index.html` file directly in your web browser.
2. The frontend is configured to interact with the backend API at `http://localhost:5000`.

## Deployment

This section provides general guidance for deploying the Flask backend. The frontend consists of static files that can be served by a web server or CDN. Deploying a web application involves several components working together.

### Backend (Flask Application)

For production environments, do not use the built-in Flask development server (`python app.py`). Instead, use a production-ready WSGI server like Gunicorn or uWSGI.

**1. WSGI Server:**
   A WSGI (Web Server Gateway Interface) server is necessary to run a Python web application in production. It handles requests from the web server and communicates with your Flask app.
   - **Gunicorn Example:**
     Navigate to the `backend` directory and run:
     ```bash
     gunicorn -w 4 -b 0.0.0.0:5000 app:app
     ```
     (This command assumes your Flask application instance is named `app` in the `app.py` file (`filename:Flask_instance_name`). The `-w 4` flag starts 4 worker processes. Adjust the host and port as needed.)
   - **uWSGI Example:**
     uWSGI is another popular choice, often used with Nginx. Configuration can be more complex and is typically done via a `.ini` file.

**2. Reverse Proxy (Nginx/Apache):**
   A reverse proxy sits in front of your application/WSGI server and handles incoming client requests. It's standard practice to put your WSGI server behind a reverse proxy like Nginx or Apache. The reverse proxy can:
   - Serve the Flask application by forwarding requests to the WSGI server.
   - Handle HTTPS (SSL/TLS termination).
   - Serve static files directly (for the backend, if any; frontend files are handled separately as described below).
   - Provide load balancing if you have multiple application instances.
   - Offer security benefits (e.g., rate limiting, IP blocking).

**3. Environment Variables:**
   While this application doesn't heavily rely on them yet, it's good practice to manage configuration (e.g., secret keys, database URIs in more complex apps) using environment variables. Flask can load configuration from environment variables at startup.

### Frontend Files
The `frontend/` directory of this project contains static HTML, CSS, and JavaScript files. These should be served by your reverse proxy (e.g., Nginx, configured to serve this directory) or uploaded to a Content Delivery Network (CDN) for better performance.

### Hosting Environment
The specific deployment steps will vary significantly based on your chosen hosting environment:
- **Docker:** Containerize the Flask application with its WSGI server. This is a common approach for portability and scalability.
- **Kubernetes:** Orchestrate Docker containers for larger deployments.
- **Platform as a Service (PaaS)** (e.g., Heroku, AWS Elastic Beanstalk, Google App Engine; many other cloud providers offer similar services): These platforms often simplify deployment by handling much of the infrastructure. You typically provide your code and some configuration.
- **Virtual Private Server (VPS):** You'll need to set up the WSGI server, reverse proxy, and manage the server environment yourself.

**Disclaimer:** The information above provides general deployment guidance. Always refer to the documentation of the specific tools and hosting platforms you choose for detailed instructions.

## Screenshots

This section is intended to showcase the look and feel of the server management panel.

<!-- Add screenshots here -->

Contributors are encouraged to add screenshots of the application in action. This will help users quickly understand what the panel looks like and its capabilities.
For example, you could include screenshots of:
* The login page
* The server list
* Server detail/status page
* File management interface
* Docker management interface
