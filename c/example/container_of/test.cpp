#include <stdio.h>
#include <container_of.h>

struct A {
    char x;
    int age;
};

int main() {
    struct A a;
    printf("             A ptr = %p\n", &a);
    printf("container_of A ptr = %p\n", container_of(&a.age, struct A, age));

    return 0;
}
