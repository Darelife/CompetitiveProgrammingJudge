#include <bits/stdc++.h>
using namespace std;
#define int long long

signed main() {
    int t;
    cin >> t;
    while(t--) {
        int n;
        cin >> n;
        vector<int> v(n + 1), cnt(n + 1);
        for(int i = 1; i <= n; i++) {
            cin >> v[i];
            cnt[v[i]]++;
        }

        vector<int> have(n + 1);
        have[n] = cnt[n];
        for(int i = n - 1; i >= 1; i--) have[i] = have[i + 1] + cnt[i];

        int init = 0, aft = 0;
        for(int i = 1; i <= n; i++) {
            init += i * v[i];
            aft += have[i] * (2 * n - have[i] + 1) / 2;
        }

        int cur = aft - init, mx = 0;
        for(int i = 1; i <= n; i++) mx = max(mx, i - n + have[v[i]] - 1);

        cout << cur + mx << endl;
    }
}
