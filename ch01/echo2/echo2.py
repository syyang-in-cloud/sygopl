import sys


def main():
    s, sep = '', ''
    for _, arg in enumerate(sys.argv[1:]) :
        s += sep + arg
        sep = ' '
    print(s)


if __name__ == "__main__":
    main()
