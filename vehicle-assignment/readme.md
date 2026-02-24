assignment - 1

Step 1 : define custom type

a. create a Vehicle struct with common properties like Brand,Model,Year and Color

b. Define three custom types: Car,Boat and Motorcycle,each embedding the Vehicle struct. Add specific properties to each type.For example:

Car : NumDoors,EngineType
Boat : Length, PropulsionType
Motorcycle : NumWheels, HasSidecar

Step 2: create an interface
a. declare an interface VehicleInterface with methods that all vehicle types must implement such as Start(),Stop(),Steer()

Step 3: implement error handling

a. define custom error type VehicleError that implements the error interface and includes additional information about the error.

b. Handle potential errors in adding a vehicle by returning the VehicleError when applicable.

Step 4 : create management system

a. Implement a function to add new vehicles to the system and return the newly added vehicle.

b. Write a function to start a vehicle and handle any errors that may occur.

c. Implement a function to stop a vehicle 



//how I initialised github?

1. went to github to create repo
2. git init on terminal
3. git config user.email "24arushisharma@gmail.com"
4. git remote add origin "<repo-link>" 
a. remote -> connection/alias that points your local git repo to the repo on github
b. origin -> new remote 


Car
 ├── Vehicle
 │     ├── Brand
 │     ├── Model
 │     ├── Year
 │     └── Color
 ├── Doors
 └── Engine

 Car{
    Vehicle{Brand:brand, Model:model, Year:year, Color:color},
    4,
    "V6",
}