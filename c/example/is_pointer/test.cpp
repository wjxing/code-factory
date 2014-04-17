#include <stdio.h>
#include <is_pointer.h>
struct A {
    char x;
    int age;
};

int main() {
    if (IsPointer<A>::value) {
        printf("A is a pointer\n");
    } else {
        printf("A is not a pointer\n");
    }
    if (IsPointer<A*>::value) {
        printf("A* is a pointer\n");
    } else {
        printf("A* is not a pointer\n");
    }
    if (IsPointer<typeof(a)>::value) {
        printf("a is a pointer\n");
    } else {
        printf("a is not a pointer\n");
    }
    return 0;
}
