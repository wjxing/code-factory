#include <iostream>
#include <SkDecode.h>

using namespace skia;

int main() {
    std::cout << "version : " << VERSION << std::endl;
    SkStream stream("png-decode");
    SkDecode* decode = SkDecode::factory(&stream);
    if (decode)
        decode->decode();
    else
        std::cout << "none decode" << std::endl;
    return 0;
}
