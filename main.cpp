#include <iostream>
#include <memory>
#include <utility>


enum Colour {
    Red,
    Green,
    Blue
};

Colour colourFromChar(const char c) {
    switch (c) {
        case 'G':
            return Colour::Green;
        case 'R':
            return Colour::Red;
        default:
            return Colour::Blue;
    }
}

char colourToChar(Colour colour) {
    switch (colour) {
        case Red:
            return 'R';
        case Green:
            return 'G';
        default:
            return 'B';
    }
}

struct Block {
    Colour colour;
    int length;
    std::shared_ptr<Block> next;
};

typedef std::shared_ptr<Block> Row;

Row stringToBlocks(const std::string &s) {
    if (s.empty()) {
        return nullptr;
    }
    Row row = std::make_unique<Block>(Block{colourFromChar(s[0]), 1, nullptr});
    std::shared_ptr<Block> currentBlock = row;
    for (int i = 1; i < s.size(); ++i) {
        Colour currentColor = colourFromChar(s[i]);
        if (currentColor == currentBlock->colour) {
            currentBlock->length++;
        } else {
            currentBlock->next = std::make_unique<Block>(Block{currentColor, 1, nullptr});
            currentBlock = currentBlock->next;
        }
    }
    return row;
}

std::string blocksToString(const Row &row) {
    Row currentBlock = row;
    std::string s;
    while (currentBlock != nullptr) {
        for (int i = 0; i < currentBlock->length; ++i) {
            s += colourToChar(currentBlock->colour);
        }
        currentBlock = currentBlock->next;
    }
    return s;
}


void transformRow(const Row &row) {
    if (row == nullptr) return;
    Row currentBlock = row;
    Row previous;
    while (currentBlock != nullptr) {
        Colour newColour;
        if (currentBlock->next != nullptr) {
            newColour = Colour(3 - currentBlock->colour - currentBlock->next->colour);
            if (currentBlock->length > 1) {
                currentBlock->length--;
                std::shared_ptr<Block> newBlock = std::make_unique<Block>(Block{newColour, 1, nullptr});
                newBlock->next = currentBlock->next;
                currentBlock->next = newBlock;
                previous = newBlock;
                currentBlock = newBlock->next;
            } else {
                if (previous != nullptr && newColour == previous->colour) {
                    previous->length++;
                    previous->next = currentBlock->next;
                } else {
                    currentBlock->colour = newColour;
                    previous = currentBlock;
                }
                currentBlock = currentBlock->next;
            }
        } else {
            if (currentBlock->length > 1) {
                currentBlock->length--;
            } else {
                previous->next = nullptr;
            }
            currentBlock = nullptr;
        }
    }
}


void printRow(const Row &row) {
    Row currentBlock = row;
    std::cout << "Row" << std::endl;
    while (currentBlock != nullptr) {
        std::cout << currentBlock->colour << "-";
        std::cout << currentBlock->length << std::endl;
        currentBlock = currentBlock->next;
    }
}

//#pragma clang diagnostic push
//#pragma ide diagnostic ignored "performance-unnecessary-value-param"
//
//char triangle(std::string row_str) {
//    std::cout << row_str << std::endl;
//    Row currentBlock = stringToBlocks(row_str);
////    int i = 1;
//    while (currentBlock->next != nullptr || currentBlock->length > 1) {
//        transformRow(currentBlock);
////        for (int j = 0; j < i; ++j) {
////            std::cout << " ";
////        }
////        i++;
////        std::cout << blocksToString(currentBlock) << std::endl;
//    }
//    return *blocksToString(currentBlock).c_str();
//}
//
//#pragma clang diagnostic pop
#include <string>
#include <vector>
#include <cstdio>
int pow(int base, int exponent){
    if(exponent == 0){
        return 1;
    }
    int out = base;
    for(int i = 0; i < exponent - 1; --exponent){
        out *= base;
    }
    return out;
}

class tree {
public:
    tree(std::string& ptr) : row(ptr) {
        convto_layer(row.size() - 1);
        max_depth = layer.size();
    };

    char solve(uint_fast16_t ini_offset, uint_fast16_t depth){
        if(depth == max_depth){
            return row[ini_offset];
        }
        return get_colour(solve(ini_offset, 1 + depth),solve(ini_offset + layer[depth], 1 + depth));
    };

    char get_colour(char a, char b){
        if(a == b){
            return a;
        }
        else{
            for(int i = 0; i < 3; ++i){
                if(!((a == colour[i]) or (b == colour[i]))){
                    return colour[i];
                }
            }

        }
        return 'O';
    };

    void convto_layer(int_fast16_t n){
        if(n == 1){
            layer.push_back(1);
            return;
        }
        int power = 0;
        while(n > pow(3, power)){
            ++power;
        }
        --power;
        while(n != 0){
            int m = pow(3, power);
            if(n - m < 0){
                --power;
            }
            else{
                layer.push_back(m);
                n -= m;
            }
        }

    };

    std::string& row;
    std::vector<uint_fast16_t> layer;
    uint_fast8_t max_depth = 0;
    char colour[3] = {'R', 'G', 'B'};
};

std::string triangle(std::string row) {
    tree three(row);
    char x = three.solve(0,0);
    std::string out = "";
    out.push_back(x);
    return out;
}

std::string generateRandomRow(int length) {
    std::string row = "";
    for (int i = 0; i < length; ++i) {
        row += colourToChar(Colour(rand() % 3));
    }
    return row;
}


int main() {
//    std::cout << triangle("RBRGBRB") << std::endl;
//    std::cout << triangle("RBRGBRBGGRRRBGBBBGG") << std::endl;
    const std::string &row = generateRandomRow(1000000);
    std::cout << row << std::endl;
    std::cout << triangle(row) << std::endl;
    return 0;


//    Assert::That(triangle("GB"), Equals("R"));
//    Assert::That(triangle("RRR"), Equals("R"));
//    Assert::That(triangle("RGBG"), Equals("B"));
//    Assert::That(triangle("RBRGBRB"), Equals("G"));
//    Assert::That(triangle("RBRGBRBGGRRRBGBBBGG"), Equals("G"));
//    Assert::That(triangle("B"), Equals("B"));
}
