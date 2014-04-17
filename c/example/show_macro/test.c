#include <stdio.h>
#include <show_macro.h>

#define HELLO           do { \
                            int a = 1; \
                            a++; \
                        } while(0)

int main() {
    printf("%s\n", SHOWMACRO(HELLO));
    return 0;
}
