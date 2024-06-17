# Goapi
Base structure for creating an API with Golang using Gorm, Echo and Docker.

---

## core
This folder is where you place the internal code of the application. It contains business logic such as controllers, use cases, repositories, and entities. The idea behind this folder is that the code here is not accessible outside the package (module), which helps to maintain encapsulation and modularity.

## api
This folder can contain everything related to the API interface, such as route definitions, HTTP request handlers, and API documentation. It provides a clear separation between the application logic and the interface for communication with the outside world.

## pkg
This folder contains reusable packages or libraries that can be used in different parts of the application or even in other projects. For example, it might contain generic utilities, configuration functions, error handlers, among others.

---

#### middleware
Contains middlewares that intercept and process requests before reaching the controllers.

#### route
Contains the application's route definitions, mapping URLs to request handlers.

#### app
This folder is where you place the application-specific logic, such as controllers, use cases, repositories, and entities. It contains most of the application's business logic.

#### controller
Here you place the HTTP controllers that receive client requests, call the appropriate use cases, and return the HTTP responses.

#### usecase
This folder contains the application's use cases, which contain the application's specific business logic. The use cases orchestrate business operations and interact with repositories to access data.

#### repository
Repositories are responsible for interacting with the data storage layer (database, cache, etc.). They encapsulate the data access logic and provide methods to retrieve, create, update, and delete entities.

#### entity
This folder contains the application's entity definitions, which represent the main domain objects the application works with. Entities are often mapped directly to database tables.

#### message
Contains the messages to be used in request responses.

#### data (request/response)
Contains the data structures to be used as responses or received in requests.

#### database
Related to database connections and configurations.

#### server
Related to the HTTP/HTTPS server and its configurations.

#### config
Settings to be used in the application.
