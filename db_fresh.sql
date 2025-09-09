echo "DROP DATABASE coffee_db" | docker exec -i postgres-dock-1 psql -U postgres -d coffee_db
echo "CREATE DATABASE coffee_db" | docker exec -i postgres-dock-1 psql -U postgres -d coffee_db