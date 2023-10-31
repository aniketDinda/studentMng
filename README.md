# studentMng

This service allows you to manage student information via HTTP APIs. Below, you will find instructions on how to use the provided APIs.

## Getting Started

To start the User Management Service, you can follow these steps:

1. Clone the repository:

   ```shell
   git clone https://github.com/aniketDinda/studentMng.git
   cd studentMng

2. Install required dependencies (if not already installed):
   
   ```shell
   go mod tidy
   

3. Run the server:

   ```shell
   go run main.go
      
        or
   
   ```
    docker build -t stumng .
    docker run -p 8085:8085 stumng


4. Endpoints - Create a New User

   ```shell
   curl -X POST -H "Content-Type: application/json" -d '{
    "name": "Aniket Dinda",
  	"address":"abcd1234",
  	"city": "Kolkata",
  	"country": "India",
  	"pin_code": "700012",
  	"marks": 1205
   }' http://localhost:8085/student/add

5. Endpoints - Fetch Students

   ```shell
   curl http://localhost:8081/students/view

6. Endpoints - Update Student Mark

   ```shell
   curl -X POST -H "Content-Type: application/json" -d '{
    "name": "Aniket Dinda",
  	"marks": 1000
   }' http://localhost:8085/student/update

7. Endpoints - Delete student

   ```shell
   curl -X POST -H "Content-Type: application/json" -d '{
    "name": "Aniket Dinda",
   }' http://localhost:8085/student/delete

8. Endpoints - GetRank

   ```shell
   curl -X POST -H "Content-Type: application/json" -d '{
    "name": "Aniket Dinda",
   }' http://localhost:8085/student/getRank

