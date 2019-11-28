import fileinput
import sys


def countLines(f, counts):
    for input in f:
        input = input.rstrip()
        if input in counts.keys():
            counts[input] += 1
        else:
            counts[input] = 1


def main():
    counts = dict()
    files = sys.argv[1:]

    if len(files):
        for file in files:
            try:
                f = open(file, 'r')
                with f:
                    countLines(f, counts)
            except Exception as err:
                sys.stderr.write('ERROR: %s\n' % str(err))
                continue
    else:
        countLines(sys.stdin, counts)

    for line, n in counts.items():
        if n > 1:
            print(f"{n}\t{line}")


if __name__ == "__main__":
    main()
