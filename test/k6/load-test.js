import http from 'k6/http';
import {check} from 'k6';

const baseURL = 'http://arch.homework/api/v1'

export const options = {
    discardResponseBodies: true,
    scenarios: {
        readUser: {
            executor: 'ramping-arrival-rate',
            exec: 'readUser',
            preAllocatedVUs: 20, // how large the initial pool of VUs would be
            maxVUs: 100, // if the preAllocatedVUs are not enough, we can initialize more
            stages: [
                {duration: '1m', target: 3500},
                {duration: '2m', target: 4200},
                {duration: '3m', target: 5300},
                {duration: '4m', target: 6350},
                {duration: '5m', target: 3400},
            ]
        },
        createUser: {
            executor: 'ramping-arrival-rate',
            exec: 'createUser',
            preAllocatedVUs: 20, // how large the initial pool of VUs would be
            maxVUs: 100, // if the preAllocatedVUs are not enough, we can initialize more
            stages: [
                {duration: '1m', target: 1250},
                {duration: '2m', target: 1320},
                {duration: '3m', target: 2130},
                {duration: '4m', target: 1535},
                {duration: '5m', target: 2140},
            ]
        },
        updateUser: {
            executor: 'ramping-arrival-rate',
            exec: 'updateUser',
            preAllocatedVUs: 20, // how large the initial pool of VUs would be
            maxVUs: 100, // if the preAllocatedVUs are not enough, we can initialize more
            stages: [
                {duration: '3m', target: 8200},
                {duration: '2m', target: 6400},
                {duration: '5m', target: 3400},
                {duration: '1m', target: 7500},
                {duration: '4m', target: 6250},
            ]
        },
        deleteUser: {
            executor: 'ramping-arrival-rate',
            exec: 'deleteUser',
            preAllocatedVUs: 20, // how large the initial pool of VUs would be
            maxVUs: 100, // if the preAllocatedVUs are not enough, we can initialize more
            stages: [
                {duration: '3m', target: 130},
                {duration: '2m', target: 320},
                {duration: '1m', target: 250},
                {duration: '5m', target: 140},
                {duration: '4m', target: 535},
            ]
        },
    },
};

export function readUser() {
    const id = getRandomInt(1, 1000)
    const res = http.get(baseURL + '/user/' + id, {tags: {name: 'readUser'}});
    check(res, {'status was 200': (r) => r.status == 200});
}

export function createUser() {
    const url = baseURL + '/user';
    const payload = JSON.stringify({
        username: "johndoe_"+getRandomInt(100, 100000),
        firstName: "John",
        lastName: "Doe",
        email: "john_doe_"+getRandomInt(100, 100000)+"@example.com",
        phone: "+7123456789"
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
        tags: {
            name: 'createUser',
        },
    };

    const res = http.post(url, payload, params);
    check(res, {'status was 201': (r) => r.status == 201});
}

export function updateUser() {
    const url = baseURL + '/user/'+getRandomInt(1, 1000);
    const payload = JSON.stringify({
        firstName: "Johnny",
        lastName: "Silverhand",
        email: "johnny_silverhand_"+getRandomInt(100, 100000)+"@example.com",
        phone: "+7123456789"
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
        tags: {
            name: 'updateUser',
        },
    };

    const res = http.put(url, payload, params);
    check(res, {'status was 200': (r) => r.status == 200});
}

export function deleteUser() {
    const id = getRandomInt(1000, 10000)
    const res = http.del(baseURL + '/user/' + id, {tags: {name: 'deleteUser'}});
    check(res, {'status was 204': (r) => r.status == 204});
}

function getRandomInt(min, max) {
    return Math.floor(Math.random() * (max - min) + min);
}
