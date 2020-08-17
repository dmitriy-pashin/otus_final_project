docker run -d --name pgsql \
-e POSTGRES_PASSWORD=postgres \
-e PGDATA=/var/lib/postgresql/data/pgdata \
-v /Users/dmitriy.pashin/Workspace/pgdata:/var/lib/postgresql/data \
-p 5432:5432 \
postgres



docker exec -ti pgsql psql -Upostgres -dpostgres

docker exec -ti pgsql psql -Uotus -dantiddos

create user otus with encrypted password 'otus';
create database antibf;

grant all privileges on database antibf to otus;



goose -dir migrations postgres "user=otus password=otus dbname=antibf sslmode=disable" up