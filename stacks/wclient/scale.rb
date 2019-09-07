25.times do |n|
    sleep 60
    system("sudo docker service scale wclient_wclient=#{(n + 1) * 2}")
end