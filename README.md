# Trucks_Go


![Preview](https://placehold.co/600x400 "Trucks_Go")


## :information_source: About  

This repository is a GO project focused on truck management in the transport industry.


## :wrench: Tools
- [Go 1.23.2](https://go.dev/ "Go official website")


## :inbox_tray: Setup for contributions
1. Open a terminal and paste these lines

```bash
git clone git@github.com:loickcherimont/trucks.git
```

2. Go in the project , fetch all dependencies and run it

```bash
cd trucks

# Choose simple login data
export TRUCKS_USERNAME=YOUR_USERNAME 
export TRUCKS_PASSWORD=YOUR_PASSWORD
go run main.go
```

3. Open your web browser and run Go server using:
```bash
http://127.0.0.1:8080
```

<!--## :warning: Prerequisites
To run correctly this project, you'll a server use : [Vite 5.3.1](https://vitejs.dev/ "Vite official website")-->

<!--## :thinking: How does it run ?
Click on the button to send a request to server and downloader 2 files at the same time

into an inner /out folder.
-->

> [!NOTE]
> A better version with more features is coming... 🏗️

<!--## :test_tube: Features

TODO definitive (for now)
> Create a type User
> Build a CRUD interface for /admin/trucks to manage trucks
  - GET /admin/trucks -> GET all trucks
  - GET /admin/truck/{id} -> Get a specific truck by its ID
  - POST /admin/truck{id} -> Create a new truck
  - PUT /admin/truck{id} -> Modify info on a truck
  - DELETE /admin/truck/{id} -> Delete a specific truck 
  
//////////////////////// FOR DATABASE ///////////////////////////
// What do I build with MySQL
// DBName: db_transport
// 2 tables:
// - trucks: Store 3 trucks
// - user_admin: Store admin info (superUser)
// For user_admin : hash the password before add it to database (https://gowebexamples.com/password-hashing/)
//////////////////////// END - FOR DATABASE ///////////////////////////
-->


## :key: License

Developed by Loick Cherimont  

Under MIT License  

Last edition : 2024-10-04


