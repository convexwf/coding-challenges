#include <iostream>
#include <string>
#include <map>

// JSON解析器类
class JsonParser {
public:
    JsonParser(const std::string& json) : json_(json) {}

    bool parse() {
        // Check whether 
        if (json_.empty()) {
            return false;
        }

        // 检查JSON格式是否正确
        if (json_[0] != '{' || json_[json_.size() - 1] != '}') {
            return false;
        }

        // 去掉首尾的大括号
        std::string data = json_.substr(1, json_.size() - 2);

        // 按逗号分割键值对
        size_t pos = 0;
        while ((pos = data.find(',')) != std::string::npos) {
            std::string pair = data.substr(0, pos);
            parsePair(pair);
            data.erase(0, pos + 1);
        }
        parsePair(data);

        return true;
    }

    void parsePair(const std::string& pair) {
        // 按冒号分割键值
        size_t pos = pair.find(':');
        std::string key = pair.substr(0, pos);
        std::string value = pair.substr(pos + 1);

        // 去掉引号
        key = key.substr(1, key.size() - 2);
        value = value.substr(1, value.size() - 2);

        // 存储键值对
        data_[key] = value;
    }

    void printData() {
        // 打印解析结果
        for (const auto& pair : data_) {
            std::cout << pair.first << " : " << pair.second << std::endl;
        }
    }

private:
    std::string json_;
    std::map<std::string, std::string> data_;
};

int main() {
    std::string json = R"({"name": "John", "age": 30, "city": "New York"})";
    JsonParser parser(json);
    if (parser.parse()) {
        parser.printData();
    } else {
        std::cout << "Failed to parse JSON" << std::endl;
    }

    return 0;
}