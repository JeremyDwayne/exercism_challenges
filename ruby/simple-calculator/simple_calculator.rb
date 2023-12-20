class SimpleCalculator
  UnsupportedOperation = Class.new(StandardError)
  ALLOWED_OPERATIONS = ['+', '/', '*'].freeze

  def self.calculate(first_operand, second_operand, operation)
    raise UnsupportedOperation, 'Operation not allowed' unless ALLOWED_OPERATIONS.include?(operation)
    raise ArgumentError, 'First operand must be a number' unless first_operand.is_a?(Numeric)
    raise ArgumentError, 'Second operand must be a number' unless second_operand.is_a?(Numeric)
    return 'Division by zero is not allowed.' if second_operand.zero?

    result = first_operand.method(operation).(second_operand)
    "#{first_operand} #{operation} #{second_operand} = #{result}"
  end
end
