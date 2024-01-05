// Build your own wc tool in C++ Language.

#include <iostream>
#include <fstream>
#include <string>
#include <vector>

class WordCounter {
private:
    int wordCount;
    int lineCount;
    int charCount;

public:
    WordCounter() : wordCount(0), lineCount(0), charCount(0) {}

    void countWords(std::string fileName) {
        std::ifstream file(fileName);
        std::string line;

        while (std::getline(file, line)) {
            lineCount++;
            charCount += line.size();

            std::string word;
            for (char c : line) {
                if (std::isspace(c)) {
                    if (!word.empty()) {
                        wordCount++;
                        word.clear();
                    }
                } else {
                    word.push_back(c);
                }
            }

            if (!word.empty()) {
                wordCount++;
            }
        }
    }

    void displayCounts() {
        std::cout << "Number of words: " << wordCount << std::endl;
        std::cout << "Number of lines: " << lineCount << std::endl;
        std::cout << "Number of characters: " << charCount << std::endl;
    }
};

int main() {
    WordCounter wc;
    wc.countWords("input.txt");
    wc.displayCounts();
    return 0;
}