import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  vus: 500 ,
  duration: '60s',
};

export default function () {
  const ipAddress = __ENV.IP_ADDRESS; 
  const payload = JSON.stringify({ 
    player1: "paper",
    player2: "rock",
  });
  const res = http.post(`http://${ipAddress}:8084/match`, payload, {
    headers: { 'Content-Type': 'application/json' }, 
  });
  check(res, { 'status was 200': (r) => r.status == 200 });
  sleep(1);
}

