import fileinput
import sys


def main():
    counts = dict()
    files = sys.argv[1:]

    for file in files:
        try:
            f = open(file, 'r')
            with f:
                #countLines(f, counts)
                data = f.readlines()
                for line in data:
                    line = line.rstrip()
                    if line in counts.keys():
                        counts[line] += 1
                    else:
                        counts[line] = 1
        except Exception as err:
            sys.stderr.write('ERROR: %s\n' % str(err))
            continue

    for line, n in counts.items():
        if n > 1:
            print(f"{n}\t{line}")


if __name__ == "__main__":
    main()
