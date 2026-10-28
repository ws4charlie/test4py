#   Coding (Logical and Maintainable)

#o  Coffee machine location

#§  There is an office with developers; we want to install a new coffee machine.
#   Please, find an optimal location for the new coffee machine.

#§  This office is represented by a bidimensional array of Booleans. (bool[][] office)

#§  optimal = min(sum(distance to each developer))

# coffee machine should be in one of these cubicals
class Solution:
    def distance(p1, p2) -> int:
        return abs(p1[0] - p1[0]) + abs(p1[1] - p1[1])

    def find_loccation(self, office: 'List[List[boolean]]') -> '(int, int)':
        row = len(office)
        col = len(office[0])

        # get all devs location
        devs = []
        for r in range(row):
            for c in range(col):
                if office[r][c]:
                    devs.append(r, c)

        optimal = None
        loc = None
        for r in range(row):
            for c in range(col):
                sum = 0
                for d in devs:
                    sum += self.distance(d, (r,c))
                if not loc:
                    optimal = sum
                    loc = (r, c)
                else:
                    if sum < optimal:
                        optimal = sum
                        loc = (r, c)
        return loc

    # both row and col should be in the best place
    def find_loccation2(self, office: 'List[List[boolean]]') -> '(int, int)':
        row = len(office)
        col = len(office[0])

        # get all devs location
        row_sum = 0
        col_sum = 0
        dev_rows = []
        dev_cols = []
        for r in range(row):
            for c in range(col):
                if office[r][c]:
                    row_sum += r
                    col_sum += c
                    dev_rows.append(r)
                    dev_cols.append(c)

        rmin = min(dev_rows)
        rmax = max(dev_rows)
        cmin = min(dev_cols)
        cmax = max(dev_cols)

