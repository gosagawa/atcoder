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

    int a, b;

    cin >> a;
    cin >> b;
    cout << "A:";
    int i;
    i = 0;
    while (i < a) {
        cout << "]";
        i++;
    }
    cout << endl;

    cout << "B:";
    i = 0;
    while (i < b) {
        cout << "]";
        i++;
    }
    cout << endl;
}
