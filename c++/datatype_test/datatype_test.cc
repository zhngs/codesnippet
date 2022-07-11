#include <stdio.h>
#include <limits>

int main()
{
    printf("char length: %d\n", sizeof(char));
    printf("short length: %d\n", sizeof(short));
    printf("int length: %d\n", sizeof(int));
    printf("long length: %d\n", sizeof(long));
    printf("long long length: %d\n", sizeof(long long));
    printf("float length: %d\n", sizeof(float));
    printf("double length: %d\n", sizeof(double));
    printf("long double length: %d\n", sizeof(long double));
    printf("char* length: %d\n", sizeof(char*));
    printf("bool length: %d\n", sizeof(bool));   

    printf("data len: %d\n", std::numeric_limits<int>::max());
}