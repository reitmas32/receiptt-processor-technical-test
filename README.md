# Receipt Processor 'Technical test by Fetch'

## Requirements

 - [Docker](https://www.docker.com/get-started/)

## Instructions by Run the API

### **Create container**

```bash
docker build -t receipt_procesor .
```

### **Run the container**

```bash
docker run -d receipt_procesor
```

### Try the API

THe API is avalible in `localhost:5000`

# Notas 
Add `sudo` th the commands of docker if the $USER not exist in `docker` group