import random
from timeit import timeit


def count_fruits(fruits: list[str]) -> dict[str, int]:
    return {fruit: fruits.count(fruit) for fruit in set(fruits)}


def main() -> None:
    assert count_fruits(
        [
            "apple",
            "banana",
            "apple",
            "cherry",
            "banana",
            "cherry",
            "apple",
            "apple",
            "cherry",
            "banana",
            "cherry",
        ]
    ) == {"apple": 4, "banana": 3, "cherry": 4}
    assert count_fruits([]) == {}

    ver_long_list = random.choices(['apple', 'banana', 'cherry'], k=1000000)
    print(timeit(lambda: count_fruits(ver_long_list), number=100) / 100)


if __name__ == "__main__":
    main()
