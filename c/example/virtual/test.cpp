#include<iostream>

typedef void(*PFun)();

class IInterface {
  public:
    virtual void asInterface() {
        std::cout << "IInterface::asInterface" << std::endl;
    }
    virtual void getDescribe() {
        std::cout << "IInterface::getDescribe" << std::endl;
    }
    virtual void getName() {
        std::cout << "IInterface::getName" << std::endl;
    }
};

class Listener {
  public:
    virtual void asInterface() {
        std::cout << "Listener::asInterface" << std::endl;
    }
};

class BpInterface : public IInterface, public Listener {
  public:
    virtual void asInterfaceSelf() {
        std::cout << "BpInterface::asInterface" << std::endl;
    }
};

int main() {
    // IInterface interface;
    BpInterface interface;
    std::cout << "sizeof int* = " << sizeof(int*) << std::endl;
    std::cout << "sizeof PFun = " << sizeof(PFun) << std::endl;
    // BpInterface interface;
    std::cout << "IInterface virtual address : " << (int*)*(int*)&interface << std::endl;
    std::cout << "IInterface first virtual function address : " << (int*)*(int*)*(int*)&interface << std::endl;
    std::cout << "IInterface second virtual function address : " << (int*)*((int*)*(int*)(&interface) + 2) << std::endl;
    std::cout << "IInterface third virtual function address : " << (int*)*((int*)*(int*)(&interface) + 4) << std::endl;
    PFun pFun = (PFun)*(int*)*(int*)&interface;
    pFun();
    pFun = (PFun)*((int*)*(int*)(&interface) + 2);
    pFun();
    pFun = (PFun)*((int*)*(int*)(&interface) + 4);
    pFun();
    pFun = (PFun)*((int*)*(int*)(&interface) + 6);
    pFun();
    std::cout << "-------------------------" << std::endl;
    Listener* listener = &interface;
    listener->asInterface();
    IInterface* iinterface = &interface;
    iinterface->asInterface();
    std::cout << "listener " << listener << std::endl;
    std::cout << "iinterface " << iinterface << std::endl;
    return 0;
}
