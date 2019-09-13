#!/usr/bin/env ruby
# https://www.graphviz.org/doc/info/colors.html

def is_domain?(line)
    line.include?('mmosandbox/domain')
end

def recolor_domain(line)
    line.gsub('paleturquoise', 'coral')
end

def is_clients?(line)
    line.include?('mmosandbox/clients')
end

def recolor_clients(line)
    line.gsub('paleturquoise', 'gold')
end

def is_lib?(line)
    line.include?('mmosandbox/lib')
end

def recolor_lib(line)
    line.gsub('paleturquoise', 'azure3')
end

def is_infra?(line)
    line.include?('mmosandbox/infra')
end

def recolor_infra(line)
    line.gsub('paleturquoise', 'springgreen')
end

def sub(line)
    return recolor_domain(line) if is_domain?(line)
    return recolor_lib(line) if is_lib?(line)
    return recolor_clients(line) if is_clients?(line)
    return recolor_infra(line) if is_infra?(line)

    return line
end

def remove_github(line)
    line.gsub('github.com/schweigert/mmosandbox/', '').gsub('github.com/', 'vendor/')
end

STDIN.read.split("\n").each do |line|
    puts remove_github(sub(line))
end
