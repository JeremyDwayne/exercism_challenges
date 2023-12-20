class LogLineParser
  attr_reader :level, :message
  def initialize(line)
    @line = line
    parse
  end

  def message
    @message
  end

  def log_level
    @level.downcase
  end

  def reformat
    "#{@message} (#{@level.downcase})"
  end

  private
  def parse
    regex = /\[(.*)\]: (.*)/
    matches = @line.match(regex)
    @level = matches[1]
    @message = matches[2].strip
  end
end
