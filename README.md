OUI Interview Challenge
Tariq Chaudhry

Problem:
```
Deliverable
Implement the following GraphQL schema:

type Set {
  members: [Int!]!
  intersectingSets: [Set!]!
}

type Query {
  sets: [Set!]
}

input SetInput {
  members: [Int!]!
}

type Mutation {
  createSet(input: SetInput!): Set!
}

Some issues to keep in mind:

● Data should be persisted durably (that is, not in memory).
● What happens when duplicate sets are added, noting that sets with the same members
(regardless of order) are considered equivalent?

If you have extra time, consider adding to the schema and implementing:
● Unique IDs for each set.
● Pagination of set lists.
```

Used the gqlgen package for the GraphQl server implemented in golang
https://github.com/99designs/gqlgen

To run using docker-compose, clone the repository and use `docker-compose up`
and then if you have psql on your host machine, you can setup the database using...
```
psql -h 127.0.0.1 -U postgres -f sql/up.sql
```
otherwise connect to the postgres container and run it from there
