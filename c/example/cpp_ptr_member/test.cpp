#include <iostream>

class A {
  public:
    typedef void (A::*Callback)();
  private:
    Callback call;
    void inner() {
        std::cout << "A::inner" << std::endl;
    }
  public:
    void initCallback() {
        call = &A::inner;
    }
    void doCallback() {
        Callback callback = call;
        (this->*callback)();
    }
    Callback getCallback() {
        std::cout << call << std::endl;
        return call;
    }
    void pubInner() {
        std::cout << "A::pubInner" << std::endl;
    }
};

class B {
  public:
    typedef void (B::*CallbackB)();
    void inner() {
        std::cout << "B::inner" << std::endl;
    }
};

class SubA : public A ,public B {
  public:
    typedef void (SubA::*Callback)();
    void pubInner() {
        std::cout << "SubA::pubInner" << std::endl;
    }
};

int main() {
    A a;
    a.initCallback();
    // a.doCallback();
    A::Callback func = a.getCallback();
    (a.*func)();
    SubA b;
    SubA::Callback sfunc = &SubA::pubInner;
    (b.*sfunc)();

    SubA::Callback funcS;
    static_cast<B::CallbackB>(funcS);
    return 0;
}
