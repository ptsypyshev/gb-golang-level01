1. Get from Flags:
main.exe --db_url "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable" --jaeger_url "http://jaeger:16681" --sentry_url "http://sentry:9001" --kafka_broker "kafka:9091" --port 8081 --some_app_id 1 --some_app_key "Idgz1PE3zO9iNc0E3oeH3CHDPX9MzZe31"

2. Get from Envs:
main.exe

3. Get from JSON:
main.exe --json "config.json"

4. Get from YAML:
main.exe --yaml "config.yaml"


Предложение по теме для консультации:
Быть может провести что-то типа мастер-класса по разработке?
Например, взять прямо задачу из реальной практики и решить её?
Мне было бы интересно