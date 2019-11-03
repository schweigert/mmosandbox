require 'json'

value = [
  'rmetrics_001/bench/rclient/create_account_operation',
  'rmetrics_001/bench/rclient/create_character_operation',
  'rmetrics_001/bench/rclient/listen_chat',
  'rmetrics_001/bench/rclient/move_character',
  'rmetrics_001/bench/rclient/send_chat',
  'rmetrics_001/bench/rclient/spawn_character',
  'rmetrics_001/bench/rclient/start_session',
  'rmetrics_002/bench/rclient/create_account_operation',
  'rmetrics_002/bench/rclient/create_character_operation',
  'rmetrics_002/bench/rclient/listen_chat',
  'rmetrics_002/bench/rclient/move_character',
  'rmetrics_002/bench/rclient/send_chat',
  'rmetrics_002/bench/rclient/spawn_character',
  'rmetrics_002/bench/rclient/start_session',
  'rmetrics_003/bench/rclient/create_account_operation',
  'rmetrics_003/bench/rclient/create_character_operation',
  'rmetrics_003/bench/rclient/listen_chat',
  'rmetrics_003/bench/rclient/move_character',
  'rmetrics_003/bench/rclient/send_chat',
  'rmetrics_003/bench/rclient/spawn_character',
  'rmetrics_003/bench/rclient/start_session',
  'smetrics_001/bench/sclient/create_account_operation',
  'smetrics_001/bench/sclient/create_character_operation',
  'smetrics_001/bench/sclient/listen_chat',
  'smetrics_001/bench/sclient/move_character',
  'smetrics_001/bench/sclient/send_chat',
  'smetrics_001/bench/sclient/spawn_character',
  'smetrics_001/bench/sclient/start_session',
  'smetrics_002/bench/sclient/create_account_operation',
  'smetrics_002/bench/sclient/create_character_operation',
  'smetrics_002/bench/sclient/listen_chat',
  'smetrics_002/bench/sclient/move_character',
  'smetrics_002/bench/sclient/send_chat',
  'smetrics_002/bench/sclient/spawn_character',
  'smetrics_002/bench/sclient/start_session',
  'smetrics_003/bench/sclient/create_account_operation',
  'smetrics_003/bench/sclient/create_character_operation',
  'smetrics_003/bench/sclient/listen_chat',
  'smetrics_003/bench/sclient/move_character',
  'smetrics_003/bench/sclient/send_chat',
  'smetrics_003/bench/sclient/spawn_character',
  'smetrics_003/bench/sclient/start_session',
  'wmetrics_001/bench/wclient/create_account_operation',
  'wmetrics_001/bench/wclient/create_character_operation',
  'wmetrics_001/bench/wclient/listen_chat',
  'wmetrics_001/bench/wclient/move_character',
  'wmetrics_001/bench/wclient/send_chat',
  'wmetrics_001/bench/wclient/spawn_character',
  'wmetrics_001/bench/wclient/start_session',
  'wmetrics_002/bench/wclient/create_account_operation',
  'wmetrics_002/bench/wclient/create_character_operation',
  'wmetrics_002/bench/wclient/listen_chat',
  'wmetrics_002/bench/wclient/move_character',
  'wmetrics_002/bench/wclient/send_chat',
  'wmetrics_002/bench/wclient/spawn_character',
  'wmetrics_002/bench/wclient/start_session',
  'wmetrics_003/bench/wclient/create_account_operation',
  'wmetrics_003/bench/wclient/create_character_operation',
  'wmetrics_003/bench/wclient/listen_chat',
  'wmetrics_003/bench/wclient/move_character',
  'wmetrics_003/bench/wclient/send_chat',
  'wmetrics_003/bench/wclient/spawn_character',
  'wmetrics_003/bench/wclient/start_session',
].map do |operation|
  {
    operation: operation,
    dirs: Dir.glob("#{operation}/**/*/*")
  }
end.map do |dir_tree|
  {
    operation: dir_tree[:operation],
    data: dir_tree[:dirs].map { |path| JSON.parse(File.read(path)) }.flatten.sort_by { |set| set['time'] }
  }
end.each do |data_payload|
  File.write("#{data_payload[:operation]}/data.json", JSON.pretty_generate(data_payload[:data]))
end