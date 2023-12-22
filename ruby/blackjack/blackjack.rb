# I hate case statements
module Blackjack
  CARD_VALUES = {
    "ace": 11, "king": 10, "queen": 10, "jack": 10,
    "ten": 10, "nine": 9, "eight": 8, "seven": 7,
    "six": 6, "five": 5, "four": 4, "three": 3,
    "two": 2
  }
  def self.parse_card(card)
    CARD_VALUES[card.to_sym] || 0
  end

  def self.card_range(card1, card2)
    case parse_card(card1) + parse_card(card2)
    when 4..11 then "low"
    when 12..16 then "mid"
    when 17..20 then "high"
    when 21 then "blackjack"
    end
  end

  def self.first_turn(card1, card2, dealer_card)
    range = card_range(card1, card2)
    case 
    when card1 == "ace" && card2 == "ace" then "P"
    when range == "blackjack" 
      ["ace", "ten", "jack", "queen", "king"].include?(dealer_card) ? "S" : "W"
    when range == "high" then "S"
    when range == "mid"
      if parse_card(dealer_card) >= 7
        "H"
      else
        "S"
      end
    when range == "low" then "H"
    end
  end
end
