# frozen_string_literal: true

module OrsumInflandi
  class Logger
    def initialize(log)
      @log = log
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

    def log(type, color)
      puts "#{add_padding(type).colorize(color)} #{@log}"
    end

    private

    def add_padding(str)
      "#{str.ljust(10, ' ')}|"
    end
  end
end
