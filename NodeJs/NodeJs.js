const axios = require('axios');
const https = require('https');

axios({
    method: 'get',
    url: 'https://databay.com/cdn-cgi/trace/',
    httpsAgent: new https.Agent({ rejectUnauthorized: false, proxy: false }),
    proxy: {
        host: 'resi-global-gateways.databay.com',
        port: 7676,
        auth: { username: 'USERNAME', password: 'PASSWORD' },
		protocol: 'http'
    }
})
.then(response => console.log(response.data))
.catch(error => console.error('Error:', error));