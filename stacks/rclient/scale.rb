system("sudo docker stack deploy --compose-file docker-compose.yml rclient")

100.times do |n|
    sleep 30
    system("sudo docker service scale rclient_rclient=#{1 + n}")
end
