import http from 'k6/http';
import { sleep } from 'k6';

//10 vus usuarios
//30 duration segundos de duração

export const options = {
   vus: 7000,
   duration: '10s',
};


export default function() {
   // SAMPLE IP
   const MY_CLUSTER_IP = '192.168.58.2';
   
   http.get(`http://${MY_CLUSTER_IP}:31500/v1/delivery/products`);
   sleep(0.5);
}
