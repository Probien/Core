
<div align="center" style="display:flex;flex-direction:column;">
    <img width="300" src="https://imgdb.net/storage/uploads/495cc30ad5b741033ede8604cb0ef566cb48b5685a252f34de460850dabb82f6.png" alt="Probien logo"/>
  <h2>Backend Service Written in Go - Services For a Pawn Shop Web Application</h2>
  <p>
    <a target="_blank" href="https://crowdin.com/project/excalidraw">
      <img src="https://img.shields.io/badge/License-GPL%20v3-yellow.svg">
    </a>
        <a target="_blank" href="https://crowdin.com/project/excalidraw">
      <img src="https://img.shields.io/github/last-commit/ThePandaDevs/Probien-Backend">
    </a>
      <h4>Phase 1/2, Im still working on this</h4>
  </p>
</div>

## Lessons Learned

What did you learn while building this project? What challenges did you face and how did you overcome them?

First, I had no idea how to structure the project, I come from java and have always used the MVC pattern to build APIs, then I decided to learn another type of desing so I create this project with domain based approach [(DDD)](https://airbrake.io/blog/software-design/domain-driven-design) (Recomendations are welcome).


## Bussiness Logic

- Administrator
  - Can see all occurred logs into application (sessions, movements and payments).
  - Can see detailed reports by branch office (money recorded, total products, employee statistics, etc).
  - Can manage user access, create new branches, create new employees and create new category products.
  
- Employee
  - Can manage the category of products
  - Can do pawn orders to customers as well as do endorsements.
  - Can manage the status of pawn orders (in course, late, paid or overdue)

- Customer
  - Can do multiple pawns, their stuff are evaluated and classified by a category product when a quote is realized.
  - Can do endorsements depending of payment modality (weekly or monthly).

Lost products
  - When a customer doesn't make any endorsement after the deadline, all the client's things become the property of probien.

## Roadmap

- Phase 1
  - JWT Authentication :heavy_check_mark:
  - Authorization Based on Roles :heavy_check_mark:
  - Database logs :heavy_check_mark:
  - Transaction | CRUD Operations :clock330:
  - HTTP Tests :heavy_check_mark:

- Phase 2
  - E-mail Notifications :heavy_minus_sign:
  - Recover Password :heavy_minus_sign:
  - Reports By Pawn Shop Branch :heavy_minus_sign:
  - Swagger Documentation :heavy_minus_sign:

- Phase 3
  - Refactor code and bussiness logic
  - (Possibly) add employee payments & extra features
  - Start frontend...

## Design

config/
├─ database.go | Connection to postgresql using ORM
|
├─ migrations/
│  ├─ stored procedures | Raw sql/
│  ├─ models | Entity struct that represent mapping to data model/
core/
|
├─ application | Write business logic/
├─ domain | Entity struct that represent I/O JSON format/
│  ├─ repository | Repository interface for infrastructure/
|
├─ infrastructure/
│  ├─ auth | Middleware and security filters/
│  ├─ persistance | Implements repository interface with database /
|
├─ interface | Expose http endpoints/
router | Routing for endpoints/
server.go

## Database Model (subject to change)

<img src="https://i.ibb.co/QPsbgy0/database.png" alt="database" border="0">

## Contributing

Contributions are always welcome!

This project is small, if you want to contribute improving code or add a new feature send a PR, make sure to add a description.


