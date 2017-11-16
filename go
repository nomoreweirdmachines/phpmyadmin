docker stop phpmyadmin
docker rm phpmyadmin
docker pull phpmyadmin/phpmyadmin

docker run --name phpmyadmin -d -e PMA_HOST=dbserver -p 8081:80 --network fortifynetwork phpmyadmin/phpmyadmin
