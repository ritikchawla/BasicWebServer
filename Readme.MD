# Basic Web Server Project

This is a simple web server project built using Go that allows you to manage a list of films. You can view the list of films, and add new films using a basic web interface.

## Prerequisites

Before running the web server, ensure you have Go installed on your machine. You also need the following dependencies for the project:

- [Bootstrap](https://getbootstrap.com): Used for styling the web interface.
- [htmx](https://htmx.org): Used to update the page content dynamically without full page reloads.

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/ritikchawla/BasicWebServer.git
   cd BasicWebServer
   ```
  2.  Build and run the server:
      ```sh
      go run main.go
      ```
    
3.  Open your web browser and navigate to [http://localhost:8000](http://localhost:8000/) to access the web interface.
    

## Usage

### Film List

-   The home page displays a list of films with their titles and directors.
-   You can view the list of existing films.

### Add Film

-   Use the "Add Film" section to add new films to the list.
-   Enter the film's title and director, then click the "Submit" button.
-   The new film will be added to the list dynamically without a full page reload.
    

## Thanks
I am passionate about exploring new technologies and creating exciting projects. If you have interesting project ideas and would like to collaborate, feel free to reach out to me via [email](mailto:chawlaritik@gmail.com).