#include <stdio.h>

struct pid{

  int val;

  //void *fn;
  char fn[10];

};

struct str{

   int val;

 struct pid id[10];

};

struct str init_str = {

    .val = 1,

   .id = {         [ 0 ... 9] = {2, "haha"}     },

};

int main(void){     int i;

    for(i = 0; i < 10; i++)         printf("init_str.id[%d].val =  %d, fn = %s\n", i, init_str.id[i].val, init_str.id[i].fn);

    return 0;

}
