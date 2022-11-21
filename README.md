# Запуск local

1. Запустить zookeeper и kafka
```
start C:\kafka\bin\windows\zookeeper-server-start.bat C:\kafka\config\zookeeper.properties

start C:\kafka\bin\windows\kafka-server-start.bat C:\kafka\config\server.properties
```

2. Для создания топика `test` использовать следующую команду 
```
.\bin\windows\kafka-topics.bat --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic test
```

3. Для просмотра топиков можно использовать команду 
```
.\bin\windows\kafka-topics.bat --list --bootstrap-server localhost:9092
```

4. Для добавления записей можно использовать команду (`producer`)
```
.\bin\windows\kafka-console-producer.bat --topic test --bootstrap-server localhost:9092
```

5. Для считывания можно использовать команду (`consumer`)
```
.\bin\windows\kafka-console-consumer.bat --topic test --bootstrap-server localhost:9092 --from-beginning
```

# Docker

1. Выполянем команду 
```
docker-compose up -d
```

2. Для создания топика `test` использовать следующую команду 
```
docker exec kafka kafka-topics --bootstrap-server kafka:9092 --create --topic test
```

3. Для просмотра топиков можно использовать команду 
```
docker exec kafka kafka-topics --bootstrap-server kafka:9092 --list
```

4. Для добавления записей можно использовать команду (`producer`)
```
docker exec --interactive --tty  kafka kafka-console-producer --bootstrap-server kafka:9092 --topic test
```

5. Для считывания можно использовать команду (`consumer`)
```
docker exec --interactive --tty  kafka kafka-console-consumer --bootstrap-server kafka:9092 --topic test --from-beginning
```