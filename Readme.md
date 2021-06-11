
# docker build
docker build . -t iot-api:v0.1.0

# docker tag for registry
docker tag iot-api:v0.1.0 192.168.0.156:5000/iot-api

# docker push to registry
docker push 192.168.0.156:5000/iot-api

# run docker
docker run -d --network=host --name iot-api -p 8080:8080 iot-api:v0.1.1
