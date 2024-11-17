#include <stdio.h>
#include <omp.h>

int main() {
#pragma omp parallel 
{
    #pragma omp single
    {
        printf("Single.\n");
    }

    printf("Not single.\n");
}

return 0;
}