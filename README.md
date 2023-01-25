# ShortURLApi

Ссылка для скачивания docker-образа: https://hub.docker.com/repository/docker/rodmul/short_link/general

Для запуска с PostgreSQL хранилищем: docker run -d -p 4000:4000 --name short_link_container rodmul/short_link:v1 "--use_db_storage"
Для запуска с Im-memory хранилащем: docker run -d -p 4000:4000 --name short_link_container rodmul/short_link:v1
