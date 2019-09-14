require 'open-uri'
require 'json'

metrics = [
    '*.metrics.*.cpu_perc',
    '*.metrics.*.*.*.cpu_perc',
    '*.metrics.*.mem_perc',
    '*.metrics.*.*.*.mem_perc',
    '*.metrics.*.pids',
    '*.metrics.*.*.*.pids',
    '*.metrics.*.block_in',
    '*.metrics.*.*.*.block_in',
    '*.metrics.*.block_out',
    '*.metrics.*.*.*.block_out',
    '*.metrics.*.net_out',
    '*.metrics.*.*.*.net_out',
    '*.metrics.*.net_in',
    '*.metrics.*.*.*.net_in',
    'worlds.*.*',
    'bench.wweb.routes.*.*',
    'bench.wgame.*.*',
    'bench.wauth.*.*',
    'bench.wclient.*.*',
    'gbeacon.*.free_memory_linux_monitor',
    'gbeacon.*.free_swap_linux_monitor',
    'gbeacon.*.total_memory_linux_monitor',
    'gbeacon.*.total_swap_linux_monitor',
    'gbeacon.*.used_memory_linux_monitor',
    'gbeacon.*.used_swap_linux_monitor',
    'gbeacon.*.uw_bandrx',
    'gbeacon.*.uw_bandtx',
    'gbeacon.*.uw_cpuused',
    'gbeacon.*.uw_diskioreads',
    'gbeacon.*.uw_diskiowrites',
    'gbeacon.*.uw_diskused',
    'gbeacon.*.uw_diskused_perc',
    'gbeacon.*.uw_httpconns',
    'gbeacon.*.uw_load',
    'gbeacon.*.uw_memused',
    'gbeacon.*.uw_tcpused',
    'gbeacon.*.uw_udpused',
]


def get(metric)
    base_url = 'http://10.20.218.237:8080/render'

    params = [
        'format=json',
        'from=-100min',
        "target=#{metric}"
    ].join('&')

    puts "\nGetting #{metric}"
    source = open("#{base_url}?#{params}", &:read)
    puts "\ndownloaded: #{ (source.size / (8 * 1024 * 1024.0)).round(2) } MB"

    data = JSON.parse(source)
    data.each do |col|
        subpaths = col['target'].split('.')
        next if subpaths.include?('--')

        acutal_path = ['metrics']

        subpaths.each do |next_path|
            acutal_path << next_path

            Dir.mkdir(acutal_path.join('/')) unless Dir.exist?(acutal_path.join('/'))
        end
        points = col['datapoints'].map do |point|
            { value: point.first, time: point.last }
        end.reject do |point|
            point[:value].nil?
        end

        puts "#{subpaths.join('.')} size: #{points.size}"

        next if points.size.zero?

        acutal_path << 'data.json'
        File.open(acutal_path.join('/'), 'w+') do |file|
            file.write JSON.pretty_generate(points)
        end
    end
end

def try_get(metric)
    get(metric)
rescue e
    puts "Error: #{e}"
    try_get(metric)
end

metrics.each { |metric| try_get(metric) }
