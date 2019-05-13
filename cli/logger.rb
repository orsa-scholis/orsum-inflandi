# frozen_string_literal: true

module OrsumInflandi
  class Logger
    attr_reader :log_message

    def initialize(log_message)
      @log_message = log_message
    end

    def info_log
      log('Info', :blue)
    end

    def frontend_log
      log('Frontend', :green)
    end

    def backend_log
      log('Backend', :light_blue)
    end

    def client_log
      log('Client', :green)
    end

    def server_log
      log('Server', :light_blue)
    end

    def log(type, color)
      puts "#{add_padding(type).colorize(color)} #{@log_message}"
    end

    private

    def add_padding(str)
      "#{str.ljust(10, ' ')}|"
    end
  end
end
