OUI Interview Challenge

graphql client
https://github.com/machinebox/graphql

```
select a.members, array_agg(b.id) from sets a 
inner join sets b on a.members && b.members
WHERE a.id <> b.id
GROUP BY a.id
