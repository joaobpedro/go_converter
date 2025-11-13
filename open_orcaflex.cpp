#include <iostream>
#include <fstream>
#include <sstream>

int main () {
    // open a file read it and spit the contents to the console in cpp
    std::ifstream source("HH2-PR3-LS-B1-SY01-L60-G26-Detail.txt");
    
    if (!source) {
        std::cerr << "Error opening the file" << std::endl;
        return 1;
    }

    // std::string line;
    // std::cout << "File Contents:" << std::endl;

    // while (std::getline(source, line)) {
    //     std::cout << line << std::endl;
    // }
    //

    std::stringstream buffer;
    buffer << source.rdbuf();
    std::string contents = buffer.str();
    source.close();

    std::cout << "Files contents all at once:";
    std::cout << contents << std::endl;

    return 0;
}
