#include <string>
#include <utility>
#include <iostream>
#include <vector>
#include <algorithm>
#include <math.h>
#include <iomanip>

using namespace std;

pair<string, int> encode(const string &s) {
    if (s.empty())
        return {"", 0};
    vector<string> matrix;
    matrix.reserve(s.size());
    for (unsigned int i = 0; i < s.size(); ++i) {
        matrix.push_back(s.substr(i, s.size() - i) + s.substr(0, i));
    }
    sort(matrix.begin(), matrix.end());
    unsigned int line = 0;
    string code;
    for (unsigned int i = 0; i < matrix.size(); ++i) {
        basic_string<char> &row = matrix[i];
        if (row == s) {
            line = i;
        }
        code += row[row.size() - 1];
    }
    return make_pair(code, line);
}

string decode(const string &s, int n) {
    vector<string> code;
    code.reserve(s.size());
    for (unsigned int i = 0; i < s.size(); ++i) {
        std::ostringstream indexString;
        indexString << setw(int(log10(s.size()))+1) << setfill('0') << i;
        code.push_back(s[i] + indexString.str());
    }
    vector<string> sorted(code.size());
    copy(code.begin(), code.end(), sorted.begin());
    sort(sorted.begin(), sorted.end());
    string solution;
    for (unsigned int i = 0; i < s.size(); ++i) {
        string letter = sorted[n];
        solution += letter[0];
        n = stoi(letter.substr(1));
    }
    std::cout << solution << std::endl;
    return solution;
}


int main() {
    const pair<string, int> &one = encode("bananabar");
    cout << one.first << " " << one.second << endl;
    decode("nnbbraaaa", 4);
    const string &two = decode("e emnllbduuHB", 2);
    cout << two << endl;
    return 0;


//    Assert::That(triangle("GB"), Equals("R"));
//    Assert::That(triangle("RRR"), Equals("R"));
//    Assert::That(triangle("RGBG"), Equals("B"));
//    Assert::That(triangle("RBRGBRB"), Equals("G"));
//    Assert::That(triangle("RBRGBRBGGRRRBGBBBGG"), Equals("G"));
//    Assert::That(triangle("B"), Equals("B"));
}
