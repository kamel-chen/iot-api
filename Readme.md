
# docker build
docker build . -t iot-api:v0.1.0

# docker tag for registry
docker tag iot-api:v0.1.0 192.168.0.156:5000/iot-api

# docker push to registry
docker push 192.168.0.156:5000/iot-api

# run docker
docker run -d --network=host --name iot-api -p 8080:8080 iot-api:v0.1.1

# influx db backup & restore
```
influx backup Backup -b iot-test -t token
influx restore Backup --full -t token
```

# NSQ start
```
nsqlookupd
nsqd --lookupd-tcp-address=127.0.0.1:4160
nsqadmin --lookupd-http-address=127.0.0.1:4161
```