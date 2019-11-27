import sys


def main():
    s, sep = '', ''
    for arg in sys.argv[1:] :
        s += sep + arg
        sep = ' '
    print(s)


if __name__ == "__main__":
    main()
