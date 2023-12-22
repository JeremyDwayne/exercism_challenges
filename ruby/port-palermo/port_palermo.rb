module Port
  IDENTIFIER = :PALE

  def self.get_identifier(city)
    city[0..3].upcase.to_sym
  end

  def self.get_terminal(ship_identifier)
    type = ship_identifier[0..2]
    ['OIL', 'GAS'].include?(type) ? :A : :B
  end
end
