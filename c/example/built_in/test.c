/*
 * =====================================================================================
 *
 *       Filename:  main.c
 *
 *    Description:  Test builtin function
 *
 *        Version:  1.0
 *        Created:  01/19/2014 12:11:49 PM
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  JiaxingWu
 *   Organization:  LTT
 *
 * =====================================================================================
 */
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

int main() {
    int x = 0b00001010;
    int b = __builtin_clz(x);
    printf("x is %d bits\n", (int)(sizeof(x) << 3));
    printf("x is %d bits\n", (int)(sizeof(x) << 3));
    printf("last 1's index in x is %d\n", b);
    return 0;
}
