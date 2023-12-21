class AssemblyLine
  CARS_PER_HOUR = 221
  SUCCESS_RATES = {
    1 => 1.0, 2 => 1.0, 3 => 1.0, 4 => 1.0,
    5 => 0.9, 6 => 0.9, 7 => 0.9, 8 => 0.9,
    9 => 0.8,
    10 => 0.77
  }

  def initialize(speed)
    @speed = speed
  end

  def production_rate_per_hour
    (@speed * CARS_PER_HOUR) * SUCCESS_RATES[@speed]
  end

  def working_items_per_minute
    (production_rate_per_hour / 60).to_i
  end
end
