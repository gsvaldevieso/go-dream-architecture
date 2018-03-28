# Go Dream Architecture
An attempt to build a maintainable architecture in go projects, based on Hexagonal Architecture, The Clean Architecture, Onion Architecture, DDD, etc.

![Gopher](https://media.giphy.com/media/oPu2IgQHwb3Qk/giphy.gif)

This project is based on the current materials:

http://jeffreypalermo.com/blog/the-onion-architecture-part-1/

https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

http://alistair.cockburn.us/Hexagonal+architecture

The whole code is based on my view about "clean architecture" and how to build maintainable projects. This repo is motivated by the absense of good base projects for the golang community and I hope it can be helpful for you as for me.

## Folder strucuture:

### ./domain/entity:
> Entities encapsulate Enterprise wide business rules. An entity can be an object with methods, or it can be a set of data structures and functions. It doesn’t matter so long as the entities could be used by many different applications in the enterprise. They encapsulate the most general and **high-level rules**. They are **the least likely to change when something external changes**. For example, you would not expect these objects to be affected by a change to page navigation, or security. No operational change to any particular application should affect the entity layer.

### ./domain/repository:
>The repository is implemented in the domain layer, because it works with domain objects. But in the domain layer we should have no idea about any database nor any storage, so the repository is just an interface.

### ./domain/service:
>Domain services hold domain logic that doesn’t naturally fit entities and value objects.

### ./domain/vo:
>A Value Object is an immutable type that is distinguishable only by the state of its properties. That is, unlike an Entity, which has a unique identifier and remains distinct even if its properties are otherwise identical, two Value Objects with the exact same properties can be considered equal. 

### ./infrastructure:
>The outermost layer is generally composed of frameworks and tools such as the Database, the Web Framework, etc. Generally you don’t write much code in this layer other than glue code that communicates to the next circle inwards.

### ./infrastructure/database ./infrastructure/queue:
>This layer is where all the details go. The Web is a detail. **The database is a detail.** We keep these things on the outside where they can do little harm.

### ./repository:
>The logic of repository is implemented in the repository layer, because it works with infrastructure objects. The structures in this layer must respect the domain interface and do any complex logic like SQL queries.

### ./usecase:
>The software in this layer contains application specific business rules. It encapsulates and implements all of the use cases of the system. These use cases orchestrate the flow of data to and from the entities, and direct those entities to use their enterprise wide business rules to achieve the goals of the use case.

# TODO (yes, you can help me!):
* Implement infrastructure objects and create adapters for drivers like MySql, using libs for go;
* Implement some entity/aggregate with domain rules;
* Define some repository interface in the domain;
* Implement some VO, for example: Money;
* Implement entity and repository interation in the use case;
* And everything you think is cool...

Thank you!