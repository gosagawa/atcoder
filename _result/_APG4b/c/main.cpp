#include <bits/stdc++.h>
#define out(X) cout << (X) << endl;
#ifdef __LOCAL
#define DBG(X) cout << #X << " = " << (X) << endl;
#else
#define DBG(X)
#endif

using namespace std;

int main() {
#ifdef __LOCAL
    freopen("input", "r", stdin);
#endif

    int x, a, b;

    cin >> x;
    cin >> a;
    cin >> b;

    x++;
    out(x);
    x *= a + b;
    out(x);
    x *= x;
    out(x);
    x--;
    out(x);
}
