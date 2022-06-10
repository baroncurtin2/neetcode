class Solution:
    def encode(self, strs: list[str]) -> str:
        def encoder(string: str) -> str:
            return f"{len(string)}#{string}"

        return "".join([encoder(string) for string in strs])

    def decode(self, string: str) -> list[str]:
        def decoder(string: str, start_idx: int, end_idx: int) -> (int, str):
            # get word length
            word_len = int(string[start_idx:end_idx])
            # get word + 1 to skip '#'
            word = string[end_idx + 1:end_idx + 1 + word_len]

            return word_len, word

        result = []

        start_idx = 0

        while start_idx < len(string):
            end_idx = start_idx

            # iterate until # to get word length
            while string[end_idx] != "#":
                end_idx += 1

            # get word
            word_len, word = decoder(string, start_idx, end_idx)

            # add word to result
            result.append(word)

            # update start_idx
            start_idx = end_idx + 1 + word_len

        return result


def print_test(test_case: list[str], encoded: str, decoded: list[str]) -> None:
    print(
        f"Test Input: {test_case} \n"
        f"Encoded: {encoded} \n"
        f"Decoded: {decoded} \n"
    )


if __name__ == "__main__":
    test1 = ["lint", "code", "love", "you"]
    test2 = ["we", "say", ":", "yes"]

    sol = Solution()
    encode_t1 = sol.encode(test1)
    decode_t1 = sol.decode(encode_t1)

    print_test(test1, encode_t1, decode_t1)

    encode_t2 = sol.encode(test2)
    decode_t2 = sol.decode(encode_t2)

    print_test(test2, encode_t2, decode_t2)
