1. install redis on your local docker docker pull redis

2. create a local redis container named my-redis and run it. docker run --name my-redis -p 6380:6380 -d redis

3. log into your local myredis container using "docker exec -it <redis-name> bash", as below, my redis container name is
   my-redis % docker exec -it my-redis bash                                      
   root@545633d3db0a:/data#

4. view db size with redis-cli at beginning after you log into your redis container. root@545633d3db0a:/data# redis-cli
   dbsize
   (integer) 0

5. set value in your redis container root@5fdb3dbb9dd7:/data# redis-cli set ramsey best OK

6. get value in your redis container root@5fdb3dbb9dd7:/data# redis-cli get ramsey     
   "best"


8. view db size with redis-cli at the end root@5fdb3dbb9dd7:/data# redis-cli dbsize
   (integer) 1
