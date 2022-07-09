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

    int n, a;

    cin >> n;
    cin >> a;
    for (int i = 0; i < n; ++i) {
        char op;
        int b;
        cin >> op;
        cin >> b;
        if (op == '/' && b == 0) {
            out("error") break;
        }
        switch (op) {
            case '+':
                a += b;
                break;
            case '-':
                a -= b;
                break;
            case '*':
                a *= b;
                break;
            case '/':
                a /= b;
                break;
        }
        cout << i + 1;
        cout << ":";
        out(a);
    }
}
