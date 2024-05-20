import requests

session = requests.Session()
session.proxies = {
    'http': 'http://USERNAME:PASSWORD@resi-global-gateways.databay.com:7676',
    'https': 'http://USERNAME:PASSWORD@resi-global-gateways.databay.com:7676'
}

response_cloudflare = session.get('https://databay.com/cdn-cgi/trace')
print(response_cloudflare.text)