#include <iostream>
template<typename T>
class FinalClass {
  private:
  public:
    FinalClass() {
        std::cout << "Final" << std::endl;
    }
    ~FinalClass() {}
    struct __class {
        typedef T type;
    };
    friend class __class::type;
};

// class IInterface : virtual FinalClass<IInterface> {
class IInterface : virtual FinalClass<IInterface> {
  public:
    IInterface() {
        std::cout << "IInterface" << std::endl;
    }
    virtual ~IInterface() {}
};

class BpInterface : public IInterface {
  public:
    BpInterface() {}
    virtual ~BpInterface() {}
};

int main() {
    new BpInterface();
    return 0;
}
