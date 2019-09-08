50.times do |n|
    sleep 60
    system("sudo docker service scale sclient_sclient=#{(n + 1) * 2}")
end