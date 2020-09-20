# salesloft-test

## Description

This project show my proposal for the technical test

There are two options to run the project

In both cases you need to export the api in the environment variable: `SALESLOFT_API_KEY`

## From line command

Tu run the project from command line, you will need to command prompt, 
in the first one run the following command:

```
make run-backend
```

In the second command prompt, un the following command:
```
make run-frontend
```

After that open in the explorer the address: `http://localhost:8080`


## Docker compose file

I attached a docker compose file, tu build and use it you will need Docker and a local egistry in the port 5000

With the following you can run the local registry:
```
docker run -d -p 5000:5000 --restart=always --name registry registry:2
```

After that you need to run the following commands, to un the docker compose:

```
make create_docker
make publish_docker
docker-compose -f docker-compose.yaml up
```

In the same way as the option above, open the address `http://localhost:8080` in the explorer.

Thank

Arnulfo José Suárez Gaekel