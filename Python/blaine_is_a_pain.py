import enum
import numpy as np


class Direction(enum.Enum):
    NORTH = 0
    NORTH_EAST = 1
    EAST = 2
    SOUTH_EAST = 3
    SOUTH = 4
    SOUTH_WEST = 5
    WEST = 6
    NORTH_WEST = 7

    def shift(self):
        return [np.array((0, -1)), np.array((1, -1)), np.array((1, 0)), np.array((1, 1)),
                np.array((0, 1)), np.array((-1, 1)), np.array((-1, 0)), np.array((-1, -1))][self.value]

    def turn(self, counter_clockwise=False):
        return Direction((self.value + len(Direction) + (1 if not counter_clockwise else -1)) % len(Direction))


def train_crash(track, a_train, a_train_pos, b_train, b_train_pos, limit):
    build_track(track)


def build_track(track):
    start_coords = None
    crossings = {}
    stations = {}
    direction = Direction.EAST
    track_list = [list(row) for row in track.split("\n") if len(row.strip()) > 0]
    for x, element in enumerate(track_list[0]):
        if element == "/":
            start_coords = np.array((x, 0))
            break
    coords = start_coords + direction.shift()
    pos = 1
    while not all(coords == start_coords):
        current_element = track_list[coords[1]][coords[0]]
        if current_element in "\\/":
            turns = {
                ("\\", Direction.NORTH): (2, True),
                ("/", Direction.NORTH): (2, False),
                ("/", Direction.NORTH_EAST): (1, False),  # TODO both possible
                ("\\", Direction.EAST): (2, False),
                ("\\", Direction.SOUTH_EAST): (1, False),
                ("\\", Direction.SOUTH): (2, True),
                ("/", Direction.SOUTH): (2, False),
                ("/", Direction.SOUTH_WEST): (2, True),
                ("\\", Direction.WEST),
                ("\\", Direction.NORTH_WEST)
            }
            direction = direction.turn(current_element == "")
        if current_element in "-|":
            coords += direction.shift()
        elif current_element:
            pass
        print(coords, current_element)
        pos += 1


if __name__ == '__main__':
    TRACK_EX = """\
                                /------------\\
/-------------\\                /             |
|             |               /              S
|             |              /               |
|        /----+--------------+------\\        |   
\\       /     |              |      |        |     
 \\      |     \\              |      |        |                    
 |      |      \\-------------+------+--------+---\\
 |      |                    |      |        |   |
 \\------+--------------------+------/        /   |
        |                    |              /    | 
        \\------S-------------+-------------/     |
                             |                   |
/-------------\\              |                   |
|             |              |             /-----+----\\
|             |              |             |     |     \\
\\-------------+--------------+-----S-------+-----/      \\
              |              |             |             \\
              |              |             |             |
              |              \\-------------+-------------/
              |                            |               
              \\----------------------------/ 
"""

    assert train_crash(TRACK_EX, "Aaaa", 147, "Bbbbbbbbbbb", 288, 1000) == 516
