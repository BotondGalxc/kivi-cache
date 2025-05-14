import grpc from 'k6/net/grpc';
import { check, sleep } from 'k6';

const client = new grpc.Client();
client.load(['definitions'], '../../cache/cache.proto');

export const options = {
  vus: 50,
  duration: '60s',
};

function makeid(length) {
    var result           = '';
    var characters       = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    var charactersLength = characters.length;
    for ( var i = 0; i < length; i++ ) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength));
    }
    return result;
}

export default () => {
    client.connect('localhost:5001', {
			plaintext: true,
			timeout: '10s'
		});

    var key = makeid(2)
    var value = makeid(10)

    const putRequest = { key: key, value: value, expiresSec: 5 };
    const response = client.invoke('cache.KiviCacheService/Put', putRequest);

    check(response, {
        'status is OK': (r) => r && r.status === grpc.StatusOK,
    });

    client.close();

    for (let step = 0; step < 4; step++) {
        client.connect('localhost:5001', {
			plaintext: true,
			timeout: '10s'
		});
        const getRequest = { key: key };
        const response = client.invoke('cache.KiviCacheService/Get', getRequest);

        check(response, {
            'status is OK': (r) => r && r.status === grpc.StatusOK,
        });

        client.close();
    }
}