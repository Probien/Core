
<div align="center" style="display:flex;flex-direction:column;">
    <img width="300" src="https://imgdb.net/storage/uploads/495cc30ad5b741033ede8604cb0ef566cb48b5685a252f34de460850dabb82f6.png" alt="Probien logo"/>
  <h2>Microservice Core Written in Go - Services For a Pawn Shop Web Application</h2>
  <p>
    <a target="_blank" href="https://crowdin.com/project/excalidraw">
      <img src="https://img.shields.io/badge/License-GPL%20v3-yellow.svg">
    </a>
        <a target="_blank" href="https://crowdin.com/project/excalidraw">
      <img src="https://img.shields.io/github/last-commit/ThePandaDevs/Probien-Backend">
    </a>
      <h4>Phase 1/3, Im still working on this when I have time</h4>
  </p>
</div>

## Bussiness Logic

- Administrator
  - Can see all occurred logs into application (sessions, movements and payments).
  - Can see detailed reports by branch office (money recorded, total products, employee statistics, etc).
  - Can manage user access, create new branches, create new employees and create new category products.
  
- Employee
  - Can manage the category of products
  - Can do pawn orders to customers as well as do endorsements.
  - Can manage the status of pawn orders (in course, overdue, paid or lost)

- Customer
  - Can do multiple pawns, their stuff are evaluated and classified by a category product when a quote is realized.
  - Can do endorsements depending of payment modality (weekly or monthly).
  - The Customer has an extension date of payment depending on the modality
    - Weekly: 1 extra day to make the payment
    - Monthly: 3 extra days to make the payment 

- Extras
  - When a customer doesn't make any endorsement after the deadline, all the client's things become property of probien.

## Design
```

config/
├─ database.go | Connection to postgresql using ORM
├─ redis.go | Connection to redis cloud
|
├─ migrations/
│  ├─ stored procedures | Raw sql/
│  ├─ models | Entity struct that represent mapping to data model/
core/
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
```

## Database Model (subject to change)

<img src="https://user-images.githubusercontent.com/67834146/177461447-59efbaa8-f04e-4003-96d8-1719af65025b.png" alt="database" border="0">

## Contributing

Contributions are always welcome!

This project is small, if you want to contribute improving code or add a new feature send a PR, make sure to add a description.


