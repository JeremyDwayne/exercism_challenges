class LocomotiveEngineer
  def self.generate_list_of_wagons(*wagon_ids)
    wagon_ids
  end

  def self.fix_list_of_wagons(each_wagons_id, missing_wagons)
    wagon_1, wagon_2, locomotive, *rest = each_wagons_id
    [locomotive, *missing_wagons, *rest, wagon_1, wagon_2]
  end

  def self.add_missing_stops(route, **missing_stops)
    route.tap { |r| r[:stops] = *missing_stops.values }
  end

  def self.extend_route_information(route, more_route_information)
    {**route, **more_route_information}
  end
end
