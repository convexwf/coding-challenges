#include <iostream>
#include <string>
#include <fstream>
#include <unordered_set>

int main(int argc, char* argv[]) {
    if(argc != 2) {
        std::cerr << "Usage: " << argv[0] << " <file>" << std::endl;
        return 1;
    }

    std::string filename = argv[1];
    std::ifstream input(filename);
    if(!input.is_open()) {
        std::cerr << "Error opening file: " << filename << std::endl;
        return 1;
    }

    std::unordered_set<std::string> lines;
    std::string line;
    while(std::getline(input, line)) {
        if(lines.find(line) == lines.end()) {
            std::cout << line << std::endl;
            lines.insert(line);
        }
    }

    return 0;
}