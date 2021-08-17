# gotodo - sample Go Project

This project serves the purpose of demonstrating the structure and ideas of business logic implementation, responsibility segregation (receiving data -> deserialization -> running business logic -> persisting data -> sending events (if any) -> serialization).   

## Structure

- `/pkg` - folder contains go modules related to current project and some other general purpose modules
  - `/http-helper` - contains general purpose http related middlewares, router, server, error system
  - `/cqrses` - contains general purpose code related to handling **commands**, sending **events**
  - `/logger` - simple wrapper around uber's zap logger
  - `/auth` - JWT parser, enriches http request context with jwt claims
  - `/utils` - some general utility functions (parsing list params, ...)
  - `/todo` - contains all business logic including code that works with database, provides http handlers, mocks, tests

## Responsibility Segregation

### Data receiving (Transport)

Create corresponding handler factory to create handlers, for example, HttpHandlerFactory creates http handlers, there might be amqp handlers which receives requests from RabbiMQ or some other message broker supporting AMQP.  

### Data parsing (Deserialization)

Deserialization is done inside handler, then data passed to commands. 

### Business logic (Command Handling)

Commands are the abstraction to bridge transport and deserialization layer from business logic layer, they also serve purpose to generate events and publish to eventbus. Commands are executed by CommandHandler which in turn passes to commands required dependecy  to command, usually in the form of Service interface (which contains all business logic).

### Persisting data (Repository)

Repository is an interface which describes methods required to work with persistent storage, there might be different implementations of that interface. Usually passed as dependency into a Service (business logic).  

### Event publishing (Publisher)

Publisher publishes event to event bus, transparently integrated as command handler middleware. 

### Responding (Serialization)

Serialization of the response, can be integrated as handler middleware.
