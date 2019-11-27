import fileinput


def main():
    counts = dict()
    for input in fileinput.input():
        key_input = input.rstrip()
        if key_input in counts.keys():
            counts[key_input] += 1
        else:
            counts[key_input] = 1

    for line, n in counts.items():
        if n > 1:
            print(f"{n}\t{line}")


if __name__ == "__main__":
    main()
