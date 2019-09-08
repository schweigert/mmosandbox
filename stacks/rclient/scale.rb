50.times do |n|
    sleep 60
    system("sudo docker service scale rclient_rclient=#{(n + 1) * 2}")
end