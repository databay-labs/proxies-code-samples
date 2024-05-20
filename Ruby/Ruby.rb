require 'uri'
require 'net/http'

uri = URI.parse('https://databay.com/cdn-cgi/trace/')
proxy = Net::HTTP::Proxy('resi-global-gateways.databay.com', 7676, 'USERNAME', 'PASSWORD')

req = Net::HTTP::Get.new(uri.path)

# Set the HTTP object to use SSL because the uri scheme is https
http = proxy.new(uri.host, uri.port)
http.use_ssl = true
http.verify_mode = OpenSSL::SSL::VERIFY_NONE

result = http.start do |h|
    h.request(req)
end

puts result.code
puts result.body