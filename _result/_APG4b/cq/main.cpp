#include <bits/stdc++.h>
#define out(X) cout << (X) << endl;
#ifdef __LOCAL
#define dbg(X) cout << #X << " = " << (X) << endl;
#else
#define dbg(X)
#endif

using namespace std;

int main() {
#ifdef __LOCAL
    freopen("input", "r", stdin);
#endif

    int a, b;
    string op;

    cin >> a >> op >> b;
    if (op == "+") {
        out(a + b);
    }
    if (op == "-") {
        out(a - b);
    }
    if (op == "*") {
        out(a * b);
    }
    if (op == "/") {
        if (b == 0) {
            out("error");
        } else {
            out(a / b);
        }
    }
    if (op == "?" || op == "=" || op == "!") {
        out("error");
    }
}
