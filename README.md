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

2. Go in the project , fetch all dependencies and run it.

<!-- Verify if "go install ." it's OK! -->
```bash
cd trucks
go install . 
```

3. Store login data into an `.env` file at the root.
```bash
# .env
TRUCKS_USERNAME=YOUR_USERNAME 
TRUCKS_PASSWORD=YOUR_PASSWORD
```

4. Open a terminal at the root of the project and run the project using `go run main.go`.

5. Go on your favorite browser at: `http://127.0.0.1:8080` to see the first page of the application.

## :warning: Prerequisites

Execute `trucks.sql` from `/database` folder into your **MySQL server**,  
To prepare initial *trucks* for database.

```bash
# MySQL Server

source path
```mysql
source ./path_to_project/trucks/database/trucks.sql
```
<!--## :thinking: How does it run ?


> [!NOTE]
> A better version with more features is coming... 🏗️

<!--## :test_tube: Features
- Implement AddTruck func logic

**********
With Gin
- Build a CRUD interface for /admin/trucks to manage trucks
- 		// database/database.go Fix: Prevent duplication of the previous query
		// about user_admin table
- Prevent database to begin with another integer than 1 for the first element in database.

- `session, err := models.Store.Get(r, "session-name")
			utils.ProcessError(err, w)` that line is repetitive in handlers/handlers.go and middlewares/middlewares.go
- Secure the app (Store into public folder only file client can visit)
- Avoid infinite session (set max time until user leave the application)
-->


## :key: License

Developed by Loick Cherimont  

Under MIT License  

Last edition : 2024-11-14


