=begin
Write your code for the 'Series' exercise in this file. Make the tests in
`series_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/series` directory.
=end
# try 1 dont like
# class Series
#   def initialize(digits)
#     @digits = digits.split("")
#   end
#
#   def slices(num)
#     raise ArgumentError if num < 1
#     raise ArgumentError if @digits.empty?
#     raise ArgumentError if @digits.length < num
#
#     output = []
#     @digits.each_with_index do |c,i|
#       break if i+num > @digits.length
#       temp = @digits[i...(i+num)]
#       output << temp.join
#     end
#     output
#   end
# end

# Try 2, much cleaner
# but now you have to know what each_cons and some ruby magic mean
class Series
  def initialize(digits)
    @digits = digits
  end

  def slices(num)
    raise ArgumentError if num < 1
    raise ArgumentError if @digits.empty?
    raise ArgumentError if @digits.length < num

    @digits.chars.each_cons(num).map(&:join)  end
end
