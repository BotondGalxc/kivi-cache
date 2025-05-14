import grpc from 'k6/net/grpc';
import { check } from 'k6';
import exec from 'k6/execution';

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

    // Items don't get deleted and let tem expire
    // we ensure, there are no duplicate keys for each run
    var key = exec.vu.idInTest + "-" + makeid(5)
    var value = makeid(10)

    const putRequest = { key: key, value: value, expiresSec: 5 };
    const response = client.invoke('cache.KiviCacheService/Put', putRequest);

    check(response, {
        'Put status OK': (r) => r && r.status === grpc.StatusOK,
    });

    client.close();

    for (let step = 0; step < 10; step++) {
        client.connect('localhost:5001', {
			plaintext: true,
			timeout: '10s'
		});
        const getRequest = { key: key };
        const response = client.invoke('cache.KiviCacheService/Get', getRequest);

        check(response, {
            'Get status OK': (r) => r && r.status === grpc.StatusOK && r.message.value === value,
        });

        client.close();
    }
}