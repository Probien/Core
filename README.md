
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
      <h4>Alpha version, Im still working on this</h4>
  </p>
</div>

## Lessons Learned

What did you learn while building this project? What challenges did you face and how did you overcome them?

First, I had no idea how to structure the project, I come from java and have always used the MVC pattern to build APIs, then I decided to learn another type of desing so I create this project with domain based approach [(DDD)](https://airbrake.io/blog/software-design/domain-driven-design) (Recomendations are welcome). 

## Design

- Application
  - Write business logic
    - employee.go (GetEmployeeById, GetAllEmployees, &...)
- Domain
  - Define interface
    - repository interface for infrastructure
  - Define struct
    - Entity struct that represent mapping to data model
      - employee.go
- Infrastructure
  - Implements repository interface
    - employee_repository.go
- Interfaces
  - HTTP handler

## Database Model (subject to change)

<img width="1200" height="400" src="https://imgdb.net/storage/uploads/25847851dc958eb3e3dd6742dd036e6c06c083a116c524c4cdcfc146cee00034.png" alt="database"/>


## Roadmap

- JWT Authentication :heavy_check_mark:
- Authentication Based on Roles :clock330:
- Transaction | CRUD Operations :clock330:
- E-mail Notifications :heavy_minus_sign:
- Recover Password :heavy_minus_sign:
- Reports By Pawn Shop Branch :heavy_minus_sign:
- Swagger Documentation :heavy_minus_sign:

## Contributing

Contributions are always welcome!

This project is small, if you want to contribute improving code or add a new feature send a PR, make sure to add a description.


