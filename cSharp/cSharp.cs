using System.Net;

var proxy = new HttpClientHandler
{
    Proxy = new WebProxy("http://resi-global-gateways.databay.com:7676")
    {
        Credentials = new NetworkCredential("USERNAME", "PASSWORD")
    }
};

using var httpClient = new HttpClient(proxy);
var response = await httpClient.GetAsync("https://databay.com/cdn-cgi/trace");
var content = await response.Content.ReadAsStringAsync();
Console.WriteLine(content);