# shopping cart

## How to use/run the code
**The code expects a running postgresql database reachable on `localhost:5432`** A [docker-compose file](docker-compose.yaml) is provided to run
the database with the correct db and credentials setup. Use the following command to start the database or have a look at
[docker-compose.yaml](docker-compose.yaml) for more information (db, credentials etc.).
```shell
docker compose up -d
```

### testing
The code starts a PostgreSQL docker container to run the the [shoppingcart/repo/postgres](shoppingcart/repo/postgres) package tests.
In order to be able to run these tests you must have a running docker compatible daemon on your machine. The tests can
be run using plain `go test`:
```shell
go test -cover ./...
```

### building
[goreleaser](https://goreleaser.com/) is used to build and package a Docker image out of the code. You can create a
snapshot release by running:

```shell
goreleaser release --snapshot --clean
```

This will build binaries for darwin/linux x86/arm64 (see the [configuration](.goreleaser.yaml) of goreleaser for more
information).

If you do not want to use `goreleaser` you can simply use plain `go` to build/test the code.

## Reasoning
### Architecture
The solution follows the ports and adapters architecture pattern (also called hexagonal architecture) to loosely couple the application components. In this case there is a single
component called `shoppingcart` which has two ports:
- Repository - for data persistence
- Reservation - to model the external item reservation system

#### Repository port ([shoppingcart/repo/](shoppingcart/repo/))
The Repository port contains two different adapters:

- `mock` - As the name says this adapter can be used for testing. The code is generated using [gomock](https://pkg.go.dev/go.uber.org/mock).
- `postgres` - Is an adapter to store data in a PostgreSQL database.

#### Reservation port ([shoppingcart/reservation/](shoppingcart/reservation/))
The Reservation port has three different adapters:
- `dummy` - As the name implies provides a dummy adapter implementation which simply returns a random number as `reservation_id`.
- `http` - Is an example how an adapter could look like that talks to an external http service.
- `mock` - Is a generated mock adapter for testing. The code is generated using [gomock](https://pkg.go.dev/go.uber.org/mock).

### Persistence
I simply picked PostgreSQL because I'm familiar with it and because I had a boilerplate application at hand. In a real
world example one would probably not use a relational database to store some shopping cart information because a simple
key-value store (e.g. redis) would be enough for that. Thanks to the ports and adapters architecture it would be easy to
implement a different persistence provider.

### Libraries

#### http
No special http framework is needed for such a simple task. Plain stdlib `net/http` provides enough functionality for what
is needed.

#### sql
Same as with http, stdlib `database/sql` provides enough functionality for what is needed. Of course an SQL driver is needed.
In the past I worked with [github.com/jackc/pgx/v5](https://pkg.go.dev/github.com/jackc/pgx/v5), which works well together with the stdlib and thus I decided to use it
for this project as well.

#### mocks/testing
For such a simple project the stdlib packages would be enough. However I decided to use [gomock](https://pkg.go.dev/go.uber.org/mock) to be able
to automatically generate mocks rather than writing them on my own. For assertions I like to use [github.com/stretchr/testify](https://pkg.go.dev/github.com/stretchr/testify)
because it helps to reduce a lot of `if err != nil {}` constructs in tests and in my opinion makes test code easier to read.
Both of these libraries aren't strictly needed but they make my life a bit easier and thus I used them.

[github.com/ory/dockertest/v3](https://pkg.go.dev/github.com/ory/dockertest/v3) is used to start a PostgreSQL docker
container for the [shoppingcart/repo/postgres](shoppingcart/repo/postgres) package tests. I did this, because I do not like
to mock tests that use SQL because I think mocking these kind of tests doesn't really provide value except for test
coverage.

## Self assessment
The solution is on the more complex side for such a simple task. However, I tried to showcase how code can be decoupled
by using the ports and adapters architecture pattern. In my opinion this is a good baseline for a well-structured codebase
that is maintainable in the long term. As I already wrote for such a simple task it is a bit overkill and I wouldn't
recommend doing this for something at this scale that is meant to be placed in production.

### What is missing?
#### metrics
Proper application metrics should be added for monitoring purposes. Thanks to the ports and adapters architecture this can
easily be done by wrapping the port interfaces with an implementation that publishes metrics. Usually these "middleware"
implementations can be generated but there was not enough time left to do this.

#### proper (access) logging
There are a few log statements spread across the code. However, there is no proper access log nor any logs of e.g. database
calls. But once again this can be solved by middlewares.

#### reservations worker pool
Currently each pending reservation is processed one after each other. In a real world example this should probably be done
in parallel using a configurable worker pool.
I decided to use a simple polling approach to asynchronously process the reservations as this is probably the simplest
solution.

#### configuration
As of know all configuration is hardcoded (database connection strings, timeouts etc.). Configuration should be possible
form "outside", e.g. through environment variables.

#### cleanup of all `TODO` in the code
There are a few `TODO` in code with small things that should be cleaned up before this could be used in production.

#### CI/CD
There is currently no CI/CD pipeline. Usually I'd use GitHub actions or GitLab CI/CD for that. The use of goreleaser makes
it very easy to add a pipeline which automatically creates releases and publishes the artifacts.

#### linting
Usually I use a special [golangci-lint](https://golangci-lint.run/) configuration and CI/CD setup in a production project.
This is missing currently. However, the default configuration of `golangci-lint` does not report any lint errors.

#### infrastructure topics
There is a lot more that could be done:
- k8s deployment
- readiness/liveness probes
- SSL/TLS
- rate limiting and security headers
- database migrations
- etc.


