#include <stdio.h>
#include <stdlib.h>

int example_1() {
  // Example 1 using malloc syntax.
  char * str;
  // Function allocates memory and leaves the memory uninitialized.
  str = (char*) malloc(15);

  free(str);
  return(0);
}

int example_2() {
  //Example 2 using calloc syntax.
  char * str;
  // Function allocates memory and initializes all bits to zero.
  str = (char*) calloc(15, sizeof(char));

  free(str);
  return(0);
}