import http from 'k6/http';
import {check} from 'k6';

const baseURL = 'http://arch.homework/api/v1'

export const options = {
    discardResponseBodies: true,
    scenarios: {
        readUser: {
            executor: 'constant-vus',
            exec: 'readUser',
            vus: 1,
            duration: '10m',
        },
        createUser: {
            executor: 'constant-vus',
            exec: 'createUser',
            vus: 1,
            duration: '10m',
        },
        updateUser: {
            executor: 'constant-vus',
            exec: 'updateUser',
            vus: 1,
            duration: '10m',
        },
        deleteUser: {
            executor: 'constant-vus',
            exec: 'deleteUser',
            vus: 1,
            duration: '10m',
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
        username: "johndoe_" + getRandomInt(100, 100000),
        firstName: "John",
        lastName: "Doe",
        email: "john_doe_" + getRandomInt(100, 100000) + "@example.com",
        phone: "+7123456789",
        age: getRandomInt(1, 99),
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
    const url = baseURL + '/user/' + getRandomInt(1, 1000);
    const payload = JSON.stringify({
        firstName: "Johnny",
        lastName: "Silverhand",
        email: "johnny_silverhand_" + getRandomInt(100, 100000) + "@example.com",
        phone: "+7123456789",
        age: getRandomInt(1, 99),
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
