# Backend

## 4 in a row server

### Docker setup

Run:

```bash
cd backend
docker build -t orsa-scholis/orsum-inflandi .
```

to build the backend docker image.

Use something like `docker run -it --rm -p 4560:4560 -h orsum-inflandi-server --name orsum-inflandi orsa-scholis/orsum-inflandi`
to start the server. 
