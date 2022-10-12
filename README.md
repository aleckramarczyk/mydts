MDT telemetry server and agent

This project is designed to run in a docker deployment. Requires a mysql database connection.

First, build the image:
  from parent directory:
  docker build . --rm --tag mydts:1.0

Example docker compose file:
```yaml
version: '3.3'

services:
  mysql:
    container_name: mysql
    image: mysql:8.0
    restart: unless-stopped
    environment:
      MYSQL_DATABASE: mdts
      MYSQL_ROOT_PASSWORD: changeme
      MYSQL_ROOT_HOST: '%'
    ports:
      - 3306:3306 #This could be ommitted if you don't want to expose the mysql service. However, we use a different service to display collected information that needs this access
    volumes:
      - mysql:/var/lib/mysql
    networks:
      - mydts
  mydts:
    container_name: mydts
    image: mydts:1.0
    restart: unless-stopped
    environment:
      DB_HOST: mysql 
      DB_PORT: 3306
      DB_TABLE: mdts
      DB_PASSWORD: changeme
      DB_USER: root
      API_PORT: 8080
    ports:
      - 8080:8080
    networks:
      - mydts

volumes:
  mysql:

networks:
  mydts:
```
The agent requires a config file named config.json in it's directory.

config.json:
```json
{
  "API_Host": "0.0.0.0:8080",
  "API_Endpoint": "/api/mdt"
}
Replace API_Host with the IP that your server will be running on 

Next build the agent:
```bash
go build -o agent.exe ./agent/cmd
zip agent.zip agent.exe config.json
rm agent.exe
```

Distrubute the zip file to as many clients as you want to track. Running the executable as a service is recommended.