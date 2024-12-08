#include <iostream>
#include <fstream>
#include <string>
#include "input.h"

int main() {
    for (unsigned int i = 0; i < __input_txt_len; ++i) {
        std::cout << __input_txt[i];
    }
    std::cout << std::endl;

    return 0;
}
