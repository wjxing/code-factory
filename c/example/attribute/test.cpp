#include<iostream>
using namespace std;
__attribute__((constructor)) void before_main()
{
    cout<<"Before Main"<<endl;
}
__attribute__((destructor)) void after_main()
{
    cout<<"After Main"<<endl;
}
class AAA{
    public:
        AAA(){
            cout<<"AAA construct"<<endl;
        }
        ~AAA(){
            cout<<"AAA destructor" <<endl;
        }
};
AAA A;
int main()
{
    cout<<"in main " <<endl;
    return 0;
}
