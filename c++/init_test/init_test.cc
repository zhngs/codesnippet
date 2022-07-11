#include <iostream>

class Test
{
public:
    Test()
      : a(1),
        b(2.1),
        c()
    {
    }

    int a;
    double b;
    char c[3];
};

void print_array(const char* p, int len)
{
    for (int i = 0; i < len; i++) {
        std::cout<< (int)p[i] << " ";
    }
    std::cout<<std::endl;
}

Test return_test()
{
    return Test();
}

int main()
{
    char ch[3] = "0";
    print_array(ch, sizeof(ch));
    uint8_t ch2[3] = "1";
    print_array((char*)ch2, sizeof(ch2));
    Test t = Test();
    std::cout<< t.a << " " << t.b << " ";
    print_array(t.c, sizeof(t.c));
    Test t2;
    std::cout<< t2.a << " " << t2.b << " ";
    print_array(t2.c, sizeof(t2.c));
    Test t3 = return_test();
    std::cout<< t3.a << " " << t3.b << " ";
    print_array(t3.c, sizeof(t3.c));
}