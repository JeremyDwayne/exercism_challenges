=begin
Write your code for the 'Acronym' exercise in this file. Make the tests in
`acronym_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/acronym` directory.
=end

class Acronym
  def self.abbreviate(string)
    acronym = ""
    string.split(/\W+/).each do |word|
      acronym += word[0].upcase unless word.empty?
    end
    acronym
  end
end
