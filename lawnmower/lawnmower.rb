def main
  lines = File.readlines(ARGV[0])
  lines = lines.collect {|line| line.strip }

  test_cases = lines.shift.to_i

  lawns = []

  test_cases.times do |test_case|
    height, width = lines.shift.split(' ').map(&:to_i)
    lawn = Lawn.new(width, height)

    height.times do |height_idx|
      data = lines.shift.split(' ').map(&:to_i)
      lawn.add_row(data)
    end
    lawns << lawn
  end

  lawns.each_with_index do |lawn, idx|
    lawn.link!
    puts "Case ##{idx + 1}: #{lawn.mowable? ? 'YES' : 'NO'}"
  end
end

class Lawn
  attr_accessor :width, :height, :data
  def initialize(width, height)
    @width = width
    @height = height
    @data = []
  end

  def pos(row_idx, col_idx)
    return nil if row_idx == height || col_idx == width
    return nil if row_idx < 0 || col_idx < 0

    data[row_idx][col_idx]
  end

  def to_s
    print "[\n"
    rows.each do |row|
      row.each_with_index do |square, col_idx|
        print "#{square.height}"
      end
      print "\n"
    end
    print "]\n"
  end

  def link!
    rows.each_with_index do |row, row_idx|
      row.each_with_index do |square, col_idx|
        square.link_left( pos(row_idx, col_idx - 1) )
        square.link_right( pos(row_idx, col_idx + 1) )
        square.link_up(   pos(row_idx-1, col_idx) )
        square.link_down( pos(row_idx+1, col_idx) )
      end
    end
  end

  def squares(&block)
    rows.each do |row|
      row.each do |square|
        yield square
      end
    end
  end

  def mowable?
    squares do |square|
      # puts "Square: #{square.inspect}"
      if square.row_idx == 2 && square.col_idx == 2
       #  puts "checking the special square"
        return false if !square.has_exit_path?
      end
    end
    return true
  end

  def rows
    data
  end

  def add_row(row)
    row_idx = data.length
    col_idx = -1
    data[ row_idx ] = row.collect do |height|
      col_idx += 1
      Square.new(height, row_idx, col_idx)
    end
  end
end

class Square
  attr_reader :height, :left, :right, :up, :down, :row_idx, :col_idx
  def initialize(height, row_idx, col_idx)
    @height = height
    @row_idx = row_idx
    @col_idx = col_idx
  end

  def has_exit_path?(visited=[])
    return true if edge?

    visited << self

    # puts "I am height #{height} at #{row_idx},#{col_idx}"

    dirs = [:up, :left, :right, :down]
    dirs.each do |dir|
      # puts "checking: #{dir}"
      square = self.send(dir)
      # puts "square is: #{square}"
      if visitable?(square, visited)
        # puts "squared is visitable"
        visited << square

        return true if square.has_exit_path?(visited)
      end
    end
    false
  end

  def visitable?(square, visited_array)
    return false if square.nil?
    return false if visited_array.include?(square)

    return square.height <= self.height
  end

  def edge?
    left.nil? || right.nil? || up.nil? || down.nil?
  end

  def link_left(square)
    @left = square
  end

  def link_right(square)
    @right = square
  end

  def link_up(square)
    @up = square
  end

  def link_down(square)
    @down = square
  end
end

class Graph
end

main
