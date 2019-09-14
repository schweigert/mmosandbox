system("sudo docker stack deploy --compose-file docker-compose.yml wclient")
system("sudo docker service scale wclient_wclient=0")
sleep 5 * 60
system("sudo docker service scale wclient_wclient=1")
sleep 5 * 60

100.times do |n|
    sleep 30
    system("sudo docker service scale wclient_wclient=#{1 + n}")
end
