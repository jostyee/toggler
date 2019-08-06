# toggler

`toggler` is a Self hosted Feature Toggle Service.

The service designed to be hosted on public web.
The service expects that public web request will be received from all kind of sources.
Such case is the combined usage from SPA, lambda service and traditional backend services.

It is goal to provide a stable, reliable and free rollout management tooling for teams.
By using feature flags you can decouple the feature release from the deployment or config change process,
and also make it simple to keep feature states in sync for all your users.

The project aims only to be just barely enough for the minimal requirement
that needed to do centralised feature release management.

Other than percentage based feature enrollment for piloting,
every custom decision logic is expected to be implemented by your company trough an HTTP API.

## Is this Service for your team/company ?

Answer the following questions, in order to determine,
if this project is needed for your team or not.

Can my team…

* apply [Dark Launching](/docs/DarkLaunch.md) practices ?
* deploy frequently the codebase independently from feature release ?
* confidently deploy to production after the automated tests are passed ?
* perform deployment during normal business hours with negligible downtime?
* complete its work without needing fine-grained communication and coordination with people outside of the team?
* deploy and release its product or service on demand, independently of other services the product or service depend upon?

If your answer yes to most of them,
then you can stop here, because adding this service to your stack would not solve too much.
else, please continue...

## Scalability

The Service follows the [12 factor app](https://12factor.net/) principles,
and scale out via the process model.
The application don't use external resource dependent implementation,
so as long the external resource you use can be scale out, you will be fine.

If you need to add a new storage implementation,
because you need to use that,
feel free to create a Issue or a PR.
If you decide to implement your own integration with a storage,
the expected behavior requirements/tests/coverage can be located under the `usecases/specs.StorageSpec` object.
For examples, you can check the already existing storage implementations as well.

## [Features](/docs/features.md)
Please visit the Features documentation for this section.

## [Design](/docs/design/README.md)
Please visit the Design documentation to read the principles and conventions applied in this project.

## Quick Start / Setup

### Configuration
The application can be configured trough either CLI option or with environment variables.
It follows the convention that works easily with SaaS platforms or containerization based solutions.

#### Storage
The storage external resource will be used to persist data,
and then using as source of facts.

The toggler doesn't depend on a certain storage system.
It use behavior based specification, and has multiple implementation that fulfil this contract.
This could potentially remove the burden on your team to introduce a new db just for the sake of the project.

You can choose from the following

* [Redis](https://github.com/antirez/redis)
* [Postgres](https://github.com/postgres/postgres)
* InMemory (for testing purposes only)

The Storage connection can be configured trough the `DATABASE_URL` environment variable
or by providing the `-database-url` cli option to the executable.

To use one of the implementation, all you have to do is
to provide the connection string in the CLI option or in the environment variable.

example connection strings:
> redis://user:passwd@ec2-111.eu-west-1.compute.amazonaws.com:17379

> postgres://user:passwd@ec2-111.eu-west-1.compute.amazonaws.com:5432/dbname

```bash
export DATABASE_URL="postgres://user:passwd@ec2-111.eu-west-1.compute.amazonaws.com:5432/dbname"
```

#### Cache
The cache external resource is an optional addition.
By default, the service don't try to be smart, and use no cache at all.

You choose to have a trade off for your storage system to use a traditional database
that can provide your fact data with cost effectiveness, stability and maintainability perspectives.
But then you don't want to sacrifice the service response times, so you can use a cache system to solve this.
The Caching system do automatic cache invalidation with TTL and on Update/Delete storage operations.

Currently only redis is available, but additional solutions in progress.

To setup the application to use the cache, either provide the `-cache-url` cli option
or the `CACHE_URL` environment variable.

To setup the cache TTL, you can use the `-cache-ttl` cli option or the `CACHE_TTL` environment variable.
A cache ttl duration in string format must be a unsigned sequence of
decimal numbers, each with optional fraction and a unit suffix,
such as "300ms", "1.5h" or "2h45m".
Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".

### Deployment
* [heroku](/docs/deploy/heroku.md)
* [on-premises](/docs/deploy/on-prem.md)
* [Docker](/docs/deploy/docker.md)

### Usage

#### Security token creation
To gain access to write and update related actions in the system,
you must create a security token that will be used even on the webGUI.

To create a token, execute the following command on the server:
```bash
./toggler -cmd create-token "token-owner-uid"
```

the uniq id of the owner could be a email address for example.
The token will be printed on the STDOUT.
The token cannot be regained if it is not saved after token creation.

#### API Documentation
* [HTTP API documentation](/docs/httpapi/README.md)
* you can find the swagger documentation at the /swagger.json endpoint.
* the webgui also provides swagger-ui out of the box on the /swagger-ui path

## For Contributors
* [Backlog](https://github.com/adamluzsi/toggler/projects)

Feel free to open an issue if you see anything

## Thank you for reading about this project! :)