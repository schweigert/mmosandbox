system("sudo docker stack deploy --compose-file docker-compose.yml sclient")

100.times do |n|
    sleep 30
    system("sudo docker service scale sclient_sclient=#{1 + n}")
end
