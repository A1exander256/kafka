### Запуск

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