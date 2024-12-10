# jokenpo
Classic jokenpo game built with Go

## Prerequisites

- go 1.23.4


## Makefile

In the root has a Makefile with all available commands.You can run usin ```make <command>```.Like example below

```
make run
```

## Run aplication

you can run using the command ```make run``` than the server will run on PORT 8084(define in .env file)

To test the ```/match``` endpoint you can run the command below

```
curl --request POST \
  --url http://0.0.0.0:8084/match \
  --header 'Content-Type: application/json' \
  --data '{
  "player1": "lizard",
  "player2": "scizor"
}'
```

The available entries are:

- paper
- rock
- scizor
- spock
- lizard

If you need add more options you can update the ```rules.yaml```

```yaml
paper: 
  - rock
  - spock
```

in the example above, **paper** will win from **rock** and **spock** 


## Load Test

### run load test

before run the load test please update ```loadtest/script.js```. with time duration and VUs 

```javascript
export const options = {
  vus: 500 ,  // here change for desire number of virtual users(VU)
  duration: '60s', // here the load test duration
};


```

### Results

For **500** VU
```
     ✓ status was 200

     checks.........................: 100.00% 29516 out of 29516
     data_received..................: 4.2 MB  69 kB/s
     data_sent......................: 5.3 MB  86 kB/s
     http_req_blocked...............: avg=1.67ms   min=750ns    med=2.5µs   max=250.58ms p(90)=8.54µs  p(95)=14.33µs 
     http_req_connecting............: avg=1.65ms   min=0s       med=0s      max=224.47ms p(90)=0s      p(95)=0s      
     http_req_duration..............: avg=15.45ms  min=357.37µs med=12.65ms max=101.79ms p(90)=29.92ms p(95)=40.97ms 
       { expected_response:true }...: avg=15.45ms  min=357.37µs med=12.65ms max=101.79ms p(90)=29.92ms p(95)=40.97ms 
     http_req_failed................: 0.00%   0 out of 29516
     http_req_receiving.............: avg=356.19µs min=4.79µs   med=10.54µs max=65.11ms  p(90)=65.39µs p(95)=146.06µs
     http_req_sending...............: avg=670.34µs min=2.45µs   med=8.83µs  max=83.04ms  p(90)=1.23ms  p(95)=4.17ms  
     http_req_tls_handshaking.......: avg=0s       min=0s       med=0s      max=0s       p(90)=0s      p(95)=0s      
     http_req_waiting...............: avg=14.42ms  min=332.37µs med=12.07ms max=99.85ms  p(90)=27.42ms p(95)=36.96ms 
     http_reqs......................: 29516   484.446377/s
     iteration_duration.............: avg=1.02s    min=1s       med=1.01s   max=1.35s    p(90)=1.04s   p(95)=1.05s   
     iterations.....................: 29516   484.446377/s
     vus............................: 500     min=500            max=500
     vus_max........................: 500     min=500            max=500


running (1m00.9s), 000/500 VUs, 29516 complete and 0 interrupted iterations
default ✓ [ 100% ] 500 VUs  1m0s
```

For **1000** VU

```
      ✗ status was 200
      ↳  99% — ✓ 50961 / ✗ 235

     checks.........................: 99.54% 50961 out of 51196
     data_received..................: 7.2 MB 119 kB/s
     data_sent......................: 9.1 MB 149 kB/s
     http_req_blocked...............: avg=37.56ms min=0s       med=3.41µs  max=13.21s   p(90)=10.62µs  p(95)=19.51µs
     http_req_connecting............: avg=37.54ms min=0s       med=0s      max=13.21s   p(90)=0s       p(95)=0s     
     http_req_duration..............: avg=12.8ms  min=0s       med=10.66ms max=132.75ms p(90)=23.1ms   p(95)=29.17ms
       { expected_response:true }...: avg=12.86ms min=418.45µs med=10.71ms max=132.75ms p(90)=23.15ms  p(95)=29.24ms
     http_req_failed................: 0.45%  235 out of 51196
     http_req_receiving.............: avg=62.52µs min=0s       med=12.62µs max=17.11ms  p(90)=48.33µs  p(95)=99.33µs
     http_req_sending...............: avg=456.4µs min=0s       med=11.7µs  max=57.84ms  p(90)=868.89µs p(95)=2.38ms 
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s       p(90)=0s       p(95)=0s     
     http_req_waiting...............: avg=12.28ms min=0s       med=10.35ms max=132.64ms p(90)=22.25ms  p(95)=27.23ms
     http_reqs......................: 51196  839.420445/s
     iteration_duration.............: avg=1.18s   min=1s       med=1.01s   max=31.03s   p(90)=1.03s    p(95)=1.05s  
     iterations.....................: 51196  839.420445/s
     vus............................: 288    min=288            max=1000
     vus_max........................: 1000   min=1000           max=1000


running (1m01.0s), 0000/1000 VUs, 51196 complete and 0 interrupted iterations
default ✓ [ 100% ] 1000 VUs  1m0s

```

For **10000** VUs

```
          ✗ status was 200
      ↳  95% — ✓ 169653 / ✗ 8309

     checks.........................: 95.33% 169653 out of 177962
     data_received..................: 24 MB  360 kB/s
     data_sent......................: 30 MB  451 kB/s
     http_req_blocked...............: avg=284.36ms min=0s         med=4.87µs   max=29.09s p(90)=45.54µs  p(95)=295.61µs
     http_req_connecting............: avg=284.05ms min=0s         med=0s       max=28.99s p(90)=0s       p(95)=0s      
     http_req_duration..............: avg=343.82ms min=0s         med=192.98ms max=9.29s  p(90)=904.07ms p(95)=1.17s   
       { expected_response:true }...: avg=360.66ms min=141.33µs   med=210.83ms max=9.29s  p(90)=919.57ms p(95)=1.18s   
     http_req_failed................: 4.66%  8309 out of 177962
     http_req_receiving.............: avg=2.78ms   min=-2164097ns med=22.33µs  max=1.44s  p(90)=148.12µs p(95)=464.36µs
     http_req_sending...............: avg=16.77ms  min=0s         med=23.33µs  max=2s     p(90)=26.81ms  p(95)=72.73ms 
     http_req_tls_handshaking.......: avg=0s       min=0s         med=0s       max=0s     p(90)=0s       p(95)=0s      
     http_req_waiting...............: avg=324.26ms min=0s         med=175.54ms max=9.29s  p(90)=864.7ms  p(95)=1.12s   
     http_reqs......................: 177962 2659.975428/s
     iteration_duration.............: avg=3.28s    min=1s         med=1.38s    max=35.22s p(90)=3.25s    p(95)=21.75s  
     iterations.....................: 177962 2659.975428/s
     vus............................: 1220   min=1220             max=10000
     vus_max........................: 10000  min=10000            max=10000


running (1m06.9s), 00000/10000 VUs, 177962 complete and 0 interrupted iterations
default ✓ [ 100% ] 10000 VUs  1m0s

```

