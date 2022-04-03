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

    // 変数a,b,cにtrueまたはfalseを代入してAtCoderと出力されるようにする。
    bool a = true;   // true または false
    bool b = false;  // true または false
    bool c = true;   // true または false

    // ここから先は変更しないこと

    if (a) {
        cout << "At";
    } else {
        cout << "Yo";
    }

    if (!a && b) {
        cout << "Bo";
    } else if (!b || c) {
        cout << "Co";
    }

    if (a && b && c) {
        cout << "foo!";
    } else if (true && false) {
        cout << "yeah!";
    } else if (!a || c) {
        cout << "der";
    }

    cout << endl;
}
