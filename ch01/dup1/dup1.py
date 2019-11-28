import fileinput


def main():
    counts = dict()
    for input in fileinput.input():
        input = input.rstrip()
        if input in counts.keys():
            counts[input] += 1
        else:
            counts[input] = 1

    for line, n in counts.items():
        if n > 1:
            print(f"{n}\t{line}")


if __name__ == "__main__":
    main()
