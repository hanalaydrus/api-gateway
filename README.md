**API Gateway**

API Gateway is the microservices endpoint, a bridge between front-end and back-end, you need to have semantic-service, volume-service, and density-service built to run this.
 - semantic-service: https://github.com/hanalaydrus/semantic
 - volume-service: https://github.com/hanalaydrus/volume-of-vehicles-cpp
 - density-service: https://github.com/hanalaydrus/density-of-vehicles

This are steps to run this code in your local:

**1. Install Docker**

 - Windows : https://docs.docker.com/docker-for-windows/install/ 

 - Mac : https://docs.docker.com/docker-for-mac/install/

 - Ubuntu : https://docs.docker.com/install/linux/docker-ce/ubuntu/

**2. Make sure docker-compose installed**, If not please check the link below
https://docs.docker.com/compose/install/

**3. Clone this repo** `git clone https://github.com/hanalaydrus/api-gateway.git`

**4. Run** `docker-compose up --build` **or** `sudo docker-compose up --build` **in the repo**

**5. The API-Gateway should already run**, you could check it through `docker ps`
