# Hello World HTTP
Hello World HTTP is a container with a simple HTTP web server.

The default serving address is `:8080`, set `HELLO_ADDR` to change the it.
The default response text is `Hello, World!`, set `HELLO_TEXT` to change the it.

## Run
Run the Hello World HTTP container using podman:
```
$ podman run -p 8080:8080 quay.io/brusdev/hello-world-http
```
The exepected output is:
```
Hello, World HTTP!
The default serving address is :8080, set HELLO_ADDR to change the it.
The default response text is Hello, World!, set HELLO_TEXT to change the it.
Serving :8080 with the text Hello, World!
```
Test the Hello World HTTP container:
```
$ curl localhost:8080/test
```
The exepected output is:
```
Hello, World!
```

### Run with a custom text
Run the Hello World HTTP container with the custom text `Hello, brusdev!` using podman:
```
$ podman run -e HELLO_TEXT='Hello, brusdev!' -p 8080:8080 quay.io/brusdev/hello-world-http
```
The exepected output is:
```
Hello, World HTTP!
The default serving address is :8080, set HELLO_ADDR to change the it.
Serving :8080 with the text Hello, brusdev!
```
Test the Hello World HTTP container with the custom text `Hello, brusdev!`:
```
$ curl localhost:8080/test
```
The exepected output is:
```
Hello, brusdev!
```

## Build
Build the container image using podman:
```
$ podman build . --tag quay.io/brusdev/hello-world-http
```

## Push
Push the container image using podman:
```
$ podman push quay.io/brusdev/hello-world-http
```

## Deploy
Create a deployment on Kubernetes using kubectl:
```
$ kubectl apply -f hello-world-http-deployment.yaml
```
Create a service on Kubernetes using kubectl:
```
$ kubectl apply -f hello-world-http-service.yaml
```
Test the deployment and the service using busybox:
```
$ kubectl run -it --rm busybox --image=busybox --restart=Never
If you don't see a command prompt, try pressing enter.
/ # 
/ # 
/ # wget -q -O - hello-world-http-service:8080
Hello, World!
/ # exit
pod "busybox" deleted
```
