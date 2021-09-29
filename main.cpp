#include <string>
#include <utility>
#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

pair<string, int> encode(const string &s) {
    if (s.empty())
        return {"", 0};
    vector<string> matrix;
    matrix.reserve(s.size());
    for (int i = 0; i < s.size(); ++i) {
        matrix.push_back(s.substr(i, s.size() - i) + s.substr(0, i));
    }
    sort(matrix.begin(), matrix.end());
    int line = 0;
    string code;
    for (int i = 0; i < matrix.size(); ++i) {
        basic_string<char> &row = matrix[i];
        if (row == s) {
            line = i;
        }
        code += row[row.size()-1];
    }
    return make_pair(code, line);
}

string decode(const string &s, int n) {
    return "";
}


int main() {
    const pair<string, int> &one = encode("bananabar");
    cout << one.first << " " << one.second << endl;
    const string &two = decode("nnbbraaaa", 4);
    cout << two << endl;
    return 0;


//    Assert::That(triangle("GB"), Equals("R"));
//    Assert::That(triangle("RRR"), Equals("R"));
//    Assert::That(triangle("RGBG"), Equals("B"));
//    Assert::That(triangle("RBRGBRB"), Equals("G"));
//    Assert::That(triangle("RBRGBRBGGRRRBGBBBGG"), Equals("G"));
//    Assert::That(triangle("B"), Equals("B"));
}
