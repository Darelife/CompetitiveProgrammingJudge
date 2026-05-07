import sys


def solve() -> None:
    data = sys.stdin.buffer.read().split()
    it = iter(data)
    t = int(next(it))
    out_lines = []
    for _ in range(t):
        n = int(next(it))
        # 1-indexed array
        a = [0] + [int(next(it)) for _ in range(n)]

        # frequency and sum of positions for each height
        cnt = [0] * (n + 2)
        pos_sum = [0] * (n + 2)
        for i in range(1, n + 1):
            val = a[i]
            cnt[val] += 1
            pos_sum[val] += i

        # m[h] = number of columns with height >= h
        # sum_pos[h] = sum of indices of those columns
        m = [0] * (n + 3)
        sum_pos = [0] * (n + 3)
        for h in range(n, 0, -1):
            m[h] = m[h + 1] + cnt[h]
            sum_pos[h] = sum_pos[h + 1] + pos_sum[h]

        # original total movement
        T = 0
        for h in range(1, n + 1):
            if m[h] == 0:
                continue
            # sum of the m[h] largest indices
            s_max = m[h] * (2 * n - m[h] + 1) // 2
            T += s_max - sum_pos[h]

        best = T
        for i in range(1, n + 1):
            h = a[i]
            # change when removing the top cube of column i
            delta = i - (n - m[h] + 1)
            cand = T + delta
            if cand > best:
                best = cand
        out_lines.append(str(best))

    sys.stdout.write("\n".join(out_lines))


if __name__ == "__main__":
    solve()
