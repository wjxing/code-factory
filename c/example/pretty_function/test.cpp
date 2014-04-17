#include <iostream>
#include <typeinfo>

namespace ltt {
class NFC {

};

class TG {

};
}

template<typename T>
class Cate {
  public:
    typedef T R;
    Cate() {
        std::cout << __PRETTY_FUNCTION__ << std::endl;
    }
};

int main() {
    Cate<ltt::NFC> nfc;
    Cate<ltt::TG> tg;
    std::cout << typeid(ltt::NFC).name() << std::endl;
    return 0;
}
