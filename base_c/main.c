#include <stdio.h>
#include <stdlib.h>

int main() {
  char * str;
  str = (char*) malloc(15);

  free(str);
  return(0);
}