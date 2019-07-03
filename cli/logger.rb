# typed: true
# frozen_string_literal: true

require 'sorbet-runtime'

module OrsumInflandi
  class Logger
    extend T::Sig

    attr_reader :log_message

    sig { params(log_message: String).void }
    def initialize(log_message)
      @log_message = log_message
    end

    sig { void }
    def info_log
      log('Info', :blue)
    end

    sig { void }
    def frontend_log
      log('Frontend', :green)
    end

    sig { void }
    def backend_log
      log('Backend', :light_blue)
    end

    sig { void }
    def client_log
      log('Client', :green)
    end

    sig { void }
    def server_log
      log('Server', :light_blue)
    end

    sig { params(type: String, color: Symbol).void }
    def log(type, color)
      puts "#{add_padding(type).colorize(color)} #{@log_message}"
    end

    private

    sig { params(str: String).returns(String) }
    def add_padding(str)
      "#{str.ljust(10, ' ')}|"
    end
  end
end
