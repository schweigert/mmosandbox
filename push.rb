# docker tag local-image:tagname new-repo:tagname
# docker push new-repo:tagname

def remove_prefix(service)
    service.gsub('mmosandbox_', '')
end

def now
    @now = Time.now
end

def lasted(service)
    [remove_prefix(service), 'lasted'].join('.')
end

def dated(service)
    [remove_prefix(service), now.year, now.month, now.day, now.hour, now.hour, now.min].join('.')
end

services = %w[
    mmosandbox_wclient
    mmosandbox_wweb
    mmosandbox_wgame
    mmosandbox_wauth
    mmosandbox_rgame
    mmosandbox_rauth
    mmosandbox_rweb
    mmosandbox_rcrud
    mmosandbox_rclient
    mmosandbox_sclient
    mmosandbox_sweb
    mmosandbox_sgame
    mmosandbox_schat
    mmosandbox_sauth
]

services.each do |service|
    [dated(service), lasted(service)].each do |tag|
        system("docker tag #{service} schweigert/mmosandbox:#{tag}")
        system("docker push schweigert/mmosandbox:#{tag}")
    end
end
