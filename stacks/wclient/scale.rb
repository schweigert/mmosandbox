system("sudo docker stack deploy --compose-file docker-compose.yml wclient")

100.times do |n|
    sleep 30
    system("sudo docker service scale wclient_wclient=#{1 + n}")
end
